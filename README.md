# Go Terminal Snake Game

A classic Snake game that runs directly in the terminal, written in Go.

## Demo

```
####################################################
|                                                  |
|                                                  |
|                    *                             |
|                                                  |
|              OOOOO                               |
|                   O                              |
|                   O                              |
|                                                  |
####################################################
Score: 4
```

## Requirements

- Go 1.22 or higher
- A Unix-like terminal (Linux or macOS)

## Installation

```bash
git clone https://github.com/patrickluzdev/go-terminal-game.git
cd go-terminal-game
go run main.go
```

## Controls

| Key | Action        |
|-----|---------------|
| `w` | Move up       |
| `s` | Move down     |
| `a` | Move left     |
| `d` | Move right    |
| `q` | Quit the game |

## How It Works

- The snake starts at the center of a 50×20 grid moving right
- Eat the apple (`*`) to grow and earn points
- The game ends if the snake hits a wall or itself
- After game over, press any key to restart or `q` to quit

## Dependencies

- [`charmbracelet/x/term`](https://github.com/charmbracelet/x) — raw terminal mode for real-time keyboard input

## Project Structure

```
go-terminal-game/
├── main.go     # Game logic, rendering, and input handling
├── go.mod
└── go.sum
```

## License

MIT
