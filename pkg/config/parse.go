package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/depender/email-sequence-service/constants"
	"github.com/depender/email-sequence-service/pkg/config/model"
	log "github.com/depender/email-sequence-service/pkg/logger"
	"time"

	"github.com/spf13/cast"
)

func handleParsingError(key string, err error) {
	log.Errorf("error while parsing configuration for key %s", key)

	pr := getApplication()
	pr.monitoring.HandleError(err)
}

func getConfigFromContext(ctx context.Context, key string) *model.Config {
	contextConfig := ctx.Value(constants.Config)
	if contextConfig == nil {
		handleParsingError(key, errors.New("config not found in context"))

		// if no configuration is embedded in context, it will fetch the latest configuration.
		// a) Refresh context by fetching it again: Will get the latest configuration
		// b) Pass old context passed by config package: Will give the old configuration
		// c) Pass normal context: Will report error and fallback to latest configuration

		app := getApplication()
		if app == nil {
			handleParsingError(key, errors.New("config not initialised"))
			return nil
		}
		return app.watcher.GetConfig()
	}

	conf, ok := contextConfig.(*model.Config)
	if !ok {
		handleParsingError(key, errors.New("type received for config invalid"))
		return nil
	}

	return conf
}

func Get(ctx context.Context, key string) interface{} {
	conf := getConfigFromContext(ctx, key)

	val, ok := conf.GetValueForKey(key)
	if !ok {
		handleParsingError(key, fmt.Errorf("key %s not found", key))
		return nil
	}

	return val
}

func GetBool(ctx context.Context, key string) bool {
	val, err := cast.ToBoolE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetString(ctx context.Context, key string) string {
	val, err := cast.ToStringE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetInt(ctx context.Context, key string) int {
	val, err := cast.ToIntE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetInt8(ctx context.Context, key string) int8 {
	val, err := cast.ToInt8E(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetInt16(ctx context.Context, key string) int16 {
	val, err := cast.ToInt16E(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetInt32(ctx context.Context, key string) int32 {
	val, err := cast.ToInt32E(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetInt64(ctx context.Context, key string) int64 {
	val, err := cast.ToInt64E(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetUint(ctx context.Context, key string) uint {
	val, err := cast.ToUintE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetUint8(ctx context.Context, key string) uint8 {
	val, err := cast.ToUint8E(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetUint16(ctx context.Context, key string) uint16 {
	val, err := cast.ToUint16E(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetUint32(ctx context.Context, key string) uint32 {
	val, err := cast.ToUint32E(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetUint64(ctx context.Context, key string) uint64 {
	val, err := cast.ToUint64E(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetFloat32(ctx context.Context, key string) float32 {
	val, err := cast.ToFloat32E(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetFloat64(ctx context.Context, key string) float64 {
	val, err := cast.ToFloat64E(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetTime(ctx context.Context, key string) time.Time {
	val, err := cast.ToTimeE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetDuration(ctx context.Context, key string) time.Duration {
	val, err := cast.ToDurationE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetSlice(ctx context.Context, key string) []interface{} {
	val, err := cast.ToSliceE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetBoolSlice(ctx context.Context, key string) []bool {
	val, err := cast.ToBoolSliceE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetIntSlice(ctx context.Context, key string) []int {
	val, err := cast.ToIntSliceE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetStringSlice(ctx context.Context, key string) []string {
	val, err := cast.ToStringSliceE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetDurationSlice(ctx context.Context, key string) []time.Duration {
	val, err := cast.ToDurationSliceE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetStringMapString(ctx context.Context, key string) map[string]string {
	val, err := cast.ToStringMapStringE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetStringMapStringSlice(ctx context.Context, key string) map[string][]string {
	val, err := cast.ToStringMapStringSliceE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetStringMapBool(ctx context.Context, key string) map[string]bool {
	val, err := cast.ToStringMapBoolE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetStringMapInt(ctx context.Context, key string) map[string]int {
	val, err := cast.ToStringMapIntE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetStringMapInt64(ctx context.Context, key string) map[string]int64 {
	val, err := cast.ToStringMapInt64E(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}

func GetStringMap(ctx context.Context, key string) map[string]interface{} {
	val, err := cast.ToStringMapE(Get(ctx, key))
	if err != nil {
		handleParsingError(key, err)
	}
	return val
}
