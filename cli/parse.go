package cli

import (
	"errors"
	"strings"
)

type definition interface {
	meta() *Definition
}

func Parse(argv []string, defs ...definition) ([]string, error) {
	for _, def := range defs {
		def.meta().reset()
	}

	for _, def := range defs {
		meta := def.meta()
		if meta.required && meta.hasDefault {
			return nil, errors.New("argument cannot be required if has default value: " + meta.names[0])
		}
	}

	var positionals []string
	for i := 0; i < len(argv); i++ {
		token := argv[i]

		// positional
		if !strings.HasPrefix(token, "-") {
			positionals = append(positionals, token)
			continue
		}

		// positional with names containing -, --
		if token == "--" {
			positionals = append(positionals, argv[i+1:]...)
			break
		}
		
		// long form: --name
		if strings.HasPrefix(token, "--") {
			name := strings.TrimPrefix(token, "--")
			def := lookupDefinitionByName(name, false, defs...)
			if def == nil {
				return nil, errors.New("unknown argument: " + token)
			}
			consumed, err := def.meta().take(argv[i+1:])
			if err != nil {
				return nil, err
			}
			def.meta().seen = true
			i += consumed
			continue
		}

		// short form: -a or cluster -abc
		name := strings.TrimPrefix(token, "-")
		var consumed int
		for _, char := range name {
			def := lookupDefinitionByName(string(char), true, defs...)
			if def == nil {
				return nil, errors.New("unknown argument: -" + string(char))
			}
			if len(name) > 1 {
				if _, isFlag := def.meta().behavior.(*Flag); !isFlag {
					return nil, errors.New("-" + string(char) + " takes a value and cannot be part of a cluster")
				}
			}
			tokensTaken, err := def.meta().take(argv[i+1:])
			if err != nil {
				return nil, err
			}
			def.meta().seen = true
			consumed += tokensTaken
		}
		i += consumed
	}

	for _, def := range defs {
		meta := def.meta()
		if meta.required && !meta.seen {
			return nil, errors.New("required argument missing: " + meta.names[0])
		}
	}
	return positionals, nil
}

func lookupDefinitionByName(name string, short bool, defs ...definition) definition {
	for _, def := range defs {
		for _, candidate := range def.meta().names {
			if short && len(candidate) == 1 && candidate == name {
				return def
			}
			if !short && len(candidate) > 1 && candidate == name {
				return def
			}
		}
	}
	return nil
}
