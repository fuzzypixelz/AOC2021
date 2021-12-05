package main
import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    line := scanner.Text()
    ones := make([]int, len(line))
    length := 0
    for {
        line = scanner.Text()
        for i := 0; i < len(line); i++ {
            if line[i] == '1' {
                ones[i]++
            }
        }
        length++
        if !scanner.Scan() {
            break
        }
    }
    gamma := 0
    epsilon := 0
    for i, b := range ones {
        shb := 1 << (len(ones) - i - 1)
        if b > length - b {
            gamma |= shb
        } else {
            epsilon |= shb
        }
    }
    fmt.Printf("The Gamma Rate is %b; The Epsilon Rate is %b\n" +
               "Hence the power consumption: %d\n",
                gamma, epsilon, gamma * epsilon)
}
