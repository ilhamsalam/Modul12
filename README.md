# Latihan modul 12
Penjelasan singkat pada tiap latihan soal modul 12

## Setiap soal
Pada setiap soal saya menggunakan beberapa komponen :
- `Package main` agar program dapat dieksekusi
- `import fmt` agar dapat menggunakan beberapa operasi dasar bahasa program `go`
- `main()` sebagai titik awal program dan dieksekusi ketika program dijalankan

## Soal 1
```go
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
```
Operasi diatas merupakan operasi mengurutkan angka dari terkecil ke terbesar dengan cara : Pertama, elemen terkecil pada array tersebut ditukar dengan elemen pertama. Lakukan hal ini sampai array terurut dari yang terkecil ke terbesar

## Soal 2
```go
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
```
Operasi diatas merupakan operasi mengurutkan angka dari terkecil ke terbesar dengan cara : Pertama, elemen terkecil pada array tersebut ditukar dengan elemen pertama. Lakukan hal ini sampai array terurut dari yang terkecil ke terbesar

```go
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}
```
Operasi diatas merupakan operasi membalik urutan elemen dalam sebuah slice. Ini dilakukan dengan menukar elemen dari ujung kiri dan kanan secara berpasangan hingga mencapai tengah slice.

```go
        var odd, even []int
        for _, num := range nums {
            if num%2 == 0 {
                even = append(even, num)
            } else {
                odd = append(odd, num)
            }
        }
```
Operasi diatas merupakan operasi memisahkan nomor-nomor genap dan ganjil ke dalam slice even dan odd masing-masing.

## Soal 3
```go
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
```
Operasi diatas merupakan operasi mengurutkan angka dari terkecil ke terbesar dengan cara : Pertama, elemen terkecil pada array tersebut ditukar dengan elemen pertama. Lakukan hal ini sampai array terurut dari yang terkecil ke terbesar

```go
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
```
Operasi diatas merupakan operasi mencari nilai median dari data yang sudah diurutkan oleh selection sort sebelumnya. Jika jumlah elemen dalam slice ganjil, median adalah elemen tengah. Jika jumlah elemen dalam slice genap, median adalah rata-rata dari dua elemen tengah.

```go
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
```
Operasi diatas merupakan operasi menginput nilai pada array. dengan beberapa ketentuan : -5313: Angka ini bertindak sebagai "penanda akhir" input. Ketika pengguna memasukkan angka ini, program akan berhenti membaca input dan melakukan perhitungan. 0: Angka ini menandakan bahwa pengguna ingin menghitung median dari data yang sudah dimasukkan.
