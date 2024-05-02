package parser

import (
	"errors"
	"testing"
	"time"

	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/service"
	"pacstall.dev/webserver/utils/expect"
)

type mockPackageLastUpdatedService struct {
	mockGetPackagesLastUpdatedCalled int
	mockGetPackagesLastUpdated       func(programsClonePath string) ([]service.PackageLastUpdatedTuple, error)
}

func (s *mockPackageLastUpdatedService) GetPackagesLastUpdated(programsClonePath string) ([]service.PackageLastUpdatedTuple, error) {
	out, err := s.mockGetPackagesLastUpdated(programsClonePath)
	return out, err
}

func Test_ParserService_setLastUpdatedAt_Empty(t *testing.T) {
	service := New(
		config.PacstallProgramsConfiguration{},
		config.ServerConfiguration{},
		config.RepologyConfiguration{},
		nil,
		nil,
		nil,
		&mockPackageLastUpdatedService{
			mockGetPackagesLastUpdated: func(programsClonePath string) ([]service.PackageLastUpdatedTuple, error) {
				return []service.PackageLastUpdatedTuple{}, nil
			},
		},
	)

	err := service.setLastUpdatedAt([]*pac.Script{}, "")
	expect.NoError(t, err)
}

func Test_ParserService_setLastUpdatedAt_Error(t *testing.T) {
	service := New(
		config.PacstallProgramsConfiguration{},
		config.ServerConfiguration{},
		config.RepologyConfiguration{},
		nil,
		nil,
		nil,
		&mockPackageLastUpdatedService{
			mockGetPackagesLastUpdated: func(programsClonePath string) ([]service.PackageLastUpdatedTuple, error) {
				return []service.PackageLastUpdatedTuple{}, errors.New("dummy error")
			},
		},
	)

	err := service.setLastUpdatedAt([]*pac.Script{}, "")
	expect.AnyError(t, err)
}

func Test_ParserService_setLastUpdatedAt_UpdatesPackages(t *testing.T) {
	pkg1Time := time.Now()
	pkg2Time := time.Now().Add(-24 * time.Hour)

	service := New(
		config.PacstallProgramsConfiguration{},
		config.ServerConfiguration{},
		config.RepologyConfiguration{},
		nil,
		nil,
		nil,
		&mockPackageLastUpdatedService{
			mockGetPackagesLastUpdated: func(programsClonePath string) ([]service.PackageLastUpdatedTuple, error) {
				return []service.PackageLastUpdatedTuple{
					{PackageName: "package-1", LastUpdated: pkg1Time},
					{PackageName: "package-2", LastUpdated: pkg2Time},
				}, nil
			},
		},
	)

	data := []*pac.Script{
		{PackageName: "package-1"},
		{PackageName: "package-2"},
	}

	err := service.setLastUpdatedAt(data, "")

	expect.NoError(t, err)
	expect.Equals(t, "last updated at set", []*pac.Script{
		{PackageName: "package-1", LastUpdatedAt: pkg1Time},
		{PackageName: "package-2", LastUpdatedAt: pkg2Time},
	}, data)

}
