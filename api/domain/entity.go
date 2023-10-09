package domain

import "github.com/google/uuid"

type Entities []Entity

type Entity struct {
	Id int64
	User UserID
	Name UserName
}

type UserID struct {
	V uuid.UUID
}

type UserName struct {
	V string
}

type CreateJson struct {
	User UserID
	Name UserName
}