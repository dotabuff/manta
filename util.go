package manta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
)

func init() {
	spew.Config.SortKeys = true
}

var debugLevel uint

func init() {
	if os.Getenv("DEBUG") != "" {
		debugLevel = 1
	}
	if os.Getenv("TRACE") != "" {
		debugLevel = 10
	}
}

var (
	_sprintf = fmt.Sprintf
)

// Convert a string to an int32
func atoi32(s string) (int32, error) {
	n, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, err
	}
	return int32(n), nil
}

// debugging level check
func v(level uint) bool {
	return level <= debugLevel
}

// printf only if debugging
func _debugf(format string, args ...interface{}) {
	if v(1) {
		args = append([]interface{}{_caller(2)}, args...)
		fmt.Printf("%s: "+format+"\n", args...)
	}
}

func _printf(format string, args ...interface{}) {
	args = append([]interface{}{_caller(2)}, args...)
	fmt.Printf("%s: "+format+"\n", args...)
}

// error with printf syntax
func _errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

// panic with printf syntax
func _panicf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

// dump named object
func _dump(label string, args ...interface{}) {
	fmt.Printf("%s: %s", _caller(2), label)
	spew.Dump(args...)
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

func _nameof(i interface{}) string {
	ss := strings.Split(strings.Replace(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".", "/", -1), "/")
	return ss[len(ss)-1]
}

func _typeof(i interface{}) string {
	if i == nil {
		return "nil"
	}
	return reflect.TypeOf(i).String()
}
