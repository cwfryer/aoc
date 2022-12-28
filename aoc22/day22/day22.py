from lark import Lark, Transformer
from collections import namedtuple
from enum import IntEnum
import numpy as np

Point = namedtuple("Point", "row col")


class Direction(IntEnum):
    up = 0
    right = 1
    down = 2
    left = 3


grammar = r"""
steps : [distance | turn]+
distance : NUMBER+
turn : "R" -> right 
     | "L" -> left
%import common.NUMBER -> NUMBER
"""


class TreeToSteps(Transformer):
    steps = list

    def distance(self, other):
        return int(other[0])

    def right(self, other):
        return "R"

    def left(self, other):
        return "L"


def printMap(map):
    minRow = 0
    maxRow = max([row for row, col in map.keys()])
    minCol = min([col for row, col in map.keys()])
    maxCol = max([col for row, col in map.keys()])

    output = [list(" " * (maxCol + 1)) for r in range(maxRow + 1)]
    for row, col in map:
        output[row][col] = map[Point(row, col)]

    for row in output:
        print("".join(row))


data = open("inputs/input22.txt").read()
mapInput, path = data.split("\n\n")
parser = Lark(grammar, start="steps")
tree = parser.parse(path)
steps = TreeToSteps().transform(tree)

map = {}
mapInput = mapInput.split("\n")
for r, row in enumerate(mapInput):
    for c, value in enumerate(row):
        if value != " ":
            map[Point(r, c)] = value

minRow = 0
minCol = min([col for row, col in map.keys() if row == 0])

loc = Point(minRow, minCol)
dir = Direction.right
map[loc] = ">"

for step in steps:
    if isinstance(step, int):
        if dir == Direction.right:
            for i in range(step):
                newLoc = Point(loc.row, loc.col + 1)
                if newLoc in map and map[newLoc] != "#":
                    map[newLoc] = ">"
                    loc = newLoc
                    continue
                elif newLoc in map and map[newLoc] == "#":
                    continue
                elif not newLoc in map:
                    minCol = min([col for row, col in map if row == loc.row])
                    newLoc = Point(loc.row, minCol)
                    if map[newLoc] != "#":
                        map[newLoc] = ">"
                        loc = newLoc
                        continue
        elif dir == Direction.left:
            for i in range(step):
                newLoc = Point(loc.row, loc.col - 1)
                if newLoc in map and map[newLoc] != "#":
                    map[newLoc] = "<"
                    loc = newLoc
                    continue
                elif newLoc in map and map[newLoc] == "#":
                    continue
                elif not newLoc in map:
                    maxCol = max([col for row, col in map if row == loc.row])
                    newLoc = Point(loc.row, maxCol)
                    if map[newLoc] != "#":
                        map[newLoc] = "<"
                        loc = newLoc
                        continue
        elif dir == Direction.down:
            for i in range(step):
                newLoc = Point(loc.row + 1, loc.col)
                if newLoc in map and map[newLoc] != "#":
                    map[newLoc] = "v"
                    loc = newLoc
                    continue
                elif newLoc in map and map[newLoc] == "#":
                    continue
                elif not newLoc in map:
                    minRow = min([row for row, col in map if col == loc.col])
                    newLoc = Point(minRow, loc.col)
                    if map[newLoc] != "#":
                        map[newLoc] = "v"
                        loc = newLoc
                        continue
        elif dir == Direction.up:
            for i in range(step):
                newLoc = Point(loc.row - 1, loc.col)
                if newLoc in map and map[newLoc] != "#":
                    map[newLoc] = "^"
                    loc = newLoc
                    continue
                elif newLoc in map and map[newLoc] == "#":
                    continue
                elif not newLoc in map:
                    maxRow = max([row for row, col in map if col == loc.col])
                    newLoc = Point(maxRow, loc.col)
                    if map[newLoc] != "#":
                        map[newLoc] = "^"
                        loc = newLoc
                        continue
    else:  # it's a turn
        if step == "R":
            dir = (dir + 1) % 4
        else:
            dir = (dir - 1) % 4

# add final location
match dir:
    case Direction.right:
        map[loc] = ">"
    case Direction.left:
        map[loc] = "<"
    case Direction.up:
        map[loc] = "^"
    case Direction.down:
        map[loc] = "v"

# Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^).
# The final password is the sum of 1000 times the row, 4 times the column, and the facing.
facing = 0
match dir:
    case Direction.right:
        facing = 0
    case Direction.left:
        facing = 2
    case Direction.up:
        facing = 3
    case Direction.down:
        facing = 1

password = (1000 * (loc.row + 1)) + (4 * (loc.col + 1)) + facing
print(f"part 1: {password}")

# part 2

data = open("inputs/input22.txt").read()
mapInput, path = data.split("\n\n")
parser = Lark(grammar, start="steps")
tree = parser.parse(path)
steps = TreeToSteps().transform(tree)

NewSide = namedtuple("NewSide", "sideNumber direction newRow newCol")

sideLength = 50
rules = [
    None,
    {  # rule1
        Direction.up: NewSide(
            6, Direction.right, lambda row, col: col, lambda row, col: 0
        ),
        Direction.down: NewSide(
            3, Direction.down, lambda row, col: 0, lambda row, col: col
        ),
        Direction.left: NewSide(
            4,
            Direction.right,
            lambda row, col: sideLength - 1 - row,
            lambda row, col: 0,
        ),
        Direction.right: NewSide(
            2, Direction.right, lambda row, col: row, lambda row, col: 0
        ),
    },
    {  # rule2
        Direction.up: NewSide(
            6, Direction.up, lambda row, col: sideLength - 1, lambda row, col: col
        ),
        Direction.down: NewSide(
            3, Direction.left, lambda row, col: col, lambda row, col: sideLength - 1
        ),
        Direction.left: NewSide(
            1, Direction.left, lambda row, col: row, lambda row, col: sideLength - 1
        ),
        Direction.right: NewSide(
            5,
            Direction.left,
            lambda row, col: sideLength - 1 - row,
            lambda row, col: sideLength - 1,
        ),
    },
    {  # rule3
        Direction.up: NewSide(
            1, Direction.up, lambda row, col: sideLength - 1, lambda row, col: col
        ),
        Direction.down: NewSide(
            5, Direction.down, lambda row, col: 0, lambda row, col: col
        ),
        Direction.left: NewSide(
            4, Direction.down, lambda row, col: 0, lambda row, col: row
        ),
        Direction.right: NewSide(
            2, Direction.up, lambda row, col: sideLength - 1, lambda row, col: row
        ),
    },
    {  # rule4
        Direction.up: NewSide(
            3, Direction.right, lambda row, col: col, lambda row, col: 0
        ),
        Direction.down: NewSide(
            6, Direction.down, lambda row, col: 0, lambda row, col: col
        ),
        Direction.left: NewSide(
            1,
            Direction.right,
            lambda row, col: sideLength - 1 - row,
            lambda row, col: 0,
        ),
        Direction.right: NewSide(
            5, Direction.right, lambda row, col: row, lambda row, col: 0
        ),
    },
    {  # rule5
        Direction.up: NewSide(
            3, Direction.up, lambda row, col: sideLength - 1, lambda row, col: col
        ),
        Direction.down: NewSide(
            6, Direction.left, lambda row, col: col, lambda row, col: sideLength - 1
        ),
        Direction.left: NewSide(
            4, Direction.left, lambda row, col: row, lambda row, col: sideLength - 1
        ),
        Direction.right: NewSide(
            2,
            Direction.left,
            lambda row, col: sideLength - 1 - row,
            lambda row, col: sideLength - 1,
        ),
    },
    {  # rule6
        Direction.up: NewSide(
            4, Direction.up, lambda row, col: sideLength - 1, lambda row, col: col
        ),
        Direction.down: NewSide(
            2, Direction.down, lambda row, col: 0, lambda row, col: col
        ),
        Direction.left: NewSide(
            1, Direction.down, lambda row, col: 0, lambda row, col: row
        ),
        Direction.right: NewSide(
            5, Direction.up, lambda row, col: sideLength - 1, lambda row, col: row
        ),
    },
]


def printSide(sideMap):
    print("-" * (sideLength + 2))
    for row in sideMap:
        output = "|" + "".join(row) + "|"
        print(output)
    print("-" * (sideLength + 2))


# create the six maps
lines = mapInput.split("\n")
lines = [list(line.ljust(150)) for line in lines]
allSides = np.array(lines, dtype=np.unicode_)

sides = [
    None,
    allSides[0:50, 50:100],
    allSides[0:50, 100:150],
    allSides[50:100, 50:100],
    allSides[100:150, 0:50],
    allSides[100:150, 50:100],
    allSides[150:200, 0:50],
]


def isWall(sideNumber, row, col):
    value = sides[sideNumber][row, col]
    return value == "#"


Location = namedtuple("Location", "side direction row, col")


def getNextMove(sideCurrent, dirCurrent, rowCurrent, colCurrent):
    rule = rules[sideCurrent][dirCurrent]
    if dirCurrent == Direction.up:
        rowNext = rowCurrent - 1
        colNext = colCurrent
        dirNext = dirCurrent
        sideNext = sideCurrent
        if rowNext < 0:
            sideNext = rule.sideNumber
            dirNext = rule.direction
            rowNext = rule.newRow(rowCurrent, colCurrent)
            colNext = rule.newCol(rowCurrent, colCurrent)
    elif dirCurrent == Direction.down:
        rowNext = rowCurrent + 1
        colNext = colCurrent
        dirNext = dirCurrent
        sideNext = sideCurrent
        if rowNext == sideLength:
            sideNext = rule.sideNumber
            dirNext = rule.direction
            rowNext = rule.newRow(rowCurrent, colCurrent)
            colNext = rule.newCol(rowCurrent, colCurrent)
    elif dirCurrent == Direction.left:
        rowNext = rowCurrent
        colNext = colCurrent - 1
        dirNext = dirCurrent
        sideNext = sideCurrent
        if colNext < 0:
            sideNext = rule.sideNumber
            dirNext = rule.direction
            rowNext = rule.newRow(rowCurrent, colCurrent)
            colNext = rule.newCol(rowCurrent, colCurrent)
    elif dirCurrent == Direction.right:
        rowNext = rowCurrent
        colNext = colCurrent + 1
        dirNext = dirCurrent
        sideNext = sideCurrent
        if colNext == sideLength:
            sideNext = rule.sideNumber
            dirNext = rule.direction
            rowNext = rule.newRow(rowCurrent, colCurrent)
            colNext = rule.newCol(rowCurrent, colCurrent)

    if sides[sideNext][rowNext, colNext] != "#":
        return Location(sideNext, dirNext, rowNext, colNext)
    else:
        return Location(sideCurrent, dirCurrent, rowCurrent, colCurrent)


mapChars = {
    Direction.up: "^",
    Direction.down: "v",
    Direction.left: "<",
    Direction.right: ">",
}

loc = Location(1, Direction.right, 0, 0)
for step in steps:
    if isinstance(step, int):
        for i in range(step):
            nextMove = getNextMove(loc.side, loc.direction, loc.row, loc.col)
            sides[loc.side][loc.row, loc.col] = mapChars[loc.direction]
            loc = nextMove
    else:
        if step == "R":
            dirNext = Direction((loc.direction + 1) % 4)
        else:
            dirNext = Direction((loc.direction - 1) % 4)
        loc = Location(loc.side, dirNext, loc.row, loc.col)

# now convert to original coordinates
match loc.side:
    case 1:
        row = loc.row
        col = loc.col + 50
    case 2:
        row = loc.row
        col = loc.col + 100
    case 3:
        row = loc.row + 50
        col = loc.col + 50
    case 4:
        row = loc.row + 100
        col = loc.col
    case 5:
        row = loc.row + 100
        col = loc.col + 50
    case 6:
        row = loc.row + 150
        col = loc.col

facing = 0
match loc.direction:
    case Direction.right:
        facing = 0
    case Direction.left:
        facing = 2
    case Direction.up:
        facing = 3
    case Direction.down:
        facing = 1

password = (1000 * (row + 1)) + (4 * (col + 1)) + facing
print(f"part 2: {password}")
