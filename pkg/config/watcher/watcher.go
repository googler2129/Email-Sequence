package watcher

import (
	"context"
	"fmt"
	"github.com/depender/email-sequence-service/pkg/config/model"
	"github.com/depender/email-sequence-service/pkg/config/monitor"
	"github.com/depender/email-sequence-service/pkg/config/provider"
	log "github.com/depender/email-sequence-service/pkg/logger"
	"sync"
	"time"
)

type Watcher struct {
	conf *model.Config
	mu   sync.Mutex
}

func (cw *Watcher) GetConfig() *model.Config {
	cw.mu.Lock()
	defer cw.mu.Unlock()
	return cw.conf
}

func NewWatcher(ctx context.Context, pr provider.Provider, monitorTool monitor.Monitoring, pollDuration time.Duration) (*Watcher, error) {
	cw := &Watcher{}
	err := cw.watchForConfigUpdates(ctx, pr, monitorTool, pollDuration)
	if err != nil {
		return nil, err
	}
	return cw, nil
}

func (cw *Watcher) watchForConfigUpdates(ctx context.Context, pr provider.Provider, monitorTool monitor.Monitoring, pollDuration time.Duration) error {
	conf, err := pr.GetConfig(ctx)
	if err != nil {
		return fmt.Errorf("error while getting internal.Config: %w", err)
	}

	cw.setConfigData(conf)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("stopped watching internal.Config [Panic]: %+w", r)
				log.Error(err.Error())
			}
		}()

		for {
			time.Sleep(pollDuration)
			conf, err = pr.GetConfig(ctx)
			if err != nil {
				log.Error("error while getting internal.Configuration")
				monitorTool.HandleError(err)
				continue
			}

			cw.setConfigData(conf)
		}
	}()
	return nil
}

func (cw *Watcher) setConfigData(cfg *model.Config) {
	cw.mu.Lock()
	defer cw.mu.Unlock()
	cw.conf = cfg
}
