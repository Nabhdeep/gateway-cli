package cmd

import (
	"log/slog"
	"os"

	"github.com/nabhdeep/gateway-cli/constants"
	"github.com/nabhdeep/gateway-cli/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var loadConfig = &cobra.Command{
	Use:   "config [path]",
	Short: "Load Services Config from the specified path",
	Long:  `Loads Services Config path yaml file. Default is (config/services_config.yaml)`,
	Run:   run,
}

func run(c *cobra.Command, args []string) {
	var configPath string = "config/services_config.yaml"
	if len(args) > 0 {
		configPath = args[0]
	}
	isPathTrue := pathValidation(configPath)
	if !isPathTrue {
		slog.Error("Path does not exist", slog.String("Path", configPath))
		return
	}
	edit_path_in_gateway_config(configPath)
	slog.Info("Loading config from", slog.String("Path", configPath))
}

func pathValidation(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}

func Get_Gatewat_server_cofig_path() string {
	var gateway_config_path string = constants.Gateway_config_path
	if _path := os.Getenv("GATEWAY_CONFIG_PATH"); len(_path) > 0 && pathValidation(_path) {
		gateway_config_path = _path
	}
	return gateway_config_path
}

func Load_gateWay_config_file(gateway_config_path string) (*config.Config, error) {
	var gateway_config config.Config
	if _path := os.Getenv("GATEWAY_CONFIG_PATH"); len(_path) > 0 && pathValidation(_path) {
		gateway_config_path = _path
	}
	yaml_file, err := os.ReadFile(gateway_config_path)
	if err != nil {
		slog.Error("Unable to read file from gatway_config ", slog.String("Path", gateway_config_path))
		return &config.Config{}, err
	}
	err = yaml.Unmarshal(yaml_file, &gateway_config)
	if err != nil {
		slog.Error("Error parsing YAML file")
		return &config.Config{}, err
	}
	return &gateway_config, nil
}

func edit_path_in_gateway_config(p string) {
	var gateway_config_path string = Get_Gatewat_server_cofig_path()

	gateway_config, err := Load_gateWay_config_file(gateway_config_path)
	if err != nil {
		return
	}

	// modfiy srivces path
	gateway_config.Services_config_path = p

	updated_config, err := yaml.Marshal(gateway_config)

	if err != nil {
		slog.Error("Error saving YAML file")
		return
	}

	err = os.WriteFile(gateway_config_path, updated_config, 0644)
	if err != nil {
		slog.Error("Error saving YAML file")
		return
	}
	slog.Info("Services path updated successfully")
}

func init() {
	loadConfig.DisableFlagsInUseLine = true
	rootCmd.AddCommand(loadConfig)
}
