package util

import "strings"

func GetGoType(sqlType string) string {
	sqlType = strings.ToLower(sqlType)
	switch sqlType {
	case "char", "varchar", "binary", "varbinary", "tinyblob", "blob", "mediumblob",
		"longblob", "tinytext", "text", "mediumtext", "longtext", "enum", "set":
		return "string"
	case "tinyint":
		return "bool"
	case "smallint", "mediumint", "int", "bigint":
		return "int64"
	case "decimal", "float", "double":
		return "float64"
	case "json":
		return "string"
	default:
		return "interface{}"
	}
}
