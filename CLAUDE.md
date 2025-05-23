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

## Performance Optimization Notes

### Completed Optimizations (32.6% total improvement achieved)

**Phase 0: Go Version Update (28.6% improvement)**
- Updated Go 1.16.3 → 1.21.13 for immediate runtime performance gains
- Zero code changes required, excellent ROI
- Always prioritize infrastructure updates first

**Phase 1: Buffer Management (5.5% additional improvement)**
- **Stream buffer pooling** (`stream.go`): Eliminated frequent buffer reallocations with intelligent 2x growth strategy
- **String table key history pooling** (`string_table.go`): Reused slices for string table parsing  
- **Compression buffer pooling** (`compression.go`): Shared Snappy decompression buffers across codebase
- **Key insight**: Pool overhead is minimal compared to allocation reduction benefits

**Phase 2: Memory Management (0.4% additional improvement)**
- **Field state pooling** (`field_state.go`): Size-class pools (8/16/32/64/128 elements) for field state objects
- **Entity field cache pooling** (`entity.go`): Reused fpCache and fpNoop maps with proper lifecycle management
- **Key insight**: Incremental improvements provide cumulative benefits under sustained load

**Phase 3: Core Optimizations (1.2% additional improvement)**
- **Field path pool optimization** (`field_path.go`): Pre-warmed with 100 field paths, optimized reset function
- **Bit reader optimizations** (`reader.go`): Pre-computed bit masks, varint fast paths, single-bit optimization
- **String interning** (`reader.go`): Automated interning for strings ≤32 chars with 10K cache limit
- **Key insight**: Core path optimizations provide compounding benefits for high-throughput scenarios

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

### Performance Impact Summary
- **Original baseline (Go 1.16.3):** 1163ms, 51 replays/minute
- **After Phase 0-3:** 784ms, 77 replays/minute  
- **Exceeded primary <800ms target with 32.6% total improvement**

### Optimization Lessons Learned

1. **Go version updates provide massive ROI** - should always be first priority
2. **Buffer pooling works well** for frequently allocated/deallocated objects
3. **sync.Pool is efficient** for reducing allocation pressure in hot paths
4. **Smart growth strategies** (2x) reduce reallocation frequency
5. **Shared utilities** (compression.go) provide consistent optimization across codebase
6. **Benchmark frequently** - small improvements compound significantly

### Memory Pool Patterns Used

```go
// Effective pool pattern used throughout optimizations
var bufferPool = &sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, initialCapacity)
    },
}

// Usage pattern
buf := bufferPool.Get().([]byte)
defer bufferPool.Put(buf)
buf = buf[:0] // Reset length, keep capacity
```

### Next Optimization Targets
- Field state memory pooling for entity updates
- Entity field cache optimization  
- Protobuf message pooling for callback system