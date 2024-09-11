package player

type PlayerJob int

const (
	Soldier PlayerJob = iota
	Wizard  PlayerJob = iota
	Archer  PlayerJob = iota
	Knight  PlayerJob = iota
)

func (job PlayerJob) String() string {
	return [...]string{"Soldier", "Wizard", "Archer", "Knight"}[job]
}

func (job PlayerJob) EnumIndex() int {
	return int(job)
}
