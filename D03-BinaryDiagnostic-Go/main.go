package main
import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func compareBits(numbers map[string]struct{}, col int) (byte, byte) {
    ones := 0
    zeros := 0
    for n := range numbers {
        if n[col] == '1' {
            ones++
        } else {
            zeros++
        }
    }
    if ones >= zeros {
        return '1', '0'
    } else {
        return '0', '1'
    }
}

func main() {
    file, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Only useful for Part II
    numbers := map[string]struct{} {}
    
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    line := scanner.Text()
    numbers[line] = struct{} {}
    ones := make([]int, len(line))
    length := 0
    for {
        line = scanner.Text()
        numbers[line] = struct{} {}
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
    fmt.Printf("(PI) The Gamma Rate is %b; The Epsilon Rate is %b\n" +
               "Hence the power consumption: %d\n",
                gamma, epsilon, gamma * epsilon)

    numbersO2 := numbers
    numbersCO2 := map[string]struct{} {}
    for n, r := range numbers {
        numbersCO2[n] = r
    }

    for i := 0; i < len(ones); i++ {
        mcb, _ := compareBits(numbersO2, i)

        for n := range numbersO2 {
            if n[i] != mcb {
                delete(numbersO2, n)
            }
        }

        if len(numbersO2) == 1 {
            break
        }
              
    }

    for i := 0; i < len(ones); i++ {
        _, lcb := compareBits(numbersCO2, i)

        for n := range numbersCO2 {
            if n[i] != lcb {
                delete(numbersCO2, n)
            }
        }

        if len(numbersCO2) == 1 {
            break
        }
              
    }

    generator := uint64(0)
    for n := range numbersO2 {
        generator, err = strconv.ParseUint(n, 2, 64)
        fmt.Printf("The O2 generator rating is %s\n", n)
    }

    scrubber := uint64(0)
    for n := range numbersCO2 {
        scrubber, err = strconv.ParseUint(n, 2, 64)
        fmt.Printf("The CO2 scrubber rating is %s\n", n)
    }

    fmt.Printf("(PII) Hence the life support rating of the submarine is %d\n",
               scrubber * generator) 
    
}
