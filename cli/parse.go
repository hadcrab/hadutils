package cli

import (
	"slices"
	"strings"
)

func Parse(argv []string, defs ...Definition) ([]string, error) {
	var positionals []string
	for i := range defs {
		arg := defs[i].argument()
		arg.seen = false
    	if arg.hasDefault {
        	if err := arg.value.Set(arg.deflt); err != nil {
            	return nil, err
         	}
     	}
	}
	for i := 0; i < len(argv); i++ {
		token := argv[i]
		if !strings.HasPrefix(token, "-") {
			positionals = append(positionals, token)
			continue
		}
		name := strings.TrimLeft(token, "-")
		var found *Argument
		for argIndex := range defs {
			arg := defs[argIndex].argument()
			if slices.Contains(arg.names, name) {
					found = arg
				}
			if found != nil {
				break
			}
		}
		if found == nil {
			return nil, errors.New("unknown argument: " + token)
		}
		consumed, err := found.behavior.Parse(argv[i+1:], found.TakeValue)
		if err != nil {
			return nil, err
		}
		found.seen = true
		i += consumed
	}
	for i := range defs {
		arg := defs[i].argument()
		if !arg.seen && arg.required {
			return nil, errors.New("required argument: " + arg.names[0])
		}
	}
	return positionals, nil
}
