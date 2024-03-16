package config

// EBayCredential - container for oAuth EBayCredentials
type EBayCredential struct {
	Environment string `mapstructure:"environment" json:"environment" yaml:"environment"`
	AppID       string `mapstructure:"appid" json:"appid" yaml:"appid"`
	CertID      string `mapstructure:"certid" json:"certid" yaml:"certid"`
	DevID       string `mapstructure:"devid" json:"devid" yaml:"devid"`
	RedirectURI string `mapstructure:"redirecturi" json:"redirecturi" yaml:"redirecturi"`
	WebEndpoint string `mapstructure:"web-endpoint" json:"web-endpoint" yaml:"web-endpoint"`
	ApiEndpoint string `mapstructure:"api-endpoint" json:"api-endpoint" yaml:"api-endpoint"`
}
