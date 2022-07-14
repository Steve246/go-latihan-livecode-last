package manager

import "go_livecode_persiapan/service"

type ServiceManager interface {
	LopeiService() *service.LopeiService
}

type serviceManager struct {
	lopeiService *service.LopeiService
}

func (r *serviceManager) LopeiService() *service.LopeiService {
	return r.lopeiService
}

func NewServiceManager(repoManager RepositoryManager) ServiceManager {
	return &serviceManager{
		lopeiService: service.NewLopeiService(repoManager.LopeiRepo()),
	}
}
