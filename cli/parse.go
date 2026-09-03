package cli

import (
	"errors"
	"strings"
)

type definition interface {
	meta() *Definition
}

func Parse(argv []string, defs ...definition) ([]string, error) {
	for _, d := range defs {
		d.meta().reset()
	}

	for _, d := range defs {
		m := d.meta()
		if m.required && m.hasDefault {
			return nil, errors.New("argument cannot be required if has default value: " + m.names[0])
		}
	} 
	
	var positionals []string
	for i := 0; i < len(argv); i++ {
		var nameLong string
		var nameShort string
		token := argv[i]
		if strings.HasPrefix(token, "--") {
			nameLong = strings.TrimPrefix(token, "--")
		} else if strings.HasPrefix(token, "-") {
        	nameShort = strings.TrimPrefix(token, "-")
		} else {
			positionals = append(positionals, token)
			continue
		}
		var found definition
		for _, d := range defs {
			for _, name := range d.meta().names {
				if name != "" && len(name) > 1 && name == nameLong {
					found = d
					break
				} 
				if name != "" && len(name) == 1 && name == nameShort {
					found = d
					break
				}
			}
   			if found != nil {
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
