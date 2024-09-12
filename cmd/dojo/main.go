package main

import (
	"dojo/internal/action"
	"dojo/internal/battle"
	game "dojo/internal/game"
	"dojo/internal/player"
	"dojo/internal/turn"
	"log"
)

func main() {
	CreateNewGame()
}

func CreateNewGame() bool {
	game, _ := game.New()
	game.RunBattle()
	printGame(game)

	return true
}

func printPlayer(message string, p *player.Player) {
	print("%s: %s - Job: %s - Health: %d - Level: %d", message, p.Name, p.Job, p.Health.Current, p.Level)
}

func printRule() {
	print("-----------------------------------------------------------------------------------")
}

func printGame(g *game.Game) {
	printRule()
	print("Welcome to Dojo RPG Combat System - Version GoLang")
	printRule()
	printPlayer("Player 1", &g.Players[0])
	printPlayer("Player 2", &g.Players[1])
	printRule()

	for index, battle := range g.Battles {
		printBattle(index, &battle)
	}
	printRule()
}

func print(format string, v ...any) {
	log.Printf(format, v...)
}

func printBattle(index int, b *battle.Battle) {

	print("Battle: %d", index)
	print("Total Number of Turns: %d", b.TotalNumberOfTurns())
	printRule()
	for index, turn := range b.Turns {
		printTurn(index, turn)
	}
	printPlayer("Winner", b.Winner)
	printPlayer("Looser", b.Looser)
}

func printTurn(index int, t turn.Turn) {
	print("Turn: %d", index)
	print("Turn Action: %s", t.Action)
	printPlayer("Actor", t.Actor)
	printPlayer("Target", t.Target)

	if t.Action == action.HealMagic {
		print("Heal: %d", t.Heal)
	} else {
		print("Damage: %d", t.Damage)
	}
	printRule()
}
