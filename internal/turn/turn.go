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

	random_act := normalDice.Roll()
	random_tgt := normalDice.Roll()
	damage := (actor.Strength * random_act) - (target.Defense * random_tgt)
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

	random_act := magicDice.Roll()
	random_tgt := magicDice.Roll()
	damage := (actor.Magic * random_act) - (target.Magic * random_tgt)
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

	random_act := magicDice.Roll()
	heal := (actor.Strength * random_act)
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
