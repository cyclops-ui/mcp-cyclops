package templates

import (
	"context"
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/api/v1alpha1"

	"github.com/mark3labs/mcp-go/mcp"
)

func (t *TemplateController) getTemplateSchemaTool() mcp.Tool {
	return mcp.NewTool("get_template_schema",
		mcp.WithDescription("Returns JSON schema for the given template. Needs to be checked before calling create_module tool"),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Type of the Template"),
			mcp.Enum("git", "helm", "oci"),
		),
		mcp.WithString("repo",
			mcp.Required(),
			mcp.Description("Template repo (Helm or Git repo)"),
		),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("In case of a git repo, folder of the template in the repository. For a helm chart, the name of the chart"),
		),
		mcp.WithString("version",
			mcp.Required(),
			mcp.Description("Semantic version of the chart, or in case of a git repo can be a reference as commit hash or branch/tag"),
		),
		mcp.WithString("icon",
			mcp.Description("Icon URL for the template"),
		),
	)
}

func (t *TemplateController) getTemplateSchema(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	templateType := request.Params.Arguments["type"].(string)
	repo := request.Params.Arguments["repo"].(string)
	path := request.Params.Arguments["path"].(string)
	version := request.Params.Arguments["version"].(string)

	template, err := t.templateRepo.GetTemplate(repo, path, version, "", v1alpha1.TemplateSourceType(templateType))
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(template.RawSchema)), nil
}
