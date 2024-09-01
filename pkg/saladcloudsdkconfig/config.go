package saladcloudsdkconfig

type Config struct {
	BaseUrl    *string
	ApiKey     *string
	HookParams map[string]string
}

func NewConfig() Config {
	baseUrl := DEFAULT_ENVIRONMENT
	newConfig := Config{
		BaseUrl:    &baseUrl,
		HookParams: make(map[string]string),
	}

	return newConfig
}

func (c *Config) SetBaseUrl(baseUrl string) {
	c.BaseUrl = &baseUrl
}

func (c *Config) GetBaseUrl() string {
	return *c.BaseUrl
}

func (c *Config) SetApiKey(apiKey string) {
	c.ApiKey = &apiKey
}

func (c *Config) GetApiKey() string {
	return *c.ApiKey
}
