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