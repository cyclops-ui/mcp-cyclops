package modules

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

func (m *ModuleController) listModulesTool() mcp.Tool {
	return mcp.NewTool("list_modules",
		mcp.WithDescription("List All Cyclops Modules"),
	)
}

func (m *ModuleController) listModules(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	module, err := m.k8sClient.ListModules()
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(module)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(b)), nil
}
