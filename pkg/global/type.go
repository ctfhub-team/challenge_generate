package global

type Config struct {
	Author      string `yaml:"author"`
	Contact     string `yaml:"contact"`
	RegistryUrl string `yaml:"registry_url"`
}

type DockerCompsoe struct {
	Version  string `yaml:"version"`
	Services struct {
		Challenge struct {
			Build       string   `yaml:"build"`
			Image       string   `yaml:"image"`
			Ports       []string `yaml:"ports"`
			Environment []string `yaml:"environment"`
		} `yaml:"challenge"`
	} `yaml:"services"`
}

type Meta struct {
	Author struct {
		Name    string `yaml:"name"`
		Contact string `yaml:"contact"`
	} `yaml:"author"`
	Task struct {
		Name        string   `yaml:"name"`
		Type        string   `yaml:"type"`
		Description string   `yaml:"description"`
		Level       string   `yaml:"level"`
		Flag        string   `yaml:"flag"`
		Hints       []string `yaml:"hints"`
	} `yaml:"task"`
	Challenge struct {
		Name  string   `yaml:"name"`
		Refer string   `yaml:"refer"`
		Tags  []string `yaml:"tags"`
	} `yaml:"challenge"`
}
