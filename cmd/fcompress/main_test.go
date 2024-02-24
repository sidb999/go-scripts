package main

import (
	"os"
	"testing"
)

func TestParseArguments(t *testing.T) {
	cases := []struct {
		desc string
		want VideoFile
		args []string
	}{
		{"nameless args", VideoFile{src: "src.mp4", out: ""}, []string{"cmd", "src.mp4"}},
		{"src test", VideoFile{src: "src.mp4", out: ""}, []string{"cmd", "-s", "src.mp4"}},
		{"named arg out", VideoFile{src: "", out: "out.mp4"}, []string{"cmd", "-o", "out.mp4"}},
		{"both named args", VideoFile{src: "src.mp4", out: "out.mp4"}, []string{"cmd", "-s", "src.mp4", "-o", "out.mp4"}},
	}
	for _, test := range cases {
		t.Run(test.desc, func(t *testing.T) {
			var got VideoFile

			os.Args = test.args
			got = parseArguments()
			if got.out != test.want.out {
				t.Errorf("'%v' results: got %v, want %v", test.desc, got, test.want)
			}
		})
	}
}
