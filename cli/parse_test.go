package cli

import (
	"slices"
	"strings"
	"testing"
)

// newDefs returns a FRESH set of definitions for each test case.
//
// Why fresh: a Definition is stateful (value, seen). Parse calls reset()
// at the start, but if we reused the same arguments across cases, a value
// from a previous case could leak into the next one. A per-case factory
// rules out cross-case interference by construction.
func newDefs() []definition {
	algo := String("algorithm", "a").Default("SHA256")
	quiet := Bool("quiet", "q")
	copy := Bool("copy", "c")
	return []definition{algo, quiet, copy}
}

// parseCase is one row of the table: input, expectations, and optionally
// its own def factory when the standard set is not enough.
type parseCase struct {
	name    string
	argv    []string
	build   func() []definition // nil → newDefs
	wantPos []string
	wantErr string // expected error substring; "" means no error expected
}

func TestParse(t *testing.T) {
	cases := []parseCase{
		{
			name:    "empty argv",
			argv:    nil,
			wantPos: nil,
		},
		{
			name:    "only positionals",
			argv:    []string{"a", "b"},
			wantPos: []string{"a", "b"},
		},
		{
			name:    "short flag",
			argv:    []string{"-q"},
			wantPos: nil,
		},
		{
			name:    "long flag",
			argv:    []string{"--quiet"},
			wantPos: nil,
		},
		{
			name:    "long option consumes its value",
			argv:    []string{"--algorithm", "sha512"},
			wantPos: nil,
		},
		{
			name:    "value is not left as positional",
			argv:    []string{"--algorithm", "sha512", "file.txt"},
			wantPos: []string{"file.txt"},
		},
		{
			name:    "short option missing value",
			argv:    []string{"-a"},
			wantErr: "algorithm: requires a value",
		},
		{
			name:    "long option missing value",
			argv:    []string{"--algorithm"},
			wantErr: "algorithm: requires a value",
		},
		{
			name:    "cluster of flags",
			argv:    []string{"-cq"},
			wantPos: nil,
		},
		{
			name:    "bare dash stays positional",
			argv:    []string{"-"},
			wantPos: []string{"-"},
		},
		{
			name:    "dash mixed with files",
			argv:    []string{"-", "file.txt"},
			wantPos: []string{"-", "file.txt"},
		},
		{
			name:    "double dash separator makes rest positional",
			argv:    []string{"--quiet", "--", "x", "-y"},
			wantPos: []string{"x", "-y"},
		},
		{
			name:    "unknown long argument",
			argv:    []string{"--nope"},
			wantErr: "unknown argument: --nope",
		},
		{
			name:    "unknown short argument",
			argv:    []string{"-z"},
			wantErr: "unknown argument: -z",
		},
		{
			name:    "option inside cluster is rejected",
			argv:    []string{"-qa"}, // q is a flag, a is an option with a value
			wantErr: "-a, algorithm takes a value and cannot be part of a cluster",
		},
		{
			name: "required argument missing",
			argv: nil,
			build: func() []definition {
				return []definition{Int("size").Required()}
			},
			wantErr: "required argument missing: size",
		},
		{
			name: "required with default is rejected",
			argv: nil,
			build: func() []definition {
				// the conflict is checked before any argument is parsed
				return []definition{String("name").Default("x").Required()}
			},
			wantErr: "argument cannot be required if has default value: name",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			build := tc.build
			if build == nil {
				build = newDefs
			}

			gotPos, err := Parse(tc.argv, build()...)

			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("Parse(%v) unexpected error: %v", tc.argv, err)
				}
			} else {
				if err == nil {
					t.Fatalf("Parse(%v): expected error containing %q, got nil", tc.argv, tc.wantErr)
				}
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("Parse(%v) error = %q, want substring %q", tc.argv, err, tc.wantErr)
				}
			}

			if !slices.Equal(gotPos, tc.wantPos) {
				t.Errorf("Parse(%v) positionals = %v, want %v", tc.argv, gotPos, tc.wantPos)
			}
		})
	}
}

// TestParseSetsValues checks argument state AFTER parsing: value picked up
// the passed-in value, and seen is set for the arguments that were met.
func TestParseSetsValues(t *testing.T) {
	algo := String("algorithm", "a")
	quiet := Bool("quiet", "q")

	_, err := Parse([]string{"--algorithm", "sha512", "-q"}, algo, quiet)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}

	if got := algo.Value(); got != "sha512" {
		t.Errorf("algo.Value() = %q, want %q", got, "sha512")
	}
	if got := quiet.Value(); got != true {
		t.Errorf("quiet.Value() = %v, want true", got)
	}
}

func TestFormatNames(t *testing.T) {
	cases := []struct {
		names []string
		want  string
	}{
		{names: []string{"quiet", "q"}, want: "-q, --quiet"},
		{names: []string{"a", "algorithm"}, want: "-a, --algorithm"},
		{names: []string{"noShort"}, want: "--noShort"},
		{names: []string{"s"}, want: "-s"},
		{names: nil, want: "    "}, // placeholder for a def with no names at all
	}

	for _, tc := range cases {
		got := formatNames(tc.names)
		if got != tc.want {
			t.Errorf("formatNames(%v) = %q, want %q", tc.names, got, tc.want)
		}
	}
}

func TestLookupDefinitionByName(t *testing.T) {
	defs := newDefs()

	// a short name is found by the short lookup
	if got := lookupDefinitionByName("q", true, defs...); got == nil {
		t.Error("lookupDefinitionByName(q, short) = nil, want definition")
	}

	// a long name is not found by the short lookup
	if got := lookupDefinitionByName("algorithm", true, defs...); got != nil {
		t.Errorf("lookupDefinitionByName(algorithm, short) = %v, want nil", got)
	}

	// a long name is found by the long lookup
	if got := lookupDefinitionByName("algorithm", false, defs...); got == nil {
		t.Error("lookupDefinitionByName(algorithm, long) = nil, want definition")
	}

	// an unknown name is not found by either lookup
	if got := lookupDefinitionByName("nope", false, defs...); got != nil {
		t.Errorf("lookupDefinitionByName(nope, long) = %v, want nil", got)
	}
}
