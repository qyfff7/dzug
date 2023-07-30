package rpc

import (
	"context"
	"dzug/discovery"
	"dzug/idl/relation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	// 启用etcd服务发现
	var endpoints = []string{"localhost:2379"}
	ser := discovery.NewServiceDiscovery(endpoints)
	defer ser.Close()
	ser.WatchService("relation")

	// grpc监听与UserClient初始化
	target := ser.GetServices()["relation"]
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials())) // grpc.WithInsecure() // 不使用TLS认证
	if err != nil {
		log.Fatalf("net.Connect err : %v", err)
	}
	defer conn.Close()
	RelationClient := relation.NewDouyinRelationActionServiceClient(conn)

	r, err := RelationClient.DouyinRelationAction(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}
