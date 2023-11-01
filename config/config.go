package config

/*
 * @Author: lwnmengjing
 * @Date: 2023/11/1 09:42:35
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/11/1 09:42:35
 */
import (
	"log/slog"
	"os"

	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/core/server/grpc"
	"github.com/mss-boot-io/mss-boot/pkg/config"
	"github.com/mss-boot-io/mss-boot/pkg/config/source"
)

var Cfg Config

type Config struct {
	Server config.GRPC    `yaml:"server" json:"server"`
	Listen *config.Listen `yaml:"listen" json:"listen"`
	Logger *config.Logger `yaml:"logger" json:"logger"`
}

func (e *Config) Init(handler func(srv *grpc.Server)) {
	opts := []source.Option{
		source.WithDir("config"),
		source.WithProvider(source.Local),
	}
	err := config.Init(e, opts...)
	if err != nil {
		slog.Error("config init failed", slog.Any("err", err))
		os.Exit(-1)
	}

	e.Logger.Init()

	runnable := []server.Runnable{
		e.Server.Init(handler, grpc.WithID("service-grpc")),
	}

	if e.Listen != nil {
		runnable = append(runnable, e.Listen.Init())
	}

	server.Manage.Add(runnable...)
}

func (e *Config) OnChange() {
	e.Logger.Init()
	slog.Info("!!! config change and reload")
}
