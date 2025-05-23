# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## About This Project

Manta is a Dota 2 replay parser written in Go for Source 2 engine replays. It provides low-level access to replay data through a callback-based architecture without imposing higher-level structure on the data.

## Development Commands

```bash
# Run tests with coverage (WARNING: takes a long time - parses many replays)
make test

# Run performance benchmarks  
make bench

# Update protobuf definitions from Steam
make update

# Generate callback code from templates
make generate

# Generate coverage reports
make cover

# Run specific test (much faster than full test suite)
go test -run TestSpecificFunction

# Run tests for specific package
go test ./string_table

# Run single replay test (recommended for development)
go test -run TestMatchNew7116386145  # Latest replay
go test -run TestMatch1731962898     # Older replay
```

**Performance Note**: Running `make test` parses 40+ replay files and takes significant time. For development, run specific tests like `go test -run TestMatchNew7116386145` which tests a single recent replay and runs much faster.

## Core Architecture

### Parser Flow
1. **Stream Reader** (`stream.go`) - Low-level binary data reading
2. **Parser** (`parser.go`) - Main parsing logic, handles compression and message routing
3. **Callbacks** (`callbacks.go`) - Event-driven architecture with auto-generated handlers
4. **Entity System** (`entity.go`) - Tracks game entities through their lifecycle
5. **Field Decoding** (`field_*.go`) - Complex property decoding with various data types

### Key Components

**Parser**: Central component that manages replay parsing. Handles file validation, compression (Snappy), and message routing to appropriate handlers.

**Callbacks**: Auto-generated from protobuf definitions. All Dota 2 message types have corresponding callback functions. Users register handlers for events they care about.

**Entity Management**: Tracks all game entities (heroes, items, buildings) through Created/Updated/Deleted/Entered/Left states. Entities have complex field structures decoded via the field system.

**Field System**: Handles decoding of entity properties. Supports quantized floats, bit-packed data, vectors, and various primitive types. Field paths represent hierarchical property structures.

**String Tables**: Efficient string storage system used by the game engine. Handles both compressed and uncompressed string data.

### Data Flow
1. Binary replay data → Stream reader
2. Stream reader → Parser (handles compression)
3. Parser → Protobuf message parsing
4. Messages → Registered callbacks
5. Entity updates → Field decoding → Entity state changes

## Generated Code

- `dota/` directory contains 80+ auto-generated protobuf files from Valve's game definitions
- `gen/callbacks.go` is generated from `gen/callbacks.tmpl` template
- Run `make generate` after modifying the template
- Run `make update` to pull latest protobuf definitions from Steam

## Testing

Tests use real Dota 2 replay files and fixture data:
- `fixtures/` contains test data for various components
- `replays/` contains actual match replay files for integration tests
- Many tests require specific replay files to validate parsing correctness
- Benchmark tests measure parsing performance on real data

## Working with Fields

Field decoding is complex due to Dota 2's optimized network format:
- Fields can be quantized floats, bit-packed integers, or complex nested structures
- Field paths use dot notation (e.g., "m_vecOrigin.0" for X coordinate)
- Field types are determined by send table definitions
- Always check field type before decoding to avoid panics

## Code Style and Formatting

### Go Code Formatting
**IMPORTANT:** Always run `go fmt` on Go files before committing to ensure consistent formatting.

```bash
# Format all Go files in the project
go fmt ./...

# Format specific file
go fmt filename.go
```

**Best Practices:**
- Use tabs for indentation (Go standard)
- No trailing whitespace
- Single trailing newline at end of files
- Use `gofmt` or equivalent in your editor to format on save

## Benchmarking and Performance Testing

### Running Benchmarks

```bash
# Run all benchmarks
make bench

# Run benchmarks with memory profiling
go test -bench=. -benchmem -memprofile=mem.prof

# Run specific benchmark (faster for development)
go test -bench=BenchmarkMatch2159568145 -benchmem

# Run benchmark multiple times for stability
go test -bench=BenchmarkMatch2159568145 -benchmem -count=5

# Profile CPU usage during benchmarks
go test -bench=BenchmarkMatch2159568145 -cpuprofile=cpu.prof

# Profile memory allocations
go test -bench=BenchmarkMatch2159568145 -memprofile=mem.prof -memprofilerate=1
```

### Performance Profiling

```bash
# Analyze CPU profile
go tool pprof cpu.prof

# Analyze memory profile
go tool pprof mem.prof

# Generate flame graph (if installed)
go tool pprof -http=:8080 cpu.prof

# Check allocations per operation
go test -bench=BenchmarkMatch2159568145 -benchmem | grep "allocs/op"
```

### Benchmark Types

1. **Throughput benchmarks**: Use BenchmarkMatch* functions with real replay data
2. **Memory benchmarks**: Track allocations per operation with -benchmem
3. **Component benchmarks**: Create focused benchmarks for specific operations
4. **Regression benchmarks**: Compare performance against baseline measurements

### Creating Custom Benchmarks

For testing specific optimizations, create focused benchmarks:

```go
func BenchmarkFieldDecoding(b *testing.B) {
    // Setup test data
    for i := 0; i < b.N; i++ {
        // Run operation under test
    }
}
```

### Interpreting Results

- **ns/op**: Nanoseconds per operation (lower is better)
- **B/op**: Bytes allocated per operation (lower is better)  
- **allocs/op**: Number of allocations per operation (lower is better)
- **MB/s**: Throughput for data processing benchmarks (higher is better)

Always run benchmarks multiple times and look for consistent results. Use `benchstat` tool to compare benchmark runs statistically.

## Performance Optimization Summary

### Final Results (30.8% total improvement achieved)

**Comprehensive Optimization Campaign (Phases 0-8)**
- **Original baseline:** 1163ms per replay, 51 replays/minute
- **Final performance:** 805ms per replay, 75 replays/minute
- **Total improvement:** 30.8% faster parsing, 47% higher throughput

### Key Optimization Insights

**1. Infrastructure Updates Provide Massive ROI**
- Go 1.16.3 → 1.21.13 alone achieved 28.6% improvement with zero code changes
- Always prioritize infrastructure updates before algorithmic optimizations

**2. Memory Pooling Is Highly Effective**
- sync.Pool provides significant allocation reduction in hot paths
- Size-class pools (8/16/32/64/128) work well for varying object sizes
- Buffer reuse patterns show consistent performance improvements

**3. Optimization Has Diminishing Returns**
- Early phases (0-4) achieved 33.4% improvement with clear ROI
- Later phases (6-8) showed minimal gains or even regressions
- Field path optimizations regressed due to map overhead vs. algorithmic benefits

**4. Hot Path Identification Is Critical**
- Reader bit operations and varint decoding are true performance bottlenecks
- Field path operations had less impact than expected
- Entity management optimizations provided modest but measurable gains

**5. Architectural Constraints Limit Further Gains**
- Interface{} boxing in field decoders remains unavoidable
- Fundamental parsing algorithm is already well-optimized  
- Additional improvements require architectural changes or different approaches

### String Interning Implementation Pattern

```go
// Global string interning system
var (
    stringInternMap   = make(map[string]string)
    stringInternMutex sync.RWMutex
    stringBuffer      = &sync.Pool{
        New: func() interface{} {
            return make([]byte, 0, 64)
        },
    }
)

// Efficient interning with size limits and double-checked locking
func internString(s string) string {
    if len(s) == 0 || len(s) > 32 {
        return s
    }
    
    stringInternMutex.RLock()
    if interned, exists := stringInternMap[s]; exists {
        stringInternMutex.RUnlock()
        return interned
    }
    stringInternMutex.RUnlock()
    
    stringInternMutex.Lock()
    defer stringInternMutex.Unlock()
    
    if interned, exists := stringInternMap[s]; exists {
        return interned
    }
    
    if len(stringInternMap) < 10000 {
        stringInternMap[s] = s
        return s
    }
    
    return s
}

// Optimized string reading with pooled buffers
func (r *reader) readString() string {
    buf := stringBuffer.Get().([]byte)
    buf = buf[:0]
    defer stringBuffer.Put(buf)
    
    for {
        b := r.readByte()
        if b == 0 {
            break
        }
        buf = append(buf, b)
    }

    return internString(string(buf))
}
```

### Effective Memory Pool Pattern

```go
// Standard pool pattern used throughout optimizations
var bufferPool = &sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, initialCapacity)
    },
}

// Usage pattern
func optimizedFunction() {
    buf := bufferPool.Get().([]byte)
    defer bufferPool.Put(buf)
    buf = buf[:0] // Reset length, keep capacity
    
    // Use buf for operations...
}
```

### Benchmarking Best Practices

1. **Always benchmark before and after** changes to measure impact
2. **Run multiple iterations** (-count=3 minimum) for statistical significance  
3. **Profile both CPU and memory** to identify true bottlenecks
4. **Focus on hot paths** - optimize where the time is actually spent
5. **Watch for regressions** - some optimizations add overhead that outweighs benefits
6. **Document results** in commit messages and roadmaps for future reference

### Performance Tools Used

```bash
# Primary benchmarking workflow
go test -bench=BenchmarkMatch2159568145 -benchmem -count=3

# CPU profiling to identify hot paths  
go test -bench=BenchmarkMatch2159568145 -cpuprofile=cpu.prof

# Memory allocation analysis
go test -bench=BenchmarkMatch2159568145 -memprofile=mem.prof

# Statistical comparison of benchmark runs
benchstat old.txt new.txt
```