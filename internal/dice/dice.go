package dice

import (
	"math/rand"
)

type DiceBuilderFunc func() Dice

func BuildDefault() DiceBuilderFunc {
	return func() Dice {
		return BuildMultipleMinMax(1, 6)()
	}
}

func BuildDefaultMagic() DiceBuilderFunc {
	return func() Dice {
		return BuildMultipleMinMax(1, 8)()
	}
}

type Dice struct {
	ValidOptions []int
}

func BuildOnlyValue(value int) DiceBuilderFunc {
	return func() Dice {
		return Dice{ValidOptions: []int{value}}

	}
}

func BuildMultipleSameValues(value int) DiceBuilderFunc {
	return func() Dice {
		return Dice{ValidOptions: []int{value, value, value, value, value, value}}

	}
}

func BuildMultipleMinMax(min int, max int) DiceBuilderFunc {
	return func() Dice {
		dice := Dice{}
		for i := min; i <= max; i++ {
			dice.ValidOptions = append(dice.ValidOptions, i)
		}
		return dice
	}
}

func New(builder ...DiceBuilderFunc) Dice {

	if len(builder) == 0 {
		return BuildDefault()()
	}

	return builder[0]()
}

func (dice *Dice) Roll() int {
	random := getRandomFromRange(dice.ValidOptions)
	return dice.ValidOptions[random]
}

func getRandomFromRange(opt []int) int {
	return rand.Intn(len(opt))
}
