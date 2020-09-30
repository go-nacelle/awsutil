package awsutil

import (
	"testing"

	mockassert "github.com/derision-test/go-mockgen/testutil/assert"
	logmocks "github.com/go-nacelle/log/mocks"
	"github.com/go-nacelle/nacelle/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLogAdapter(t *testing.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log("Hello")
	mockassert.CalledOnceWith(t, logger.DebugFunc, func(t assert.TestingT, call interface{}) bool {
		return call.(logmocks.LoggerDebugFuncCall).Arg0 == "Hello" // TODO - ergonomics
	})
}

func TestLogAdapterWithArgs(t *testing.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log("Hello, %s and %s", "bar", "baz")
	mockassert.CalledOnceWith(t, logger.DebugFunc, func(t assert.TestingT, call interface{}) bool {
		c := call.(logmocks.LoggerDebugFuncCall)
		return c.Arg0 == "Hello, %s and %s" // TODO - ergonomics
		// TODO - && c.Arg1 == "bar" && c.Arg2 == "baz"
	})
}

func TestLogAdapterNonStringFormat(t *testing.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log([]string{"this", "is", "dangerous"})
	mockassert.CalledOnceWith(t, logger.DebugFunc, func(t assert.TestingT, call interface{}) bool {
		return call.(logmocks.LoggerDebugFuncCall).Arg0 == "[this is dangerous]" // TODO - ergonomics
	})
}

func TestLogAdapterNoArgs(t *testing.T) {
	logger := mocks.NewMockLogger()
	adapter := NewAWSLogAdapter(logger)
	adapter.Log()
	mockassert.NotCalled(t, logger.DebugFunc)
}
