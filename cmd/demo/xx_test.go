package main

import "testing"

func TestXx(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "xx", want: "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Xx(tt.args.x); got != tt.want {
				t.Errorf("Xx() = %v, want %v", got, tt.want)
			}
		})
	}
}
