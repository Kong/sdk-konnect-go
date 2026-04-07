package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func findRepoRoot() (string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("resolve script location")
	}

	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(currentFile), "..", ".."))
	if _, err := os.Stat(filepath.Join(repoRoot, "go.mod")); err != nil {
		return "", fmt.Errorf("locate repository root from %s: %w", currentFile, err)
	}
	return repoRoot, nil
}

func readModulePath(goModPath string) (string, error) {
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return "", fmt.Errorf("read go.mod: %w", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "module ") {
			modulePath := strings.TrimSpace(strings.TrimPrefix(trimmed, "module "))
			if modulePath == "" {
				break
			}
			return modulePath, nil
		}
	}

	return "", fmt.Errorf("module path not found in %s", goModPath)
}

func relativePath(basePath, targetPath string) string {
	relative, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		return targetPath
	}
	return relative
}
