package uuid

import (
	"github.com/google/uuid"
)

type UUID struct {
	uuid uuid.UUID
}

func (u *UUID) UnmarshalJSON(in []byte) error {
	st := string(in)
	if st == "" {
		u.uuid = uuid.Nil
		return nil
	}

	uid, err := uuid.Parse(st)
	if err != nil {
		return err
	}

	u.uuid = uid
	return nil
}

func (u *UUID) MarshalJSON() (out []byte, err error) {
	if u.uuid == uuid.Nil {
		return []byte(""), nil
	}

	return []byte(u.uuid.String()), err
}

func New() UUID {
	return UUID{
		uuid: uuid.New(),
	}
}

func (u UUID) String() string {
	return u.uuid.String()
}

func Parse(s string) (parsed UUID, err error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return UUID{
			uuid.Nil,
		}, err
	}

	return UUID{
		id,
	}, nil
}
