package helper

import (
	"path/filepath"
)

func Include(path string) []string {
	files, _ := filepath.Glob("web/views/partials/*.html")
	pathFiles, _ := filepath.Glob("web/views/" + path + "/*.html")

	files = append(files, pathFiles...)
	return files
}
