package main

import (
	"fmt"
	"github.com/xavividal/ddd-go-example/internal/app/player"
	"github.com/xavividal/ddd-go-example/internal/infra/container"
)

func main() {
	container.New()
	c := container.Instance()

	rq := player.CreatePlayerRQ{Name: "Han Solo"}
	rs := c.App.Player.CreatePlayer.Execute(rq)

	fmt.Println(rs)
}
