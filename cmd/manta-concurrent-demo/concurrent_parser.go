package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
	
	"github.com/dotabuff/manta"
)

// ConcurrentParser provides high-throughput parsing using pipeline concurrency
type ConcurrentParser struct {
	// Configuration
	NumWorkers    int // Number of worker goroutines for parsing
	BufferSize    int // Size of pipeline buffers
	MaxBatchSize  int // Maximum replays to process in parallel
	
	// Pipeline stages
	readChan    chan *ReplayJob
	parseChan   chan *ReplayJob
	resultChan  chan *ReplayResult
	
	// Worker management
	ctx       context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
	
	// Statistics
	stats *ConcurrentStats
}

// ReplayJob represents a single replay to be processed
type ReplayJob struct {
	ID       string
	Data     []byte
	Callback func(*ReplayResult) error
	StartTime time.Time
}

// ReplayResult contains the parsed result and timing information
type ReplayResult struct {
	Job       *ReplayJob
	Parser    *manta.Parser
	Error     error
	Duration  time.Duration
	Entities  int
	Ticks     uint32
}

// ConcurrentStats tracks performance metrics for the concurrent parser
type ConcurrentStats struct {
	mu                sync.RWMutex
	ProcessedReplays  int64
	TotalDuration     time.Duration
	AverageRPS        float64 // Replays per second
	PeakRPS           float64
	ActiveWorkers     int
	QueuedJobs        int
	lastUpdateTime    time.Time
	lastProcessed     int64
}

// NewConcurrentParser creates a new concurrent parser with optimal defaults
func NewConcurrentParser() *ConcurrentParser {
	numWorkers := runtime.GOMAXPROCS(0) // Use all available cores
	if numWorkers < 2 {
		numWorkers = 2
	}
	
	ctx, cancel := context.WithCancel(context.Background())
	
	cp := &ConcurrentParser{
		NumWorkers:   numWorkers,
		BufferSize:   numWorkers * 4, // 4x buffer for smooth pipeline flow
		MaxBatchSize: numWorkers * 2, // 2x workers for batching
		
		// Pipeline channels
		readChan:   make(chan *ReplayJob, numWorkers*4),
		parseChan:  make(chan *ReplayJob, numWorkers*4),
		resultChan: make(chan *ReplayResult, numWorkers*4),
		
		// Context
		ctx:    ctx,
		cancel: cancel,
		
		// Statistics
		stats: &ConcurrentStats{
			lastUpdateTime: time.Now(),
		},
	}
	
	return cp
}

// Start initializes the concurrent processing pipeline
func (cp *ConcurrentParser) Start() error {
	// Start reader stage (single goroutine for IO coordination)
	cp.wg.Add(1)
	go cp.readerStage()
	
	// Start parsing workers (CPU-intensive stage)
	for i := 0; i < cp.NumWorkers; i++ {
		cp.wg.Add(1)
		go cp.parsingWorker(i)
	}
	
	// Start result collector (single goroutine for output coordination)
	cp.wg.Add(1)
	go cp.resultCollector()
	
	// Start statistics updater
	cp.wg.Add(1)
	go cp.statsUpdater()
	
	return nil
}

// Stop gracefully shuts down the concurrent parser
func (cp *ConcurrentParser) Stop() error {
	// Signal shutdown
	cp.cancel()
	
	// Close input channel to drain pipeline
	close(cp.readChan)
	
	// Wait for all workers to finish
	cp.wg.Wait()
	
	return nil
}

// ProcessReplay submits a single replay for concurrent processing
func (cp *ConcurrentParser) ProcessReplay(id string, data []byte, callback func(*ReplayResult) error) error {
	job := &ReplayJob{
		ID:        id,
		Data:      data,
		Callback:  callback,
		StartTime: time.Now(),
	}
	
	select {
	case cp.readChan <- job:
		cp.updateQueueStats(1)
		return nil
	case <-cp.ctx.Done():
		return fmt.Errorf("concurrent parser is shutting down")
	default:
		return fmt.Errorf("pipeline buffer full - too many concurrent replays")
	}
}

// ProcessBatch processes multiple replays concurrently with optimal batching
func (cp *ConcurrentParser) ProcessBatch(replays []ReplayData, callback func(*ReplayResult) error) error {
	batchSize := len(replays)
	if batchSize > cp.MaxBatchSize {
		// Process in smaller batches to avoid overwhelming the pipeline
		for i := 0; i < batchSize; i += cp.MaxBatchSize {
			end := i + cp.MaxBatchSize
			if end > batchSize {
				end = batchSize
			}
			
			batch := replays[i:end]
			for _, replay := range batch {
				if err := cp.ProcessReplay(replay.ID, replay.Data, callback); err != nil {
					return err
				}
			}
			
			// Small delay to prevent overwhelming the system
			time.Sleep(10 * time.Millisecond)
		}
		return nil
	}
	
	// Process entire batch at once
	for _, replay := range replays {
		if err := cp.ProcessReplay(replay.ID, replay.Data, callback); err != nil {
			return err
		}
	}
	
	return nil
}

// ReplayData represents input data for batch processing
type ReplayData struct {
	ID   string
	Data []byte
}

// GetStats returns current performance statistics
func (cp *ConcurrentParser) GetStats() ConcurrentStats {
	cp.stats.mu.RLock()
	defer cp.stats.mu.RUnlock()
	return *cp.stats
}

// readerStage handles the reading/queueing stage of the pipeline
func (cp *ConcurrentParser) readerStage() {
	defer cp.wg.Done()
	defer close(cp.parseChan)
	
	for {
		select {
		case job := <-cp.readChan:
			if job == nil {
				return // Channel closed
			}
			
			// Forward to parsing stage
			select {
			case cp.parseChan <- job:
				cp.updateQueueStats(-1)
			case <-cp.ctx.Done():
				return
			}
			
		case <-cp.ctx.Done():
			return
		}
	}
}

// parsingWorker handles CPU-intensive parsing in the worker pool
func (cp *ConcurrentParser) parsingWorker(workerID int) {
	defer cp.wg.Done()
	
	for {
		select {
		case job := <-cp.parseChan:
			if job == nil {
				return // Channel closed
			}
			
			// Update active worker count
			cp.updateWorkerStats(1)
			
			// Parse the replay
			result := cp.parseReplay(job)
			
			// Forward result
			select {
			case cp.resultChan <- result:
			case <-cp.ctx.Done():
				cp.updateWorkerStats(-1)
				return
			}
			
			cp.updateWorkerStats(-1)
			
		case <-cp.ctx.Done():
			return
		}
	}
}

// parseReplay performs the actual parsing work
func (cp *ConcurrentParser) parseReplay(job *ReplayJob) *ReplayResult {
	startTime := time.Now()
	
	// Create parser instance for this replay
	parser, err := manta.NewParser(job.Data)
	if err != nil {
		return &ReplayResult{
			Job:      job,
			Error:    err,
			Duration: time.Since(startTime),
		}
	}
	
	// Parse the replay
	err = parser.Start()
	duration := time.Since(startTime)
	
	return &ReplayResult{
		Job:      job,
		Parser:   parser,
		Error:    err,
		Duration: duration,
		Entities: 0, // Entity count not accessible from external packages
		Ticks:    parser.Tick,
	}
}

// resultCollector handles the output stage of the pipeline
func (cp *ConcurrentParser) resultCollector() {
	defer cp.wg.Done()
	
	for {
		select {
		case result := <-cp.resultChan:
			if result == nil {
				return // Channel closed
			}
			
			// Update statistics
			cp.updateProcessingStats(result)
			
			// Call user callback
			if result.Job.Callback != nil {
				if err := result.Job.Callback(result); err != nil {
					// Log callback error but continue processing
					fmt.Printf("Callback error for replay %s: %v\n", result.Job.ID, err)
				}
			}
			
		case <-cp.ctx.Done():
			return
		}
	}
}

// statsUpdater periodically updates performance statistics
func (cp *ConcurrentParser) statsUpdater() {
	defer cp.wg.Done()
	
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			cp.updateRPSStats()
		case <-cp.ctx.Done():
			return
		}
	}
}

// Helper methods for statistics tracking
func (cp *ConcurrentParser) updateQueueStats(delta int) {
	cp.stats.mu.Lock()
	cp.stats.QueuedJobs += delta
	cp.stats.mu.Unlock()
}

func (cp *ConcurrentParser) updateWorkerStats(delta int) {
	cp.stats.mu.Lock()
	cp.stats.ActiveWorkers += delta
	cp.stats.mu.Unlock()
}

func (cp *ConcurrentParser) updateProcessingStats(result *ReplayResult) {
	cp.stats.mu.Lock()
	cp.stats.ProcessedReplays++
	cp.stats.TotalDuration += result.Duration
	cp.stats.mu.Unlock()
}

func (cp *ConcurrentParser) updateRPSStats() {
	cp.stats.mu.Lock()
	defer cp.stats.mu.Unlock()
	
	now := time.Now()
	elapsed := now.Sub(cp.stats.lastUpdateTime).Seconds()
	if elapsed > 0 {
		processed := cp.stats.ProcessedReplays - cp.stats.lastProcessed
		currentRPS := float64(processed) / elapsed
		
		cp.stats.AverageRPS = float64(cp.stats.ProcessedReplays) / time.Since(cp.stats.lastUpdateTime).Seconds()
		if currentRPS > cp.stats.PeakRPS {
			cp.stats.PeakRPS = currentRPS
		}
		
		cp.stats.lastProcessed = cp.stats.ProcessedReplays
	}
}