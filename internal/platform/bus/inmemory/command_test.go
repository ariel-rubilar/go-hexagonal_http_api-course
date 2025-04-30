package inmemory_test

import (
	"context"
	"testing"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/bus/inmemory"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/command"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCommandBus_Dispatch_Success(t *testing.T) {
	bus := inmemory.NewCommand()
	ctx := context.Background()

	cmd := new(mocks.CommandMock)
	cmd.On("Type").Return(command.Type("test_command"))
	handler := new(mocks.HandlerMock)

	handler.On("Handle", ctx, cmd).Return("result", nil)

	bus.Register(cmd.Type(), handler)

	r, err := bus.Dispatch(ctx, cmd)
	assert.Equal(t, "result", r)

	assert.NoError(t, err)
	handler.AssertExpectations(t)
	cmd.AssertExpectations(t)
}

func TestCommandBus_Dispatch_Fail(t *testing.T) {
	bus := inmemory.NewCommand()
	ctx := context.Background()

	cmd := new(mocks.CommandMock)
	cmd.On("Type").Return(command.Type("test_command"))
	handler := new(mocks.HandlerMock)

	handler.On("Handle", ctx, cmd).Return(nil, assert.AnError)

	bus.Register(cmd.Type(), handler)

	r, err := bus.Dispatch(ctx, cmd)

	assert.Nil(t, r)

	assert.Error(t, err)
	handler.AssertExpectations(t)
	cmd.AssertExpectations(t)
}

func TestCommandBus_Dispatch_Handler_Not_Found(t *testing.T) {
	bus := inmemory.NewCommand()
	ctx := context.Background()

	cmd := new(mocks.CommandMock)
	cmd.On("Type").Return(command.Type("test_command"))
	handler := new(mocks.HandlerMock)

	bus.Register(command.Type("test_command_2"), handler)

	_, err := bus.Dispatch(ctx, cmd)

	assert.NoError(t, err)
	handler.AssertNotCalled(t, "Handle", ctx, cmd)
	cmd.AssertExpectations(t)
}
