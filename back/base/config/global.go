package config

type Config struct {
	Server struct {
		Port int    `mapstructure:"port"`
		Name string `mapstructure:"name"`
	}
	JWT struct {
		SecretKey  string `mapstructure:"secret_key"`
		Expiration string `mapstructure:"expiration"`
		Issuer     string `mapstructure:"issuer"`
	}
	Github struct {
		AuthURL             string `mapstructure:"auth_url"`
		TokenURL            string `mapstructure:"token_url"`
		UserInfoURL         string `mapstructure:"user_info_url"`
		ClientID            string `mapstructure:"client_id"`
		ClientSecret        string `mapstructure:"client_secret"`
		RedirectURL         string `mapstructure:"redirect_url"`
		AuthorizationHeader string `mapstructure:"authorization_header"`
		UserAgentHeader     string `mapstructure:"user_agent_header"`
		UserAgentValue      string `mapstructure:"user_agent_value"`
	}
	Consul struct {
		Address     string `mapstructure:"address"`
		ServiceName string `mapstructure:"service_name"`
		ServiceID   string `mapstructure:"service_id"`
		ServicePort int    `mapstructure:"service_port"`
	}
}

var GlobalConfig Config
