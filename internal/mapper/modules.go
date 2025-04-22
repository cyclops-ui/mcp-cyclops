package mapper

import (
	"encoding/json"
	"fmt"
	"github.com/cyclops-ui/cyclops/cyclops-ctrl/api/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateModule(moduleName, repo, path, version, sourceType string, values []byte) v1alpha1.Module {
	return v1alpha1.Module{
		TypeMeta: v1.TypeMeta{
			Kind:       "Module",
			APIVersion: "cyclops-ui.com/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: moduleName,
		},
		Spec: v1alpha1.ModuleSpec{
			TemplateRef: v1alpha1.TemplateRef{
				URL:        repo,
				Path:       path,
				Version:    version,
				SourceType: v1alpha1.TemplateSourceType(sourceType),
			},
			Values: apiextensionsv1.JSON{
				Raw: values,
			},
		},
	}
}

func UpdateModuleValues(module *v1alpha1.Module, values map[string]interface{}) (*v1alpha1.Module, error) {
	// region values
	var curr map[string]interface{}
	if len(module.Spec.Values.Raw) > 0 {
		if err := json.Unmarshal(module.Spec.Values.Raw, &curr); err != nil {
			return nil, fmt.Errorf("failed to parse current values: %w", err)
		}
	} else {
		curr = make(map[string]interface{})
	}

	merged := DeepMerge(curr, values)
	mergedBytes, err := json.Marshal(merged)
	if err != nil {
		return nil, fmt.Errorf("failed to encode merged values: %w", err)
	}
	// endregion

	// region history
	history := module.History
	if module.History == nil {
		history = make([]v1alpha1.HistoryEntry, 0)
	}

	history = append([]v1alpha1.HistoryEntry{{
		Generation:      module.Generation,
		TargetNamespace: module.Spec.TargetNamespace,
		TemplateRef: v1alpha1.HistoryTemplateRef{
			URL:        module.Spec.TemplateRef.URL,
			Path:       module.Spec.TemplateRef.Path,
			Version:    module.Status.TemplateResolvedVersion,
			SourceType: module.Spec.TemplateRef.SourceType,
		},
		Values: module.Spec.Values,
	}}, history...)

	if len(module.History) > 10 {
		module.History = module.History[:len(module.History)-1]
	}
	// endregion

	return &v1alpha1.Module{
		TypeMeta: v1.TypeMeta{
			Kind:       "Module",
			APIVersion: "cyclops-ui.com/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:            module.Name,
			Annotations:     module.Annotations,
			ResourceVersion: module.GetResourceVersion(),
		},
		Spec: v1alpha1.ModuleSpec{
			TargetNamespace: module.Spec.TargetNamespace,
			TemplateRef:     module.Spec.TemplateRef,
			Values: apiextensionsv1.JSON{
				Raw: mergedBytes,
			},
		},
		History: history,
		Status:  module.Status,
	}, nil
}

func DeepMerge(dst, src map[string]interface{}) map[string]interface{} {
	for k, v := range src {
		if vMap, ok := v.(map[string]interface{}); ok {
			if dstMap, found := dst[k].(map[string]interface{}); found {
				dst[k] = DeepMerge(dstMap, vMap)
				continue
			}
		}
		dst[k] = v
	}
	return dst
}
