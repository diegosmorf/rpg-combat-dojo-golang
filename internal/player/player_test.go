package player

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validPlayerNames = []string{
	"Player-1",
	"Player-100",
	"Player-10000",
	"Player-1234567890",
	"Player-123456789012",
	"Player-1234567890123",
}

func Test_WhenCreateValidSoldier_Then_Success(t *testing.T) {

	for i := 0; i < len(validPlayerNames); i++ {
		// arrange
		expectedName := validPlayerNames[i]
		expectedPlayer := BuildSoldier(expectedName)()
		//act
		player, err := New(BuildDefault())
		//assert
		assert.Equal(t, err, nil)
		assert.Equal(t, expectedPlayer.Job, player.Job)
		assert.Equal(t, expectedPlayer.Health, player.Health)
		assert.Equal(t, expectedPlayer.Level, player.Level)
		assert.Equal(t, expectedPlayer.Strength, player.Strength)
		assert.Equal(t, expectedPlayer.IsAlive(), player.IsAlive())
	}
}

func Test_WhenCreatePlayerInvalidName_Then_Error(t *testing.T) {

	//arrange
	inputTestList := []struct {
		name    string
		message string
	}{
		{"", PlayerMessageEmpty},
		{"1", PlayerMessageMin},
		{"12", PlayerMessageMin},
		{"Player-12345678901234", PlayerMessageMax},
		{"Player-1234567890123456789012345678901234567890", PlayerMessageMax},
	}

	for i := 0; i < len(inputTestList); i++ {
		// arrange
		expectedName := inputTestList[i].name
		expectedMessage := fmt.Errorf(inputTestList[i].message, expectedName)
		//act
		player, err := New(BuildSoldier(expectedName))
		//assert
		assert.NotNil(t, err)
		assert.Nil(t, player)
		assert.Equal(t, err.Error(), expectedMessage.Error())
	}
}

func Test_WhenCreateValidArcher_Then_Success(t *testing.T) {

	for i := 0; i < len(validPlayerNames); i++ {
		// arrange
		expectedName := validPlayerNames[i]
		expectedPlayer := BuildAllFields(expectedName, Archer, 450, 1, 45, 20, 40)()
		//act
		player, err := New(BuildArcher(expectedName))
		//assert
		assert.Equal(t, err, nil)
		assert.Equal(t, expectedPlayer.Job, player.Job)
		assert.Equal(t, expectedPlayer.Health, player.Health)
		assert.Equal(t, expectedPlayer.Level, player.Level)
		assert.Equal(t, expectedPlayer.Strength, player.Strength)
		assert.Equal(t, expectedPlayer.IsAlive(), player.IsAlive())
	}
}

func Test_WhenCreateValidWizard_Then_Success(t *testing.T) {

	for i := 0; i < len(validPlayerNames); i++ {
		// arrange
		expectedName := validPlayerNames[i]
		expectedPlayer := BuildAllFields(expectedName, Wizard, 400, 1, 10, 15, 55)()
		//act
		player, err := New(BuildWizard(expectedName))
		//assert
		assert.Equal(t, err, nil)
		assert.Equal(t, expectedPlayer.Job, player.Job)
		assert.Equal(t, expectedPlayer.Health, player.Health)
		assert.Equal(t, expectedPlayer.Level, player.Level)
		assert.Equal(t, expectedPlayer.Strength, player.Strength)
		assert.Equal(t, expectedPlayer.IsAlive(), player.IsAlive())
	}
}

func Test_WhenCreateValidKnight_Then_Success(t *testing.T) {

	for i := 0; i < len(validPlayerNames); i++ {
		// arrange
		expectedName := validPlayerNames[i]
		expectedPlayer := BuildAllFields(expectedName, Knight, 500, 1, 50, 30, 10)()
		//act
		player, err := New(BuildKnight(expectedName))
		//assert
		assert.Equal(t, err, nil)
		assert.Equal(t, expectedPlayer.Job, player.Job)
		assert.Equal(t, expectedPlayer.Health, player.Health)
		assert.Equal(t, expectedPlayer.Level, player.Level)
		assert.Equal(t, expectedPlayer.Strength, player.Strength)
		assert.Equal(t, expectedPlayer.IsAlive(), player.IsAlive())
	}
}

func Test_WhenIncreaseExperience_Then_LevelUp(t *testing.T) {

	var inputTestList = []struct {
		experience     int
		expectedHealth int
		expectedLevel  int
	}{
		{0, 500, 1},
		{50, 500, 1},
		{100, 700, 2},
		{110, 700, 2},
		{199, 700, 2},
	}

	for i := 0; i < len(inputTestList); i++ {
		// arrange
		experience := inputTestList[i].experience
		expectedHealth := inputTestList[i].expectedHealth
		expectedLevel := inputTestList[i].expectedLevel
		player1, _ := New(BuildKnight("Player-1"))
		//act
		player1.IncreaseExperience(experience)

		//assert
		assert.Equal(t, expectedHealth, player1.Health.Current)
		assert.Equal(t, expectedLevel, player1.Level)
	}
}

func Test_WhenCopyWizard_Then_AllInfoEquals(t *testing.T) {

	for i := 0; i < len(validPlayerNames); i++ {
		// arrange
		expectedName := validPlayerNames[i]
		expectedPlayer, err := New(BuildWizard(expectedName))
		//act
		player := expectedPlayer.Copy()
		//assert
		assert.Equal(t, err, nil)
		assert.Equal(t, expectedPlayer.Job, player.Job)
		assert.Equal(t, expectedPlayer.Health.Current, player.Health.Current)
		assert.Equal(t, expectedPlayer.Level, player.Level)
		assert.Equal(t, expectedPlayer.Strength, player.Strength)
		assert.Equal(t, expectedPlayer.Defense, player.Defense)
		assert.Equal(t, expectedPlayer.Magic, player.Magic)
		assert.Equal(t, expectedPlayer.IsAlive(), player.IsAlive())
	}
}
