package gtris

import (

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) transferPieceToGameZone() {
	piece := g.currentPiece
	piecePos := g.piecePosition
	for dy, row := range piece.Blocks {
		for dx, value := range row {
			if value != pieceBlockMarker {
				continue
			}

			gameZonePos := &Position{
				X: piecePos.X + dx,
				Y: piecePos.Y + dy,
			}

			g.gameZone[gameZonePos.Y][gameZonePos.X] = piece.Image
		}
	}
    g.pieceCount ++ 
}

func (g *Game) insideGameZone(deltaPos Position) bool {
	piecePos := *g.piecePosition
	piecePos.Add(deltaPos)
	return g.pieceInsideGameZone(g.currentPiece, piecePos)
}

func (g *Game) pieceInsideGameZone(piece *Piece, piecePos Position) bool {

	for dy, row := range piece.Blocks {
		for dx, value := range row {
			if value == pieceBlockMarker {
				screenPos := &Position{
					X: piecePos.X + dx,
					Y: piecePos.Y + dy,
				}

				if screenPos.X < 0 || screenPos.X >= int(g.gameZoneSize.Width) {
					return false
				}
				if screenPos.Y < 0 || screenPos.Y >= int(g.gameZoneSize.Height) {
					return false
				}

				if g.gameZone[screenPos.Y][screenPos.X] != nil {
					return false
				}
			}
		}
	}

	return true
}

func (g *Game) fetchNextPiece() {
	getNextPiece := func() *Piece {
		np := g.pieces[0]
		g.pieces = append(g.pieces[1:], np)
		return np
	}

	if g.nextPiece == nil {
		g.nextPiece = getNextPiece()
	}

	g.currentPiece = g.nextPiece
    dy := len(g.currentPiece.Blocks) - 1
	g.piecePosition = &Position{X: 2, Y: int(g.gameZoneSize.Height) - g.maxHeight() - 4 - dy}
	g.nextPiece = getNextPiece()
}

func (g *Game) maxHeight() int {
    for y, row := range g.gameZone {
        for _, cellImage := range row {
            if cellImage != nil {
                return int(g.gameZoneSize.Height) - y
            }
        }
    }
    return 0
}

func (g *Game) checkForLines() int {
	lines := []int{}
	for y, row := range g.gameZone {
		var full = false
		for _, cellImage := range row {
			if cellImage == nil {
				full = false
				break
			}
		}
		if full {
			lines = append(lines, y)
		}
	}

	for _, y := range lines {
		emptyRow := [][]*ebiten.Image{
			make([]*ebiten.Image, g.gameZoneSize.Width),
		}
		g.gameZone = append(append(emptyRow, g.gameZone[0:y]...), g.gameZone[(y+1):]...)
	}

	return len(lines)
}

func (g *Game) rotatePiece() *Piece {
	newPiece := *g.currentPiece

	newPiece.Blocks = make([][]int, len(g.currentPiece.Blocks[0]))
	for y := range newPiece.Blocks {
		newPiece.Blocks[y] = make([]int, len(g.currentPiece.Blocks))
	}

	for y, row := range g.currentPiece.Blocks {
		for x := range row {
			newPiece.Blocks[x][len(g.currentPiece.Blocks)-1-y] = g.currentPiece.Blocks[y][x]
		}
	}

	return &newPiece
}
