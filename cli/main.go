package main

import (
	"slices"
	"errors"
	"strings"
)

func (o *Option) Parse(args []string) (int, error) {
	if len(args) == 0 {
		return 0, errors.New("option requires a value")
	}
	if err := o.Value.Set(args[0]); err != nil {
		return 0, err
	}
	return 1, nil
}

func ParseArgs(input []string, arguments []Argument) ([]string, error) {
	var positionals []string
	for i := 0; i < len(input); i++ {
		token := input[i]
		if !strings.HasPrefix(token, "-") {
			positionals = append(positionals, token)
			continue
		}
		name := strings.TrimLeft(token, "-")
		var found *Argument
		for argIndex := range arguments {
			if slices.Contains(arguments[argIndex].Names, name) {
					found = &arguments[argIndex]
				}
			if found != nil {
				break
			}
		}
		if found == nil {
			return nil, errors.New("unknown argument: " + token)
		}
		consumed, err := found.Behavior.Parse(input[i+1:])
		if err != nil {
			return nil, err
		}
		i += consumed
	}
	return positionals, nil
}

func main() {
	
}