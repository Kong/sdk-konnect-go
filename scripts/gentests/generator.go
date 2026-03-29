package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"path"
	"sort"
	"strings"
)

func writeImportBlock(builder *strings.Builder, imports map[string]string) {
	specs := make([]importSpec, 0, len(imports))
	for alias, importPath := range imports {
		specs = append(specs, importSpec{Alias: alias, Path: importPath})
	}
	sort.Slice(specs, func(left, right int) bool {
		return specs[left].Path < specs[right].Path
	})

	var stdImports, extImports []importSpec
	for _, spec := range specs {
		if !strings.Contains(spec.Path, ".") {
			stdImports = append(stdImports, spec)
		} else {
			extImports = append(extImports, spec)
		}
	}

	builder.WriteString("import (\n")
	writeImportSpecs(builder, stdImports)
	if len(stdImports) > 0 && len(extImports) > 0 {
		builder.WriteString("\n")
	}
	writeImportSpecs(builder, extImports)
	builder.WriteString(")\n\n")
}

func writeImportSpecs(builder *strings.Builder, specs []importSpec) {
	for _, spec := range specs {
		builder.WriteString("\t")
		if defaultAlias := path.Base(spec.Path); spec.Alias != defaultAlias {
			builder.WriteString(spec.Alias)
			builder.WriteString(" ")
		}
		fmt.Fprintf(builder, "\"%s\"\n", spec.Path)
	}
}

func generateTestFile(target *typeInfo) ([]byte, error) {
	renderer := &typeRenderer{
		target: target,
		requiredImports: map[string]string{
			"testing":       "testing",
			target.Pkg.Name: target.Pkg.ImportPath,
		},
	}

	assertions, err := buildAssertions(target, renderer)
	if err != nil {
		return nil, err
	}

	var builder strings.Builder
	builder.WriteString(generatedHeader)
	builder.WriteString("\n")
	builder.WriteString("package fields_test\n\n")
	writeImportBlock(&builder, renderer.requiredImports)
	fmt.Fprintf(&builder, "func Test%s%sFieldTypes(t *testing.T) {\n", exportName(target.Pkg.Name), target.Name)
	builder.WriteString("\tt.Parallel()\n\n")
	fmt.Fprintf(&builder, "\tvar value %s.%s\n", target.Pkg.Name, target.Name)
	if len(assertions) == 0 {
		builder.WriteString("\t_ = value\n")
	} else {
		for _, assertion := range assertions {
			fmt.Fprintf(&builder, "\tvar _ %s = value.%s\n", assertion.TypeName, assertion.FieldName)
		}
	}
	builder.WriteString("}\n")

	formatted, err := format.Source([]byte(builder.String()))
	if err != nil {
		return nil, fmt.Errorf("format source: %w", err)
	}
	return formatted, nil
}

func buildAssertions(target *typeInfo, renderer *typeRenderer) ([]fieldAssertion, error) {
	assertions := make([]fieldAssertion, 0, len(target.Struct.Fields.List))
	for _, field := range target.Struct.Fields.List {
		typeName, err := renderer.render(field.Type)
		if err != nil {
			return nil, err
		}

		fieldNames, err := fieldNames(field)
		if err != nil {
			return nil, err
		}
		for _, fieldName := range fieldNames {
			if !ast.IsExported(fieldName) {
				return nil, fmt.Errorf("field %s.%s is not exported; generated tests in test/fields cannot access it", target.Name, fieldName)
			}
			assertions = append(assertions, fieldAssertion{FieldName: fieldName, TypeName: typeName})
		}
	}
	return assertions, nil
}

func fieldNames(field *ast.Field) ([]string, error) {
	if len(field.Names) > 0 {
		names := make([]string, 0, len(field.Names))
		for _, name := range field.Names {
			names = append(names, name.Name)
		}
		return names, nil
	}

	name := embeddedFieldName(field.Type)
	if name == "" {
		return nil, fmt.Errorf("unsupported embedded field type %T", field.Type)
	}
	return []string{name}, nil
}

func embeddedFieldName(expr ast.Expr) string {
	switch value := expr.(type) {
	case *ast.Ident:
		return value.Name
	case *ast.SelectorExpr:
		return value.Sel.Name
	case *ast.StarExpr:
		return embeddedFieldName(value.X)
	case *ast.IndexExpr:
		return embeddedFieldName(value.X)
	case *ast.IndexListExpr:
		return embeddedFieldName(value.X)
	default:
		return ""
	}
}

func (renderer *typeRenderer) render(expr ast.Expr) (string, error) {
	switch value := expr.(type) {
	case *ast.Ident:
		if _, ok := builtinIdentifiers[value.Name]; ok {
			return value.Name, nil
		}
		if _, ok := renderer.target.Pkg.TypeNames[value.Name]; ok {
			if err := renderer.addImport(renderer.target.Pkg.Name, renderer.target.Pkg.ImportPath); err != nil {
				return "", err
			}
			return renderer.target.Pkg.Name + "." + value.Name, nil
		}
		return value.Name, nil
	case *ast.SelectorExpr:
		if baseIdentifier, ok := value.X.(*ast.Ident); ok {
			if importPath, imported := renderer.target.File.ImportAliases[baseIdentifier.Name]; imported {
				if err := renderer.addImport(baseIdentifier.Name, importPath); err != nil {
					return "", err
				}
				return baseIdentifier.Name + "." + value.Sel.Name, nil
			}
		}

		left, err := renderer.render(value.X)
		if err != nil {
			return "", err
		}
		return left + "." + value.Sel.Name, nil
	case *ast.StarExpr:
		x, err := renderer.render(value.X)
		if err != nil {
			return "", err
		}
		return "*" + x, nil
	case *ast.ArrayType:
		length := ""
		if value.Len != nil {
			formattedLength, err := renderer.render(value.Len)
			if err != nil {
				return "", err
			}
			length = formattedLength
		}
		elt, err := renderer.render(value.Elt)
		if err != nil {
			return "", err
		}
		return "[" + length + "]" + elt, nil
	case *ast.MapType:
		key, err := renderer.render(value.Key)
		if err != nil {
			return "", err
		}
		entryValue, err := renderer.render(value.Value)
		if err != nil {
			return "", err
		}
		return "map[" + key + "]" + entryValue, nil
	case *ast.ParenExpr:
		x, err := renderer.render(value.X)
		if err != nil {
			return "", err
		}
		return "(" + x + ")", nil
	case *ast.IndexExpr:
		x, err := renderer.render(value.X)
		if err != nil {
			return "", err
		}
		index, err := renderer.render(value.Index)
		if err != nil {
			return "", err
		}
		return x + "[" + index + "]", nil
	case *ast.IndexListExpr:
		x, err := renderer.render(value.X)
		if err != nil {
			return "", err
		}
		indices := make([]string, 0, len(value.Indices))
		for _, indexExpr := range value.Indices {
			index, err := renderer.render(indexExpr)
			if err != nil {
				return "", err
			}
			indices = append(indices, index)
		}
		return x + "[" + strings.Join(indices, ", ") + "]", nil
	case *ast.ChanType:
		valueType, err := renderer.render(value.Value)
		if err != nil {
			return "", err
		}
		switch value.Dir {
		case ast.RECV:
			return "<-chan " + valueType, nil
		case ast.SEND:
			return "chan<- " + valueType, nil
		default:
			return "chan " + valueType, nil
		}
	case *ast.Ellipsis:
		if value.Elt == nil {
			return "...", nil
		}
		elt, err := renderer.render(value.Elt)
		if err != nil {
			return "", err
		}
		return "..." + elt, nil
	default:
		return "", fmt.Errorf("unsupported field type expression %T", expr)
	}
}

func (renderer *typeRenderer) addImport(alias, importPath string) error {
	if existingPath, ok := renderer.requiredImports[alias]; ok && existingPath != importPath {
		return fmt.Errorf("import alias conflict for %q: %s and %s", alias, existingPath, importPath)
	}
	renderer.requiredImports[alias] = importPath
	return nil
}

func exportName(value string) string {
	if value == "" {
		return value
	}
	return strings.ToUpper(value[:1]) + value[1:]
}

func toSnakeCase(value string) string {
	if value == "" {
		return value
	}

	var builder strings.Builder
	runes := []rune(value)
	for index, current := range runes {
		if index > 0 && isSnakeBoundary(runes, index) {
			builder.WriteByte('_')
		}
		builder.WriteRune(toLower(current))
	}
	return builder.String()
}

func isSnakeBoundary(runes []rune, index int) bool {
	current := runes[index]
	previous := runes[index-1]
	if isUpper(current) && (isLower(previous) || isDigit(previous)) {
		return true
	}
	if isUpper(current) && index+1 < len(runes) && isLower(runes[index+1]) && isUpper(previous) {
		return true
	}
	return false
}

func isUpper(value rune) bool {
	return value >= 'A' && value <= 'Z'
}

func isLower(value rune) bool {
	return value >= 'a' && value <= 'z'
}

func isDigit(value rune) bool {
	return value >= '0' && value <= '9'
}

func toLower(value rune) rune {
	if isUpper(value) {
		return value + ('a' - 'A')
	}
	return value
}
