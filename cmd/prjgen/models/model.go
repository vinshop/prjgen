package models

type ModelField struct {
	Name       string
	DataType   string
	GoType     string
	Tags       map[string][]string
	IsNullable bool
	IsPrimary  bool
}

type Model struct {
	Base
	Name  string
	Table string
}
