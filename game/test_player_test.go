package game

import (
	"game/util"
	"testing"
)

func TestPlayer_RecvSetIcon(t *testing.T) {
	player := CreateTestPlayer()

	player.RecvSetIcon(1)
	if !(player.modPlayer.Icon == 1) {
		t.Errorf("recv set icon error %d", player.modPlayer.Icon)
	}
}

func TestModPlayer_RecvSetCard(t *testing.T) {
	player := CreateTestPlayer()
	player.RecvSetCard(4)
	if !(player.modPlayer.Card == 4) {
		t.Errorf("recv set card error %d", player.modPlayer.Card)
	}
}

func TestPlayer_RecvSetName(t *testing.T) {
	player := CreateTestPlayer()

	player.RecvSetName("我操")
	if !(player.modPlayer.Name == "我操") {
		t.Errorf("recv set name error %s", player.modPlayer.Name)
	}

	player.RecvSetName("fzdwx")
	if !(player.modPlayer.Name == "我操") {
		t.Errorf("recv set name error %s", player.modPlayer.Name)
	}

	player.RecvSetName("外挂")
	if !(player.modPlayer.Name == "我操") {
		t.Errorf("recv set name error %s", player.modPlayer.Name)
	}

	player.RecvSetName("你好")
	if !(player.modPlayer.Name == "你好") {
		t.Errorf("recv set name error %s", player.modPlayer.Name)
	}
}

func TestPlayer_RecvSetSign(t *testing.T) {
	player := CreateTestPlayer()

	player.RecvSetSign("我操")
	if !(player.modPlayer.Sign == "我操") {
		t.Errorf("recv set sign   error %s", player.modPlayer.Sign)
	}

	player.RecvSetSign("fzdwx")
	if !(player.modPlayer.Sign == "我操") {
		t.Errorf("recv set sign error %s", player.modPlayer.Sign)
	}

	player.RecvSetSign("外挂")
	if !(player.modPlayer.Sign == "我操") {
		t.Errorf("recv set sign error %s", player.modPlayer.Sign)
	}

	player.RecvSetSign("你好")
	if !(player.modPlayer.Sign == "你好") {
		t.Errorf("recv set sign error %s", player.modPlayer.Sign)
	}
}

func TestModPlayer_AddExp(t *testing.T) {
	player := CreateTestPlayer()

	player.modPlayer.AddExp(5000)

	if !util.Assert(player.modPlayer.PlayerExp, int64(975)) {
		t.Errorf("add exp error exp not match")
	}

	if !util.Assert(player.modPlayer.PlayerLevel, 6) {
		t.Errorf("add exp error level not match")
	}
}
