package dto

type Service struct {
	FilePath    string      `yaml:"-" json:"file_path,omitempty"`
	Build       interface{} `yaml:"-" json:"-,omitempty"`
	DomainName  string      `yaml:"-" json:"domain_name,omitempty"`
	Image       string      `yaml:"image" json:"image,omitempty"`
	Environment []string    `yaml:"environment" json:"environment,omitempty"`
	Ports       []string    `yaml:"ports" json:"ports,omitempty"`
	Networks    []string    `yaml:"networks" json:"networks,omitempty"`
	Volumes     []string    `yaml:"volumes" json:"volumes,omitempty"`
	Type        string      `yaml:"-" json:"type,omitempty"`
}

type Network struct {
	Name     string `yaml:"name" json:"name,omitempty"`
	External bool   `yaml:"external" json:"external,omitempty"`
}

type DockerCompose struct {
	Version  string              `yaml:"version"`
	Services map[string]*Service `yaml:"services"`
	Networks map[string]*Network `yaml:"networks"`
}

type RegistryData struct {
	SystemServices map[string]*Service `json:"system_services,omitempty"`
	UserServices   map[string]*Service `json:"user_services,omitempty"`
}
