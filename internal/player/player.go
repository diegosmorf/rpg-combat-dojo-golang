package player

import (
	"dojo/internal/utils"
	"fmt"
)

type Player struct {
	Name       string
	Job        PlayerJob
	Health     HealthPoints
	Level      int
	Experience int
	Strength   int
	Defense    int
	Magic      int
}

const (
	NameMin             int    = 3
	NameMax             int    = 20
	PlayerMessageEmpty  string = "Player Name cannot be empty. Current:'%s'"
	PlayerMessageMin    string = "Player Name length cannot be lower than 3 characters. Current:'%s'"
	PlayerMessageMax    string = "Player Name length cannot be greater than 20 characters. Current:'%s'"
	ExperienceToLevelUp int    = 100
	HealthLevelUp       int    = 200
	StrengthLevelUp     int    = 10
	DefenseLevelUp      int    = 10
	MagicLevelUp        int    = 10
)

func New(builder ...PlayerBuilderFunc) (*Player, error) {
	player := BuildDefault()()

	if len(builder) == 1 {
		player = builder[0]()
	}

	err := player.IsValid()

	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (player *Player) LevelUp() {
	player.Level++
	player.Health.LevelUp(HealthLevelUp)
	player.Strength += StrengthLevelUp
	player.Defense += DefenseLevelUp
	player.Magic += MagicLevelUp
}

func (player *Player) IncreaseExperience(experience int) {
	player.Experience += experience
	if player.Experience >= ExperienceToLevelUp {
		player.LevelUp()
		player.Experience = 0
	}
}

func (player *Player) Copy() *Player {
	return &Player{
		Name:       player.Name,
		Job:        player.Job,
		Health:     player.Health,
		Level:      player.Level,
		Experience: player.Experience,
		Strength:   player.Strength,
		Defense:    player.Defense,
		Magic:      player.Magic}
}

func (player *Player) IsAlive() bool {
	return player.Health.Current > 0
}

func (player *Player) IsValid() error {

	if player.Name == "" {
		return fmt.Errorf(PlayerMessageEmpty, player.Name)
	}

	if len(player.Name) < NameMin {
		return fmt.Errorf(PlayerMessageMin, player.Name)
	}

	if len(player.Name) > NameMax {
		return fmt.Errorf(PlayerMessageMax, player.Name)
	}

	return nil
}

type PlayerBuilderFunc func() Player

func BuildAllFields(name string, job PlayerJob, health int, level int, strength int, defense int, magic int) PlayerBuilderFunc {
	return func() Player {
		return Player{
			Name:       name,
			Job:        job,
			Health:     BuildHealth(health),
			Level:      level,
			Experience: 0,
			Strength:   strength,
			Defense:    defense,
			Magic:      magic}
	}
}

func BuildWizard(name string) PlayerBuilderFunc {
	return func() Player {
		return BuildAllFields(name, Wizard, 400, 1, 10, 15, 55)()
	}
}

func BuildArcher(name string) PlayerBuilderFunc {
	return func() Player {
		return BuildAllFields(name, Archer, 450, 1, 45, 20, 40)()
	}
}

func BuildKnight(name string) PlayerBuilderFunc {
	return func() Player {
		return BuildAllFields(name, Knight, 500, 1, 50, 30, 10)()
	}
}
func BuildSoldier(name string) PlayerBuilderFunc {
	return func() Player {
		return BuildAllFields(name, Soldier, 500, 1, 35, 20, 20)()
	}
}

func BuildDefault() PlayerBuilderFunc {
	return func() Player {
		return BuildSoldier("Player-" + utils.GetUniqueId())()
	}
}
