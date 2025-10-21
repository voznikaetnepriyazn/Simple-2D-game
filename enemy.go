package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	x, y   int
	sprite *ebiten.Image
}

func InitEnemy(x, y int, sprite *ebiten.Image) *Enemy {
	return &Enemy{
		x:      x,
		y:      y,
		sprite: sprite,
	}
}

func Update() err {

}
