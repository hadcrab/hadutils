package main

import (
	"os"
	"strconv"
	"strings"
)

type Value interface {
    String() string
    Set(string) error
}

type IntValue struct {
	Value int
}

func (v *IntValue) String() string {
    return strconv.Itoa(v.Value)
}

func (v *IntValue) Set(value string) error {
    parsed, err := strconv.Atoi(value)
    if err != nil {
        return err
    }

    v.Value = parsed
    return nil
}

type BoolValue struct {
    Value bool
}

func (v *BoolValue) String() string {
    return strconv.FormatBool(v.Value)
}

func (v *BoolValue) Set(value string) error {
    parsed, err := strconv.ParseBool(value)
    if err != nil {
        return err
    }

    v.Value = parsed
    return nil
}

type StringValue struct {
    Value string
}

func (v *StringValue) String() string {
    return v.Value
}

func (v *StringValue) Set(value string) error {
    v.Value = value
    return nil
}

type Argument struct {
	Names      []string
	Description string
	Required bool
	Value Value
}

func DefineArgument(names []string, desc string, isReq bool, defValue Value) Argument {
	return Argument{
		Names: names,
		Description: desc,
		Required: isReq,
		Value: defValue,
	}
}

func ParseArgs(arguments []Argument) ([]string, error) {
	args := os.Args[1:]
	for argIndex, arg := range args {
		if strings.HasPrefix(arg, "--") {
			arg = strings.TrimPrefix(strings.TrimPrefix(arg, "--"), "-")
			for definitionIndex := range arguments {
				for _, argName := range arguments[definitionIndex].Names {
					if argName == arg {
						if argIndex+1 < len(args) {
							arguments[definitionIndex].Value.Set(args[argIndex+1])
						}  
					}
				}
			}
		}
	}
}

func main() {
	algorithm := DefineArgument(
    []string{"algorithm", "a"},
    "Hash algorithm",
    false,
    &StringValue{Value: "sha256"},
	)
	
}