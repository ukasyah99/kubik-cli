package config

type IDCloudHostConfig struct {
	APIKey string
	Host   string
}

func IDCloudHost() *IDCloudHostConfig {
	return &IDCloudHostConfig{}
}
