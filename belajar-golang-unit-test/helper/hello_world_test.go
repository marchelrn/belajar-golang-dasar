package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Marchel)",
			request:  "Marchel",
			expected: "Hello World Marchel",
		},

		{
			name:     "HelloWorld(Matthew)",
			request:  "Matthew",
			expected: "Hello World Matthew",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("subtest", func(t *testing.T) {
		result := HelloWorld("Marchel")
		require.Equal(t, "Hello World Marchel", result, "Result must be 'Hello World Marchel'")
	})

	t.Run("Matthew", func(t *testing.T) {
		result := HelloWorld("Matthew")
		require.Equal(t, "Hello World Matthew", result, "Result must be 'Hello World Marchel'")
	})
}

func TestMain(m *testing.M) {
	go fmt.Println("Before Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows")
	}
	result := HelloWorld("Marchel")
	require.Equal(t, result, "Hello World")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Marchel")
	if result != "Hello World Marchel" {
		t.Errorf("Hello World Marchel got %s", result)
	}
	fmt.Println("Test HelloWorld passed")
}

func TestHelloWorldMarchel(t *testing.T) {
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
	result := HelloWorld("Marchel")
	require.Equal(t, "Hello World Marchel", result, "Result must be 'Hello World Marchel'")
	fmt.Println("Test HelloWorldAssertion passed")
}
