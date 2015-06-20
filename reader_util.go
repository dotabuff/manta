package manta

import (
	"regexp"
	"strings"

	"github.com/fatih/color"
)

var (
	colorBold   = color.New(color.Bold).SprintFunc()
	colorError  = color.New(color.FgRed).SprintFunc()
	colorValue  = color.New(color.FgCyan).SprintFunc()
	colorZero   = color.New(color.Faint).SprintFunc()
	printableRe = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
)

func isPrintable(s string) bool {
	return printableRe.MatchString(s)
}

// Something capable of dumping data in a given format.
type readerDumper struct {
	name string
	fmt  string
	zero interface{}
	fn   func(r *reader) interface{}
}

// The list of name columns with dumpers that will be dumped.
var readerDumpers = []readerDumper{
	{"binary", "%-1v", "0", func(r *reader) interface{} { return r.readBits(1) }},
	{"uint8", "%-3v", "0", func(r *reader) interface{} { return r.readBits(8) }},
	{"var32", "%-11v", "0", func(r *reader) interface{} { return r.readVarInt32() }},
	{"varu32", "%-10v", "0", func(r *reader) interface{} { return r.readVarUint32() }},
	{"varu64", "%-20v", "0", func(r *reader) interface{} { return r.readVarUint64() }},
	//{"uint64", "%-30v", "0", func(r *reader) interface{} { return r.readLeUint64() }},
	//{"uint4", "%-5v", "0", func(r *reader) interface{} { return r.readBits(4) }},
	//{"uint5", "%-5v", "0", func(r *reader) interface{} { return r.readBits(5) }},
	//{"uint6", "%-5v", "0", func(r *reader) interface{} { return r.readBits(6) }},
	//{"uint10", "%-5v", "0", func(r *reader) interface{} { return r.readBits(10) }},
	//{"uint11", "%-5v", "0", func(r *reader) interface{} { return r.readBits(11) }},
	{"float32", "%-12v", "0", func(r *reader) interface{} { return r.readFloat32() }},
	{"string", "%v", "-", func(r *reader) interface{} {
		if s := r.readString(); isPrintable(s) {
			return s
		}
		return "-"
	}},
}

// Dumps the rest of the buffer.
func (r *reader) dumpRemaining() {
	r.dumpBits(r.remBits())
}

// Dumps a given number of bits.
func (r *reader) dumpBits(n int) {
	o := r.pos
	for i := r.pos; i < (o + n); i++ {
		r.pos = i
		x := r.pos
		line := _sprintf("@ bit %05d (byte %03d + %d) ", r.pos, r.pos/8, r.pos%8)
		for _, d := range readerDumpers {
			val := func() (out string) {
				var v interface{}
				defer func() {
					if err := recover(); err != nil {
						v = "ERR"
					}
					r.pos = x
					out = _sprintf(d.fmt, v)
				}()
				v = d.fn(r)
				return
			}()

			colorFn := colorValue
			if val == "ERR" {
				colorFn = colorError
			} else if strings.TrimSpace(val) == d.zero {
				colorFn = colorZero
			}

			line += _sprintf(" | %s: %s", d.name, colorFn(val))
		}
		_debugf(line)
	}
	r.pos = o
}
