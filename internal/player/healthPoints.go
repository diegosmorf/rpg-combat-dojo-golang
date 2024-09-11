package player

type HealthPoints struct {
	Current int
	Min     int
	Max     int
}

func BuildHealth(max int) HealthPoints {
	return BuildHealthAllFields(max, 0, max)
}

func BuildHealthAllFields(current int, min int, max int) HealthPoints {
	return HealthPoints{Current: current, Min: min, Max: max}
}

func (health *HealthPoints) ApplyDamage(damage int) {

	if damage > 0 {
		health.Current = health.Current - damage
		if health.Current < health.Min {
			health.Current = health.Min
		}
	}
}

func (health *HealthPoints) ApplyHeal(heal int) {
	health.Current = health.Current + heal

	if health.Current > health.Max {
		health.Current = health.Max
	}
}

func (health *HealthPoints) Reset() {
	health.Current = health.Max
}

func (health *HealthPoints) LevelUp(increase int) {
	health.Max = health.Max + increase
	health.Reset()
}
