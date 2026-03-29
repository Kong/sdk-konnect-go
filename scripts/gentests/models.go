package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

func loadModelPackages(repoRoot, modulePath string) (map[string]*packageInfo, error) {
	modelsRoot := filepath.Join(repoRoot, "models")
	packageDirs, err := discoverPackageDirs(modelsRoot)
	if err != nil {
		return nil, err
	}

	packages := make(map[string]*packageInfo, len(packageDirs))
	fset := token.NewFileSet()

	for _, packageDir := range packageDirs {
		goFiles, err := listGoFiles(packageDir)
		if err != nil {
			return nil, err
		}
		if len(goFiles) == 0 {
			continue
		}

		pkg := &packageInfo{
			Dir:        packageDir,
			ImportPath: path.Join(modulePath, filepath.ToSlash(relativePath(repoRoot, packageDir))),
			Files:      make(map[string]*fileInfo, len(goFiles)),
			Types:      make(map[string]*typeInfo),
			TypeNames:  make(map[string]struct{}),
		}

		for _, goFile := range goFiles {
			parsedFile, err := parser.ParseFile(fset, goFile, nil, 0)
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", goFile, err)
			}

			if pkg.Name == "" {
				pkg.Name = parsedFile.Name.Name
			}

			file := &fileInfo{
				Path:          goFile,
				ImportAliases: collectImports(parsedFile),
			}
			pkg.Files[goFile] = file

			for _, decl := range parsedFile.Decls {
				genDecl, ok := decl.(*ast.GenDecl)
				if !ok || genDecl.Tok != token.TYPE {
					continue
				}

				for _, spec := range genDecl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}

					pkg.TypeNames[typeSpec.Name.Name] = struct{}{}
					pkg.Types[typeSpec.Name.Name] = &typeInfo{
						Name:   typeSpec.Name.Name,
						Pkg:    pkg,
						File:   file,
						Expr:   typeSpec.Type,
						Struct: structType(typeSpec.Type),
					}
				}
			}
		}

		packages[pkg.ImportPath] = pkg
	}

	return packages, nil
}

func discoverPackageDirs(modelsRoot string) ([]string, error) {
	dirs := map[string]struct{}{}
	err := filepath.WalkDir(modelsRoot, func(currentPath string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		if !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}
		dirs[filepath.Dir(currentPath)] = struct{}{}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk model packages: %w", err)
	}

	packageDirs := make([]string, 0, len(dirs))
	for dir := range dirs {
		packageDirs = append(packageDirs, dir)
	}
	sort.Strings(packageDirs)
	return packageDirs, nil
}

func listGoFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", dir, err)
	}

	goFiles := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		goFiles = append(goFiles, filepath.Join(dir, name))
	}
	sort.Strings(goFiles)
	return goFiles, nil
}

func collectImports(file *ast.File) map[string]string {
	imports := make(map[string]string, len(file.Imports))
	for _, importSpec := range file.Imports {
		importPath := strings.Trim(importSpec.Path.Value, "\"")
		alias := path.Base(importPath)
		if importSpec.Name != nil {
			alias = importSpec.Name.Name
		}
		if alias == "_" || alias == "." {
			continue
		}
		imports[alias] = importPath
	}
	return imports
}

func structType(expr ast.Expr) *ast.StructType {
	structExpr, ok := expr.(*ast.StructType)
	if !ok {
		return nil
	}
	return structExpr
}

func resolveTargets(typeNames []string, packages map[string]*packageInfo) ([]*typeInfo, error) {
	resolved := make([]*typeInfo, 0, len(typeNames))
	for _, rawType := range typeNames {
		packageHint, typeName := splitTypeName(rawType)
		matches := findTypeMatches(packageHint, typeName, packages)

		switch len(matches) {
		case 0:
			if packageHint == "" {
				return nil, fmt.Errorf("type %q not found under models", typeName)
			}
			return nil, fmt.Errorf("type %q not found in package %q", typeName, packageHint)
		case 1:
			if matches[0].Struct == nil {
				return nil, fmt.Errorf("type %q in package %q is not a struct", matches[0].Name, matches[0].Pkg.Name)
			}
			resolved = append(resolved, matches[0])
		default:
			candidates := make([]string, 0, len(matches))
			for _, match := range matches {
				candidates = append(candidates, fmt.Sprintf("%s.%s", match.Pkg.Name, match.Name))
			}
			sort.Strings(candidates)
			return nil, fmt.Errorf("type %q is ambiguous; use one of: %s", rawType, strings.Join(candidates, ", "))
		}
	}

	return resolved, nil
}

func splitTypeName(rawType string) (string, string) {
	separator := strings.LastIndex(rawType, ".")
	if separator < 0 {
		return "", rawType
	}
	return rawType[:separator], rawType[separator+1:]
}

func findTypeMatches(packageHint, typeName string, packages map[string]*packageInfo) []*typeInfo {
	packageKeys := make([]string, 0, len(packages))
	for packageKey := range packages {
		packageKeys = append(packageKeys, packageKey)
	}
	sort.Strings(packageKeys)

	matches := make([]*typeInfo, 0, 1)
	for _, packageKey := range packageKeys {
		pkg := packages[packageKey]
		if packageHint != "" && !matchesPackageHint(pkg, packageHint) {
			continue
		}

		typeInfo, ok := pkg.Types[typeName]
		if !ok {
			continue
		}
		matches = append(matches, typeInfo)
	}

	return matches
}

func matchesPackageHint(pkg *packageInfo, packageHint string) bool {
	normalizedHint := filepath.ToSlash(packageHint)
	if pkg.Name == normalizedHint {
		return true
	}
	if filepath.Base(pkg.Dir) == normalizedHint {
		return true
	}
	if strings.HasSuffix(pkg.ImportPath, normalizedHint) {
		return true
	}
	return false
}
