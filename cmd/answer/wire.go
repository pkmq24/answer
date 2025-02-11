//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"github.com/segmentfault/answer/internal/base/conf"
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/base/middleware"
	"github.com/segmentfault/answer/internal/base/server"
	"github.com/segmentfault/answer/internal/base/translator"
	"github.com/segmentfault/answer/internal/controller"
	"github.com/segmentfault/answer/internal/controller_backyard"
	"github.com/segmentfault/answer/internal/repo"
	"github.com/segmentfault/answer/internal/router"
	"github.com/segmentfault/answer/internal/service"
	"github.com/segmentfault/answer/internal/service/service_config"
	"github.com/segmentfault/pacman"
	"github.com/segmentfault/pacman/log"
)

// initApplication init application.
func initApplication(
	debug bool,
	serverConf *conf.Server,
	dbConf *data.Database,
	cacheConf *data.CacheConf,
	i18nConf *translator.I18n,
	swaggerConf *router.SwaggerConfig,
	serviceConf *service_config.ServiceConfig,
	logConf log.Logger) (*pacman.Application, func(), error) {
	panic(wire.Build(
		server.ProviderSetServer,
		router.ProviderSetRouter,
		controller.ProviderSetController,
		controller_backyard.ProviderSetController,
		service.ProviderSetService,
		repo.ProviderSetRepo,
		translator.ProviderSet,
		middleware.ProviderSetMiddleware,
		newApplication,
	))
}
