package hw8

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}

func TestSingleError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))

	expectedMessage := "1 errors occured:\n\t* error 1\n"
	assert.EqualError(t, err, expectedMessage)
}

func TestNilError(t *testing.T) {
	err := &MultiError{}
	expectedMessage := ""
	assert.EqualError(t, err, expectedMessage)
}

func TestNilErrorAndEmptyAppend(t *testing.T) {
	err := &MultiError{}
	err = Append(err)
	expectedMessage := ""
	assert.EqualError(t, err, expectedMessage)
}

func TestFirstMultiError(t *testing.T) {
	var err error
	errMul := &MultiError{errors: []error{errors.New("error 1")}}
	err = Append(err, errMul)

	expectedMessage := "1 errors occured:\n\t* error 1\n"
	assert.EqualError(t, err, expectedMessage)
}

func TestMultiMultiError(t *testing.T) {
	var err error
	errMul := &MultiError{errors: []error{errors.New("error 2")}}
	err = Append(err, errors.New("error 1"))
	err = Append(err, errMul)

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
