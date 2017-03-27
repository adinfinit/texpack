package main

import (
	"os"
	"path/filepath"
	"strings"
)

func Walk(folder string, exts []string, found func(name, filename string)) {
	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if !MatchesAny(exts, ext) {
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

func MatchesAny(xs []string, s string) bool {
	for _, x := range xs {
		if x == s {
			return true
		}
	}
	return false
}
