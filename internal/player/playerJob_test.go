package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validJobs = []struct {
	job   PlayerJob
	name  string
	index int
}{
	{Soldier, "Soldier", 0},
	{Wizard, "Wizard", 1},
	{Archer, "Archer", 2},
	{Knight, "Knight", 3},
}

func Test_WhenUseValidJob_Then_GetIndexNameSuccessfully(t *testing.T) {

	for i := 0; i < len(validJobs); i++ {
		// arrange
		expectedType := validJobs[i].job
		expectedName := validJobs[i].name
		expectedIndex := validJobs[i].index
		//act
		index := expectedType.EnumIndex()
		currentName := expectedType.String()
		//assert
		assert.Equal(t, expectedIndex, index)
		assert.Equal(t, expectedName, currentName)
	}
}
