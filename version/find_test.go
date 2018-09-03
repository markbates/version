package version

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Find(t *testing.T) {
	table := []struct {
		name string
		in   string
		v    string
	}{
		{name: "constCap", in: constCap, v: "v1.0.3"},
		{name: "constLow", in: constLow, v: "1.2.3-beta.1"},
		{name: "varCap", in: varCap, v: "v2.3.4"},
		{name: "varLow", in: varLow, v: "4.5.6"},
	}

	for _, tt := range table {
		t.Run(tt.name, func(st *testing.T) {
			r := require.New(st)
			in := strings.NewReader(tt.in)
			v, err := Find(in, false)
			r.NoError(err)
			r.Equal(tt.v, v)
		})
	}
}

func Test_Find_AllowDev(t *testing.T) {
	table := []struct {
		name string
		in   string
		v    string
	}{
		{name: "constCap", in: constCap, v: "v1.0.3"},
		{name: "constLow", in: constLow, v: "1.2.3-beta.1"},
		{name: "varCap", in: varCap, v: "v2.3.4"},
		{name: "varLow", in: varLow, v: "4.5.6"},
	}

	for _, tt := range table {
		t.Run(tt.name, func(st *testing.T) {
			r := require.New(st)
			in := strings.NewReader(strings.Replace(tt.in, tt.v, "development", -1))
			v, err := Find(in, true)
			r.NoError(err)
			r.Equal("development", v)
		})
	}
}

func Test_Find_Bad(t *testing.T) {
	table := []string{
		`const Version = "development"`,
		`var version = "development"`,
		`var MyVersion = "v1.0.0"`,
		``,
		`var Version = "v!dx()"`,
	}

	for _, tt := range table {
		t.Run(tt, func(st *testing.T) {
			r := require.New(st)
			in := strings.NewReader(tt)
			_, err := Find(in, false)
			r.Error(err)
		})
	}
}

var constCap = `package runtime

// Version is the current version of the buffalo binary
// const Version = "v0.12.6"

// Version is the current version of the buffalo binary
const Version = "v1.0.3"`

var constLow = `package runtime

// Version is the current version of the buffalo binary
// const Version = "v0.12.6"

// Version is the current version of the buffalo binary
const version = "1.2.3-beta.1"`

var varCap = `package runtime

// Version is the current version of the buffalo binary
// const Version = "v0.12.6"

// Version is the current version of the buffalo binary
var Version = "v2.3.4"`

var varLow = `package runtime

// Version is the current version of the buffalo binary
// const Version = "v0.12.6"

// Version is the current version of the buffalo binary
var version = "4.5.6"`
