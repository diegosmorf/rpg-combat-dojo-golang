package turn

import (
	action "dojo/internal/action"
	dice "dojo/internal/dice"
	player "dojo/internal/player"
)

type Turn struct {
	Actor  *player.Player
	Target *player.Player
	Action action.ActionType
	Damage int
	Heal   int
}

func RunPhysicalAttack(actor *player.Player, target *player.Player, normalDice dice.Dice) Turn {

	actorDiceValue := normalDice.Roll()
	targetDiceValue := normalDice.Roll()
	damage := (actor.Strength * actorDiceValue) - (target.Defense * targetDiceValue)
	experience := calculateExperience(target.Health.Current, damage)
	target.Health.ApplyDamage(damage)
	actor.IncreaseExperience(experience)

	return Turn{
		Actor:  actor.Copy(),
		Target: target.Copy(),
		Damage: damage,
		Action: action.Attack}
}

func RunMagicFireBall(actor *player.Player, target *player.Player, magicDice dice.Dice) Turn {

	actorDiceValue := magicDice.Roll()
	targetDiceValue := magicDice.Roll()
	damage := (actor.Magic * actorDiceValue) - (target.Magic * targetDiceValue)
	experience := calculateExperience(target.Health.Current, damage)
	target.Health.ApplyDamage(damage)
	actor.IncreaseExperience(experience)

	return Turn{
		Actor:  actor.Copy(),
		Target: target.Copy(),
		Damage: damage,
		Action: action.FireBallMagic}
}

func RunMagicHeal(actor *player.Player, target *player.Player, magicDice dice.Dice) Turn {

	actorDiceValue := magicDice.Roll()
	heal := (actor.Strength * actorDiceValue)
	target.Health.ApplyHeal(heal)
	experience := calculateExperience(actor.Health.Current, heal)
	actor.IncreaseExperience(experience)

	return Turn{
		Actor:  actor.Copy(),
		Target: target.Copy(),
		Heal:   heal,
		Action: action.HealMagic}
}

func calculateExperience(current int, change int) int {
	return (int)(change/current) * 20
}
