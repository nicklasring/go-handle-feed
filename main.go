package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/mmcdole/gofeed"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("net", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView("net"); err != nil {
			return err
		}
		v.Editable = false
		v.Highlight = true
		v.Wrap = true
		v.BgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Title = " | Go - Handle Feed | "
		g.Cursor = true

		fp := gofeed.NewParser()
		url := "URL"
		feed, _ := fp.ParseURL(url)

		for _, item := range feed.Items {
			fmt.Fprintf(v, "Item: %s\n", item.Title)
		}
	}
	return nil
}

func keyEnter(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		var l string
		var err error
		_, cy := v.Cursor()
		if l, err = v.Line(cy); err != nil {
			l = ""
		}
		log.Print(l)

	}
	return nil
}

//Moves cursor in list one line up
func cursorUp(v *gocui.View, step int) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()
		var oystep int
		var cystep int
		if cy > 0 {
			if cy-step < 0 {
				cystep = 0
			} else {
				cystep = cy - step
			}
			if err := v.SetCursor(cx, cystep); err != nil {
				return err
			}
		}
		if oy > 0 && cy == 0 {
			if oy-step < 0 {
				oystep = 0
			} else {
				oystep = oy - step
			}
			if err := v.SetOrigin(ox, oystep); err != nil {
				return err
			}
		}
	}

	return nil
}

func cursorDown(v *gocui.View, step int) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()
		if cy+oy < (len(strings.Split(v.ViewBuffer(), "\n")) - 3) {
			if err := v.SetCursor(cx, cy+step); err != nil {
				if err := v.SetOrigin(ox, oy+step); err != nil {
					return nil
				}
			}
		}
	}
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

	g.Mouse = false
	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return cursorUp(v, 1)
	}); err != nil {

		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return cursorDown(v, 1)
	}); err != nil {

		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyPgup, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return cursorUp(v, 10)
	}); err != nil {

		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyPgdn, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return cursorDown(v, 10)
	}); err != nil {

		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, keyEnter); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
