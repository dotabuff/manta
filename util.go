package manta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
)

var debugMode, traceMode bool

func init() {
	if os.Getenv("DEBUG") != "" {
		debugMode = true
	}
}

var (
	_sprintf = fmt.Sprintf
	_sdump   = spew.Sdump
)

// Convert a string to an int32
func atoi32(s string) (int32, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return int32(n), nil
}

// printf only if debugging
func _debugf(format string, args ...interface{}) {
	if debugMode {
		args = append([]interface{}{_caller(2)}, args...)
		fmt.Printf("%s: "+format+"\n", args...)
	}
}

// printf only if tracing
func _tracef(format string, args ...interface{}) {
	if traceMode {
		args = append([]interface{}{_caller(2)}, args...)
		fmt.Printf("%s: "+format+"\n", args...)
	}
}

// error with printf syntax
func _errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

// panic with printf syntax
func _panicf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

// dump named object only if debugging
func _dump(label string, args ...interface{}) {
	if debugMode {
		fmt.Printf("%s: %s", _caller(2), label)
		spew.Dump(args...)
	}
}

// dumps a given byte buffer to the given fixture filename
func _dump_fixture(filename string, buf []byte) {
	fmt.Printf("writing fixture %s...\n", filename)
	if err := ioutil.WriteFile("./fixtures/"+filename, buf, 0644); err != nil {
		panic(err)
	}
}

// reads a byte buffer from the given fixture filename
func _read_fixture(filename string) []byte {
	buf, err := ioutil.ReadFile("./fixtures/" + filename)
	if err != nil {
		panic(err)
	}
	return buf
}

// marshal a proto.Message to bytes
func _proto_marshal(obj proto.Message) []byte {
	buf, err := proto.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return buf
}

// marshal an interface{} to JSON bytes
func _json_marshal(obj interface{}) []byte {
	buf, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}
	return buf
}

// Returns the name of the calling function
func _caller(n int) string {
	if pc, _, _, ok := runtime.Caller(n); ok {
		fns := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		return fns[len(fns)-1]
	}

	return "unknown"
}

func log2(n int) int {
	return int(math.Log(float64(n))/math.Log(2)) + 1
}
