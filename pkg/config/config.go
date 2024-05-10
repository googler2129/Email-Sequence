package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/depender/email-sequence-service/constants"
	"github.com/depender/email-sequence-service/pkg/config/monitor"
	"github.com/depender/email-sequence-service/pkg/config/provider"
	"github.com/depender/email-sequence-service/pkg/config/watcher"
	log "github.com/depender/email-sequence-service/pkg/logger"
	"os"
	"strings"
	"time"
)

type application struct {
	watcher    *watcher.Watcher
	monitoring monitor.Monitoring
}

var globalApplication *application

func Init(pollDuration time.Duration) (err error) {
	if pollDuration < time.Second*10 {
		err = errors.New("invalid poll duration in app config")
		return
	}

	if getApplication() != nil {
		return
	}

	ctx := context.TODO()

	configSource, ok := os.LookupEnv("CONFIG_SOURCE")
	if !ok {
		err = errors.New("CONFIG_SOURCE environment variable not found")
		log.Panic("failed to initialise config: ", err)
	}

	if configSource != constants.LocalSource && !strings.HasPrefix(configSource, "appconfig:") {
		err = fmt.Errorf("CONFIG_SOURCE %s not valid", configSource)
		log.Panic("failed to initialise config: ", err)
		return
	}

	var pr provider.Provider
	var monitoringTool monitor.Monitoring

	if configSource == constants.LocalSource {
		log.Info("Reading local configuration files")
		pr, err = provider.NewLocalProvider(ctx, constants.LocalFreeFormPath)
		if err != nil {
			log.Panic("failed to initialise config: ", err)
			return
		}

		monitoringTool, err = monitor.NewPanicMonitoring()
		if err != nil {
			log.Panic("failed to initialise config: ", err)
			return
		}
	} else {
		appName := strings.TrimPrefix(configSource, "appconfig:")
		pr, err = provider.NewRemoteProvider(ctx, constants.RemoteFreeformProfile, appName)
		if err != nil {
			log.Panic("failed to initialise config: ", err)
			return
		}

		monitoringTool, err = monitor.NewCloudwatchMonitoring(ctx, appName)
		if err != nil {
			log.Panic("failed to initialise config: ", err)
			return
		}
	}

	wr, err := watcher.NewWatcher(ctx, pr, monitoringTool, pollDuration)
	if err != nil {
		log.Panic("failed to initialise config: ", err)
		return
	}

	app := &application{
		watcher:    wr,
		monitoring: monitoringTool,
	}

	setApplication(app)
	return
}

func getApplication() *application {
	return globalApplication
}

func setApplication(app *application) {
	globalApplication = app
}
