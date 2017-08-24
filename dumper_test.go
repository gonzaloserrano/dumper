package dumper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type structTest struct {
	A int
	B map[string]string
}

func TestDumps(t *testing.T) {
	param1 := "string"
	param2 := 3
	param3 := map[string]string{
		"one": "two",
	}
	param4 := &structTest{
		1,
		param3,
	}

	R(param1, param2, param3, param4)
	Y(param1, param2, param3, param4)
	G(param1, param2, param3, param4)
	B(param1, param2, param3, param4)
	C(param1, param2, param3, param4)
	M(param1, param2, param3, param4)
}

func TestName(t *testing.T) {
	assert := assert.New(t)

	// assert that returns a string if input is nil
	assert.IsType("", Name(nil))

	type testStruct struct{}
	ts := &testStruct{}
	assert.Equal("dumper.testStruct", Name(ts))

	type bar int
	var b bar = 1
	assert.Equal("bar", Name(b))
}
