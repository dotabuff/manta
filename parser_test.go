package manta

import (
	"bytes"
	"io"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type timeoutReader struct {
	r         io.Reader
	maxBytes  int
	readBytes int
}

var _ io.Reader = &timeoutReader{}

func newTimeoutReader(data []byte, maxBytes int) *timeoutReader {
	return &timeoutReader{r: bytes.NewReader(data), maxBytes: maxBytes}
}

func (r *timeoutReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.readBytes += n
	if r.readBytes >= r.maxBytes {
		return n, timeoutError{}
	}
	return n, err
}

type timeoutError struct{}

var _ net.Error = timeoutError{}

func (e timeoutError) Error() string {
	return "timed out"
}

func (e timeoutError) Timeout() bool {
	return true
}

func (e timeoutError) Temporary() bool {
	return true
}

func TestStreamingReaderTimeout(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping test in CI environment")
	}

	assert := assert.New(t)

	// get data
	data := mustGetReplayData("2159568145", "https://s3-us-west-2.amazonaws.com/manta.dotabuff/2159568145.dem")
	reader := newTimeoutReader(data, 10000)

	// create a new parser with the streaming response
	parser, err := NewStreamParser(reader)
	if err != nil {
		t.Fatalf("unable to create parser: %s", err)
	}

	// expect the parser to fail after a bit under 2 seconds with a timeout error
	err = parser.Start()
	netErr, ok := err.(net.Error)
	assert.True(ok)
	assert.True(netErr.Timeout())
}

func TestStreamingReaderUnexpectedEOF(t *testing.T) {
	assert := assert.New(t)

	// get data
	data := mustGetReplayData("2159568145", "https://s3-us-west-2.amazonaws.com/manta.dotabuff/2159568145.dem")

	// cut the data at a random point
	reader := bytes.NewReader(data[:6666])

	// begin parsing the incomplete data
	parser, err := NewStreamParser(reader)
	if err != nil {
		t.Fatalf("unable to create parser: %s", err)
	}

	// expect an unexpected EOF error
	err = parser.Start()
	assert.Equal(io.ErrUnexpectedEOF, err)
}
