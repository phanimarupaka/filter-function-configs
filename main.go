package main

import (
	"fmt"
	"os"

	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/command"
	"sigs.k8s.io/kustomize/kyaml/kio/filters"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

//nolint
func main() {
	fp := FilterLocalConfigProcessor{}
	cmd := command.Build(&fp, command.StandaloneEnabled, false)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (fp *FilterLocalConfigProcessor) Process(resourceList *framework.ResourceList) error {
	var res []*yaml.RNode
	for _, node := range resourceList.Items {
		if node.IsNilOrEmpty() {
			continue
		}
		// only add the resources which are not local configs
		if node.GetAnnotations()[filters.LocalConfigAnnotation] != "true" {
			res = append(res, node)
		}
	}
	resourceList.Result = &framework.Result{
		Name: "filter-local-configs",
	}
	resourceList.Items = res
	return nil
}

type FilterLocalConfigProcessor struct{}
