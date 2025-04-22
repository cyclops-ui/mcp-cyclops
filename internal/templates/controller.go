package templates

import (
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/template"
	"github.com/mark3labs/mcp-go/server"
)

type TemplateController struct {
	templateRepo template.ITemplateRepo
}

func NewController(templateRepo template.ITemplateRepo) *TemplateController {
	return &TemplateController{
		templateRepo: templateRepo,
	}
}

func (t *TemplateController) RegisterTemplateStoreTools(mcp *server.MCPServer) {
	mcp.AddTool(t.getTemplateSchemaTool(), t.getTemplateSchema)
}
