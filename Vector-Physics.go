package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/gdamore/tcell"
)

const RAD_TO_DEG = 180/math.Pi

var (
    screen tcell.Screen
    err error
)

type Vector struct {
    X, Y, Mag, Deg float64
}

func XY(X, Y float64) Vector {
    Deg := math.Atan2(Y, X) * RAD_TO_DEG
    Mag := math.Hypot(X, Y)
    return Vector{X, Y, Mag, Deg}
}

func MagDeg(Mag, Deg float64, North bool) Vector {
    if !North {
        Deg = 180-Deg
    }
    if Deg < 0 {
        Deg = Deg + (360 * math.Ceil(Deg/360))
    }
    X := math.Cos(Deg) * RAD_TO_DEG
    Y := math.Sin(Deg) * RAD_TO_DEG
    return Vector{X, Y, Mag, Deg}
}

func (v *Vector) String() string {
    return fmt.Sprintf(`X: %v, Y: %v
Magnitude: %v
Degrees: %v`, v.X, v.Y, v.Mag, v.Deg)
}

func listen(events chan tcell.Event) {
    for {
        ev := screen.PollEvent()
        events <- ev
    }
}

func WriteString(input string, pos [2]int, bg, fg tcell.Color) {
    lines := strings.Split(input, "\n")

    for lineIndex, line := range lines {
        for charIndex, char := range line {
            X := charIndex+pos[0]
            Y := lineIndex+pos[1]
            style := tcell.StyleDefault.Background(bg).Foreground(fg)

            screen.SetContent(X, Y, char, nil, style)
        }
    }
}

func menu[T any](options []T) uint {
    events := make(chan tcell.Event)
    go listen(events)

    var choice uint

    for ev := range events {
        switch ev.(type) {
        case *tcell.EventKey:

        }
    }
    return choice
}

func main() {
    screen, err = tcell.NewScreen()
    if err != nil {
        log.Fatal(err)
    }

    if err = screen.Init(); err != nil {
        log.Fatal(err)
    }
    defer screen.Fini()

    defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
    screen.SetStyle(defStyle)

    screen.SetContent(0, 0, 'H', nil, defStyle)
    
    time.Sleep(time.Second*5)
}