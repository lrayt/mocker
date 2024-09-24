package utils

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Log(strings.Replace("${WorkDir}/assets", "${WorkDir}11", "1234", 1))
}
