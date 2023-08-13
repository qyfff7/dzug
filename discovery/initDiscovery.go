package discovery

import (
	"dzug/conf"
	"dzug/protos/favorite"
	"dzug/protos/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	SerDiscovery serviceDiscovery

	UserClient     user.ServiceClient
	FavoriteClient favorite.DouyinFavoriteActionServiceClient
)

// InitDiscovery 初始化一个服务发现程序
func InitDiscovery() {
	endpoints := conf.Config.EtcdConfig.Addr              // etcd地址
	SerDiscovery = serviceDiscovery{EtcdAddrs: endpoints} // 放入etcd地址
	err := SerDiscovery.newServiceDiscovery()             // 实例化
	if err != nil {
		zap.L().Error("启动服务发现失败: " + err.Error())
		return
	}
}

// LoadClient 加载etcd客户端调用实例，每一次客户端调用一个方法都会调用这个方法
// 先去etcd中拿去现在的链接，再去通过grpc进行远程调用
func LoadClient(serviceName string, client any) {
	conn, err := connectService(serviceName) // 找到grpc连接链接
	if err != nil {
		zap.L().Error("grpc连接服务: " + serviceName + "失败, error: " + err.Error())
		return
	}

	switch c := client.(type) {
	case *user.ServiceClient:
		*c = user.NewServiceClient(conn)
	case *favorite.DouyinFavoriteActionServiceClient:
		*c = favorite.NewDouyinFavoriteActionServiceClient(conn)
	default:
		zap.L().Info("没有这种类型的服务")
	}
}

// connectService 通过服务名字找到对应的链接
// 比如，传入user，会找到etcd上存储的user的链接
func connectService(serviceName string) (conn *grpc.ClientConn, err error) {
	err = SerDiscovery.watchService("") // ！！！监视所有的服务
	if err != nil {
		zap.L().Error("未找到服务地址：" + err.Error())
		return nil, err
	}
	addr, err := SerDiscovery.getServiceByKey(serviceName)
	if err != nil {
		zap.L().Error("未找到服务地址：" + err.Error())
		return nil, err
	}
	conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return
}
