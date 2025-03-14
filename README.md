# Game of Life (Ebiten)

This is an implementation of Conway's Game of Life using the Ebiten game engine in Golang.

## Conway's Game of Life

Conway's Game of Life is a cellular automaton devised by mathematician John Conway. It follows these simple rules:

1. **Any live cell with fewer than two live neighbors dies** (underpopulation).
2. **Any live cell with two or three live neighbors survives**.
3. **Any live cell with more than three live neighbors dies** (overpopulation).
4. **Any dead cell with exactly three live neighbors becomes a live cell** (reproduction).

This simulation follows these rules and allows for user interaction to create custom patterns.

## Features

- Click and drag to toggle cell states.
- Infinite grid behavior (cells wrap around the edges).
- Start/Pause simulation.
- Reset the grid to an empty state.
- Generate a new random seed.
- Displays generation count.

## Installation and Running the Game

### Prerequisites

You need Go installed on your system. If you haven't installed it yet, download and install it from [Go's official site](https://go.dev/).

### Install Ebiten

You need to install the Ebiten library before running the game:

```sh
go get github.com/hajimehoshi/ebiten/v2
```

### Running the Game

Clone the repository and navigate into the directory:

```sh
git clone https://github.com/dhawal-pandya/Game-Of-Life-Ebiten.git
cd Game-Of-Life-Ebiten
```

Run the game using:

```sh
go run main.go
```

### Building the Game

To compile the game as an executable:

```sh
go build -o gameoflife
```

### Platform-Specific Instructions

#### Linux / macOS

```sh
./gameoflife
```

#### Windows

```sh
gameoflife.exe
```

## License

This project is open-source under the MIT License. Feel free to modify and distribute it.

---

Enjoy experimenting with Conway's Game of Life!
