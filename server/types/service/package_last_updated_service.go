package service

import "time"

type PackageLastUpdatedTuple struct {
	PackageName string
	LastUpdated time.Time
}

type PackageLastUpdatedService interface {
	GetPackagesLastUpdated(programsClonePath string) ([]PackageLastUpdatedTuple, error)
}
