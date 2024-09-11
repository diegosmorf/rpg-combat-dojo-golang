package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DiceTestOption struct {
	Value  int
	Result int
}

var validOptions = []DiceTestOption{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 4},
	{5, 5},
	{6, 6},
}

func Test_CreateDiceWithOnlyValue_Then_GetValidResult(t *testing.T) {

	for i := 0; i < len(validOptions); i++ {
		// arrange
		value := validOptions[i].Value
		expectedResult := validOptions[i].Result
		dice := New(BuildOnlyValue(value))
		//act
		currentResult := dice.Roll()
		//assert
		assert.Equal(t, expectedResult, currentResult)
	}
}

func Test_CreateDiceMultipleSameValues_Then_GetValidResult(t *testing.T) {

	for i := 0; i < len(validOptions); i++ {
		// arrange
		value := validOptions[i].Value
		expectedResult := validOptions[i].Result
		dice := New(BuildMultipleSameValues(value))
		//act
		currentResult := dice.Roll()
		//assert
		assert.Equal(t, expectedResult, currentResult)
	}
}

func Test_CreateDiceMultipleMinMax_Then_GetValidResult(t *testing.T) {

	for i := 0; i < len(validOptions); i++ {
		// arrange
		value := validOptions[i].Value
		expectedResult := validOptions[i].Result
		dice := New(BuildMultipleMinMax(value, value))
		//act
		currentResult := dice.Roll()
		//assert
		assert.Equal(t, expectedResult, currentResult)
	}
}

func Test_CreateDiceDefault_Then_GetValidResult(t *testing.T) {

	// arrange
	expectedMin := 1
	expectedMax := 6
	dice := New(BuildDefault())
	//act
	currentResult := dice.Roll()
	//assert
	assert.LessOrEqual(t, expectedMin, currentResult)
	assert.GreaterOrEqual(t, expectedMax, currentResult)
}

func Test_CreateDiceNil_Then_GetValidResult(t *testing.T) {

	// arrange
	expectedMin := 1
	expectedMax := 6
	dice := New()
	//act
	currentResult := dice.Roll()
	//assert
	assert.LessOrEqual(t, expectedMin, currentResult)
	assert.GreaterOrEqual(t, expectedMax, currentResult)

}

func Test_CreateDiceMagic_Then_GetValidResult(t *testing.T) {

	// arrange
	expectedMin := 1
	expectedMax := 8
	dice := New(BuildDefaultMagic())
	//act
	currentResult := dice.Roll()
	//assert
	assert.LessOrEqual(t, expectedMin, currentResult)
	assert.GreaterOrEqual(t, expectedMax, currentResult)

}
