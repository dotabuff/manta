package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/dotabuff/manta"
)

func main() {
	var (
		replayDir     = flag.String("dir", "", "Directory containing .dem replay files")
		workers       = flag.Int("workers", 0, "Number of worker goroutines (0 = auto-detect)")
		sequential    = flag.Bool("sequential", false, "Use sequential processing instead of concurrent")
		maxReplays    = flag.Int("max", 0, "Maximum number of replays to process (0 = all)")
		showStats     = flag.Bool("stats", true, "Show processing statistics")
		showProgress  = flag.Bool("progress", true, "Show progress during processing")
	)
	flag.Parse()

	if *replayDir == "" {
		log.Fatal("Please specify a replay directory with -dir")
	}

	// Find all replay files
	replayFiles, err := findReplayFiles(*replayDir)
	if err != nil {
		log.Fatalf("Error finding replay files: %v", err)
	}

	if len(replayFiles) == 0 {
		log.Fatal("No .dem files found in the specified directory")
	}

	if *maxReplays > 0 && len(replayFiles) > *maxReplays {
		replayFiles = replayFiles[:*maxReplays]
	}

	fmt.Printf("Found %d replay files to process\n", len(replayFiles))

	if *sequential {
		processSequentially(replayFiles, *showStats, *showProgress)
	} else {
		processConcurrently(replayFiles, *workers, *showStats, *showProgress)
	}
}

func findReplayFiles(dir string) ([]string, error) {
	var files []string
	
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".dem") {
			files = append(files, path)
		}
		
		return nil
	})
	
	return files, err
}

func processSequentially(files []string, showStats, showProgress bool) {
	fmt.Printf("\n🔄 Processing %d replays sequentially...\n", len(files))
	
	start := time.Now()
	var processed int
	var totalTicks uint32
	var totalEntities int
	var errors int
	
	for i, file := range files {
		if showProgress && i%10 == 0 {
			fmt.Printf("Progress: %d/%d (%.1f%%)\n", i, len(files), float64(i)/float64(len(files))*100)
		}
		
		data, err := os.ReadFile(file)
		if err != nil {
			errors++
			continue
		}
		
		parser, err := manta.NewParser(data)
		if err != nil {
			errors++
			continue
		}
		
		err = parser.Start()
		if err != nil {
			errors++
			continue
		}
		
		processed++
		totalTicks += parser.Tick
		// Count entities by iterating through entity map
		entityCount := 0
		for i := int32(0); i < 2048; i++ {
			if parser.FindEntity(i) != nil {
				entityCount++
			}
		}
		totalEntities += entityCount
	}
	
	duration := time.Since(start)
	
	if showStats {
		printStats("Sequential Processing", processed, errors, duration, totalTicks, totalEntities)
	}
}

func processConcurrently(files []string, workers int, showStats, showProgress bool) {
	fmt.Printf("\n⚡ Processing %d replays concurrently...\n", len(files))
	
	cp := NewConcurrentParser()
	if workers > 0 {
		cp.NumWorkers = workers
	}
	
	fmt.Printf("Using %d workers\n", cp.NumWorkers)
	
	if err := cp.Start(); err != nil {
		log.Fatalf("Failed to start concurrent parser: %v", err)
	}
	defer cp.Stop()
	
	start := time.Now()
	var wg sync.WaitGroup
	var mu sync.Mutex
	var processed int
	var totalTicks uint32
	var totalEntities int
	var errors int
	
	// Progress tracking
	var progressCount int
	if showProgress {
		go func() {
			ticker := time.NewTicker(2 * time.Second)
			defer ticker.Stop()
			
			for range ticker.C {
				stats := cp.GetStats()
				mu.Lock()
				current := progressCount
				mu.Unlock()
				
				if current >= len(files) {
					break
				}
				
				fmt.Printf("Progress: %d/%d (%.1f%%) - %.1f RPS - Active: %d\n", 
					current, len(files), 
					float64(current)/float64(len(files))*100,
					stats.AverageRPS,
					stats.ActiveWorkers)
			}
		}()
	}
	
	// Process all files
	for i, file := range files {
		wg.Add(1)
		
		// Read file data
		data, err := os.ReadFile(file)
		if err != nil {
			mu.Lock()
			errors++
			progressCount++
			mu.Unlock()
			wg.Done()
			continue
		}
		
		// Submit for concurrent processing
		err = cp.ProcessReplay(fmt.Sprintf("replay-%d", i), data, func(result *ReplayResult) error {
			defer wg.Done()
			
			mu.Lock()
			defer mu.Unlock()
			
			progressCount++
			
			if result.Error != nil {
				errors++
				return nil
			}
			
			processed++
			totalTicks += result.Ticks
			totalEntities += result.Entities
			
			return nil
		})
		
		if err != nil {
			mu.Lock()
			errors++
			progressCount++
			mu.Unlock()
			wg.Done()
			log.Printf("Failed to submit replay %s: %v", file, err)
		}
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	if showStats {
		printStats("Concurrent Processing", processed, errors, duration, totalTicks, totalEntities)
		
		// Show concurrent-specific stats
		stats := cp.GetStats()
		fmt.Printf("Peak RPS: %.2f\n", stats.PeakRPS)
		fmt.Printf("Average Parse Duration: %.2fms\n", float64(stats.TotalDuration.Nanoseconds())/float64(stats.ProcessedReplays)/1e6)
	}
}

func printStats(method string, processed, errors int, duration time.Duration, totalTicks uint32, totalEntities int) {
	fmt.Printf("\n📊 %s Results:\n", method)
	fmt.Printf("═══════════════════════════════════════\n")
	fmt.Printf("Processed: %d replays\n", processed)
	fmt.Printf("Errors: %d\n", errors)
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Throughput: %.2f replays/second\n", float64(processed)/duration.Seconds())
	fmt.Printf("Throughput: %.2f replays/minute\n", float64(processed)/duration.Minutes())
	
	if processed > 0 {
		fmt.Printf("Avg Time/Replay: %.2fms\n", float64(duration.Nanoseconds())/float64(processed)/1e6)
		fmt.Printf("Total Ticks: %d (avg: %.0f/replay)\n", totalTicks, float64(totalTicks)/float64(processed))
		fmt.Printf("Total Entities: %d (avg: %.0f/replay)\n", totalEntities, float64(totalEntities)/float64(processed))
	}
	fmt.Printf("═══════════════════════════════════════\n")
}