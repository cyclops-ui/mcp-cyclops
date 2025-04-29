package modules

import (
	"github.com/pkg/errors"

	"github.com/mark3labs/mcp-go/server"
	"github.com/xeipuuv/gojsonschema"

	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/cluster/k8sclient"
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/pkg/template"
)

type ModuleController struct {
	k8sClient    *k8sclient.KubernetesClient
	templateRepo template.ITemplateRepo
}

func NewController(k8sClient *k8sclient.KubernetesClient, templateRepo template.ITemplateRepo) *ModuleController {
	return &ModuleController{
		k8sClient:    k8sClient,
		templateRepo: templateRepo,
	}
}

func (m *ModuleController) RegisterModuleTools(mcp *server.MCPServer) {
	mcp.AddTool(m.getModuleByNameTool(), m.getModuleByName)
	mcp.AddTool(m.listModulesTool(), m.listModules)
	mcp.AddTool(m.createModuleTool(), m.createModule)
	mcp.AddTool(m.updateModuleTool(), m.updateModule)
}

func (m *ModuleController) validateModuleValues(schema []byte, values map[string]interface{}) (bool, error, error) {
	schemaLoader := gojsonschema.NewBytesLoader(schema)
	valuesLoader := gojsonschema.NewGoLoader(values)

	res, err := gojsonschema.Validate(schemaLoader, valuesLoader)
	if err != nil {
		return false, nil, err
	}

	if !res.Valid() {
		validationErr := errors.New("invalid values for schema")

		for _, resultError := range res.Errors() {
			validationErr = errors.Wrap(validationErr, resultError.String())
		}

		return false, validationErr, nil
	}

	return true, nil, nil
}
