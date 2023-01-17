package day2_1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHelloTom(t *testing.T) {

	output := HelloTom("a")
	expectOutput := "tom"
	assert.Equal(t, expectOutput, output)
}
func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
