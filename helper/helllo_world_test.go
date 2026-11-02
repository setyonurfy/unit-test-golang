package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("world")

	if result != "helloworld" {
		t.Error("result not hello world")
		// t.Fail()
		// panic("Test failed: expected helloworld")
	}
	fmt.Println("Test passed: expected hello world")
}
func TestHelloTayo(t *testing.T) {
	result := HelloTayo("tayo")

	if result != "hellotayo" {
		t.Fatal("result not hello tayo")
		// t,FailNow()
		// panic("Test failed: expected hellotayo")
	}

	fmt.Println("Test passed: expected hellotayo")
}

// assertion

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld(" world")
	assert.Equal(t, "hello world", result, "result should be helloworld")
	fmt.Println("Test passed with assertion: expected hello world")
}

// require
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld(" world")
	require.Equal(t, "hello world", result, "result should be helloworld")
	fmt.Println("Test passed with require: expected hello world")
}

// skip test

func TestSkip(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping test on non-windows OS")
	}
	result := HelloWorld(" world")
	require.Equal(t, "hello world", result, "result should be helloworld")
}

// main test function
func TestMain(m *testing.M) {
	fmt.Println("Before test")

	m.Run()

	fmt.Println("After test")

	// os.Exit(m.Run())
}

// subtest
func TestSubTest(t *testing.T) {
	t.Run("subtest1", func(t *testing.T) {
		result := HelloWorld(" world")
		require.Equal(t, "hello world", result, "result should be hello world")
	})
	t.Run("subtest2", func(t *testing.T) {
		result := HelloTayo("tayo")
		require.Equal(t, "hellotayo", result, "result should be hellotayo")
	})
}

// table test
func TestTableHelloWorld(t *testing.T) {
	testCases := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "subtest1",
			request:  " world",
			expected: "hello world",
		},
		{
			name:     "subtest2",
			request:  " tayo",
			expected: "hello tayo",
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "result should be "+test.expected)
		})
	}
}

// benchmak
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld(" world")
	}
}

func BenchmarkHelloTayo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloTayo("tayo")
	}
}

// sub-benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("sub-benchmark1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld(" world")
		}
	})
	b.Run("sub-benchmark2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloTayo("tayo")
		}
	})
}

// table benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "sub-benchmark1",
			request:  " world",
			expected: "hello world",
		},
		{
			name:     "sub-benchmark2",
			request:  " tayo",
			expected: "hello tayo",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result := HelloWorld(benchmark.request)
				require.Equal(b, benchmark.expected, result, "result should be "+benchmark.expected)
			}
		})
	}
}
