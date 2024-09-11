package battle

type BattleStatus int

const (
	Finished   BattleStatus = iota
	Draw       BattleStatus = iota
	Unfinished BattleStatus = iota
)

func (action BattleStatus) String() string {
	return [...]string{"Finished", "Draw", "Unfinished"}[action]
}

func (action BattleStatus) EnumIndex() int {
	return int(action)
}
