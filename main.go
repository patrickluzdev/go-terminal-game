package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"time"

	"github.com/charmbracelet/x/term"
)

const (
	Width  = 50
	Height = 20
)

type Point struct {
	X, Y int
}

type GameState struct {
	Snake     []Point
	Direction Point
	Apple     Point
	Score     int
}

func readInput(ch chan<- byte) {
	buf := make([]byte, 1)
	for {
		os.Stdin.Read(buf)
		ch <- buf[0]
	}
}

func spawnApple(state *GameState) {
	for {
		candidate := Point{rand.Intn(Width), rand.Intn(Height)}
		occupied := slices.Contains(state.Snake, candidate)
		if !occupied {
			state.Apple = candidate
			return
		}
	}
}

func update(state *GameState) bool {
	head := state.Snake[0]
	newHead := Point{
		X: head.X + state.Direction.X,
		Y: head.Y + state.Direction.Y,
	}

	if newHead.X < 0 || newHead.X >= Width || newHead.Y < 0 || newHead.Y >= Height {
		return true
	}

	if slices.Contains(state.Snake[:len(state.Snake)-1], newHead) {
		return true
	}

	state.Snake = append([]Point{newHead}, state.Snake...)

	if newHead == state.Apple {
		state.Score++
		spawnApple(state)
	} else {
		state.Snake = state.Snake[:len(state.Snake)-1]
	}

	return false
}

func render(state *GameState) {
	fmt.Print("\033[H\033[2J")

	fmt.Print("#")
	for range Width {
		fmt.Print("#")
	}
	fmt.Print("#\r\n")

	for y := range Height {
		fmt.Print("|")
		for x := range Width {
			cell := Point{x, y}
			switch {
			case slices.Contains(state.Snake, cell):
				fmt.Print("O")
			case state.Apple == cell:
				fmt.Print("*")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Print("|\r\n")
	}

	fmt.Print("#")
	for range Width {
		fmt.Print("#")
	}
	fmt.Print("#\r\n")

	fmt.Printf("Score: %d\r\n", state.Score)
}

func main() {
	oldState, err := term.MakeRaw(uintptr(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(uintptr(os.Stdin.Fd()), oldState)

	input := make(chan byte)
	go readInput(input)

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	state := &GameState{
		Snake:     []Point{{X: Width / 2, Y: Height / 2}},
		Direction: Point{1, 0},
	}
	spawnApple(state)

	fmt.Print("\033[?1049h")
	defer fmt.Print("\033[?1049l")
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")
	for {
		select {
		case <-ticker.C:
			if update(state) {
				fmt.Print("\033[H\033[2J")
				fmt.Printf("Game Over!\r\nScore: %d\r\n\r\n'q' para sair, qualquer outra para jogar de novo...\r\n", state.Score)
				key := <-input
				if key == 'q' {
					return
				}
				state = &GameState{
					Snake:     []Point{{X: Width / 2, Y: Height / 2}},
					Direction: Point{1, 0},
				}
				spawnApple(state)
			}
			render(state)
		case key := <-input:
			switch key {
			case 'w':
				state.Direction = Point{0, -1}
			case 's':
				state.Direction = Point{0, 1}
			case 'a':
				state.Direction = Point{-1, 0}
			case 'd':
				state.Direction = Point{1, 0}
			case 'q':
				return
			}
		}
	}
}
