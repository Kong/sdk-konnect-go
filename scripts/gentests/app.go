package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func run(rawArgs []string) int {
	args, err := parseArgs(rawArgs)
	if err != nil {
		return fail(err)
	}

	repoRoot, err := findRepoRoot()
	if err != nil {
		return fail(err)
	}

	modulePath, err := readModulePath(filepath.Join(repoRoot, "go.mod"))
	if err != nil {
		return fail(err)
	}

	packages, err := loadModelPackages(repoRoot, modulePath)
	if err != nil {
		return fail(err)
	}

	targets, err := resolveTargets(args.TypeNames, packages)
	if err != nil {
		return fail(err)
	}

	outputDir := filepath.Join(repoRoot, "test", "fields")
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fail(fmt.Errorf("create output directory: %w", err))
	}

	generatedCount := 0
	for _, target := range targets {
		output, err := generateTestFile(target)
		if err != nil {
			return fail(fmt.Errorf("generate %s.%s: %w", target.Pkg.Name, target.Name, err))
		}

		fileName := fmt.Sprintf("%s_%s_fields_test.go", target.Pkg.Name, toSnakeCase(target.Name))
		filePath := filepath.Join(outputDir, fileName)
		if err := os.WriteFile(filePath, output, 0o644); err != nil {
			return fail(fmt.Errorf("write %s: %w", filePath, err))
		}

		generatedCount++
		fmt.Println(filepath.ToSlash(relativePath(repoRoot, filePath)))
	}

	fmt.Printf("generated %d field guard test(s)\n", generatedCount)
	return 0
}

func fail(err error) int {
	fmt.Fprintln(os.Stderr, "error:", err)
	return 1
}
