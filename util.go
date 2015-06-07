package manta

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var debugMode bool

func init() {
	if os.Getenv("DEBUG") != "" {
		debugMode = true
	}
}

// printf only if debugging
func _debugf(format string, args ...interface{}) {
	if debugMode {
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

// Returns the name of the calling function
func _caller(n int) string {
	if pc, _, _, ok := runtime.Caller(n); ok {
		fns := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		return fns[len(fns)-1]
	}

	return "unknown"
}
