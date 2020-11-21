package main

import (
	"fmt"
	"strings"
)

func sorting(arr []string) {
    for i := 0; i < len(arr); i++ {
        tmp := arr[i]
        j := i
        for j > 0 && arr[j-1] > tmp {
            arr[j] = arr[j-1]
            j--
        }
        arr[j] = tmp
    }
}

func main () {
    var n string
    fmt.Println("Masukan nomer yang akan di urutkan dengan koma(,) sebagai pemisah")
    fmt.Print("number : ")
    fmt.Scan(&n)
    arr := strings.Split(n, ",")
    fmt.Print("Isi Element Array: ", arr, "\n", "Hasil Pengurutan: ")
    sorting(arr)
    for i := 0; i < len(arr); i++ {
        fmt.Print(arr[i], " ")
    }
    fmt.Print("\n")
}
