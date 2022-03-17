package game

import (
	"fmt"
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

	player.modPlayer.addExp(5000, player)

	if !util.Assert(player.modPlayer.PlayerExp, int64(975)) {
		t.Errorf("add exp error exp not match")
	}

	if !util.Assert(player.modPlayer.PlayerLevel, 6) {
		t.Errorf("add exp error level not match")
	}
}

func TestModPlayer_AddExp2(t *testing.T) {
	p := CreateTestPlayer()
	p.modPlayer.addExp(5000000000, p)

	fmt.Println(p.modPlayer.PlayerLevel)
}

func TestPlayer_SetShowCard(t *testing.T) {
	p := CreateTestPlayer()

	ints := []int{1, 2, 34, 5, 6, 7, 8, 10, 11, 12, 11}
	p.SetShowCard(ints)
	if !util.Assert(len(p.modPlayer.ShowCard), 0) {
		t.Errorf("")
	}
	ints = []int{1, 2, 34, 5, 6, 7, 8, 10, 11}
	p.SetShowCard(ints)
	if !util.Assert(len(p.modPlayer.ShowCard), 9) {
		t.Errorf("")
	}
	ints = []int{1, 2, 34, 5, 6, 7, 8, 11, 11}
	p.SetShowCard(ints)
	if !util.Assert(len(p.modPlayer.ShowCard), 8) {
		t.Errorf("")
	}
}

func TestPlayer_ReduceWorldLevel(t *testing.T) {
	p := CreateTestPlayer()

	p.modPlayer.WorldLevel = 6
	p.modPlayer.WorldLevelNow = 6

	p.ReduceWorldLevel()

	if !util.Assert(p.modPlayer.WorldLevel, int(6)) {
		t.Errorf("")
	}
	if !util.Assert(p.modPlayer.WorldLevelNow, int(5)) {
		t.Errorf("")
	}

	cool := p.modPlayer.WorldLevelCool

	p.ReduceWorldLevel()

	if !util.Assert(p.modPlayer.WorldLevelCool, cool) {
		t.Errorf("")
	}
}

func TestPlayer_SetBirth(t *testing.T) {
	player := CreateTestPlayer()

	player.SetBirth(1310)
	player.SetBirth(1010)
	player.SetBirth(1111)
}
