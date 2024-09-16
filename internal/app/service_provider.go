package app

import (
	"tzAvito/internal/api/tender"
	"tzAvito/internal/db"
	"tzAvito/internal/repository"
	"tzAvito/internal/service"
)

type ServiceProvider struct {
	tenderRepository repository.TenderRepository
	tenderService    service.TenderService
	tenderImpl       *tender.Implementation
}

func newServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) TenderRepository() repository.TenderRepository {
	if s.tenderRepository == nil {
		s.tenderRepository = repository.NewRepository(db.DB)
	}
	return s.tenderRepository
}
func (s *ServiceProvider) TenderService() service.TenderService {
	if s.tenderService == nil {
		s.tenderService = service.NewTenderService(
			s.TenderRepository(),
		)
	}

	return s.tenderService
}
func (s *ServiceProvider) TenderImpl() *tender.Implementation {
	if s.tenderImpl == nil {
		s.tenderImpl = tender.NewImplementation(s.TenderService())
	}

	return s.tenderImpl
}
