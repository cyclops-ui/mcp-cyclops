package modules

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/cyclops-ui/mcp-cyclops/internal/mapper"
)

func (m *ModuleController) updateModuleTool() mcp.Tool {
	return mcp.NewTool("update_module",
		mcp.WithDescription("Update Module by Name"),
		mcp.WithString("module_name",
			mcp.Required(),
			mcp.Description("Name of the Module to update"),
		),
		mcp.WithString("values",
			mcp.Required(),
			mcp.Description("Helm-like values in JSON string format to update the module with"),
		),
	)
}

func (m *ModuleController) updateModule(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	moduleName := request.Params.Arguments["module_name"].(string)
	valuesRaw := request.Params.Arguments["values"].(string)

	// Parse JSON string into apiextensionsv1.JSON
	var updateValues map[string]interface{}
	if err := json.Unmarshal([]byte(valuesRaw), &updateValues); err != nil {
		return nil, fmt.Errorf("failed to parse JSON values: %w", err)
	}

	// Get the current module
	curr, err := m.k8sClient.GetModule(moduleName)
	if err != nil {
		return nil, fmt.Errorf("failed to get module: %w", err)
	}

	module, err := mapper.UpdateModuleValues(curr, updateValues)
	if err := m.k8sClient.UpdateModule(module); err != nil {
		return nil, fmt.Errorf("failed to update module: %w", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("Module %s updated successfully", moduleName)), nil
}
