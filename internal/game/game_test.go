package game

import (
	b "dojo/internal/battle"
	p "dojo/internal/player"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenGameStarts_Then_2Players(t *testing.T) {

	// arrange
	expectedPlayers := 2
	//act
	game, _ := New()

	var currentValue = game.NumberOfPlayers()
	//assert
	assert.Equal(t, expectedPlayers, currentValue)

}

var validPlayerNames = []string{
	"Player-1",
	"Player-100",
	"Player-10000",
	"Player-1234567890",
	"Player-123456789012",
	"Player-1234567890123",
}

func Test_WhenGameAddValidPlayers_Then_NumberIsCorrect(t *testing.T) {

	game := Game{}
	expectedPlayers := 0

	for i := 0; i < len(validPlayerNames); i++ {
		// arrange
		expectedPlayers++
		expectedName := validPlayerNames[i]
		//act
		p1, _ := p.New(p.BuildSoldier(expectedName))
		game.AddPlayer(p1)
	}

	//assert
	assert.Equal(t, expectedPlayers, game.NumberOfPlayers())
}

func Test_WhenGameRunBattle_Soldier_Soldier_Then_ShowWinner(t *testing.T) {
	// arrange
	maxTurns := 100
	game := Game{}
	player1, _ := p.New(p.BuildSoldier("Player-1"))
	player2, _ := p.New(p.BuildSoldier("Player-2"))

	//act
	game.AddPlayer(player1)
	game.AddPlayer(player2)
	game.RunBattle()
	battle := game.Battles[0]
	//assert
	assert.Greater(t, battle.TotalNumberOfTurns(), 0)
	assert.LessOrEqual(t, battle.TotalNumberOfTurns(), maxTurns)
	assert.Greater(t, battle.Winner.Health.Current, 0)
	assert.True(t, battle.Winner.IsAlive())
	assert.LessOrEqual(t, battle.Looser.Health.Current, 0)
	assert.False(t, battle.Looser.IsAlive())
	assert.Equal(t, battle.Status, b.Finished)
}

func Test_WhenGameRunBattle_Soldier_Archer_Then_ShowWinner(t *testing.T) {
	// arrange
	maxTurns := 100
	game := Game{}
	player1, _ := p.New(p.BuildSoldier("Player-1"))
	player2, _ := p.New(p.BuildArcher("Player-2"))

	//act
	game.AddPlayer(player1)
	game.AddPlayer(player2)
	game.RunBattle()
	battle := game.Battles[0]
	//assert
	assert.Greater(t, battle.TotalNumberOfTurns(), 0)
	assert.LessOrEqual(t, battle.TotalNumberOfTurns(), maxTurns)
	assert.Greater(t, battle.Winner.Health.Current, 0)
	assert.True(t, battle.Winner.IsAlive())
	assert.LessOrEqual(t, battle.Looser.Health.Current, 0)
	assert.False(t, battle.Looser.IsAlive())
	assert.Equal(t, battle.Status, b.Finished)
}
