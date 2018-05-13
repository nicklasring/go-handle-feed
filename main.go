package main

import (
	"fmt"
	"log"
	"path"

	"github.com/jroimartin/gocui"
	"github.com/mmcdole/gofeed"
)

var links map[string]string

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("main", 0, 0, maxX-1, maxY-5); err != nil {
		links = make(map[string]string)
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView("main"); err != nil {
			return err
		}

		settings := getSettings()

		v.Editable = false
		v.Highlight = true
		v.Wrap = true
		v.BgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Title = fmt.Sprintf(" | %s - Feed | ", settings.API.Name)
		g.Cursor = true

		fp := gofeed.NewParser()

		feed, _ := fp.ParseURL(settings.API.URL)

		for _, item := range feed.Items {
			links[item.Title] = item.Link
			fmt.Fprintf(v, "%s\n", item.Title)
		}
	}
	if v2, err := g.SetView("bottom", 0, maxY-5, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v2.Editable = false
		v2.Title = "Status"
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
		err = downloadFile(links[l])
		if err != nil {
			log.Fatal(err)
		}

		fileName := path.Base(links[l])
		err = transferFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		transferMessage := fmt.Sprintf("Sucessfully transferred: %s", fileName)
		for _, view := range g.Views() {
			if view.Name() == "bottom" {
				view.Clear()
				view.Write([]byte(transferMessage))
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
	setKeyBindings(g)

	/* Main loop */
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
