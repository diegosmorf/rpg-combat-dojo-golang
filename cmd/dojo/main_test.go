package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenCreateNewGame_Then_Success(t *testing.T) {

	// arrange
	expectedValue := true
	currentValue := false

	//act
	currentValue = CreateNewGame()

	//assert
	if expectedValue != currentValue {
		t.Errorf("Expected: %t, Got: %t", expectedValue, currentValue)
	}

	assert.Equal(t, expectedValue, currentValue)
}

func Test_WhenRunMain_Then_Success(t *testing.T) {

	//act
	main()
}
