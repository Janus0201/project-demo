package main

import (
	"log"
	"net"
	"time"

	"github.com/MrLittle05/Gomall/app/user/biz/dal"
	"github.com/MrLittle05/Gomall/app/user/conf"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		klog.Error(err.Error())
	}

	// 初始化mysql和redis
	dal.Init()

	opts := kitexInit()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// 使用 consul 创建一个新的服务注册中心，需要从配置文件（conf.yaml）传入注册中心的地址 (IP : Port)
	/* 这个地址让你的服务 (微服务的地址通过框架传给Consul) 能够连接到 Consul，进而完成服务的注册和发现功能。
	   这使得微服务架构中的各个服务可以动态地发现彼此，并有效地进行通信。*/
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		log.Fatal(err)
	}

	/* server.WithRegistry(r)：这是一个选项函数（option function），用于配置服务器实例。
	   它告诉服务器使用指定的 Consul 注册器 r 来进行服务注册与发现。*/
	/* opts 列表：opts 是一个选项列表，通常用于收集各种配置选项。
	   通过 append 函数，我们将新的选项添加到这个列表中。*/
	// 最终，在创建服务器实例时，可以将这个选项列表传递给服务器，以便应用所有配置。
	opts = append(opts, server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
