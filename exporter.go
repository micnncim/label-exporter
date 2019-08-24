package exporter

import (
	"encoding/json"

	"sigs.k8s.io/yaml"
)

func LabelsToJSON(labels []*Label) ([]byte, error) {
	return json.Marshal(labels)
}

func LabelsToYAML(labels []*Label) ([]byte, error) {
	return yaml.Marshal(labels)
}
