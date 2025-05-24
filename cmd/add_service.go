package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/nabhdeep/gateway-cli/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var addService = &cobra.Command{
	Use:   "service",
	Short: "Add/Remove/List services to the services config yaml file",
	Long: `Add, remove, or list services in the services config yaml file. 
If the services yaml file does not exist, it creates one at the default path (config/services_config.yaml).
- Use -a <service name> to add a service.
- Use -r <service name> to remove a service.
- Use -l to list all services.`,
	Run: func(cmd *cobra.Command, args []string) {
		addVal, _ := cmd.Flags().GetString("a")
		removeVal, _ := cmd.Flags().GetString("r")
		listVal, _ := cmd.Flags().GetBool("l")

		if addVal != "" {
			fmt.Println("Adding service:", addVal)
			// logic to add service
		}
		if removeVal != "" {
			fmt.Println("Removing service:", removeVal)
			// logic to remove service
		}
		if listVal {
			listServices()
		}
		if addVal == "" && removeVal == "" && !listVal {
			fmt.Println("No operation specified. Use -a to add, -r to remove, or -l to list services.")
		}
	},
}

func get_services_config_path() (string, error) {
	gateway_config_path := Get_Gatewat_server_cofig_path()
	gateway_config, err := Load_gateWay_config_file(gateway_config_path)
	if err != nil {
		return "", err
	}
	return gateway_config.Services_config_path, nil
}

func Load_services_from_config() (*config.ServicesConfig, error) {
	var servicesConfig config.ServicesConfig
	p, err := get_services_config_path()
	if err != nil {
		slog.Error("Error gettign services path")
		return &servicesConfig, err
	}
	yaml_file, err := os.ReadFile(p)
	if err != nil {
		return &servicesConfig, err
	}
	err = yaml.Unmarshal(yaml_file, &servicesConfig)
	if err != nil {
		slog.Debug(err.Error())
		slog.Error("Error loding the sevices yml file")
		return &servicesConfig, err
	}
	return &servicesConfig, nil
}

func listServices() {
	services, err := Load_services_from_config()
	if err != nil {
		slog.Debug(err.Error())
		return
	}
	for _, i := range services.Services {
		fmt.Println("========================================")
		fmt.Println("🚀 Service Name :", i.Name)
		fmt.Println("----------------------------------------")
		fmt.Printf("🔑 API Key      : %s\n", i.Api_Key)
		fmt.Printf("🌐 URL          : %s\n", i.Baseurl)
		for _, r := range i.Routes {
			fmt.Printf("📜 Method, Endpoint   : %v %v\n", r.Method, r.Endpoint)
		}
		fmt.Printf("🚦 Rate Limit   : %d requests/sec\n", i.Rate_Limits)
		fmt.Printf("📝 Allow List   : %v\n", i.Allow_List)
		fmt.Printf("✅ Enabled      : %t\n", i.Enabled)
		fmt.Println("========================================")
	}

}

func init() {
	addService.Flags().StringP("a", "a", "", "Add a new service")
	addService.Flags().StringP("r", "r", "", "Remove a service")
	addService.Flags().BoolP("l", "l", false, "List all services")
	rootCmd.AddCommand(addService)
}
