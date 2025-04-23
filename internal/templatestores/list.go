package templatestores

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

func (t *TemplateStoreController) listTemplateStoresTool() mcp.Tool {
	return mcp.NewTool("list_template_store",
		mcp.WithDescription("List Template Stores from cluster"),
	)
}

func (t *TemplateStoreController) listTemplateStores(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	templateStores, err := t.k8sClient.ListTemplateStore()
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(templateStores)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(b)), nil
}
