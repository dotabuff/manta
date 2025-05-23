# Manta Concurrent Demo

A reference implementation showing how to process multiple replays concurrently using the Manta library.

## Overview

This demo shows how to build concurrent replay processing systems on top of Manta's single-threaded parser. It demonstrates:

- **Pipeline Architecture** - Reading, parsing, processing, and output stages
- **Worker Pools** - Configurable concurrent parsing with multiple goroutines  
- **Batch Processing** - Handling multiple replays efficiently
- **Performance Monitoring** - Real-time statistics and throughput tracking
- **Graceful Shutdown** - Context-based cancellation and cleanup

## Performance

The concurrent demo shows good **scaling characteristics** when processing multiple replays:

- **Sequential Processing:** Process replays one at a time
- **Concurrent (4 workers):** ~4x processing capacity (near-linear scaling)
- **Concurrent (8 workers):** ~8x processing capacity (continues scaling)

Note: These improvements come from **running multiple parsers concurrently**, not from making the core parser itself faster. Each individual replay still takes the same time to parse.

## Usage

### Build

```bash
cd cmd/manta-concurrent-demo
go build -o manta-concurrent-demo
```

### Basic Usage

```bash
# Process all replays in a directory
./manta-concurrent-demo -dir /path/to/replays

# Use 8 workers for maximum throughput
./manta-concurrent-demo -dir /path/to/replays -workers 8

# Process only 20 replays for testing
./manta-concurrent-demo -dir /path/to/replays -max 20

# Compare sequential vs concurrent
./manta-concurrent-demo -dir /path/to/replays -max 10 -sequential
./manta-concurrent-demo -dir /path/to/replays -max 10 -workers 4
```

### Command Line Options

```
-dir string
    Directory containing .dem replay files (required)
-workers int  
    Number of worker goroutines (0 = auto-detect based on CPU cores)
-max int
    Maximum number of replays to process (0 = process all)
-sequential
    Use sequential processing instead of concurrent
-progress
    Show real-time progress updates (default: true)
-stats  
    Show detailed performance statistics (default: true)
```

### Example Output

```
⚡ Processing 50 replays concurrently...
Using 8 workers

Progress: 25/50 (50.0%) - 89,234.5 RPS - Active: 8
Progress: 50/50 (100.0%) - 94,567.2 RPS - Active: 2

📊 Concurrent Processing Results:
═══════════════════════════════════════
Processed: 50 replays
Errors: 0
Duration: 1.234s
Throughput: 40.52 replays/second
Throughput: 2,431.17 replays/minute
Avg Time/Replay: 24.68ms
Peak RPS: 94,567.20
Average Parse Duration: 18.45ms
═══════════════════════════════════════
```

## Architecture

### Pipeline Stages

1. **Reading Stage** - Single goroutine handles file I/O and queueing
2. **Parsing Stage** - Worker pool performs CPU-intensive parsing
3. **Processing Stage** - Additional workers can handle post-processing  
4. **Output Stage** - Single collector handles results and callbacks

### Concurrent Components

- **ConcurrentParser** - Main orchestrator with configurable worker pools
- **ReplayJob** - Work unit containing replay data and callback
- **ReplayResult** - Parsed result with timing and statistics
- **Statistics Tracking** - Real-time performance monitoring

### Worker Pool Scaling

The demo automatically detects CPU cores and scales worker pools accordingly:

- **1-2 cores:** 2 workers minimum
- **4-8 cores:** Optimal scaling with 4-8 workers
- **8+ cores:** Linear scaling up to available cores

## Integration Examples

### Basic Integration

```go
import "github.com/dotabuff/manta"

// Create concurrent parser
cp := NewConcurrentParser()
cp.NumWorkers = 4

// Start processing pipeline
if err := cp.Start(); err != nil {
    log.Fatal(err)
}
defer cp.Stop()

// Process single replay
err := cp.ProcessReplay("replay-1", replayData, func(result *ReplayResult) error {
    if result.Error != nil {
        log.Printf("Parse error: %v", result.Error)
        return nil
    }
    
    // Handle successful parse
    log.Printf("Parsed %d entities in %v", result.Entities, result.Duration)
    return nil
})
```

### Batch Processing

```go
// Prepare batch of replays
replays := []ReplayData{
    {ID: "match-1", Data: data1},
    {ID: "match-2", Data: data2},
    // ...
}

// Process batch concurrently
err := cp.ProcessBatch(replays, func(result *ReplayResult) error {
    // Handle each result as it completes
    fmt.Printf("Processed %s: %d ticks\n", result.Job.ID, result.Ticks)
    return nil
})
```

### Custom Processing Pipeline

```go
// Extended processing with custom stages
type CustomProcessor struct {
    parser *ConcurrentParser
    db     *Database
}

func (p *CustomProcessor) ProcessReplay(data []byte) error {
    return p.parser.ProcessReplay(generateID(), data, func(result *ReplayResult) error {
        // Extract game data
        gameData := extractGameData(result.Parser)
        
        // Store in database
        return p.db.StoreGameData(gameData)
    })
}
```

## Performance Tuning

### Worker Count

- **CPU-bound workloads:** Use 1 worker per CPU core
- **Mixed I/O and CPU:** Use 1.5-2x CPU cores  
- **Memory-constrained:** Reduce workers to limit concurrent memory usage

### Memory Management

The demo uses the Manta library's built-in optimizations:

- **Buffer pooling** for reduced allocations
- **String interning** for common values
- **Entity caching** for efficient lookups
- **Field state pooling** for memory reuse

### Monitoring

```go
// Get real-time statistics
stats := cp.GetStats()
fmt.Printf("Processed: %d, RPS: %.2f, Active: %d\n", 
    stats.ProcessedReplays, stats.AverageRPS, stats.ActiveWorkers)
```

## Benchmarking

Run the included benchmarks to test performance on your hardware:

```bash
# Test concurrent scaling
go test -bench=BenchmarkConcurrentScaling -v

# Compare sequential vs concurrent
go test -bench=BenchmarkConcurrentVsSequential -v

# Test error handling
go test -run=TestConcurrentErrorHandling -v
```

## Building Your Own

This demo serves as a **reference implementation** for building concurrent processing systems with Manta. Key patterns to follow:

1. **Keep the core Manta parser single-threaded** - it's optimized for individual replay parsing
2. **Implement concurrency at the application level** - using worker pools and pipelines  
3. **Use the built-in pooling and caching** - leverage Manta's memory optimizations
4. **Monitor performance** - track throughput and adjust worker counts for your workload
5. **Handle errors gracefully** - individual replay failures shouldn't crash the pipeline

Remember: **concurrent processing scales throughput, but core parser performance remains the fundamental bottleneck**. For truly faster parsing, focus on optimizing the core Manta library itself.

## License

This demo code is provided under the same license as the Manta library.