package inspectorConfig

import (
	"log"

	"github.hpe.com/pablo-gon-sanchez/configurationHelper"
	yaml "gopkg.in/yaml.v2"
)

// InspectorConfig s
type InspectorConfig struct {
	Commands []Command `yaml:"commands"`
}

// Command is
type Command struct {
	Name       string `yaml:"name"`
	Parameters string `yaml:"parameters"`
}

func (ic InspectorConfig) unmarshallYaml(url string) InspectorConfig {
	btsConfig, err := configurationHelper.LoadRemote(url)
	log.Printf("raw config %v", string(btsConfig))
	if err != nil {
		log.Fatalf("error reading remote URL: %v\n", err)
	}

	err = yaml.Unmarshal(btsConfig, &ic)
	log.Printf("parsed config %v", ic)
	if err != nil {
		log.Fatalf("Error parsing yaml %v \n", err)
	}

	return ic
}
