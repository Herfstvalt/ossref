package refs

// File represents the top-level .references.yml structure.
type File struct {
	References []Reference `yaml:"references"`
}

// Reference is a single architectural reference to another project.
type Reference struct {
	Project string `yaml:"project"`
	Learned string `yaml:"learned"`
	Applied string `yaml:"applied,omitempty"`
}
