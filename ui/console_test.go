package ui

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Console(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "twoMowers",
			args: args{
				fileName: "../resources/testdata/input.txt",
			},
			want: "1 3 N\n5 1 E\n",
		},
		{
			name: "Three mowers, third collides with second",
			args: args{
				fileName: "../resources/testdata/input_three_mowers.txt",
			},
			want: "2 4 N\n5 0 W\n1 5 N\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := bytes.Buffer{}
			console := Console{}
			console.ConsoleExit(&out, tt.args.fileName)
			assert.Equal(t, tt.want, out.String())
		})
	}
}
