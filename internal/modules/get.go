package modules

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

func (m *ModuleController) getModuleByNameTool() mcp.Tool {
	return mcp.NewTool("get_module",
		mcp.WithDescription("Fetch Module by Name"),
		mcp.WithString("module_name",
			mcp.Required(),
			mcp.Description("Name of the Module"),
		),
	)
}

func (m *ModuleController) getModuleByName(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	moduleName := request.Params.Arguments["module_name"].(string)

	module, err := m.k8sClient.GetModule(moduleName)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(module)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(b)), nil
}
