package delivery

import (
	"go_livecode_persiapan/manager"
	"go_livecode_persiapan/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type LopeiGrpcServer struct {
	netListen      net.Listener
	server         *grpc.Server
	serviceManager manager.ServiceManager
}

func (p *LopeiGrpcServer) Run() {
	service.RegisterLopeiPaymentServer(p.server, p.serviceManager.LopeiService())
	log.Println("Server runs", p.netListen.Addr().String())
	if err := p.server.Serve(p.netListen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// func ServerLopei() *LopeiGrpcServer {
// 	lopeiGrpcServer := new(LopeiGrpcServer)
// 	c := config.NewConfig()
// 	lis, err := net.Listen("tcp", c.Url)

// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	repoManager := manager.NewRepositoryManager()

// 	lopeiGrpcServer.netListen = lis
// 	lopeiGrpcServer.server = grpcServer
// 	lopeiGrpcServer.serviceManager = manager.NewServiceManager(repoManager)
// 	return lopeiGrpcServer
// }
