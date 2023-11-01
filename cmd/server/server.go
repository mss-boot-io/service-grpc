package server

import (
	"context"
	"service-grpc/handlers"
	pb "service-grpc/proto"

	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/core/server/grpc"
	"github.com/spf13/cobra"

	"service-grpc/config"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/10/31 16:37:31
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/10/31 16:37:31
 */

var (
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "start server",
		Long:    "start service-http server",
		Example: "service-http server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func setup() error {
	// setup config
	config.Cfg.Init(func(srv *grpc.Server) {
		pb.RegisterHelloworldServer(srv.Server(), handlers.NewHandler())
	})
	return nil
}

func run() error {
	ctx := context.Background()

	return server.Manage.Start(ctx)
}
