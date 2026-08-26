package cli

import (
	"errors"
	"strconv"
)

type Value interface {
	String() string
	Set(string) error
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

type Behavior interface {
	Parse(args []string, set func(string) error) (consumed int, err error)
}

type Flag struct{}

func (f Flag) Parse(_ []string, set func(string) error) (int, error) {
	return 0, set("true")
}

type Option struct{}

func (o Option) Parse(args []string, set func(string) error) (int, error) {
	if len(args) == 0 {
		return 0, errors.New("requires a value")
	}
	return 1, set(args[0])
}

type Definition interface {
	argument() *Argument
}

type StringArgument struct {
	*Argument
}

type BoolArgument struct {
	*Argument
}

type IntArgument struct {
	*Argument
}

func (v *BoolArgument) argument() *Argument {
	return v.Argument
}

func (v *IntArgument) argument() *Argument {
	return v.Argument
}

func (v *StringArgument) argument() *Argument {
	return v.Argument
} 

type Argument struct {
	names       []string
	description string
	required    bool

	deflt      string
	hasDefault bool
	seen       bool
	
	value      Value
	behavior   Behavior
}

func String(names ...string) *StringArgument {
	arg := argument(&StringValue{}, &Option{}, names...)
	return &StringArgument{
		Argument: arg,
	}
}

func Int(names ...string) *IntArgument {
	arg := argument(&IntValue{}, &Option{}, names...)
	return &IntArgument{
		Argument: arg,
	}
}

func Bool(names ...string) *BoolArgument {
	arg := argument(&BoolValue{}, &Option{}, names...)
	return &BoolArgument{
		Argument: arg,
	}
}

func argument(value Value, behavior Behavior, names ...string) *Argument {
	return &Argument{
		names:    names,
		value:    value,
		behavior: behavior,
	}
}

func (a *Argument) TakeValue(val string) error {
	return a.value.Set(val)
}

func (a *Argument) Description(description string) *Argument {
	a.description = description
	return a
}

func (a *Argument) Default(value string) *Argument {
	a.deflt = value
	a.hasDefault = true
	return a
}

func (a *Argument) Required() *Argument {
	a.required = true
	return a
}

func (a *StringArgument) Value() string {
	return a.value.(*StringValue).Value
}

func (a *IntArgument) Value() int {
	return a.value.(*IntValue).Value
}

func (a *BoolArgument) Value() bool {
	return a.value.(*BoolValue).Value
}
