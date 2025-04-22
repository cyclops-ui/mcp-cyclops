package templatestores

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

func (t *TemplateStoreController) getTemplateStoreByNameTool() mcp.Tool {
	return mcp.NewTool("get_template_store",
		mcp.WithDescription("Fetch Template Store by Name"),
		mcp.WithString("template_store_name",
			mcp.Required(),
			mcp.Description("Name of the Template Store"),
		),
	)
}

func (t *TemplateStoreController) getTemplateStoreByName(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	templateStoreName := request.Params.Arguments["template_store_name"].(string)

	templateStore, err := t.k8sClient.GetTemplateStore(templateStoreName)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(templateStore)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(b)), nil
}
