package scim_filtering

import (
	"reflect"
	"strings"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	var tests = []struct {
		s    string
		stmt *Statement
		err  string
	}{
		// eq operator
		{
			s: `userName Eq "john"`,
			stmt: &Statement{
				Name:     "username",
				Operator: EQ,
				Value:    "john",
			},
		},

		{
			s: `Username eq "john"`,
			stmt: &Statement{
				Name:     "username",
				Operator: EQ,
				Value:    "john",
			},
		},
		{
			s: `name.formatted eq "john doe"`,
			stmt: &Statement{
				Name:     "name.formatted",
				Operator: EQ,
				Value:    "john doe",
			},
		},

		// other operators
		{
			s: `username ne "john"`,
			stmt: &Statement{
				Name:     "username",
				Operator: NE,
				Value:    "john",
			},
		},
		{
			s: `name.familyName co "doe"`,
			stmt: &Statement{
				Name:     "name.familyname",
				Operator: CO,
				Value:    "doe",
			},
		},
		{
			s: `urn:ietf:params:scim:schemas:core:2.0:User:userName sw "j"`,
			stmt: &Statement{
				Name:     "username",
				Operator: SW,
				Value:    "j",
			},
		},
		{
			s: `username ew "n"`,
			stmt: &Statement{
				Name:     "username",
				Operator: EW,
				Value:    "n",
			},
		},
		{
			s: `title pr`,
			stmt: &Statement{
				Name:     `title`,
				Operator: PR,
			},
		},

		// empty value
		{
			s: `Username eq`,
			stmt: &Statement{
				Name:     "username",
				Operator: EQ,
			},
		},

		{s: `error x "value"`, err: `found "x", expected operator`},
	}

	for i, test := range tests {
		stmt, err := NewParser(strings.NewReader(test.s)).Parse()
		if !reflect.DeepEqual(test.err, errToString(err)) {
			t.Errorf("%d. %q: wrong error:\n  exp=%s\n  got=%s\n\n", i, test.s, test.err, err)
		} else if test.err == "" && !reflect.DeepEqual(test.stmt, stmt) {
			t.Errorf("%d. %q\n\nwrong stmt:\n\nexp=%#v\n\ngot=%#v\n\n", i, test.s, test.stmt, stmt)
		}
	}
}

func errToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
