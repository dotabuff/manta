package manta

import (
	"compress/bzip2"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type bzip2Reader struct {
	f  io.ReadCloser
	br io.Reader
}

var _ io.ReadCloser = &bzip2Reader{}

func (r *bzip2Reader) Read(p []byte) (int, error) {
	return r.br.Read(p)
}

func (r *bzip2Reader) Close() error {
	return r.f.Close()
}

func getReplayReader(name string) (io.ReadCloser, error) {
	path := fmt.Sprintf("fixtures/replays/%s.dem.bz2", name)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &bzip2Reader{f: f, br: bzip2.NewReader(f)}, nil
}

func mustGetReplayReader(name string) io.ReadCloser {
	r, err := getReplayReader(name)
	if err != nil {
		panic(err)
	}
	return r
}

func getReplayData(name string) ([]byte, error) {
	r, err := getReplayReader(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

func mustGetReplayData(name string) []byte {
	buf, err := getReplayData(name)
	if err != nil {
		panic(err)
	}
	return buf
}
