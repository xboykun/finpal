package id

import "github.com/google/uuid"

var (
	UUID = UUIDImp{}
)

type (
	UUIDImp struct{}
)

func (id UUIDImp) NewV7() (uuid.UUID, error) {
	return uuid.NewV7()
}

func (id UUIDImp) MustNewV7() uuid.UUID {
	return uuid.Must(uuid.NewV7())
}

func (id UUIDImp) Parse(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func (id UUIDImp) MustParse(s string) uuid.UUID {
	return uuid.MustParse(s)
}
