package cli

import (
	"errors"
	"strconv"
)

func parseInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func parseBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

func parseString(value string) (string, error) {
	return value, nil
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

type Definition struct {
	names       []string
	description string
	required    bool
	hasDefault  bool
	seen        bool
	behavior    Behavior

	reset func()
	take  func(tokens []string) (int, error)
}

type Argument[T any] struct {
	definition Definition
	deflt      T
	value      T
	parseFn    func(string) (T, error)
}

func String(names ...string) *Argument[string] {
	return argument("", &Option{}, parseString, names...)
}

func Int(names ...string) *Argument[int] {
	return argument(0, &Option{}, parseInt, names...)
}

func Bool(names ...string) *Argument[bool] {
	return argument(false, &Flag{}, parseBool, names...)
}

func argument[T any](deflt T, behavior Behavior, parseFn func(string) (T, error), names ...string) *Argument[T] {
	a := &Argument[T]{
		definition: Definition{
			names:    names,
			behavior: behavior,
		},
		deflt:   deflt,
		parseFn: parseFn,
	}
	a.value = deflt
	a.definition.reset = func() {
		a.value = a.deflt
		a.definition.seen = false
	}
	a.definition.take = func(tokens []string) (int, error) {
		return behavior.Parse(tokens, a.takeValue)
	}
	return a
}

func (a *Argument[T]) takeValue(val string) error {
	parsed, err := a.parseFn(val)
	if err != nil {
		return err
	}
	a.value = parsed
	return nil
}

func (a *Argument[T]) Description(description string) *Argument[T] {
	a.definition.description = description
	return a
}

func (a *Argument[T]) Default(value T) *Argument[T] {
	a.deflt = value
	a.definition.hasDefault = true
	return a
}

func (a *Argument[T]) Required() *Argument[T] {
	a.definition.required = true
	return a
}

func (a *Argument[T]) Value() T {
	return a.value
}

func (a *Argument[T]) meta() *Definition {
	return &a.definition
}
