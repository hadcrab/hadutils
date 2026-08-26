package cli

import (
	"slices"
	"strings"
	"errors"
)

func Parse(argv []string, defs ...*Argument) ([]string, error) {
	var positionals []string
	for i := range defs {
		defs[i].seen = false
    	if defs[i].hasDefault {
        	if err := defs[i].value.Set(defs[i].deflt); err != nil {
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
			if slices.Contains(defs[argIndex].names, name) {
					found = defs[argIndex]
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
		if !defs[i].seen && defs[i].required {
			return nil, errors.New("required argument: " + defs[i].names[0])
		}
	}
	return positionals, nil
}
