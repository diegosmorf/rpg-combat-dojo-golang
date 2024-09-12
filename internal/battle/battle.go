package battle

import (
	dice "dojo/internal/dice"
	player "dojo/internal/player"
	turn "dojo/internal/turn"
)

type Battle struct {
	Turns      []turn.Turn
	Players    []player.Player
	Winner     *player.Player
	Looser     *player.Player
	NormalDice dice.Dice
	MagicDice  dice.Dice
	Status     BattleStatus
	MaxTurns   int
}

func New(players []player.Player, dice dice.Dice, magicDice dice.Dice, maxTurns int) *Battle {
	return &Battle{
		Players:    players,
		NormalDice: dice,
		MagicDice:  magicDice,
		MaxTurns:   maxTurns}
}

func (battle *Battle) TotalNumberOfTurns() int {
	return len(battle.Turns)
}

func (battle *Battle) RunTurns() {
	actor := battle.Players[0]
	target := battle.Players[1]

	for i := 0; i < battle.MaxTurns; i++ {

		turn := battle.RunTurn(&actor, &target)
		battle.Turns = append(battle.Turns, turn)

		if !target.IsAlive() {
			break
		}

		temp := actor
		actor = target
		target = temp
	}

	if !target.IsAlive() {
		battle.Status = Finished
		battle.Winner = &actor
		battle.Looser = &target
	} else {
		battle.Status = Unfinished
		if actor.Health.Current == target.Health.Current {
			battle.Status = Draw
			battle.Winner = nil
			battle.Looser = nil
		} else if actor.Health.Current > target.Health.Current {
			battle.Winner = &actor
			battle.Looser = &target
		} else {
			battle.Winner = &target
			battle.Looser = &actor
		}
	}
}

func (battle *Battle) RunTurn(actor *player.Player, target *player.Player) turn.Turn {

	if actor.Health.Current < 150 {
		return turn.RunMagicHeal(actor, actor, battle.MagicDice)
	}

	if actor.Job == player.Wizard {
		return turn.RunMagicFireBall(actor, target, battle.MagicDice)
	}

	if actor.Job == player.Archer && target.Job != player.Wizard {
		return turn.RunMagicFireBall(actor, target, battle.MagicDice)
	}

	return turn.RunPhysicalAttack(actor, target, battle.NormalDice)
}
