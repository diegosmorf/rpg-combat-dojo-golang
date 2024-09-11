package utils

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenGetSubString_Then_Success(t *testing.T) {

	var validOptions = []struct {
		Value          string
		ExpectedResult string
		Start          int
		Length         int
	}{
		{"", "", 0, 1},
		{"Player", "P", 0, 1},
		{"Player", "Pla", 0, 3},
		{"Player", "Player", 0, 10},
		{"Player", "", 9, 10},
		{"Player", "er", 4, 100},
	}

	for i := 0; i < len(validOptions); i++ {
		// arrange
		value := validOptions[i].Value
		start := validOptions[i].Start
		length := validOptions[i].Length
		expectedResult := validOptions[i].ExpectedResult

		//act
		currentResult := GetSubString(value, start, length)

		//assert
		assert.NotNil(t, currentResult)
		assert.Equal(t, expectedResult, currentResult)
	}
}

func Test_WhenGetLeftString_Then_Success(t *testing.T) {

	var validOptions = []struct {
		Value          string
		ExpectedResult string
		Length         int
	}{
		{"", "", 1},
		{"Player", "P", 1},
		{"Player", "Pla", 3},
		{"Player", "", 0},
		{"Player", "Player", 10},
		{"Player", "Player", 100},
	}

	for i := 0; i < len(validOptions); i++ {
		// arrange
		value := validOptions[i].Value
		length := validOptions[i].Length
		expectedResult := validOptions[i].ExpectedResult

		//act
		currentResult := GetLeftString(value, length)

		//assert
		assert.NotNil(t, currentResult)
		assert.Equal(t, expectedResult, currentResult)
	}
}

func Test_WhenGetUniqueId_Then_Success(t *testing.T) {

	uniqueList := []string{}
	for i := 0; i < 100; i++ {
		//arrange
		expectedLength := 5
		//act
		currentResult := GetUniqueId()
		exists := slices.Contains(uniqueList, currentResult)
		//assert
		assert.NotNil(t, currentResult)
		assert.Equal(t, len(currentResult), expectedLength)
		assert.False(t, exists)
		uniqueList = append(uniqueList, currentResult)
	}
}
