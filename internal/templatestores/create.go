package templatestores

import (
	"context"
	"fmt"
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/api/v1alpha1"
	"github.com/mark3labs/mcp-go/mcp"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (t *TemplateStoreController) createTemplateStoreTool() mcp.Tool {
	return mcp.NewTool("create_template_store",
		mcp.WithDescription("Create Template Stores"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the Template Stores to create"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Type of the Template Stores to create"),
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

func (t *TemplateStoreController) createTemplateStore(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name := request.Params.Arguments["name"].(string)
	templateType := request.Params.Arguments["type"].(string)
	repo := request.Params.Arguments["repo"].(string)
	path := request.Params.Arguments["path"].(string)
	version := request.Params.Arguments["version"].(string)

	annotations := map[string]string{}
	if iconParam, ok := request.Params.Arguments["icon"]; ok {
		annotations = map[string]string{
			v1alpha1.IconURLAnnotation: iconParam.(string),
		}
	}

	if err := t.k8sClient.CreateTemplateStore(&v1alpha1.TemplateStore{
		TypeMeta: v1.TypeMeta{
			Kind:       "TemplateStore",
			APIVersion: "cyclops-ui.com/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:        name,
			Annotations: annotations,
		},
		Spec: v1alpha1.TemplateRef{
			URL:        repo,
			Path:       path,
			Version:    version,
			SourceType: v1alpha1.TemplateSourceType(templateType),
		},
	}); err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(fmt.Sprintf("Template Store %s updated successfully", name)), nil
}
