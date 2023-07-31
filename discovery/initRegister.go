package discovery

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

// InitRegister 放入服务名称和服务的链接，返回etcdServiceRegister和grpc服务器
func InitRegister(key, value string) (*ServiceRegister, *grpc.Server) {
	endpoints := []string{"localhost:2379"}
	serviceRegister := &ServiceRegister{
		EtcdAddrs: endpoints,
		Lease:     5,
		Key:       key,
		Value:     value,
	}
	err := serviceRegister.NewServiceRegister()
	if err != nil {
		log.Println(err)
	}

	server := grpc.NewServer()
	return serviceRegister, server
}

// GrpcListen 监听grpc链接
func GrpcListen(server *grpc.Server, value string) {
	lis, err := net.Listen("tcp", value)
	if err != nil {
		log.Println(err)
	}
	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
