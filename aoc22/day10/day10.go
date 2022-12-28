package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processInstruction(instrList map[int]int, tick int, hPos int) int {
    newPos := hPos + instrList[tick]
    return newPos
}

func drawCRT(pixel int, hPos int) {
    if pixel % 40 >= hPos - 1 && pixel % 40 <= hPos + 1 {
        fmt.Print("#")
    } else { fmt.Print(".")}

    if (pixel+1) % 40 == 0 {
        fmt.Print("\n")
    }
}

func drawSpritePosition(hPos int) {
    for i:=0;i<240;i++ {
        if i%40 >= hPos -1 && i%40 <= hPos + 1 {
            fmt.Print("#")
        } else {fmt.Print(".")}
        if (i+1) % 40 == 0 {
            fmt.Print("\n")
        }
    }
    fmt.Println()
}

func main() {

	sc := bufio.NewScanner(os.Stdin)

    iMap := make(map[int]int)
    i := 1
    for sc.Scan() {
        switch _,instr,_ := strings.Cut(sc.Text()," "); instr {
        case "":
            iMap[i] += 0
            i++
        default:
            i++
            num,_ := strconv.Atoi(instr)
            i++
            iMap[i] += num
        }
    }

    hPosMap := make(map[int]int)
    hPos := 1
    hPosMap[0] = hPos
    for i:=1;i<250;i++ {
        hPos = processInstruction(iMap,i,hPos)
        hPosMap[i] = hPos
    }
    fmt.Println(hPosMap)
    fmt.Println()

    for i:=0;i<240;i++ {
        hPos = hPosMap[i+1]
        drawCRT(i,hPos)
    }
}
	
