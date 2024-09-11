package game

import (
	battle "dojo/internal/battle"
	"dojo/internal/dice"
	player "dojo/internal/player"
)

type Game struct {
	Players []player.Player
	Battles []battle.Battle
}

func New() (*Game, error) {

	p1, _ := player.New(player.BuildKnight("Arthur King"))
	p2, _ := player.New(player.BuildWizard("Merlin the Magician"))

	game := Game{}
	game.AddPlayer(p1)
	game.AddPlayer(p2)

	return &game, nil
}

func (game *Game) AddPlayer(p *player.Player) {
	game.Players = append(game.Players, *p)
}

func (game *Game) RunBattle() {
	maxTurns := 100
	normalDice := dice.New(dice.BuildDefault())
	magicDice := dice.New(dice.BuildDefaultMagic())
	battle := battle.New(game.Players, normalDice, magicDice, maxTurns)
	battle.RunTurns()
	game.Battles = append(game.Battles, *battle)
}

func (game Game) NumberOfPlayers() int {
	return len(game.Players)
}
