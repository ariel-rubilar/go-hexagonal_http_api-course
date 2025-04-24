package uuid

import "github.com/google/uuid"

type UUID string

var Nil = UUID(uuid.Nil.String())

func New() UUID {
	uuid := uuid.New()
	return UUID(uuid.String())
}

func (u UUID) String() string {
	return string(u)
}

func Parse(s string) (UUID, error) {
	uuid, err := uuid.Parse(s)
	if err != nil {
		return Nil, err
	}
	return UUID(uuid.String()), nil
}
