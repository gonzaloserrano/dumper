// package dumper adds colors to "github.com/davecgh/go-spew/spew" using
// "github.com/fatih/color"
package dumper

import (
	"fmt"
	"os"
	"reflect"
	"runtime"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
)

// Dump dumps the parameters without color
func Dump(params ...interface{}) {
	dump(params...)
}

// R dumps the parameters in red
func R(params ...interface{}) {
	colorDump(color.FgRed, params...)
}

// Y dumps the parameters in yellow
func Y(params ...interface{}) {
	colorDump(color.FgYellow, params)
}

// G dumps the parameters in green
func G(params ...interface{}) {
	colorDump(color.FgGreen, params)
}

// B dumps the parameters in blue
func B(params ...interface{}) {
	colorDump(color.FgBlue, params)
}

// C dumps the parameters in cyan
func C(params ...interface{}) {
	colorDump(color.FgCyan, params)
}

// M dumps the parameters in magenta
func M(params ...interface{}) {
	colorDump(color.FgMagenta, params)
}

// RK dumps the parameters in red and exit
func RK(params ...interface{}) {
	colorDump(color.FgRed, params...)
	os.Exit(1)
}

// YK dumps the parameters in yellow and exit
func YK(params ...interface{}) {
	colorDump(color.FgYellow, params)
	os.Exit(1)
}

// GK dumps the parameters in green and exit
func GK(params ...interface{}) {
	colorDump(color.FgGreen, params)
	os.Exit(1)
}

// BK dumps the parameters in blue and exit
func BK(params ...interface{}) {
	colorDump(color.FgBlue, params)
	os.Exit(1)
}

// CK dumps the parameters in cyan and exit
func CK(params ...interface{}) {
	colorDump(color.FgCyan, params)
	os.Exit(1)
}

// MK dumps the parameters in magenta and exit
func MK(params ...interface{}) {
	colorDump(color.FgMagenta, params)
	os.Exit(1)
}

func colorDump(col color.Attribute, params ...interface{}) {
	color.Set(col)
	_, fn, line, _ := runtime.Caller(2)
	fmt.Printf("[%s:%d]\n", fn, line)
	dump(params...)
	color.Unset()
}

func dump(params ...interface{}) {
	spew.Dump(params...)
}

// String returns a string representation of the parameters
func String(params ...interface{}) string {
	return spew.Sprintf("%v", params)
}

// Name returns the struct or pointer to struct name, or unknown for other types
func Name(i interface{}) string {
	if i == nil {
		return "unknown"
	}

	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		return val.Elem().Type().String()
	}
	return reflect.TypeOf(i).Name()
}
