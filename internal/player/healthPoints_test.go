package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenApplyDamage_Then_ReduceHealth(t *testing.T) {

	var inputTestList = []struct {
		expectedDamage int
		expectedHealth int
	}{
		{0, 500},
		{50, 450},
		{250, 250},
		{499, 1},
		{500, 0},
		{510, 0},
	}

	const max = 500
	const min = 0

	for i := 0; i < len(inputTestList); i++ {
		// arrange
		expectedDamage := inputTestList[i].expectedDamage
		expectedHealth := inputTestList[i].expectedHealth
		//act
		health := BuildHealth(max)
		health.ApplyDamage(expectedDamage)
		//assert
		assert.Equal(t, expectedHealth, health.Current)
		assert.Equal(t, max, health.Max)
		assert.Equal(t, min, health.Min)
	}
}

func Test_WhenApplyReset_Then_CurrentEqualsMax(t *testing.T) {

	var inputTestList = []int{
		1,
		50,
		250,
		499,
		500,
		510,
	}

	const min = 0
	const current = 1

	for i := 0; i < len(inputTestList); i++ {
		// arrange
		max := inputTestList[i]
		//act
		health := BuildHealthAllFields(current, min, max)
		health.Reset()
		//assert
		assert.Equal(t, max, health.Current)
		assert.Equal(t, max, health.Max)
		assert.Equal(t, min, health.Min)
	}
}

func Test_WhenApplyLevelUp_Then_IncreaseMaxAndCurrent(t *testing.T) {

	var inputTestList = []int{
		0,
		50,
		250,
		499,
		500,
		510,
	}

	const max = 500
	const min = 0
	const increase = 100

	for i := 0; i < len(inputTestList); i++ {
		// arrange
		expectedDamage := inputTestList[i]
		//act
		health := BuildHealth(max)
		health.ApplyDamage(expectedDamage)
		health.LevelUp(increase)
		//assert
		assert.Equal(t, max+increase, health.Current)
		assert.Equal(t, max+increase, health.Max)
		assert.Equal(t, min, health.Min)
	}
}

func Test_WhenApplyHeal_Then_IncreaseHealth(t *testing.T) {

	var inputTestList = []struct {
		heal           int
		expectedHealth int
	}{
		{0, 1},
		{50, 51},
		{250, 251},
		{499, 500},
		{500, 500},
		{510, 500},
	}

	const min = 0
	const max = 500
	const current = 1

	for i := 0; i < len(inputTestList); i++ {
		// arrange
		heal := inputTestList[i].heal
		expectedHealth := inputTestList[i].expectedHealth
		//act
		health := BuildHealthAllFields(current, min, max)
		health.ApplyHeal(heal)
		//assert
		assert.Equal(t, expectedHealth, health.Current)
		assert.Equal(t, max, health.Max)
		assert.Equal(t, min, health.Min)
	}
}
