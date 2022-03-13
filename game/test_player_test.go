package game

import (
	"testing"
)

func TestRecvSetIcon(t *testing.T) {
	player := CreateTestPlayer()

	player.RecvSetIcon(1)
}