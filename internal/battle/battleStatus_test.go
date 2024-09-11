package battle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validTypes = []struct {
	status BattleStatus
	name   string
	index  int
}{
	{Finished, "Finished", 0},
	{Draw, "Draw", 1},
	{Unfinished, "Unfinished", 2},
}

func Test_WhenUseValidActions_Then_GetIndexNameSuccessfully(t *testing.T) {

	for i := 0; i < len(validTypes); i++ {
		// arrange
		expectedType := validTypes[i].status
		expectedName := validTypes[i].name
		expectedIndex := validTypes[i].index

		//act
		index := expectedType.EnumIndex()
		currentName := expectedType.String()

		//assert
		assert.Equal(t, expectedIndex, index)
		assert.Equal(t, expectedName, currentName)
	}
}
