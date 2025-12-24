package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Marchel")
	if result != "Hello World Marchel" {
		t.Errorf("Hello World Marchel got %s", result)
	}
	fmt.Println("Test HelloWorld passed")
}

func TestCallName(t *testing.T) {
	result := CallName("Marchel")
	if result != "Hello World Marchel" {
		t.Fatal("Hello World Marchel got", result)
	}
	fmt.Println("Test CallNow Marchel passed")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Marchel")
	assert.Equal(t, "Hello World Marchel", result, "Result must be 'Hello World Marchel'")
	fmt.Println("Test HelloWorldAssertion passed")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Marche")
	require.Equal(t, "Hello World Marchel", result, "Result must be 'Hello World Marchel'")
	fmt.Println("Test HelloWorldAssertion passed")
}
