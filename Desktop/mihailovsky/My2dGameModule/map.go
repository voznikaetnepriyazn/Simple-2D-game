package main

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

var gameMap = [][]rune{
	{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	{'1', '0', 'C', '0', '0', '0', 'C', '0', '0', '0', '1'},
	{'1', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1'},
	{'1', '0', '0', '0', 'P', '0', '0', '0', 'C', '0', '1'},
	{'1', 'C', '0', '0', 'C', '0', '0', 'E', '0', '0', '1'},
	{'1', '0', '0', '0', '0', '0', '0', '0', '0', 'C', '1'},
	{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
}

type Maps struct {
	Mmap [][]rune
}

func InitMap() *Maps {
	return &Maps{
		Mmap: gameMap,
	}
}

func (m *Maps) ReturnContentOfCell(x, y int) any { //содержимое клетки по координатам
	if x < 0 || x > len(m.Mmap[0]) || y < 0 || y >= len(m.Mmap) { //границы карты
		return errors.New("выход за пределы карты")
	}
	return m.Mmap[x][y]
}

func (m *Maps) Update() error {
	return nil
}

func (m *Maps) Draw(screen *ebiten.Image, wallSprite, emptySprite, itemSprite, exitSprite *ebiten.Image) {
	for y, row := range m.Mmap {
		for x, cell := range row {
			var sprite *ebiten.Image
			switch cell {
			case '1':
				sprite = wallSprite
			case '0':
				sprite = emptySprite
			case 'C':
				sprite = itemSprite
			case 'E':
				sprite = exitSprite
			default:
				continue
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x)*32, float64(y)*32)
			screen.DrawImage(sprite, op)
		}
	}
}
