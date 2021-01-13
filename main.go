package main

import (
	"github.com/xavividal/ddd-go-example/internal/app/player"
	"github.com/xavividal/ddd-go-example/internal/infra/container"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	container.New()
	c := container.Instance()

	rq := player.CreatePlayerRQ{Name: "Han Solo"}
	rs := c.App.Player.CreatePlayer.Execute(rq)

	spew.Dump(rs)
}
