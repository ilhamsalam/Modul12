package main

import (
    "fmt"
)

// Fungsi untuk melakukan selection sort
func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        // Tukar posisi elemen
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}

// Fungsi untuk menghitung median
func calculateMedian(arr []int) int {
    n := len(arr)
    if n%2 == 1 {
        // Jika jumlah elemen ganjil, median adalah elemen tengah
        return arr[n/2]
    } else {
        // Jika jumlah elemen genap, median adalah rerata dari dua elemen tengah
        return (arr[n/2-1] + arr[n/2]) / 2
    }
}

func main() {
    var data []int
    var input int

    for {
        fmt.Scan(&input)
        if input == -5313 {
            break
        } else if input == 0 {
            // Urutkan data dan cetak median
            selectionSort(data)
            fmt.Println(calculateMedian(data))
        } else {
            // Tambahkan input ke dalam data
            data = append(data, input)
        }
    }
}
