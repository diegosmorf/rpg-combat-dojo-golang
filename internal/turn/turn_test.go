package turn

import (
	"dojo/internal/dice"
	"dojo/internal/player"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenSoldier1Attack_Then_ApplyDamageToPlayer2(t *testing.T) {

	// arrange
	var validOptions = []struct {
		Value           int
		ExpectedDamaged int
	}{
		{1, 15},
		{2, 30},
		{3, 45},
		{4, 60},
		{5, 75},
		{6, 90},
	}

	for i := 0; i < len(validOptions); i++ {
		player1, _ := player.New(player.BuildSoldier("Player-1"))
		player2, _ := player.New(player.BuildSoldier("Player-2"))
		dice := dice.New(dice.BuildOnlyValue(validOptions[i].Value))
		expectedDamaged := validOptions[i].ExpectedDamaged
		initialP2Health := player2.Health.Current
		//act
		turn := RunPhysicalAttack(player1, player2, dice)

		//assert
		assert.Equal(t, expectedDamaged, turn.Damage)
		assert.Equal(t, initialP2Health-turn.Damage, player2.Health.Current)
	}
}

func Test_WhenSoldier1Heal_Then_ApplyHealToPlayer1(t *testing.T) {

	// arrange
	var validOptions = []struct {
		Value        int
		ExpectedHeal int
	}{
		{1, 35},
		{2, 70},
		{3, 105},
		{4, 140},
		{5, 175},
		{6, 210},
	}

	for i := 0; i < len(validOptions); i++ {
		initialHealthP1 := 1
		player1, _ := player.New(player.BuildSoldier("Player-1"))
		player1.Health.Current = initialHealthP1
		dice := dice.New(dice.BuildOnlyValue(validOptions[i].Value))
		expectedHeal := validOptions[i].ExpectedHeal
		//act
		turn := RunMagicHeal(player1, player1, dice)
		//assert
		assert.Equal(t, expectedHeal, turn.Heal)
		assert.Equal(t, initialHealthP1+turn.Heal, player1.Health.Current)
	}
}

func Test_WhenWizard1Fireball_Then_ApplyDamageToPlayer2(t *testing.T) {

	// arrange
	var validOptions = []struct {
		Value           int
		ExpectedDamaged int
	}{
		{1, 35},
		{2, 70},
		{3, 105},
		{4, 140},
		{5, 175},
		{6, 210},
	}

	for i := 0; i < len(validOptions); i++ {
		player1, _ := player.New(player.BuildWizard("Player-1"))
		player2, _ := player.New(player.BuildSoldier("Player-2"))
		dice := dice.New(dice.BuildOnlyValue(validOptions[i].Value))
		expectedDamaged := validOptions[i].ExpectedDamaged
		initialP2Health := player2.Health.Current
		//act
		turn := RunMagicFireBall(player1, player2, dice)

		//assert
		assert.Equal(t, expectedDamaged, turn.Damage)
		assert.Equal(t, initialP2Health-turn.Damage, player2.Health.Current)
	}
}

func Test_WhenRunPhysicalAttack_Then_CheckDamage(t *testing.T) {
	// arrange
	var validOptions = []struct {
		Value           int
		ExpectedDamaged int
	}{
		{1, 15},
		{2, 30},
		{3, 45},
		{4, 60},
		{5, 75},
		{6, 90},
	}

	for i := 0; i < len(validOptions); i++ {
		player1, _ := player.New(player.BuildSoldier("Player-1"))
		player2, _ := player.New(player.BuildSoldier("Player-1"))
		dice := dice.New(dice.BuildOnlyValue(validOptions[i].Value))
		expectedDamaged := validOptions[i].ExpectedDamaged
		initialP2Health := player2.Health.Current
		//act
		turn := RunPhysicalAttack(player1, player2, dice)
		//assert
		assert.Equal(t, expectedDamaged, turn.Damage)
		assert.Equal(t, initialP2Health-turn.Damage, player2.Health.Current)
	}
}

func Test_WhenRunFireball_Then_CheckDamage(t *testing.T) {
	// arrange
	var validOptions = []struct {
		Value           int
		ExpectedDamaged int
	}{
		{1, 35},
		{2, 70},
		{3, 105},
		{4, 140},
		{5, 175},
		{6, 210},
	}

	for i := 0; i < len(validOptions); i++ {
		player1, _ := player.New(player.BuildWizard("Player-1"))
		player2, _ := player.New(player.BuildSoldier("Player-1"))
		dice := dice.New(dice.BuildOnlyValue(validOptions[i].Value))
		expectedDamaged := validOptions[i].ExpectedDamaged
		initialP2Health := player2.Health.Current
		//act
		turn := RunMagicFireBall(player1, player2, dice)
		//assert
		assert.Equal(t, expectedDamaged, turn.Damage)
		assert.Equal(t, initialP2Health-turn.Damage, player2.Health.Current)
	}
}
