package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sum := 0
    for sc.Scan() {
        // fmt.Println(sc.Text(),decode(sc.Text()))
        sum += decode(sc.Text())
    }
    fmt.Println(encode(sum))
}

func decode(in string) int {
    n := float64(0)
    for i:=len(in)-1;i>=0;i-- {
        place := math.Pow(float64(5),float64((len(in)-1) - i))
        switch in[i] {
        case '2':
            n += place * float64(2)
        case '1':
            n += place * float64(1)
        case '0':
            n += place * float64(0)
        case '-':
            n += place * float64(-1)
        case '=':
            n += place * float64(-2)
        }
    }
    return int(n)
}

func encode(in int) string {
    f := float64(in)

    rem := f
    b5 := ""
    for rem > 0 {
        carry := math.Floor(rem/5)
        v := int(rem) % 5
        b5 += fmt.Sprintf("%d",v)
        rem = carry
    }
    // b5 := reverse(outs)
    out := ""
    carry := 0
    for _,c := range b5 {
        i,_ := strconv.Atoi(string(c))
        i += carry
        switch i {
        case 0:
            carry = 0
            out += "0"
        case 1:
            carry = 0
            out += "1"
        case 2:
            carry = 0
            out += "2"
        case 3:
            carry = 1
            out += "="
        case 4:
            carry = 1
            out += "-"
        case 5:
            carry = 1
            out += "0"
        }
    }
    return reverse(out)
}

func reverse(in string) string {
    out := ""
    for _,v := range in {
        out = string(v) + out
    }
    return out
}
