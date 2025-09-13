package modules

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func (m *ModuleController) deleteModuleByNameTool() mcp.Tool {
	return mcp.NewTool("delete_module",
		mcp.WithDescription("Delete Module from cluster"),
		mcp.WithString("module_name",
			mcp.Required(),
			mcp.Description("Name of the Module to delete"),
		),
	)
}

func (m *ModuleController) deleteModule(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	moduleName := request.Params.Arguments["module_name"].(string)

	err := m.k8sClient.DeleteModule(moduleName)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(fmt.Sprintf("Module %s deleted successfully", moduleName)), nil
}
