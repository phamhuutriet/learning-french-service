package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AddOne(a int) int {
	return a + 1
}

// func TestAddOne(t *testing.T) {
// 	var (
// 		input  = 1
// 		output = 2
// 	)

// 	assert.Equal(t, output, AddOne(input+1))
// }

func TestRequire(t *testing.T) {
	require.Equal(t, 1, 2)
	fmt.Println("After require")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 1, 2)
	fmt.Println("After assert")
}
