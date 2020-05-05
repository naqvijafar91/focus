package focus

import (
	"strconv"
	"testing"
)

func TestFourDigitCodeGen(t *testing.T) {
	codeGen := NewFourDigitCodeGenerator()
	for i := 0; i < 10; i++ {
		code := codeGen.Generate()
		codeInt, err := strconv.Atoi(code)
		if err != nil {
			t.Error(err)
		}
		if codeInt < 1000 || codeInt>9999 {
			t.Error("Range is not correct")
		}
	}
}
