package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/xavividal/ddd-go-example/internal/app/player"
	"github.com/xavividal/ddd-go-example/internal/infra/container"
)

func main() {
	c := container.Instance()

	rq := player.CreatePlayerRQ{Name: "Han Solo"}
	rs := c.App.Player.CreatePlayer.Execute(rq)

	spew.Dump(rs)
}
