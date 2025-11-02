package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player *Player
	maps   *Maps
	items  []*Item
}

func InitGame(player *Player, maps *Maps, items []*Item) *Game {
	return &Game{
		player: player,
		maps:   maps,
		items:  items,
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsWindowBeingClosed() {
		//закрытие программы
		return errors.New("ты вышел из игры")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	playerSprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\player.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение игрока", err)
	}

	itemSprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\item.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение предмета", err)
	}

	wallSprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\wall.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение стены", err)
	}

	exitSprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\exit.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение выхода", err)
	}

	emptySprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\empty.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение пустого пространства", err)
	}

	enemySprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\enemy.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение врага", err)
	}

	g.maps.Draw(playerSprite, wallSprite, emptySprite, itemSprite, exitSprite, enemySprite)

	ebitenutil.DebugPrint(screen, "start")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{
		player: &Player{},
		maps:   &Maps{},
		items:  []*Item{},
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("servive in murino")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
