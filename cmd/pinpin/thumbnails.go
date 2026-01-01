package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

var exts = []string{".png", ".jpg"}

func pickUserThumbnailRaw(filePath string) []byte {
	fh, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	raw, err := io.ReadAll(fh)
	if err != nil {
		panic(err)
	}

	return raw
}

func findFirstThumbnailWithExts(dir, baseFile string) (string, error) {
	baseName := strings.TrimSuffix(baseFile, filepath.Ext(baseFile))

	for _, ext := range exts {
		fileName := baseName + ext
		fullPath := filepath.Join(dir, fileName)

		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, nil
		} else if !os.IsNotExist(err) {
			return "", err
		}
	}

	return "", os.ErrNotExist
}
