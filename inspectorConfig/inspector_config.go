package inspectorConfig

var (
	// InspectorConfiguration is the global configuration file
	InspectorConfiguration InspectorConfig
)

// LoadConfig returns a map of the config
func LoadConfig() {
	InspectorConfiguration = InspectorConfiguration.unmarshallYaml("https://raw.github.hpe.com/pablo-gon-sanchez/inspector-gadget/master/inspectorConfig/config.yml")
}
