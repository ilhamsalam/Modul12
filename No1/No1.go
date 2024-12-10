package main

import "fmt"

func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        // Find the minimum element in unsorted array
        minIndex := i
        for j := i+1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        // Swap the found minimum element with the first element
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}

func main() {
    var n int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        var m int
        fmt.Scan(&m)

        // Baca nomor rumah kerabat
        numbers := make([]int, m)
        for j := 0; j < m; j++ {
            fmt.Scan(&numbers[j])
        }

        // Urutkan menggunakan selection sort
        selectionSort(numbers)

        // Cetak hasil
        for _, num := range numbers {
            fmt.Print(num, " ")
        }
        fmt.Println()
    }
}