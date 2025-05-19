package modules

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

func (m *ModuleController) listModuleResourcesTool() mcp.Tool {
	return mcp.NewTool("list_module_resources",
		mcp.WithDescription("Lists all Kubernetes resources owned by the given Module"),
		mcp.WithString("module_name",
			mcp.Required(),
			mcp.Description("Name of the Module"),
		),
	)
}

func (m *ModuleController) listModuleResources(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	moduleName := request.Params.Arguments["module_name"].(string)

	resources, err := m.k8sClient.GetResourcesForModule(moduleName)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(resources)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(b)), nil
}
