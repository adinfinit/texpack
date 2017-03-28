package walk

import (
	"os"
	"path/filepath"
	"strings"
)

func Exts(folder string, exts []string, found func(name, filename string)) {
	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if !matchesAny(exts, filepath.Ext(path)) {
			return nil
		}

		name := path
		if rel, err := filepath.Rel(folder, path); err == nil {
			name = rel
		}

		found(name, path)

		return nil
	})
}

func matchesAny(xs []string, s string) bool {
	s = strings.ToLower(s)
	for _, x := range xs {
		if strings.ToLower(x) == s {
			return true
		}
	}
	return false
}
