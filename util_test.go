package manta

import (
	"compress/bzip2"
	"fmt"
	"io/ioutil"
	"net/http"
)

func mustGetReplayData(name string, url string) []byte {
	buf, err := getReplayData(name, url)
	if err != nil {
		panic(err)
	}
	return buf
}

func getReplayData(name string, url string) ([]byte, error) {
	path := fmt.Sprintf("replays/%s.dem", name)
	if data, err := ioutil.ReadFile(path); err == nil {
		return data, nil
	}

	fmt.Printf("downloading replay %s from %s...\n", name, url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Return an error if we don't get a 200
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid status %d", resp.StatusCode)
	}

	var data []byte
	if url[len(url)-3:] == "bz2" {
		data, err = ioutil.ReadAll(bzip2.NewReader(resp.Body))
	} else {
		data, err = ioutil.ReadAll(resp.Body)
	}

	if err != nil {
		return nil, err
	}

	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return nil, err
	}

	fmt.Printf("downloaded replay %s from %s to %s\n", name, url, path)

	return data, nil
}
