package battle

import (
	"dojo/internal/dice"
	"dojo/internal/player"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenRun2Turns_Then_ShowWinner(t *testing.T) {
	// arrange
	expectedNumberOfTurns := 2
	player1, _ := player.New(player.BuildSoldier("Player-1"))
	player2, _ := player.New(player.BuildSoldier("Player-2"))
	players := []player.Player{*player1, *player2}
	dice := dice.New(dice.BuildDefault())
	battle := New(players, dice, dice, expectedNumberOfTurns)
	//act
	battle.RunTurns()
	//assert
	assert.Equal(t, expectedNumberOfTurns, battle.TotalNumberOfTurns())
	assert.Greater(t, battle.Winner.Health.Current, battle.Looser.Health.Current)
	assert.True(t, battle.Players[0].IsAlive())
	assert.True(t, battle.Players[1].IsAlive())
	assert.Equal(t, battle.Status, Unfinished)

}

func Test_WhenRun2Turns_Then_ShowDraw(t *testing.T) {
	// arrange
	expectedNumberOfTurns := 2
	player1, _ := player.New(player.BuildSoldier("Player-1"))
	player2, _ := player.New(player.BuildSoldier("Player-2"))
	players := []player.Player{*player1, *player2}
	dice := dice.New(dice.BuildOnlyValue(1))
	battle := New(players, dice, dice, expectedNumberOfTurns)
	//act
	battle.RunTurns()
	//assert
	assert.Equal(t, expectedNumberOfTurns, battle.TotalNumberOfTurns())
	assert.Nil(t, battle.Winner)
	assert.Nil(t, battle.Looser)
	assert.True(t, battle.Players[0].IsAlive())
	assert.True(t, battle.Players[1].IsAlive())
	assert.Equal(t, battle.Status, Draw)
}

func Test_WhenRun4Turns_Then_ShowWinner(t *testing.T) {
	// arrange
	expectedNumberOfTurns := 4
	player1, _ := player.New(player.BuildSoldier("Player-1"))
	player2, _ := player.New(player.BuildSoldier("Player-2"))
	players := []player.Player{*player1, *player2}
	dice := dice.New(dice.BuildDefault())
	battle := New(players, dice, dice, expectedNumberOfTurns)
	//act
	battle.RunTurns()
	//assert
	assert.Equal(t, expectedNumberOfTurns, battle.TotalNumberOfTurns())
	assert.Greater(t, battle.Winner.Health.Current, battle.Looser.Health.Current)
	assert.Equal(t, battle.Status, Unfinished)
}

func Test_WhenRun10Turns_Then_ShowDraw(t *testing.T) {
	// arrange
	expectedNumberOfTurns := 4
	player1, _ := player.New(player.BuildSoldier("Player-1"))
	player2, _ := player.New(player.BuildSoldier("Player-2"))
	players := []player.Player{*player1, *player2}
	dice := dice.New(dice.BuildOnlyValue(1))
	battle := New(players, dice, dice, expectedNumberOfTurns)
	//act
	battle.RunTurns()
	//assert
	assert.Equal(t, expectedNumberOfTurns, battle.TotalNumberOfTurns())
	assert.Nil(t, battle.Winner)
	assert.Nil(t, battle.Looser)
	assert.Equal(t, battle.Status, Draw)
}

func Test_WhenRunTurns_UntilPlayerDies_Then_ShowWinner(t *testing.T) {
	// arrange
	maxTurns := 100
	player1, _ := player.New(player.BuildSoldier("Player-1"))
	player2, _ := player.New(player.BuildSoldier("Player-2"))
	players := []player.Player{*player1, *player2}
	dice := dice.New(dice.BuildDefault())
	battle := New(players, dice, dice, maxTurns)
	//act
	battle.RunTurns()
	//assert
	assert.Greater(t, battle.TotalNumberOfTurns(), 0)
	assert.LessOrEqual(t, battle.TotalNumberOfTurns(), maxTurns)
	assert.Greater(t, battle.Winner.Health.Current, 0)
	assert.True(t, battle.Winner.IsAlive())
	assert.LessOrEqual(t, battle.Looser.Health.Current, 0)
	assert.False(t, battle.Looser.IsAlive())
	assert.Equal(t, battle.Status, Finished)
}

func Test_WhenRun3Turns_Knight_Wizard_Then_ShowWinner(t *testing.T) {
	// arrange
	maxTurns := 3
	player1, _ := player.New(player.BuildKnight("Player-1"))
	player2, _ := player.New(player.BuildWizard("Player-2"))
	players := []player.Player{*player1, *player2}
	normalDice := dice.New(dice.BuildDefault())
	magicDice := dice.New(dice.BuildDefaultMagic())
	battle := New(players, normalDice, magicDice, maxTurns)
	//act
	battle.RunTurns()
	//assert
	assert.Greater(t, battle.TotalNumberOfTurns(), 0)
	assert.LessOrEqual(t, battle.TotalNumberOfTurns(), maxTurns)
	assert.Greater(t, battle.Winner.Health.Current, 0)
	assert.True(t, battle.Winner.IsAlive())
	assert.Equal(t, battle.Status, Unfinished)
}

func Test_WhenRun3Turns_Knight_Archer_Then_ShowWinner(t *testing.T) {
	// arrange
	maxTurns := 3
	player1, _ := player.New(player.BuildKnight("Player-1"))
	player2, _ := player.New(player.BuildArcher("Player-2"))
	players := []player.Player{*player1, *player2}
	normalDice := dice.New(dice.BuildDefault())
	magicDice := dice.New(dice.BuildDefaultMagic())
	battle := New(players, normalDice, magicDice, maxTurns)
	//act
	battle.RunTurns()
	//assert
	assert.Greater(t, battle.TotalNumberOfTurns(), 0)
	assert.LessOrEqual(t, battle.TotalNumberOfTurns(), maxTurns)
	assert.Greater(t, battle.Winner.Health.Current, 0)
	assert.True(t, battle.Winner.IsAlive())
	assert.Equal(t, battle.Status, Unfinished)
}
