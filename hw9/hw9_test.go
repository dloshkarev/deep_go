package hw9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialization(t *testing.T) {
	tests := map[string]struct {
		person     Person
		result     string
		maybeError string
	}{
		"test case with empty fields": {
			result: "name=\nage=0\nmarried=false",
		},
		"test case with fields": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
			},
			result: "name=John Doe\nage=30\nmarried=true",
		},
		"test case with omitempty field": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := Serialize(test.person)
			if test.maybeError != "" {
				assert.EqualError(t, err, test.maybeError)
			}
			assert.Equal(t, test.result, result)
		})
	}
}

type BrokenPerson struct {
	Name string
}

func TestNoProperties(t *testing.T) {
	tests := map[string]struct {
		person     BrokenPerson
		result     string
		maybeError string
	}{
		"test case with empty fields": {
			maybeError: "tag `properties` not found",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := Serialize(test.person)
			if test.maybeError != "" {
				assert.EqualError(t, err, test.maybeError)
			}
			assert.Equal(t, test.result, result)
		})
	}
}

type BrokenPerson2 struct {
	Name string `properties:""`
}

func TestEmptyProperties(t *testing.T) {
	tests := map[string]struct {
		person     BrokenPerson2
		result     string
		maybeError string
	}{
		"test case with empty fields": {
			maybeError: "tag `properties` is empty",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := Serialize(test.person)
			if test.maybeError != "" {
				assert.EqualError(t, err, test.maybeError)
			}
			assert.Equal(t, test.result, result)
		})
	}
}
