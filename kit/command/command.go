package command

import "context"

type Type string

type Command interface {
	Type() Type
}

type Handler interface {
	Handle(context.Context, Command) (any, error)
}

type Bus interface {
	Dispatch(context.Context, Command) (any, error)
	Register(Type, Handler)
}
