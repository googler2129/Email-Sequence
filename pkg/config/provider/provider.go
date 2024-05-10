package provider

import (
	"context"
	"fmt"
	"github.com/depender/email-sequence-service/pkg/config/model"
	"strings"

	"gopkg.in/yaml.v3"
)

type Provider interface {
	GetConfig(ctx context.Context) (*model.Config, error)
}

type kv struct {
	key   string
	value map[string]interface{}
}

func yamlToConfigMap(yamlStr string) (map[string]interface{}, error) {
	var yamlMap map[string]interface{}
	var configMap = make(map[string]interface{})
	err := yaml.Unmarshal([]byte(yamlStr), &yamlMap)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling config: %w", err)
	}

	queue := []kv{
		{
			key:   "",
			value: yamlMap,
		},
	}

	for len(queue) > 0 {
		kvPair := queue[0]
		queue = queue[1:]

		for key, v := range kvPair.value {
			newKey := kvPair.key + "." + key
			if kvPair.key == "" {
				newKey = key
			}
			newKey = strings.ToLower(newKey)

			switch v := v.(type) {
			case map[string]interface{}:
				queue = append(queue, kv{
					key:   newKey,
					value: v,
				})
			case map[interface{}]interface{}:
				for interfaceKey := range v {
					switch interfaceKey.(type) {
					case string:
						continue
					default:
						return nil, fmt.Errorf("only string keys allowed. Got %T for key %v", interfaceKey, interfaceKey)
					}
				}
			default:
				if v == nil {
					return nil, fmt.Errorf("received nil value for %s", newKey)
				}
				if _, ok := configMap[newKey]; ok {
					return nil, fmt.Errorf("duplicate key %s. Please note that keys are case insensitive", newKey)
				}
				configMap[newKey] = v
			}
		}
	}
	return configMap, nil
}
