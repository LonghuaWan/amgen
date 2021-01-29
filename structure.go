package amgen

type ModelGenerator struct {
	ConfigName  string  `yaml:"-"`
	FileName    string  `yaml:"-"`
	PackageName string  `yaml:"packageName,omitempty"`
	DbName		string	`yaml:"dbName,omitempty"`
	Models      []Model `yaml:"models,omitempty"`
	Raw         string  `yaml:"raw,omitempty"`
}

type Model struct {
	Name           string   `yaml:"name,omitempty"`
	CollectionName string   `yaml:"collectionName,omitempty"`
	CRUD           bool     `yaml:"CRUD,omitempty"`
	Types          []string `yaml:"types,omitempty"`
	States         []string `yaml:"states,omitempty"`
	Fields         []Field  `yaml:"fields,omitempty"`
	Indexes         []Index `yaml:"indexes,omitempty"`
}

type Index struct {
	Name   []string `yaml:"name,omitempty"`
	Unique bool   `yaml:"unique,omitempty"`
}

type Field struct {
	Name   string `yaml:"name,omitempty"`
	Type   string `yaml:"type,omitempty"`
	Unique bool   `yaml:"unique,omitempty"`
	Valid  string `yaml:"valid,omitempty"`
}
