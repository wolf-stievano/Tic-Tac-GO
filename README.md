# Tic Tac GO

**It's tic tac toe but on steroids and made with Go!**

## Overview

Tic Tac GO is an advanced version of the classic Tic Tac Toe game. Instead of just one grid, you get to play on 9 mini grids, all contained within a larger 3x3 grid. The twist? Each move you make in one of the mini grids determines the next mini grid your opponent has to play on. The game brings a strategic depth that challenges your foresight and planning skills in new ways.

## How to Play

1. **Game Setup:**
   - The game board consists of 9 mini 3x3 Tic Tac Toe grids arranged in a 3x3 larger grid.
   - Players take turns playing as X or O.

2. **Gameplay Mechanics:**
   - On your turn, you can place your mark (X or O) in any available cell in the mini grid corresponding to the last move your opponent made.
   - If your opponent's last move was in the bottom-right cell of a mini grid, your next move must be in the corresponding mini grid in the bottom-right of the large grid.
   - If a mini grid is won by a player or results in a tie, that grid is considered "closed," and no more moves can be made there. However, if the corresponding mini grid is already closed, the player can choose any open cell on the entire board for their move.

3. **Winning the Game:**
   - The game is won by the first player to win three mini grids in a row, either horizontally, vertically, or diagonally.
   - If all 9 mini grids are filled and no player has won three in a row, the game ends in a draw.

## Features

- **Terminal-Based Gameplay:** Tic Tac GO is entirely text-based and runs in your terminal, providing a nostalgic and straightforward gaming experience.
- **Intelligent Move Handling:** The game logic ensures that all rules are followed, guiding players to the correct mini grid based on previous moves.
- **Strategic Depth:** With the added layer of choosing the next grid for your opponent, every move counts.

## Installation

To get started with Tic Tac GO, clone this repository and build the project using Go:

```bash
git clone https://github.com/wolf-stievano/Tic-Tac-GO.git
cd Tic-Tac-GO
go build

```

## Running the Game

After building the project, you can start the game by running:

```bash
go run Tic-Tac-GO.go
```

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request with your improvements.

## License

This project is licensed under the MIT License.

