package app

import (
	"context"
	"fmt"
	"github.com/dsxack/go/v2/config"
	kitconfig "github.com/dsxack/go/v2/kit/config"
	kitcli "github.com/dsxack/go/v2/kit/transport/cli"
	kithttp "github.com/dsxack/go/v2/kit/transport/http"
	"github.com/google/wire"
	"github.com/mitchellh/go-homedir"
	"github.com/oklog/run"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/trace"
	"gocloud.dev/server"
	"gocloud.dev/server/health"
	"gocloud.dev/server/requestlog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Context context.Context

type App struct{}

func Init(
	_ kitcli.Commands,
	logger logrus.FieldLogger,
	server *server.Server,
	config kitconfig.Config,
) (*App, error) {
	group := run.Group{}
	group.Add(serverRunGroup(logger, server, config.HTTP))
	group.Add(handleInterruptRunGroup())

	return &App{}, group.Run()
}

var Set = wire.NewSet(
	Init,
	server.Set,
	kitconfig.Set,
	// TODO: logrus logger?
	wire.InterfaceValue(new(requestlog.Logger), requestlog.NewNCSALogger(os.Stdout, nil)),
	wire.Value([]health.Checker{}),
	wire.InterfaceValue(new(trace.Exporter), trace.Exporter(nil)),
	wire.Value(trace.Sampler(nil)),
	ProvideHTTPHandler,
	ProvideApplicationContext,
	ProvideConfigLayer,
)

type ConfigName string
type ConfigEnvPrefix string

func ProvideConfigLayer(configName ConfigName, prefix ConfigEnvPrefix) (config.Layer, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	return config.NewMergeLayerWithSkipError(
		config.NewYAMLLayer(config.NewFallbackSource(
			config.NewFSSource(os.DirFS(cwd), string(configName)),
			config.NewFSSource(os.DirFS(home), string(configName)),
		)),
		config.NewEnvLayer(string(prefix)),
	), nil
}

func ProvideApplicationContext() (Context, func(), error) {
	ctx, cancel := context.WithCancel(context.Background())
	return ctx, cancel, nil
}

func ProvideHTTPHandler(handler kithttp.Handler) http.Handler { return handler }

func serverRunGroup(
	logger logrus.FieldLogger,
	srv *server.Server,
	cfg kitconfig.HTTP,
) (func() error, func(err error)) {
	init := func() error {
		addr := cfg.ListenAddr
		if addr == "" {
			addr = ":8080"
		}
		logger.Infof("start listen http server on %s", addr)
		return srv.ListenAndServe(addr)
	}
	cleanup := func(err error) {
		// TODO: get duration from config
		duration := time.Second * 15
		logger.Infof("start shutdown http server with timeout %s", duration)
		ctx, cancel := context.WithTimeout(context.Background(), duration)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}
	return init, cleanup
}

func handleInterruptRunGroup() (func() error, func(err error)) {
	sigs := make(chan os.Signal, 1)

	init := func() error {
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		// TODO: log signal
		sig, ok := <-sigs
		if !ok {
			return nil
		}
		return fmt.Errorf("signal received: %s", sig.String())
	}
	cleanup := func(err error) {
		signal.Stop(sigs)
		close(sigs)
	}

	return init, cleanup
}
