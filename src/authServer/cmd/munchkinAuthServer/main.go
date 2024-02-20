package main

import (
	"fmt"
	"log"

	"github.com/PenguinCats/MunchkinAuth/src/authServer/pkg"
	"gopkg.in/ini.v1"
)

type LocalConfig struct {
	KeyVaultName string

	OAuthAzure_TenantId            string
	OAuthAzure_ClientId            string
	OAuthAzure_ClientSecret        string
	OAuthAzure_ClientSecretFromAKV bool
	OAuthAzure_ClientSecretAKVName string
}

func ParseLocalConfig() LocalConfig {
	// use dependency injection to pass the globalConfig to the authServer
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
	}

	// Azure config
	AzureConfigSection := cfg.Section("Azure")
	KeyVaultName := AzureConfigSection.Key("keyvault_name").String()

	// OAuth2withAzure config
	OAuth2withAzureSection := cfg.Section("OAuth2withAzure")
	OAuthAzure_TenantId := OAuth2withAzureSection.Key("tenant_id").String()
	OAuthAzure_ClientId := OAuth2withAzureSection.Key("client_id").String()
	OAuthAzure_ClientSecret := OAuth2withAzureSection.Key("client_secret").String()
	OAuthAzure_ClientSecretFromAKV := OAuth2withAzureSection.Key("client_secret").MustBool(false)
	OAuthAzure_ClientSecretAKVName := OAuth2withAzureSection.Key("client_secret_akvname").String()

	return LocalConfig{
		KeyVaultName: KeyVaultName,

		OAuthAzure_TenantId:            OAuthAzure_TenantId,
		OAuthAzure_ClientId:            OAuthAzure_ClientId,
		OAuthAzure_ClientSecret:        OAuthAzure_ClientSecret,
		OAuthAzure_ClientSecretFromAKV: OAuthAzure_ClientSecretFromAKV,
		OAuthAzure_ClientSecretAKVName: OAuthAzure_ClientSecretAKVName,
	}
}

func main() {
	LocalConfig := ParseLocalConfig()
	fmt.Printf("LocalConfig: %v\n", LocalConfig)

	authServer := pkg.NewAuthServer()

	authServer.Start()
}
