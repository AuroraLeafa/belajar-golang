package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_HelloWorldTrue(t *testing.T) {
	result := HelloWorld("World")
	if result != "Hello World" {
		t.Error("Result Must Be Hello World")
	}
	fmt.Println("Test_HelloWorld_True DONE")
}
func Test_HelloReffTrue(t *testing.T) {
	result := HelloWorld("Reff")
	if result != "Hello Reff" {
		t.Fatal("Result Must Be Hello Reff")
	}
	fmt.Println("Test_HelloReff_True DONE")
}

func Test_HelloWorld(t *testing.T) {
	result := HelloWorld("World")
	if result != "Helloz World" {
		t.Error("Result Must Be Hello World")
	}
	fmt.Println("Test_HelloWorld DONE")
}
func Test_HelloReff(t *testing.T) {
	result := HelloWorld("Reff")
	if result != "Hellos Reff" {
		t.Fatal("Result Must Be Hello Reff")
	}
	fmt.Println("Test_HelloReff DONE")
}

func Test_HelloWorldAssert(t *testing.T) {
	result := HelloWorld("Reff")
	assert.Equal(t, "Hello Reff", result, "Result Must be Equal!")
	fmt.Println("Test Done")
}

func Test_HelloWorldRequire(t *testing.T) {
	result := HelloWorld("Reff")
	require.Equal(t, "Hello Reff", result, "Result Must be Equal!")
	fmt.Println("Test Done")
}

func Test_Skip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip()
	}
	result := HelloWorld("xD")
	assert.Equal(t, "Hello Reff", result, "Result Must be Equal!")
	fmt.Println("Test Done")
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE Unit Test")
	m.Run()
	fmt.Println("AFTER Unit Test")
}

func TestSubTest(t *testing.T) {
	t.Run("Reff", func(t *testing.T) {
		result := HelloWorld("Reff")
		assert.Equal(t, "Hello Reff", result, "Result must be Hello Reff")
	})
	t.Run("Sugianto", func(t *testing.T) {
		result := HelloWorld("Sugianto")
		assert.Equal(t, "Hello xxx", result, "Result must be Hello Sugianto")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Reff",
			request:  "Reff",
			expected: "Hello Reff",
		},
		{
			name:     "Sugg",
			request:  "Sugg",
			expected: "Hello Sugg",
		},
		{
			name:     "Sugianto",
			request:  "Sugianto",
			expected: "Hello Sugianto",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloworld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Oi")
	}
}
func BenchmarkHelloReff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Reff")
	}
}
func BenchmarkSub(b *testing.B) {
	b.Run("Reff", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Reff")
		}
	})
	b.Run("Dragunov", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dragunov")
		}
	})
	b.Run("Svv", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Svv")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Reff",
			request: "Reff",
		},
		{
			name:    "Sugg",
			request: "Sugg",
		},
		{
			name:    "Sugianto",
			request: "Sugianto",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}

}
