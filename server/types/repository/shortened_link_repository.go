package repository

import "time"

type ShortenedLink interface {
	GetLinkId() string
	SetLinkId(string)

	GetLink() string
	SetLink(string)

	GetVisits() uint
	SetVisits(uint)

	GetCreatedAt() time.Time
	SetCreatedAt(time.Time)

	GetUpdatedAt() time.Time
	SetUpdatedAt(time.Time)

	GetID() uint
	SetID(uint)
}

type ShortenedLinkRepository interface {
	FindOneByLinkId(linkId string) (ShortenedLink, error)
	IncrementVisits(id uint) error
}
