package modules

import (
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/cluster/k8sclient"
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/template"
	"github.com/mark3labs/mcp-go/server"
)

type ModuleController struct {
	k8sClient    *k8sclient.KubernetesClient
	templateRepo template.ITemplateRepo
}

func NewController(k8sClient *k8sclient.KubernetesClient, templateRepo template.ITemplateRepo) *ModuleController {
	return &ModuleController{
		k8sClient:    k8sClient,
		templateRepo: templateRepo,
	}
}

func (m *ModuleController) RegisterModuleTools(mcp *server.MCPServer) {
	mcp.AddTool(m.getModuleByNameTool(), m.getModuleByName)
	mcp.AddTool(m.listModulesTool(), m.listModules)
	mcp.AddTool(m.createModuleTool(), m.createModule)
	mcp.AddTool(m.updateModuleTool(), m.updateModule)
}
