package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mark3labs/mcp-go/server"

	_ "github.com/joho/godotenv/autoload"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/auth"
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/cluster/k8sclient"
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/template"
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/template/cache"

	"github.com/cyclops-ui/mcp-cyclops/internal/modules"
	"github.com/cyclops-ui/mcp-cyclops/internal/templates"
	"github.com/cyclops-ui/mcp-cyclops/internal/templatestores"
)

type Config struct {
	version string

	transport string
	address   string

	kubeconfigPath string
	kubeContext    string

	moduleNamespace       string
	helmReleaseNamespace  string
	moduleTargetNamespace string
}

func main() {
	config := loadConfig()

	s := server.NewMCPServer(
		"Cyclops",
		"0.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	k8sClient, err := k8sclient.NewWithConfig(
		k8sclient.ClientConfig{
			KubeconfigPath:        config.kubeconfigPath,
			Context:               config.kubeContext,
			ModuleNamespace:       config.moduleNamespace,
			HelmReleaseNamespace:  config.helmReleaseNamespace,
			ModuleTargetNamespace: config.moduleTargetNamespace,
		},
		zap.New(),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Failed to create Kubernetes client: %v", err))
		os.Exit(1)
	}

	templatesRepo := template.NewRepo(
		auth.NewTemplatesResolver(k8sClient),
		cache.NewInMemoryTemplatesCache(),
	)

	moduleController := modules.NewController(k8sClient, templatesRepo)
	moduleController.RegisterModuleTools(s)

	templateStoresController := templatestores.NewController(k8sClient)
	templateStoresController.RegisterTemplateStoreTools(s)

	templatesController := templates.NewController(templatesRepo)
	templatesController.RegisterTemplateStoreTools(s)

	switch config.transport {
	case "stdio":
		stdioServer := server.NewStdioServer(s)
		if err := stdioServer.Listen(context.Background(), os.Stdin, os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, "failed to start stdio server")
			panic(err)
		}
	case "sse":
		sseServer := server.NewSSEServer(s)
		if err := sseServer.Start(config.address); err != nil {
			fmt.Fprintln(os.Stderr, "failed to start sse server")
			panic(err)
		}
	default:
		fmt.Fprintln(os.Stderr, "invalid transport type - should be stdio or sse")
		panic("invalid transport type - should be stdio or sse")
	}
}

func loadConfig() *Config {
	config := &Config{
		version:               getEnvOrDefault("CYCLOPS_MCP_VERSION", "0.0.0"),
		transport:             getEnvOrDefault("CYCLOPS_MCP_TRANSPORT", "stdio"),
		address:               getEnvOrDefault("CYCLOPS_MCP_SSE_ADDRESS", "127.0.0.1:8000"),
		moduleNamespace:       getEnvOrDefault("CYCLOPS_MODULE_NAMESPACE", "cyclops"),
		helmReleaseNamespace:  os.Getenv("CYCLOPS_HELM_RELEASE_NAMESPACE"),
		moduleTargetNamespace: os.Getenv("CYCLOPS_MODULE_TARGET_NAMESPACE"),
		kubeContext:           os.Getenv("CYCLOPS_KUBE_CONTEXT"),
	}

	config.kubeconfigPath = os.Getenv("KUBECONFIG")
	if config.kubeconfigPath == "" {
		if home, err := os.UserHomeDir(); err == nil {
			defaultKubeconfig := filepath.Join(home, ".kube", "config")
			if _, err := os.Stat(defaultKubeconfig); err == nil {
				config.kubeconfigPath = defaultKubeconfig
			}
		}
	}

	return config
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
