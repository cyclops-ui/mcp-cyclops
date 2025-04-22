package main

import (
	"log"
	"os"

	"github.com/mark3labs/mcp-go/server"

	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/auth"
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/cluster/k8sclient"
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/template"
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/template/cache"

	"github.com/cyclops-ui/mcp-cyclops/internal/modules"
	"github.com/cyclops-ui/mcp-cyclops/internal/templates"
	"github.com/cyclops-ui/mcp-cyclops/internal/templatestores"
)

func main() {
	s := server.NewMCPServer(
		"Cyclops",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	k8sClient, err := k8sclient.New(
		"cyclops",
		"",
		"",
		zap.New(),
	)
	if err != nil {
		panic(err)
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

	if err := server.ServeStdio(s); err != nil {
		log.Printf("Server error: %v\n", err)
		os.Exit(1)
	}
}
