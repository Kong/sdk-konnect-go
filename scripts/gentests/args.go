package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type cliArgs struct {
	TypeNames []string
}

func parseArgs(rawArgs []string) (*cliArgs, error) {
	flagSet := flag.NewFlagSet("gentests", flag.ContinueOnError)
	flagSet.SetOutput(os.Stderr)

	typesFlag := flagSet.String("types", "", "comma-separated list of type names to generate tests for")
	if err := flagSet.Parse(rawArgs); err != nil {
		return nil, err
	}

	input := strings.TrimSpace(*typesFlag)
	if input == "" && flagSet.NArg() > 0 {
		input = flagSet.Arg(0)
	}
	if input == "" {
		return nil, fmt.Errorf("missing type list; pass --types=Consumer,Certificate or a positional comma-separated argument")
	}

	parts := strings.Split(input, ",")
	seen := make(map[string]struct{}, len(parts))
	typeNames := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			continue
		}
		seen[trimmed] = struct{}{}
		typeNames = append(typeNames, trimmed)
	}

	if len(typeNames) == 0 {
		return nil, fmt.Errorf("no type names found in input %q", input)
	}

	return &cliArgs{TypeNames: typeNames}, nil
}
