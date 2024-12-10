package main

import "fmt"

func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIndex := i
        for j := i+1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}

func main() {
    var n int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        var m int
        fmt.Scan(&m)

        var nums []int
        for j := 0; j < m; j++ {
            var num int
            fmt.Scan(&num)
            nums = append(nums, num)
        }

        var odd, even []int
        for _, num := range nums {
            if num%2 == 0 {
                even = append(even, num)
            } else {
                odd = append(odd, num)
            }
        }

        selectionSort(odd)
        selectionSort(even)
        reverse(even) // Membalik urutan even agar terurut menurun

        fmt.Println(odd)
        fmt.Println(even)
    }
}

func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}