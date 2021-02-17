package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"moul.io/banner"
)

func layout(g *gocui.Gui) error {
	//maxX, maxY := g.Size()
	titulo, err := g.SetView("Titulo", 0, 0, 79, 5)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	titulo.Clear()
	fmt.Fprint(titulo, banner.Inline("Tic-Tac-Term"))

	playerInfo, err := g.SetView("playerInfo", 2, 6, 18, 11)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	playerInfo.Clear()
	playerInfo.Title = "Player info"
	fmt.Fprint(playerInfo, "Name: breh")

	cambiarNombreBtn, err := g.SetView("cambiarNombreBtn", 4, 8, 12, 10)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	cambiarNombreBtn.Clear()
	g.SetViewOnTop("cambiarNombreBtn")
	fmt.Fprint(cambiarNombreBtn, "Change")
	cambiarNombreBtn.Highlight = true

	oponnentInfo, err := g.SetView("opponentInfo", 2, 12, 18, 14)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	oponnentInfo.Clear()
	oponnentInfo.Title = "Opponent info"
	fmt.Fprint(oponnentInfo, "Name: bruh")

	playMode, err := g.SetView("playMode", 2, 15, 18, 18)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	playMode.Clear()
	playMode.Highlight = true
	playMode.Title = "Play with: "
	fmt.Fprintln(playMode, "A bot")
	fmt.Fprintln(playMode, "A human")

	board, err := g.SetView("board", 22, 6, 76, 22)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	board.Clear()
	board.Title = "Board"
	board.Highlight = true

	spot11, err := g.SetView("spot11", 42, 7, 52, 11)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	spot11.Highlight = true

	spot12, err := g.SetView("spot12", 53, 7, 63, 11)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	spot12.Highlight = true

	spot13, err := g.SetView("spot13", 64, 7, 74, 11)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	spot13.Highlight = true

	spot21, err := g.SetView("spot21", 42, 12, 52, 16)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	spot21.Highlight = true

	spot22, err := g.SetView("spot22", 53, 12, 63, 16)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	spot22.Highlight = true

	spot23, err := g.SetView("spot23", 64, 12, 74, 16)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	spot23.Highlight = true

	spot31, err := g.SetView("spot31", 42, 17, 52, 21)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	spot31.Highlight = true

	spot32, err := g.SetView("spot32", 53, 17, 63, 21)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	spot32.Highlight = true

	spot33, err := g.SetView("spot33", 64, 17, 74, 21)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	spot33.Highlight = true

	turnIndicator, err := g.SetView("turnIndicator", 25, 7, 40, 9)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	turnIndicator.Frame = false
	turnIndicator.Clear()
	fmt.Fprintln(turnIndicator, "Breh's turn")

	shapeIndicator, err := g.SetView("shapeIndicator", 25, 12, 40, 15)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	shapeIndicator.Frame = false
	shapeIndicator.Clear()
	fmt.Fprintln(shapeIndicator, "Playing as\n     X")

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Mouse = true
	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
