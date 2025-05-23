package manta

import (
	"sync"
	
	"github.com/golang/snappy"
)

// Pool for compression/decompression buffers to reduce allocations
var compressionPool = &sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 1024*64) // 64KB initial capacity
	},
}

// DecodeSnappy decompresses data using a pooled buffer
func DecodeSnappy(src []byte) ([]byte, error) {
	buf := compressionPool.Get().([]byte)
	defer compressionPool.Put(buf)
	
	result, err := snappy.Decode(buf[:0], src)
	if err != nil {
		return nil, err
	}
	
	// Copy result since we're returning the buffer to pool
	output := make([]byte, len(result))
	copy(output, result)
	return output, nil
}