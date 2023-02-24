package yaml

import "sigs.k8s.io/yaml"

func MustToYaml(obj any) string {
	data, err := yaml.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return string(data)
}
