package modules

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"sigs.k8s.io/yaml"

	"github.com/cyclops-ui/cyclops/cyclops-ctrl/api/v1alpha1"

	"github.com/cyclops-ui/mcp-cyclops/internal/mapper"
)

func (m *ModuleController) createModuleManifestTool() mcp.Tool {
	return mcp.NewTool("create_module_manifest",
		mcp.WithDescription("Create a Module manifest based on values. Before calling this tool, make sure to call get_template_schema to validate values for the given template"),
		mcp.WithString("module_name",
			mcp.Required(),
			mcp.Description("Name of the Module to update"),
		),
		mcp.WithString("template_type",
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
		mcp.WithString("values",
			mcp.Required(),
			mcp.Description("Helm-like values in JSON string format to create the module with"),
		),
	)
}

func (m *ModuleController) createModuleManifestModule(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	moduleName := request.Params.Arguments["module_name"].(string)
	valuesRaw := request.Params.Arguments["values"].(string)
	templateType := request.Params.Arguments["template_type"].(string)
	repo := request.Params.Arguments["repo"].(string)
	path := request.Params.Arguments["path"].(string)
	version := request.Params.Arguments["version"].(string)

	initialValues, err := m.templateRepo.GetTemplateInitialValues(repo, path, version, v1alpha1.TemplateSourceType(templateType))
	if err != nil {
		return nil, err
	}

	var values map[string]interface{}
	if len(valuesRaw) > 0 {
		if err := json.Unmarshal([]byte(valuesRaw), &values); err != nil {
			return nil, fmt.Errorf("failed to parse current values: %w", err)
		}
	} else {
		values = make(map[string]interface{})
	}

	values = mapper.DeepMerge(initialValues, values)

	template, err := m.templateRepo.GetTemplate(repo, path, version, "", v1alpha1.TemplateSourceType(templateType))
	if err != nil {
		return nil, err
	}

	valid, validationError, err := m.validateModuleValues(template.RawSchema, values)
	if err != nil {
		return nil, err
	}

	if !valid {
		return mcp.NewToolResultError(validationError.Error()), nil
	}

	valuesBytes, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	module := mapper.CreateModule(moduleName, repo, path, version, templateType, valuesBytes)

	moduleData, err := yaml.Marshal(module)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(moduleData)), nil
}
