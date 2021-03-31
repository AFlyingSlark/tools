package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EchoTmpl(t *testing.T) {
	tmpl := &Base{}
	result, err := EchoTmpl(tmpl)

	t.Log(result)
	require.NoError(t, err, "tmpl")
}
