package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/nabhdeep/gateway-cli/pkg/constants"
)

type Service struct {
	Name             string   `yaml:"name" env-required:"true"`
	Service_Endpoint string   `yaml:"service_endpoint" env-required:"true"`
	Baseurl          string   `yaml:"baseUrl" env-required:"true"`
	Routes           []Route  `yaml:"routes" env-required:"true"`
	Rate_Limits      int      `yaml:"rate_limits" env-required:"true"`
	Api_Key          string   `yaml:"api_key" env-required:"true"`
	Allow_List       []string `yaml:"allowlist" env-required:"true"`
	Enabled          bool     `yaml:"enabled" env-required:"true"`
}

type Route struct {
	Method   string `yaml:"method"`
	Endpoint string `yaml:"endpoint"`
}

type HttpServer struct {
	Address string `yaml:"address" env-required:"true"`
}
type ServicesConfig struct {
	Services []Service `yaml:"services"`
}

type Config struct {
	Env                  string     `yaml:"env" env-required:"true"`
	HttpServer           HttpServer `yaml:"http_server" env-required:"true"`
	Services_config_path string     `yaml:"services_config_path"`
}

func MustLoad() (*Config, *ServicesConfig) {
	configPath := constants.Gateway_config_path
	_path := os.Getenv("GATEWAY_CONFIG_PATH")
	if _path != "" {
		configPath = _path
	}
	var cfg Config
	var service_config ServicesConfig
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	} else {
		err := cleanenv.ReadConfig(cfg.Services_config_path, &service_config)
		if err != nil {
			log.Fatalf("can not read config file: %s", err.Error())
		}
	}
	return &cfg, &service_config
}
