package awsutil

import (
	"testing"

	mockassert "github.com/derision-test/go-mockgen/testutil/assert"
	"github.com/go-nacelle/nacelle/mocks"
)

func TestLogAdapter(t *testing.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log("Hello")
	mockassert.CalledOnceWith(t, logger.DebugFunc, mockassert.Values("Hello"))
}

func TestLogAdapterWithArgs(t *testing.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log("Hello, %s and %s", "bar", "baz")
	mockassert.CalledOnceWith(t, logger.DebugFunc, mockassert.Values(
		"Hello, %s and %s",
		"bar",
		"baz",
	))
}

func TestLogAdapterNonStringFormat(t *testing.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log([]string{"this", "is", "dangerous"})
	mockassert.CalledOnceWith(t, logger.DebugFunc, mockassert.Values("[this is dangerous]"))
}

func TestLogAdapterNoArgs(t *testing.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log()
	mockassert.NotCalled(t, logger.DebugFunc)
}
