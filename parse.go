package wire

import (
	"github.com/BurntSushi/toml"
)

type Field struct {
	Name string `toml:"name"`
	Type string `toml:"type"`
}

type Message struct {
	Name   string  `toml:"name"`
	Fields []Field `toml:"field"`
}

type Definition struct {
	Package  string    `toml:"package"`
	Messages []Message `toml:"message"`
}

func parse(path string) (Definition, error) {
	definition := Definition{}
	_, err := toml.DecodeFile(path, &definition)
	if err != nil {
		return Definition{}, err
	}
	return definition, nil
}
