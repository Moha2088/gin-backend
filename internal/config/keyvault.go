package config

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"go.uber.org/zap"
)

type SecretClientConfig struct {
	logger *zap.Logger
}

func NewSecretClientConfig(logger *zap.Logger) *SecretClientConfig {
	return &SecretClientConfig{logger: logger}

}

func (s *SecretClientConfig) GetSecretClient() *azsecrets.Client {
	vaultUri := os.Getenv("VaultUri")
	credential, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		s.logger.Warn("Failed to create credential")
	}

	client, err := azsecrets.NewClient(vaultUri, credential, nil)

	if err != nil {
		s.logger.Warn("Failed to create client")
	}

	return client
}
