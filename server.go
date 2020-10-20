package main

import (
	"google.golang.org/grpc"
)

func main() {
    srvCfg := server.Config{Protocol: "tcp", Host: "localhost", Port: "9000"}
    srv := grpc.NewServer(srvCfg, creatingService, addingService, listingService)
    log.Printf("gRPC server running at %s://%s:%s ...\n", srvCfg.Protocol, srvCfg.Host, srvCfg.Port)
    log.Fatal(srv.Serve())
}


type grpcServer struct {} {

// implementacion llamadas remotas
func NewServerLogistica() grpc.ServerServicioLogistica {
	return &ServerLogisticaHandler{}
}

func (s ServerLogisticaHandler) CPyme(ctx context.Context, req *grpc.ComprarEnPyme) (*grpc.ComprarEnPymeResp, error) {
	// TODO: Implement the create action
}

func (s ServerLogisticaHandler) CRetail(ctx context.Context, req *grpc.ComprarEnRetail) (*grpc.ComprarEnPymeResp, error) {
	// TODO: Implement the add action
}

}

type Server interface {
	// implementacion
	Serve() error
}

// configuracion server
type Config struct {
	Protocol string
	Host     string
	Port     string
}

type grpcServer struct {
	config          server.Config
}

func NewServer(
	config server.Config,
) server.Server {
	return &grpcServer{config: config}
}

func (s *grpcServer) Serve() error {
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	listener, err := net.Listen(s.config.Protocol, addr)
	if err != nil {
		return err
	}

	srv := googlegrpc.NewServer()
	serviceServer := NewServerLogistica()
	grpc.RegisterServerServicioLogistica(srv, serviceServer)

	if err := srv.Serve(listener); err != nil {
		return err
	}

	return nil
}