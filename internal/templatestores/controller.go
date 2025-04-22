package templatestores

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/cluster/k8sclient"
)

type TemplateStoreController struct {
	k8sClient *k8sclient.KubernetesClient
}

func NewController(k8sClient *k8sclient.KubernetesClient) *TemplateStoreController {
	return &TemplateStoreController{
		k8sClient: k8sClient,
	}
}

func (t *TemplateStoreController) RegisterTemplateStoreTools(mcp *server.MCPServer) {
	mcp.AddTool(t.listTemplateStoresTool(), t.listTemplateStores)
	mcp.AddTool(t.getTemplateStoreByNameTool(), t.getTemplateStoreByName)
	//mcp.AddTool(t.createTemplateStoreTool(), t.createTemplateStore)
}
