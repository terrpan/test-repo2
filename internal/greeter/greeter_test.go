package greeter

import (
	"errors"
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr error
	}{
		{name: "named", in: "Scrapyard", want: "Hello, Scrapyard!"},
		{name: "blank", in: "", want: "Hello, world!"},
		{name: "whitespace trimmed", in: "  Ada  ", want: "Hello, Ada!"},
		{name: "too long", in: strings.Repeat("a", MaxNameLen+1), wantErr: ErrNameTooLong},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Greet(tt.in)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
