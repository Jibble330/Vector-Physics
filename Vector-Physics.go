package main

import (
    "os"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/gdamore/tcell"
)

const RAD_TO_DEG = 180/math.Pi

var (
    screen tcell.Screen
    err error

    defStyle = tcell.StyleDefault
    invertStyle = tcell.StyleDefault.Reverse(true)
    Copied Vector
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

func (v *Vector) Add(u Vector) Vector {
    return XY(v.X+u.X, v.Y+u.Y)
}

func (v *Vector) Scaled(mult float64) Vector {
    return Vector{v.X*mult, v.Y*mult, v.Mag*mult, v.Deg}
}

func listen(events chan tcell.Event) {
    for {
        ev := screen.PollEvent()
        events <- ev
        switch ev := ev.(type) {
        case *tcell.EventKey:
            if ev.Key() == tcell.KeyESC {
                Exit()
            }
        }
    }
}

func WriteString(input string, x, y int, style tcell.Style) {
    lines := strings.Split(input, "\n")

    for lineIndex, line := range lines {
        for charIndex, char := range line {
            X := charIndex+x
            Y := lineIndex+y

            screen.SetContent(X, Y, char, nil, style)
        }
    }
}

func menu(options []string) uint {
    events := make(chan tcell.Event)
    go listen(events)

    var choice uint

    for ev := range events {
        switch ev := ev.(type) {
        case *tcell.EventKey:
            key := ev.Key()
            switch key {
            case tcell.KeyUp:
                if choice > 0 {
                    choice--
                }
            case tcell.KeyDown:
                if choice < uint(len(options))-1 {
                    screen.Clear()
                    choice++
                }
            case tcell.KeyEnter:
                return choice
            }
            centerX, centerY := screen.Size()
            centerX /= 2
            centerY /= 2
            startY := centerY-len(options)
            for index, option := range options {
                startX := centerX-(3+len(option)/2)
                var numStyle tcell.Style
                if choice == uint(index) {
                    numStyle = invertStyle
                } else {
                    numStyle = defStyle
                }
                WriteString(fmt.Sprint(index+1, "."), startX, startY+index, numStyle)
                WriteString(" "+ option, startX+3, startY+index, defStyle)
                screen.Show()
            }
        }
    }
    return choice
}

func Exit() {
    screen.Clear()
    screen.Fini()
    os.Exit(0)
}

func main() {
    screen, err = tcell.NewScreen()
    if err != nil {
        log.Fatal(err)
    }

    if err = screen.Init(); err != nil {
        log.Fatal(err)
    }
    defer Exit()

    screen.SetStyle(defStyle)

    choice := menu([]string{"Add Vectors", "Scale Vectors", "Dot Product", "Cross Product", "Project Vectors"})
    switch choice {
        
    }
}
