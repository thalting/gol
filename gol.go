// This software is in the public domain.

package main

import (
    "bufio"
    "math"
    "fmt"
    "io"
    "os"
)

func rainbow(i int) (int, int, int) {
    red := int(math.Sin(0.1 * float64(i) + 0) * 127 + 128)
    green := int(math.Sin(0.1 * float64(i) + 2 * math.Pi / 3) * 127 + 128)
    blue := int(math.Sin(0.1 * float64(i) + 4 * math.Pi / 3) * 127 + 128)

    return red, green, blue
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    i := 0
    for {
        input, err := reader.ReadByte()
        if err != nil && err == io.EOF {
            break
        }
        r, g, b := rainbow(i)
        fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, input)
        i++
    }
}
