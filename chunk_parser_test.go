package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChunkParser(t *testing.T) {
	tokens := ScanChunks("pre{%if 1%}left{{x}}right{%endif%}post", "")
	// fmt.Println("tokens =", tokens)
	ast, err := Parse(tokens)
	require.NoError(t, err)
	fmt.Println("ast =", ast)
	err = ast.Render(os.Stdout, nil)
	require.NoError(t, err)
	fmt.Println()
	return

	for _, test := range chunkTests {
		tokens := ScanChunks(test.in, "")
		ast, err := Parse(tokens)
		require.NoError(t, err)
		actual := ast.Render(os.Stdout, nil)
		require.Equal(t, test.expected, actual)
	}
}

var chunkTests = []struct{in, expected string}{
	{"{{var}}", "value"},
	{"{{x}}", "1"},
}