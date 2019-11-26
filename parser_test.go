package manta

import (
	"bytes"
	"compress/bzip2"
	"io"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStreamingReaderTimeout(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping test in CI environment")
	}

	assert := assert.New(t)

	// create a http client with a timeout too low to get the entire replay
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// request the replay, the timeout should be enough for it to begin successfully
	resp, err := client.Get("http://replay111.valve.net/570/2526759078_1821332888.dem.bz2")
	if err != nil {
		t.Fatalf("unable to get remote replay: %s", err)
	}
	defer resp.Body.Close()

	// create a new parser with the streaming response
	parser, err := NewStreamParser(bzip2.NewReader(resp.Body))
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
