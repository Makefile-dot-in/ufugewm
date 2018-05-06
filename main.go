package main

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xwindow"

	"github.com/alecthomas/kingpin.v2"

	"log"
	"os"
)

var (
	app          = kingpin.New("ufugewm", "A window manager written in Go that draws inspiration from ratpoison and vim.")
	displayflag  = app.Flag("display", "Specifies display to use. Defaults to DISPLAY environment variable").Envar("DISPLAY").String()
	Display      = xgbutil.NewConnDisplay(displayflag)
	RootWin         = display.RootWin()
)

const (
	ModeFloating = iota
	ModeTiling   = iota
)

func main() {
	InitError()
}
