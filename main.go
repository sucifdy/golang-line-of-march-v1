package main

import (
    "fmt"
    "sort"
)


func Sortheight(height []int) []int {
    if len(height) == 0 {
        return []int{} 
    }

    sort.Ints(height)
    return height
}

func main() {

    fmt.Println(Sortheight([]int{}))                                                                                    // Expected output: []
    fmt.Println(Sortheight([]int{100, 120, 130, 140, 150, 160, 170, 175, 180, 181, 182, 183}))                          // Expected output: [100, 120, 130, 140, 150, 160, 170, 175, 180, 181, 182, 183]
    fmt.Println(Sortheight([]int{180, 170, 160, 150, 140, 130, 120, 110}))                                              // Expected output: [110, 120, 130, 140, 150, 160, 170, 180]
    fmt.Println(Sortheight([]int{144, 150, 150, 155, 159, 165, 165, 165, 166, 167, 169, 169, 170, 170, 172, 175, 180})) // Expected output: [144, 150, 150, 155, 159, 165, 165, 165, 166, 167, 169, 169, 170, 170, 172, 175, 180]
}