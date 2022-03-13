package game

import (
	"fmt"
	"testing"
)

func TestPlayer_RecvSetIcon(t *testing.T) {
	player := CreateTestPlayer()

	player.RecvSetIcon(1)
	fmt.Println(player.ModPlayer.Icon == 1)
}

func TestModPlayer_RecvSetCard(t *testing.T) {
	player := CreateTestPlayer()
	player.RecvSetCard(4)
	fmt.Println(player.ModPlayer.Card == 4)
}
