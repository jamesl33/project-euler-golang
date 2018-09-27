package main

import "fmt"

func is_prime(n int) bool {
    if n <= 1 {
        return false
    } else if n <= 3 {
        return true
    } else if n % 2 == 0 || n % 3 == 0 {
        return false
    }

    i := 5

    for i * i <= n {
        if n % i == 0 || n % (i + 2) == 0 {
            return false
        }

        i += 6
    }

    return true
}

func main() {
    total := 0

    for i := 0; i < 2000000; i++ {
        if is_prime(i) {
            total += i
        }
    }

    fmt.Println(total)
}
