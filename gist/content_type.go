package gist

import (
	"net/http"
	"path"
	"strings"
)

func detectContentTypeAndLanguage(fileName string, content []byte) (string, string) {
	contentType := http.DetectContentType(content)
	language := "text"
	switch strings.ToLower(path.Ext(fileName)) {
	case ".go":
		language = "go"
		contentType = "application/go"
	case ".py":
		language = "python"
		contentType = "application/python"
	case ".js":
		language = "javascript"
		contentType = "application/javascript"

	case ".sh":
		language = "shell"
		contentType = "application/shell"

	case ".c":
		language = "c"
		contentType = "application/c"

	case ".cpp":
		language = "cpp"
		contentType = "application/cpp"

	case ".java":
		language = "java"
		contentType = "application/java"

	case ".html":
		language = "html"
		contentType = "text/html"

	case ".css":
		language = "css"
		contentType = "text/css"

	case ".json":
		language = "json"
		contentType = "application/json"
	case ".md":
		language = "markdown"
		contentType = "text/markdown"
	case ".toml", ".tml":
		language = "toml"
		contentType = "application/toml"

	case ".yml", ".yaml":
		language = "yaml"
		contentType = "application/yaml"

	case ".xml":
		language = "xml"
		contentType = "application/xml"

	case ".txt":
		language = "text"
		contentType = "text/plain"

	case ".sql":
		language = "sql"
		contentType = "application/sql"

	case ".rb":
		language = "ruby"
		contentType = "application/ruby"

	case ".rs":
		language = "rust"
		contentType = "application/rust"

	case ".php":
		language = "php"
		contentType = "application/php"

	case ".pl":
		language = "perl"
		contentType = "application/perl"

	case ".lua":
		language = "lua"

	case ".kt", ".ktm":
		language = "kotlin"

	}

	return contentType, language
}
