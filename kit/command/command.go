package command

import "context"

type Type string

type Command interface {
	Type() Type
}

type Handler interface {
	Handle(context.Context, Command) error
}

type Bus interface {
	Dispatch(context.Context, Command) error
	Register(Type, Handler)
}
