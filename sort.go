package main

import (
    "fmt"
    "strings"
)

func bubble (arr []string) {
    a := strings.Join(arr, " ")
    fmt.Println(a)
    for {
        disortir := true
        for i := 0; i < len(arr)-1; i++ {
            if arr[i] > arr[i+1] {
                tmp := arr[i]
                arr[i] = arr[i+1]
                arr[i+1] = tmp
                disortir = false
            }
        }
        if disortir == true {
            break
        }
        b := strings.Join(arr, " ")
        fmt.Println(b)
    }
}

func main () {
    var n string
    fmt.Println("Masukan nomer yang akan di urutkan dengan koma(,) sebagai pemisah")
    fmt.Print("number : ")
    fmt.Scan(&n)
    arr := strings.Split(n, ",")
    bubble(arr)
}