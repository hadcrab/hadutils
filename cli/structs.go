package main

import (
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
	Parse(args []string) (consumed int, err error)
}

type Flag struct {
	Value bool
}

func (f *Flag) Parse(args []string) (int, error) {
	f.Value = true
	return 0, nil
}

type Option struct {
	Value Value
}

func (o *Option) Parse(args []string) (int, error) {
	if len(args) == 0 {
		return 0, errors.New("option requires a value")
	}
	if err := o.Value.Set(args[0]); err != nil {
		return 0, err
	}
	return 1, nil
}

type Argument struct {
	Names       []string
	Description string
	Required    bool
	Behavior    Behavior
}