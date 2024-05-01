package service

import (
	"time"

	"pacstall.dev/webserver/types/pac"
)

type PackageCacheService interface {
	FindByName(name string) (*pac.Script, error)
	FindByMaintainer(maintainer string) (*pac.Script, error)
	GetAll() []*pac.Script
	LastModified() time.Time
	Update(scripts []*pac.Script)
}
