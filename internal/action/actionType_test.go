package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validTypes = []struct {
	actionType ActionType
	name       string
	index      int
}{
	{Attack, "Attack", 0},
	{FireBallMagic, "FireBallMagic", 1},
	{HealMagic, "HealMagic", 2},
	{Item, "Item", 3},
}

func Test_WhenUseValidActions_Then_GetIndexNameSuccessfully(t *testing.T) {

	for i := 0; i < len(validTypes); i++ {
		// arrange
		expectedType := validTypes[i].actionType
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
