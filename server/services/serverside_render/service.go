package ssr

import "pacstall.dev/webserver/types/service"

type ServerSideRenderService struct {
	packageCacheService service.PackageCacheService
}

func New(packageCacheService service.PackageCacheService) *ServerSideRenderService {
	s := &ServerSideRenderService{}

	s.packageCacheService = packageCacheService

	return s
}

func (s *ServerSideRenderService) EnableServerSideRendering() {
	s.registerPacscriptSSRData()
	s.registerPacscriptListSSRData()
}
