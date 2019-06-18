/* This file is part of project-euler

Copyright (C) 2019, James Lee <jamesl33info@gmail.com>.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>. */

package main

import (
    "fmt"
)

func IsPrime(n int) bool {
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

func EratosthenesSieve(lowerBound, upperBound int) []int {
    sieve := make([]bool, upperBound)

    for i := 0; i < upperBound; i++ {
        sieve[i] = true
    }

    for i := 2; i < upperBound; i++ {
        if sieve[i] {
            for j := i * i; j < upperBound; j += i {
                sieve[j] = false
            }
        }
    }

    var primes []int

    for i := 2; i < upperBound; i++ {
        if sieve[i] && i > lowerBound {
            primes = append(primes, i)
        }
    }

    return primes
}

func Contains(s []int, e int) bool {
    for i := 0; i < len(s); i++ {
        if s[i] == e {
            return true
        }
    }

    return false
}

func main() {
    primes := EratosthenesSieve(0, 1000000)
    sumPrimes := append(primes[:0:0], primes...)

    for i := 1; i < len(sumPrimes); i++ {
        sumPrimes[i] = sumPrimes[i - 1] + primes[i]
    }

    largestConsecutive := 0
    largestPrime := 0

    for i := 0; i < len(primes); i++ {
        if i > len(primes) - largestConsecutive {
            break
        }

        for j := i; j < len(primes); j++ {
            sum := sumPrimes[j] - sumPrimes[i]

            if sum > 1000000 {
                break
            }

            if !IsPrime(sum) {
                continue
            }

            if j - i < largestConsecutive {
                continue
            }

            largestConsecutive = j - i
            largestPrime = sum
        }
    }

    fmt.Println(largestPrime)
}
