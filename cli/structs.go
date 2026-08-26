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

type Argument struct {
	names       []string
	description string
	required    bool

	deflt      string
	hasDefault bool
	seen bool
	
	value      Value
	behavior   Behavior
}

func String(names ...string) *Argument {
	return argument(&StringValue{}, &Option{}, names...)
}

func Int(names ...string) *Argument {
	return argument(&IntValue{}, &Option{}, names...)
}

func Bool(names ...string) *Argument {
	return argument(&BoolValue{}, &Flag{}, names...)
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

func (a *Argument) StringValue() string {
	return a.value.(*StringValue).Value
}

func (a *Argument) IntValue() int {
	return a.value.(*IntValue).Value
}

func (a *Argument) BoolValue() bool {
	return a.value.(*BoolValue).Value
}
