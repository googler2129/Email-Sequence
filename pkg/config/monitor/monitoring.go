package monitor

import (
	"context"
	"fmt"
	"github.com/depender/email-sequence-service/constants"
	log "github.com/depender/email-sequence-service/pkg/logger"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
)

var cloudwatchErrorCount int64

type Monitoring interface {
	HandleError(err error)
}

type panicMonitoring struct {
}

func (pm *panicMonitoring) HandleError(err error) {
	panic(err)
}

func NewPanicMonitoring() (Monitoring, error) {
	return &panicMonitoring{}, nil
}

type cloudwatchMonitoring struct {
}

func (cm *cloudwatchMonitoring) HandleError(err error) {
	atomic.AddInt64(&cloudwatchErrorCount, 1)
}

func NewCloudwatchMonitoring(ctx context.Context, providerName string) (Monitoring, error) {
	cfg, err := awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while loading default aws config for cloudwatch client: %w", err)
	}

	client := cloudwatch.NewFromConfig(cfg)
	metricName := constants.CloudwatchErrorMetric
	dimensionName := constants.CloudwatchErrorDimension
	namespace := constants.CloudwatchNamespace
	dimension := types.Dimension{
		Name:  &dimensionName,
		Value: &providerName,
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("stopped pushing metrics [Panic]: %+w", r)
				log.Error(err.Error())
			}
		}()

		for {
			time.Sleep(constants.CloudwatchPutMetricInterval)
			if atomic.LoadInt64(&cloudwatchErrorCount) == 0 {
				continue
			}
			value := atomic.LoadInt64(&cloudwatchErrorCount)
			atomic.AddInt64(&cloudwatchErrorCount, -1*value)

			valueFloat := float64(value)

			timeStamp := time.Now()
			metricData := types.MetricDatum{
				MetricName: &metricName,
				Timestamp:  &timeStamp,
				Value:      &valueFloat,
				Unit:       types.StandardUnitCount,
				Dimensions: []types.Dimension{dimension},
			}

			_, err = client.PutMetricData(ctx, &cloudwatch.PutMetricDataInput{
				MetricData: []types.MetricDatum{metricData},
				Namespace:  &namespace,
			})
			if err != nil {
				log.Error("error while putting cloudwatch metric")
				atomic.AddInt64(&cloudwatchErrorCount, value)
			}
		}
	}()
	return &cloudwatchMonitoring{}, nil
}
