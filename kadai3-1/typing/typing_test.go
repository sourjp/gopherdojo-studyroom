package typing_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/sourjp/gopherdojo-studyroom/kadai3-1/typing"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  []string
	}{
		{name: "lowercase", in: "../testdata/test.csv", out: []string{"gopher", "kawaii", "great"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := typing.Setup(test.in)
			if err != nil {
				t.Errorf("Setup(%s) got error: %s", test.in, err)
			}
			for i := range out {
				if out[i] != test.out[i] {
					t.Errorf("Setup(%s) got: %s, expect: %s", test.in, out[i], test.out[i])
				}
			}
		})
	}
}

func TestStart(t *testing.T) {
	tests := []struct {
		name string
		time time.Duration
		quiz []string
		out  int
	}{
		{name: "Test time.After()", time: 1, quiz: []string{"gopher", "kawaii"}, out: 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := typing.Start(test.time, test.quiz)
			if out != test.out {
				t.Errorf("Start(%s, %s) got: %d, expect: %d", test.time, test.quiz, out, test.out)
			}
		})
	}
}

func TestScanner(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{name: "lowercase", in: "gopher", out: "gopher"},
		{name: "Uppercase", in: "GOPHER", out: "GOPHER"},
		{name: "inculude space", in: "Go Pher", out: "Go Pher"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ch := typing.Scanner(bytes.NewBufferString(test.in))
			if ans := <-ch; ans != test.out {
				t.Errorf("Scanner(%s) got: %s, expect: %s", test.in, ans, test.out)
			}
		})
	}
}
