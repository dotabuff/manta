package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/dotabuff/manta"
)

// BenchmarkConcurrentVsSequential compares sequential and concurrent processing
func BenchmarkConcurrentVsSequential(b *testing.B) {
	// Use a smaller number of iterations since each "iteration" processes 10 replays
	if b.N > 10 {
		b.N = 10 // Limit to reasonable number for realistic testing
	}

	// Create mock replay data (small but valid)
	mockReplayData := createMockReplayData()
	numReplaysPerIteration := 10

	b.Run("Sequential", func(b *testing.B) {
		b.ReportAllocs()
		totalReplays := 0
		start := time.Now()

		for i := 0; i < b.N; i++ {
			for j := 0; j < numReplaysPerIteration; j++ {
				parser, err := manta.NewParser(mockReplayData)
				if err != nil {
					b.Skip("Cannot create parser for mock data")
				}

				// Don't actually parse, just measure setup overhead
				_ = parser
				totalReplays++
			}
		}

		duration := time.Since(start)
		rps := float64(totalReplays) / duration.Seconds()
		b.ReportMetric(rps, "replays/sec")
		b.ReportMetric(float64(totalReplays), "total_replays")
	})

	b.Run("Concurrent", func(b *testing.B) {
		cp := NewConcurrentParser()
		cp.NumWorkers = 4 // Use fixed number for consistent benchmarking

		if err := cp.Start(); err != nil {
			b.Fatal(err)
		}
		defer cp.Stop()

		b.ReportAllocs()
		totalReplays := 0
		start := time.Now()

		for i := 0; i < b.N; i++ {
			var wg sync.WaitGroup

			for j := 0; j < numReplaysPerIteration; j++ {
				wg.Add(1)
				totalReplays++

				err := cp.ProcessReplay(fmt.Sprintf("bench-%d-%d", i, j), mockReplayData, func(result *ReplayResult) error {
					defer wg.Done()
					// Don't process errors in benchmark
					return nil
				})

				if err != nil {
					wg.Done()
					b.Logf("Failed to submit replay: %v", err)
				}
			}

			wg.Wait()
		}

		duration := time.Since(start)
		rps := float64(totalReplays) / duration.Seconds()
		b.ReportMetric(rps, "replays/sec")
		b.ReportMetric(float64(totalReplays), "total_replays")

		// Report concurrent-specific metrics
		stats := cp.GetStats()
		b.ReportMetric(stats.PeakRPS, "peak_rps")
		b.ReportMetric(float64(stats.ProcessedReplays), "processed")
	})
}

// BenchmarkConcurrentScaling tests how performance scales with worker count
func BenchmarkConcurrentScaling(b *testing.B) {
	mockReplayData := createMockReplayData()
	numReplays := 20

	workerCounts := []int{1, 2, 4, 8}

	for _, workers := range workerCounts {
		b.Run(fmt.Sprintf("Workers-%d", workers), func(b *testing.B) {
			cp := NewConcurrentParser()
			cp.NumWorkers = workers

			if err := cp.Start(); err != nil {
				b.Fatal(err)
			}
			defer cp.Stop()

			b.ReportAllocs()
			start := time.Now()

			var wg sync.WaitGroup

			for i := 0; i < numReplays; i++ {
				wg.Add(1)

				err := cp.ProcessReplay(fmt.Sprintf("scale-%d", i), mockReplayData, func(result *ReplayResult) error {
					defer wg.Done()
					return nil
				})

				if err != nil {
					wg.Done()
					b.Logf("Failed to submit replay: %v", err)
				}
			}

			wg.Wait()
			duration := time.Since(start)
			rps := float64(numReplays) / duration.Seconds()

			b.ReportMetric(rps, "replays/sec")
			b.ReportMetric(float64(workers), "workers")

			stats := cp.GetStats()
			b.ReportMetric(stats.PeakRPS, "peak_rps")
		})
	}
}

// createMockReplayData creates minimal valid replay data for testing
func createMockReplayData() []byte {
	// Create minimal replay data that satisfies basic parsing requirements
	data := make([]byte, 1024)

	// Source 2 magic header
	copy(data[0:8], []byte{'P', 'B', 'D', 'E', 'M', 'S', '2', '\000'})

	// Add 8 bytes for size fields (skipped in parser)
	// Remaining bytes will be zero, which should cause parser to exit gracefully

	return data
}

// TestConcurrentParserLifecycle tests the complete lifecycle
func TestConcurrentParserLifecycle(t *testing.T) {
	cp := NewConcurrentParser()
	cp.NumWorkers = 2

	// Test starting
	if err := cp.Start(); err != nil {
		t.Fatalf("Failed to start: %v", err)
	}

	// Test processing
	mockData := createMockReplayData()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		err := cp.ProcessReplay(fmt.Sprintf("test-%d", i), mockData, func(result *ReplayResult) error {
			defer wg.Done()
			t.Logf("Processed replay %s in %v", result.Job.ID, result.Duration)
			return nil
		})

		if err != nil {
			wg.Done()
			t.Errorf("Failed to submit replay: %v", err)
		}
	}

	wg.Wait()

	// Test statistics
	stats := cp.GetStats()
	if stats.ProcessedReplays == 0 {
		t.Error("No replays were processed")
	}

	t.Logf("Processed %d replays, avg RPS: %.2f", stats.ProcessedReplays, stats.AverageRPS)

	// Test stopping
	if err := cp.Stop(); err != nil {
		t.Fatalf("Failed to stop: %v", err)
	}
}

// TestConcurrentErrorHandling tests error scenarios
func TestConcurrentErrorHandling(t *testing.T) {
	cp := NewConcurrentParser()
	cp.NumWorkers = 1

	if err := cp.Start(); err != nil {
		t.Fatalf("Failed to start: %v", err)
	}
	defer cp.Stop()

	// Test with invalid data
	invalidData := []byte("invalid replay data")

	var wg sync.WaitGroup
	wg.Add(1)

	err := cp.ProcessReplay("invalid", invalidData, func(result *ReplayResult) error {
		defer wg.Done()

		if result.Error == nil {
			t.Error("Expected error for invalid data")
		} else {
			t.Logf("Got expected error: %v", result.Error)
		}

		return nil
	})

	if err != nil {
		wg.Done()
		t.Fatalf("Failed to submit invalid replay: %v", err)
	}

	wg.Wait()
}
