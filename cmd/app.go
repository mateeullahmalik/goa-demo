package cmd

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mateeullahmalik/goa-demo/api"
	"github.com/mateeullahmalik/goa-demo/api/services"
	"github.com/mateeullahmalik/goa-demo/configs"
	"github.com/mateeullahmalik/goa-demo/internal/cli"
	"github.com/mateeullahmalik/goa-demo/internal/configurer"
	"github.com/mateeullahmalik/goa-demo/internal/errors"
	"github.com/mateeullahmalik/goa-demo/internal/log"
	"github.com/mateeullahmalik/goa-demo/internal/log/hooks"
	"github.com/mateeullahmalik/goa-demo/internal/storage/memory"
	"github.com/mateeullahmalik/goa-demo/internal/sys"
	"github.com/mateeullahmalik/goa-demo/internal/version"
	loanSrvc "github.com/mateeullahmalik/goa-demo/services/loan"
	userSrvc "github.com/mateeullahmalik/goa-demo/services/user"
)

const (
	appName  = "goa-demo"
	appUsage = "goa-demo" // TODO: Write a clear description.
)

var (
	defaultPath = configurer.DefaultPath()

	defaultTempDir    = filepath.Join(os.TempDir(), appName)
	defaultConfigFile = filepath.Join(defaultPath, appName+".yml")
)

// NewApp inits a new command line interface.
func NewApp() *cli.App {
	var configFile string

	config := configs.New()

	app := cli.NewApp(appName)
	app.SetUsage(appUsage)
	app.SetVersion(version.Version())

	app.AddFlags(
		// Main
		cli.NewFlag("config-file", &configFile).SetUsage("Set `path` to the config file.").SetDefaultText(defaultConfigFile).SetAliases("c"),
		cli.NewFlag("temp-dir", &config.TempDir).SetUsage("Set `path` for storing temp data.").SetValue(defaultTempDir),
		cli.NewFlag("log-level", &config.LogLevel).SetUsage("Set the log `level`.").SetValue(config.LogLevel),
		cli.NewFlag("log-file", &config.LogFile).SetUsage("The log `file` to write to."),
		cli.NewFlag("quiet", &config.Quiet).SetUsage("Disallows log output to stdout.").SetAliases("q"),
		// API
		cli.NewFlag("swagger", &config.API.Swagger).SetUsage("Enable Swagger UI."),
	)

	app.SetActionFunc(func(ctx context.Context, args []string) error {
		ctx = log.ContextWithPrefix(ctx, "app")

		if configFile != "" {
			if err := configurer.ParseFile(configFile, config); err != nil {
				return err
			}
		}

		if config.Quiet {
			log.SetOutput(ioutil.Discard)
		} else {
			log.SetOutput(app.Writer)
		}

		if config.LogFile != "" {
			log.AddHook(hooks.NewFileHook(config.LogFile))
		}
		log.AddHook(hooks.NewDurationHook())

		if err := log.SetLevelName(config.LogLevel); err != nil {
			return errors.Errorf("--log-level %q, %w", config.LogLevel, err)
		}

		if err := os.MkdirAll(config.TempDir, os.ModePerm); err != nil {
			return errors.Errorf("could not create temp-dir %q, %w", config.TempDir, err)
		}

		return runApp(ctx, config)
	})

	return app
}

func runApp(ctx context.Context, config *configs.Config) error {
	log.WithContext(ctx).Info("Start")
	defer log.WithContext(ctx).Info("End")

	log.WithContext(ctx).Infof("Config: %s", config)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sys.RegisterInterruptHandler(func() {
		cancel()
		log.WithContext(ctx).Info("Interrupt signal received. Gracefully shutting down...")
	})

	// entities

	db := memory.NewKeyValue()

	// business logic services
	userSrvc := userSrvc.NewService(db)
	loanSrvc := loanSrvc.NewService(db)

	// api service
	server := api.NewServer(config.API, services.NewUserSrvc(userSrvc),
		services.NewLoanSrvc(loanSrvc),
		services.NewSwagger())

	return runServices(ctx, server)
}
