package Database

import (
	uuid "github.com/satori/go.uuid"
)

type UserEntity struct {
	Id uuid.UUID `db:"Id"`
}
