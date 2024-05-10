package provider

import (
	"context"
	"fmt"
	"github.com/depender/email-sequence-service/pkg/config/model"
	"os"
)

type localProvider struct {
	freeformPath string
}

func NewLocalProvider(ctx context.Context, freeformPath string) (Provider, error) {
	lp := &localProvider{
		freeformPath: freeformPath,
	}

	return lp, nil
}

func (lp *localProvider) GetConfig(ctx context.Context) (*model.Config, error) {
	configMap, err := readLocalFile(lp.freeformPath)
	if err != nil {
		return nil, err
	}

	return model.NewConfig(configMap), nil
}

func readLocalFile(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading local file: %w", err)
	}

	configMap, err := yamlToConfigMap(string(data))
	if err != nil {
		return nil, err
	}

	return configMap, nil
}
