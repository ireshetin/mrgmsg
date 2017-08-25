package main

import (
	//"fmt"
	//"io"
	//"io/ioutil"

	"net"
	//"strings"
	"fmt"

	"encoding/json"

	"github.com/jroimartin/gocui"
)

type tParamsLoginStruct struct {
	Username string
	Password string
	Icon     string
}

var data tParamsLoginStruct

var l string

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

// layout to enter new Password
func layoutForgotPassword(g *gocui.Gui) error {

	maxX, maxY := g.Size()

	if v, err := g.SetView("newPassword", maxX/2-20, maxY/2-5, maxX/2+20, maxY/2+4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Editable = false
		v.SelBgColor = gocui.ColorWhite
		v.Title = "Enter your new Password"
	}
	if v, err := g.SetView("Username", maxX/2-18, maxY/2-4, maxX/2+18, maxY/2-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.Wrap = true
		v.Title = "Username"
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
	}

	if v, err := g.SetView("passwordForgot", maxX/2-18, maxY/2, maxX/2+18, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.Wrap = true
		v.Title = "Password"

		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorGreen
	}

	if v, err := g.SetView("but1", maxX/2-10, maxY/2+6, maxX/2+10, maxY/2+8); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlue
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Login")
	}

	return nil
}

// Login
func layoutLogin(g *gocui.Gui) error {

	maxX, maxY := g.Size()

	if v, err := g.SetView("Authtorization", maxX/2-20, maxY/2-5, maxX/2+20, maxY/2+4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Editable = false
		v.Title = "Authtorization"

	}

	if v, err := g.SetView("Username", maxX/2-18, maxY/2-4, maxX/2+18, maxY/2-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.Wrap = true
		v.Title = "Username"
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
	}

	if v, err := g.SetView("Password", maxX/2-18, maxY/2, maxX/2+18, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.Wrap = true
		v.Title = "Password"

		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorGreen
	}

	if v, err := g.SetView("but1", maxX/2-20, maxY/2+6, maxX/2+20, maxY/2+11); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlue
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Login")
		fmt.Fprintln(v, "Register")
		fmt.Fprintln(v, "Forgot Password?")
		fmt.Fprintln(v, "Exit")
	}

	return nil
}

func layoutLogon(g *gocui.Gui) error {

	maxX, maxY := g.Size()

	if v, err := g.SetView("back", 0, 0, 10, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorWhite
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Exit")

	}
	if v, err := g.SetView("profile", 11, 3, maxX/2-1, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorMagenta
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, data.Username)
	}

	if v, err := g.SetView("Icon", 0, 3, 10, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, data.Icon)
	}
	if v, err := g.SetView("ListFriends", 11, 7, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Editable = false
		v.BgColor = gocui.ColorWhite
		v.Title = "Friends"
		// cсписок
		fmt.Fprintln(v, "Katya\nPetya\nSasha\nVasya\nNastya\nMasha\nVanya")

	}

	if v, err := g.SetView("ListIcons", 0, 7, 10, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = false
		v.Editable = false
		v.BgColor = gocui.ColorWhite
		fmt.Fprintln(v, " :) \n༼ つ ◕_◕ ༽つ\n(づ｡◕‿‿◕｡)づ\n(◕‿◕✿)\n◉_◉\n(｡◕‿‿◕｡)\n(｡◕‿◕｡)\n(ʘᗩʘ')")

	}

	if v, err := g.SetView("nameFriend", maxX/2, 3, maxX-1, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorMagenta

		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Choose a friend")

	}

	if v, err := g.SetView("history", maxX/2, 7, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.BgColor = gocui.ColorWhite
		v.Wrap = true
		v.Title = "Diolog"
		// v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		//fmt.Fprintln(v, "{Nastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\n")
	}

	return nil
}

func layoutHistory(g *gocui.Gui) error {

	//maxX, maxY := g.Size()
	maxX, maxY := g.Size()

	if v, err := g.SetView("back", 0, 0, 10, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorWhite
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Exit")

	}
	if v, err := g.SetView("profile", 11, 3, maxX/2-1, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorMagenta
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, data.Username)
	}

	if v, err := g.SetView("Icon", 0, 3, 10, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, data.Icon)
	}

	if v, err := g.SetView("ListFriends", 11, 7, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Editable = false
		v.BgColor = gocui.ColorWhite
		v.Title = "Friends"
		fmt.Fprintln(v, "Katya\nPetya\nSasha\nVasya\nNastya\nMasha\nVanya")

	}

	if v, err := g.SetView("ListIconsHistory", 0, 7, 10, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = false
		v.Editable = false
		v.BgColor = gocui.ColorWhite
		fmt.Fprintln(v, " :) \n༼ つ ◕_◕ ༽つ\n(づ｡◕‿‿◕｡)づ\n(◕‿◕✿)\n◉_◉\n(｡◕‿‿◕｡)\n(｡◕‿◕｡)")

	}

	if v, err := g.SetView("history", maxX/2, 7, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.BgColor = gocui.ColorWhite
		v.Wrap = true
		v.Title = "Dialog"
		// v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		if l == "Katya" {
			fmt.Fprintln(v, "{Nastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\nNastya, [24.08.17 20:00]\n Привет\nKatya, [24.08.17 20:00]\n Привет!\nNastya, [24.08.17 20:00]\n Как дела?\nKastya, [24.08.17 20:00]\n Хорошо, а твои?\n")
		}
	}

	if v, err := g.SetView("nameFriend", maxX/2, 3, maxX-1, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorMagenta

		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, l)

	}

	return nil
}

func layoutListAuth(g *gocui.Gui) error {

	maxX, maxY := g.Size()

	if v, err := g.SetView("listAuth", 0, 3, maxX-1, 11); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Editable = false
		v.Title = "Profile"

	}

	if v, err := g.SetView("Icon", 2, 4, 15, 7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, data.Icon)
	}

	if v, err := g.SetView("iconEdit", 2, 8, 15, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.Wrap = true
		//v.SelBgColor = gocui.ColorMagenta
		v.SelFgColor = gocui.ColorBlack
		//fmt.Fprintln(v, "Enter")
	}

	if v, err := g.SetView("save", 16, 8, maxX/2, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.Wrap = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Enter CtrlS to save")
	}

	if v, err := g.SetView("profile", 16, 4, maxX/2, 7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorMagenta
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, data.Username)
	}

	if v, err := g.SetView("backToLogon", 0, 0, 10, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorWhite
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "back")

	}
	if v, err := g.SetView("ListAuth", 0, 12, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Editable = false
		v.BgColor = gocui.ColorWhite
		v.Title = "List of authorization"
		fmt.Fprintln(v, "02.10.2017 16:30\n02.10.2017 16:36")

	}

	return nil
}

func readBuffer(g *gocui.Gui, v *gocui.View) error {

	if v.Name() == "Username" {
		data.Username = v.Buffer()
		//v.Clear()
	}

	if v.Name() == "Password" {
		data.Password = v.Buffer()
		//v.Clear()
	}

	if v.Name() == "newPassword" {
		data.Password = v.Buffer()
		//v.Clear()
		// передать новый пароль серверу
	}
	if v.Name() == "iconEdit" {

		data.Icon = v.Buffer()
		//v.Clear()
		// передать новый пароль серверу
	}

	return nil
}

func delViews(g *gocui.Gui, v *gocui.View, s string) error {

	if err := g.DeleteView(s); err != nil {
		return err
	}
	//fmt.Print("\n del - ", s)
	return nil

}

func getListAuthUI(g *gocui.Gui, v *gocui.View) error {

	g.SetManagerFunc(layoutListAuth)
	if err := keybindings(g); err != nil {
		//log.Panicln(err)
	}
	return nil
}

func getLogonUI(g *gocui.Gui, v *gocui.View) error {

	go delViews(g, v, "but1")
	go delViews(g, v, "Authrorization")
	go delViews(g, v, "Username")
	go delViews(g, v, "Password")

	g.SetManagerFunc(layoutLogon)

	if err := keybindings(g); err != nil {
		//log.Panicln(err)
	}

	/*if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}*/

	// передача Username && Password серверу

	//fmt.Print("login")
	return nil
}

func keybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlS, gocui.ModNone, readBuffer); err != nil {
		return err
	}

	if err := g.SetKeybinding("but1", gocui.MouseLeft, gocui.ModNone, getMethod); err != nil {
		return err
	}

	if err := g.SetKeybinding("but2", gocui.MouseLeft, gocui.ModNone, getLoginUI); err != nil {
		return err
	}

	if err := g.SetKeybinding("ListFriends", gocui.MouseLeft, gocui.ModNone, getFriendDialog); err != nil {
		return err
	}

	for _, n := range []string{"Username", "Password", "passwordForgot", "iconEdit"} {
		if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, showMsg); err != nil {
			return err
		}
	}

	if err := g.SetKeybinding("back", gocui.MouseLeft, gocui.ModNone, getLoginUI); err != nil {
		return err
	}

	if err := g.SetKeybinding("backToLogon", gocui.MouseLeft, gocui.ModNone, getLogonUI); err != nil {
		return err
	}

	if err := g.SetKeybinding("Icon", gocui.MouseLeft, gocui.ModNone, getListAuthUI); err != nil {
		return err
	}

	if err := g.SetKeybinding("ListFriends", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("ListFriends", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	if err := g.SetKeybinding("history", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("history", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	return nil
}

func getFriendDialog(g *gocui.Gui, v *gocui.View) error {

	//var l string
	var err error

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	// l - имя друга

	g.SetManagerFunc(layoutHistory)

	if err := keybindings(g); err != nil {
		//log.Panicln(err)
	}

	return nil
}

func showMsg(g *gocui.Gui, v *gocui.View) error {

	//fmt.Println(v.Name())
	//var l string
	var err error

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	//if l, err = v.Line(cy); err != nil {
	//	l = ""
	//}
	//if err == nil {
	//log.Printf("error cursor %s", l)
	//}
	_, _ = v.Line(cy)
	return err
}

func getRegisterUI(g *gocui.Gui, v *gocui.View) error {

	g.SetManagerFunc(layoutLogon)

	if err := keybindings(g); err != nil {
		//log.Panicln(err)
	}

	//if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
	//	log.Panicln(err)
	//}

	return nil
}

func getForgotPasswordUI(g *gocui.Gui, v *gocui.View) error {

	/*go delViews(g,v,"but1")
	go delViews(g,v,"Authrorization")
	go delViews(g,v,"Username")
	go delViews(g,v,"Password")
	*/
	g.SetManagerFunc(layoutForgotPassword)

	if err := keybindings(g); err != nil {
		//log.Panicln(err)
	}

	//if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
	//	log.Panicln(err)
	//}

	return nil
}

func getLoginUI(g *gocui.Gui, v *gocui.View) error {

	g.SetManagerFunc(layoutLogin)

	if err := keybindings(g); err != nil {
		//log.Panicln(err)
	}

	//if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
	//	log.Panicln(err)
	//}

	return nil
}

func getMethod(g *gocui.Gui, v *gocui.View) error {

	var l string
	var err error

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	if l == "Login" {
		var r tRequestLogin

		r.Command = "login"

		r.Params = data
		dataToSend, _ := json.Marshal(r)

		sendToServer(dataToSend)

		getLogonUI(g, v)

		// передать данные о пользователе серверу
	}

	if l == "Register" {
		getRegisterUI(g, v)
		// передать данные о пользователе серверу
	}
	if l == "Forgot Password?" {
		getForgotPasswordUI(g, v)
		// передать данные о пользователе серверу
	}
	if l == "Exit" {
		return gocui.ErrQuit
	}
	return nil
}

type tRequestLogin struct {
	Command string             `json:"command"`
	Params  tParamsLoginStruct `json:"params"`
}

func sendToServer(msg []byte) {
	fmt.Fprintf(network.conn, string(msg))
}

func login() {

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		//log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(layoutLogin)

	if err := keybindings(g); err != nil {
		//log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		//log.Panicln(err)
	}
}

type Network struct {
	conn net.Conn
}

var network Network

func main() {

	data.Password = "aaaaa"

	network.conn, _ = net.Dial("tcp", "127.0.0.1:9001")

	data.Username = "Vasya"
	data.Password = "1234"
	data.Icon = "༼ つ ◕_◕ ༽つ"

	login()

	//g.Close()

	/*
		fmt.Print("Username:", data.Username)
		fmt.Printf("\n")
		fmt.Print("Password:", data.Password)
	*/
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
