package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

func Test_parseArgs_new(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want Op
	}{
		{name: "nil Args",
			args: nil,
			want: ListOp{}},
		{name: "empty Args",
			args: []string{},
			want: ListOp{}},
		{name: "help shorthand",
			args: []string{"-h"},
			want: HelpOp{}},
		{name: "help long form",
			args: []string{"--help"},
			want: HelpOp{}},
		{name: "current shorthand",
			args: []string{"-c"},
			want: CurrentOp{}},
		{name: "current long form",
			args: []string{"--current"},
			want: CurrentOp{}},
		{name: "unset shorthand",
			args: []string{"-u"},
			want: UnsetOp{}},
		{name: "unset long form",
			args: []string{"--unset"},
			want: UnsetOp{}},
		{name: "switch by name",
			args: []string{"foo"},
			want: SwitchOp{Target: "foo"}},
		{name: "switch by swap",
			args: []string{"-"},
			want: SwitchOp{Target: "-"}},
		{name: "delete - without contexts",
			args: []string{"-d"},
			want: DeleteOp{[]string{}}},
		{name: "delete - current context",
			args: []string{"-d", "."},
			want: DeleteOp{[]string{"."}}},
		{name: "delete - multiple contexts",
			args: []string{"-d", ".", "a", "b"},
			want: DeleteOp{[]string{".", "a", "b"}}},
		{name: "rename context",
			args: []string{"a=b"},
			want: RenameOp{"a", "b"}},
		{name: "rename context with old=current",
			args: []string{"a=."},
			want: RenameOp{"a", "."}},
		{name: "unrecognized flag",
			args: []string{"-x"},
			want: UnsupportedOp{Err: errors.Errorf("unsupported option \"-x\"")}},
		// TODO add more UnsupportedOp cases

		// TODO consider these cases
		// - kubectx foo --help
		// - kubectx -h --help
		// - kubectx -d foo --h
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseArgs(tt.args)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("parseArgs(%#v) diff: %s", tt.args, diff)
			}
		})
	}
}
