package main

type kubeResource struct {
	ApiVersion string `yaml:"apiVersion"`
	kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml: "name"`
		Namespace string `yaml: "namespace"`
	} `yaml: "metadata"`
}
