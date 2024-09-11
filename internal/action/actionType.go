package action

type ActionType int

const (
	Attack        ActionType = iota
	FireBallMagic ActionType = iota
	HealMagic     ActionType = iota
	Item          ActionType = iota
)

func (action ActionType) String() string {
	return [...]string{"Attack", "FireBallMagic", "HealMagic", "Item"}[action]
}

func (action ActionType) EnumIndex() int {
	return int(action)
}
