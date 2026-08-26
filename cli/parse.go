package cli

import (
	"errors"
	"slices"
	"strings"
)

type definition interface {
	meta() *Definition
}

func Parse(argv []string, defs ...definition) ([]string, error) {
	for _, d := range defs {
		d.meta().reset()
	}

	var positionals []string
	for i := 0; i < len(argv); i++ {
		token := argv[i]
		if !strings.HasPrefix(token, "-") {
			positionals = append(positionals, token)
			continue
		}
		name := strings.TrimLeft(token, "-")
		var found definition
		for _, d := range defs {
			if slices.Contains(d.meta().names, name) {
				found = d
				break
			}
		}
		if found == nil {
			return nil, errors.New("unknown argument: " + token)
		}
		consumed, err := found.meta().take(argv[i+1:])
		if err != nil {
			return nil, err
		}
		found.meta().seen = true
		i += consumed
	}

	for _, d := range defs {
		m := d.meta()
		if m.required && !m.seen {
			return nil, errors.New("required argument missing: " + m.names[0])
		}
	}
	return positionals, nil
}
