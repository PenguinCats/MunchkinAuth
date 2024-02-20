package common

type AzureAuthConfig struct {
	ClientSecretFromAKV bool
	TenantID            string
	ClientID            string
	ClientSecret        string
}

type GlobalConfig struct {
	AzureAuthConfig AzureAuthConfig
}
