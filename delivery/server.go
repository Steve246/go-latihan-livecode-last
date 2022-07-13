package delivery

import (
	"go_livecode_persiapan/config"
	"go_livecode_persiapan/delivery/controller"
	"go_livecode_persiapan/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func Server() *appServer {
	router := gin.Default()
	appConfig := config.NewConfig()

	infra := manager.NewInfra(appConfig)

	repoManager := manager.NewRepositoryManager(infra)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	host := appConfig.Url

	return &appServer{
		useCaseManager: useCaseManager,
		engine:         router,
		host:           host,
	}
}

func (a *appServer) iniController() {

	controller.NewMenuController(
		a.engine,
		a.useCaseManager.CrudMenuUseCase(),
	)

	controller.NewMenuPriceController(
		a.engine,
		a.useCaseManager.CrudMenuPriceUseCase(),
	)
}



func (a *appServer) Run() {
	a.iniController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
