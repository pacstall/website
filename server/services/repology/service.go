package repology

import (
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/types/repository"
)

type RepologyService struct {
	repologyConfiguration             config.RepologyConfiguration
	repologyProjectRepository         repository.RepologyProjectRepository
	repologyProjectProviderRepository repository.RepologyProjectProviderRepository
}

func New(
	repologyConfiguration config.RepologyConfiguration,
	repologyProjectRepository repository.RepologyProjectRepository,
	repologyProjectProviderRepository repository.RepologyProjectProviderRepository,
) *RepologyService {
	repologyService := &RepologyService{}

	repologyService.repologyConfiguration = repologyConfiguration
	repologyService.repologyProjectRepository = repologyProjectRepository
	repologyService.repologyProjectProviderRepository = repologyProjectProviderRepository

	return repologyService
}
