package provider

import (
	"context"
	"fmt"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appconfigdata"
	"github.com/depender/email-sequence-service/constants"
	"github.com/depender/email-sequence-service/pkg/config/model"
	log "github.com/depender/email-sequence-service/pkg/logger"
)

type remoteProvider struct {
	freeform *remoteConfigProfile
	client   *appconfigdata.Client
}

type remoteConfigProfile struct {
	token *string
	name  string
	data  string
}

func NewRemoteProvider(ctx context.Context, freeformProfile, appName string) (Provider, error) {
	cfg, err := awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while loading default aws config for appconfig client: %w", err)
	}
	client := appconfigdata.NewFromConfig(cfg)

	rp := &remoteProvider{
		client:   client,
		freeform: &remoteConfigProfile{name: freeformProfile},
	}

	err = rp.startConfigurationSession(ctx, rp.freeform, appName)
	if err != nil {
		return nil, err
	}

	return rp, nil
}

func (rp *remoteProvider) GetConfig(ctx context.Context) (*model.Config, error) {
	yamlStr, err := rp.fetchDataFromRemote(ctx, rp.freeform)
	if err != nil {
		return nil, err
	}
	if yamlStr != "" {
		rp.freeform.data = yamlStr
		log.Info("Updated freeform configuration")
	}

	return rp.createConfig()
}

func (rp *remoteProvider) createConfig() (*model.Config, error) {
	configMap, err := yamlToConfigMap(rp.freeform.data)
	if err != nil {
		return nil, err
	}

	return model.NewConfig(configMap), nil
}

func (rp *remoteProvider) startConfigurationSession(ctx context.Context,
	rcp *remoteConfigProfile, appIdentifier string) error {
	env := constants.DeployedEnv
	out, err := rp.client.StartConfigurationSession(ctx, &appconfigdata.StartConfigurationSessionInput{
		ApplicationIdentifier:          &appIdentifier,
		ConfigurationProfileIdentifier: &rcp.name,
		EnvironmentIdentifier:          &env,
	})
	if err != nil {
		return fmt.Errorf("error while starting configuration session: %w", err)
	}

	rcp.token = out.InitialConfigurationToken
	return nil
}

func (rp *remoteProvider) fetchDataFromRemote(ctx context.Context,
	rcp *remoteConfigProfile) (string, error) {
	out, err := rp.client.GetLatestConfiguration(ctx, &appconfigdata.GetLatestConfigurationInput{
		ConfigurationToken: rcp.token,
	})
	if err != nil {
		return "", fmt.Errorf("error while getting configuration from remote: %w", err)
	}

	rcp.token = out.NextPollConfigurationToken
	return string(out.Configuration), nil
}
