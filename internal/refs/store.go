package refs

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const FileName = ".references.yml"

// Find walks up from dir looking for .references.yml, returns the path.
func Find(dir string) (string, error) {
	for {
		p := filepath.Join(dir, FileName)
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("no .references.yml found")
		}
		dir = parent
	}
}

// Load reads and parses a .references.yml file.
func Load(path string) (*File, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var f File
	if err := yaml.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	return &f, nil
}

// Save writes the file back to disk.
func Save(path string, f *File) error {
	data, err := yaml.Marshal(f)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// Init creates a new empty .references.yml in the given directory.
func Init(dir string) (string, error) {
	p := filepath.Join(dir, FileName)
	if _, err := os.Stat(p); err == nil {
		return p, errors.New(".references.yml already exists")
	}
	f := &File{References: []Reference{}}
	if err := Save(p, f); err != nil {
		return "", err
	}
	return p, nil
}

// Add appends a reference and saves.
func Add(path string, ref Reference) error {
	f, err := Load(path)
	if err != nil {
		return err
	}
	f.References = append(f.References, ref)
	return Save(path, f)
}
