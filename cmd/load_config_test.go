package cmd_test

import (
	"os"
	"testing"

	"github.com/nabhdeep/gateway-cli/cmd"
	"github.com/nabhdeep/gateway-cli/pkg/config"
	"github.com/nabhdeep/gateway-cli/pkg/constants"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func assertHandler[T comparable](t *testing.T, got T, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPathValidation(t *testing.T) {
	t.Run("Check with invalid path in PathValidation", func(t *testing.T) {
		fake_path := ""
		got := cmd.PathValidation(fake_path)
		want := false
		assertHandler(t, got, want)
	})
	t.Run("Check valid path in PathValidation", func(t *testing.T) {
		fake_path := "testdata/test_config.yaml"
		got := cmd.PathValidation(fake_path)
		want := false
		assertHandler(t, got, want)
	})

}
func TestLoadConfig(t *testing.T) {
	t.Run("Should Load Config from default path", func(t *testing.T) {
		get := cmd.Get_Gatewat_server_cofig_path()
		want := constants.Gateway_config_path
		assertHandler(t, get, want)
	})
	t.Run("Should Load Config from default path as ENV has invalid path ", func(t *testing.T) {
		some_fake_env := ""
		os.Setenv("GATEWAY_CONFIG_PATH", some_fake_env)
		defer os.Unsetenv("GATEWAY_CONFIG_PATH")
		get := cmd.Get_Gatewat_server_cofig_path()
		want := constants.Gateway_config_path
		assertHandler(t, get, want)
	})

	t.Run("Should Load Config from default path as ENV has invalid path ", func(t *testing.T) {
		envPath := "/tmp/test_gateway_config.yaml"

		// creating a dummy file
		_, err := os.Create(envPath)
		if err != nil {
			t.Fatalf("Error creating temporary file: %v", err)
		}
		defer os.Remove(envPath)

		os.Setenv("GATEWAY_CONFIG_PATH", envPath)
		defer os.Unsetenv("GATEWAY_CONFIG_PATH")
		get := cmd.Get_Gatewat_server_cofig_path()
		want := envPath
		assertHandler(t, get, want)
	})

	t.Run("Should Load config from path", func(t *testing.T) {
		configNilGateway := config.Config{}
		envPath := "fakepath"
		config_gatway, err := cmd.Load_gateWay_config_file(envPath)
		assert.Error(t, err, "")
		assert.Equal(t, &configNilGateway, config_gatway, "Expected config_gatway to be nil")
	})

	t.Run("Should Load config from invalid path", func(t *testing.T) {
		configNilGateway := config.Config{}
		envPath := "fakepath"
		config_gatway, err := cmd.Load_gateWay_config_file(envPath)
		assert.Error(t, err, "")
		assert.Equal(t, &configNilGateway, config_gatway, "Expected config_gatway to be nil")
	})

	t.Run("Should Load config from path", func(t *testing.T) {
		testConfigHttpServer := config.HttpServer{
			Address: "localhost:4001",
		}
		configNilGateway := config.Config{
			Env:                  "dev",
			HttpServer:           testConfigHttpServer,
			Services_config_path: "testdata/services_config.yaml",
		}
		config_gatway, err := cmd.Load_gateWay_config_file("../testdata/config.yaml")
		assert.Nil(t, err, "Error shold return nil from valid path for valid config path ")
		assert.Equal(t, &configNilGateway, config_gatway, "Expected config_gatway to be nil")
	})
	t.Run("Should Load config from path", func(t *testing.T) {
		testConfigHttpServer := config.HttpServer{
			Address: "localhost:4001",
		}
		configNilGateway := config.Config{
			Env:                  "dev",
			HttpServer:           testConfigHttpServer,
			Services_config_path: "testdata/services_config.yaml",
		}
		config_gatway, err := cmd.Load_gateWay_config_file("../testdata/config.yaml")
		assert.Nil(t, err, "Error shold return nil from valid path for valid config path ")
		assert.Equal(t, &configNilGateway, config_gatway, "Expected config_gatway to be nil")
	})

	t.Run("Should Return error in invalid config path", func(t *testing.T) {
		envPath := "fake path "
		err := cmd.Edit_path_in_gateway_config(envPath, envPath)
		assert.Error(t, err, "Error should be returned on invalid path ")
		assert.Contains(t, err.Error(), "open")
		assert.Contains(t, err.Error(), envPath)

	})

	t.Run("Should Return nil and update the config path", func(t *testing.T) {
		// creating a dummy file
		conf := config.Config{
			Env:                  "test",
			HttpServer:           config.HttpServer{Address: "localhost:4002"},
			Services_config_path: "testdata/services_config.yaml",
		}
		envPath := "../testdata/test_config.yaml"
		yml, _ := yaml.Marshal(conf)
		os.Create(envPath)
		os.WriteFile(envPath, yml, 0644)
		defer os.Remove(envPath)

		cmd.Edit_path_in_gateway_config("fakepath", envPath)
		gateway_config, err := cmd.Load_gateWay_config_file(envPath)
		assert.Nil(t, err, "nil should be returned ")
		assert.NotEqual(t, gateway_config.Services_config_path, conf.Services_config_path)
		assert.Equal(t, gateway_config.Services_config_path, "fakepath")
	})

}
