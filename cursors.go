package main

import (
	"strings"

	"github.com/jroimartin/gocui"
)

func cursorUp(v *gocui.View, step int) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()
		cystep := cy - step
		oystep := oy - step

		if cystep < 0 {
			cystep = 0
		}
		if cy > 0 {
			if err := v.SetCursor(cx, cystep); err != nil {
				return err
			}
		}

		if oystep < 0 {
			oystep = 0
		}
		if oy > 0 && cy == 0 {
			if err := v.SetOrigin(ox, oystep); err != nil {
				return err
			}
		}
	}

	return nil
}

//Moves cursor in list one line down
func cursorDown(v *gocui.View, step int) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()
		cystep := cy + step
		oystep := oy + step
		maxLen := len(strings.Split(v.ViewBuffer(), "\n"))
		if cy+oy < (maxLen - 3) {
			if err := v.SetCursor(cx, cystep); err != nil {
				if err := v.SetOrigin(ox, oystep); err != nil {
					return nil
				}
			}
		}
	}
	return nil
}
