package main

import (
	"fmt"
	// "strings"
)

func sorting(arr []int) {
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
 //    var number string
	// fmt.Scanln(&number)
 //    arr := []{number}
	// s := strings.Split(number, ",")
 //    // sorting()
	// fmt.Println(s, arr)
    arr := []int{1, 4, 5, 2}
    fmt.Print("Isi Element Array: ", arr, "\n", "Hasil Pengurutan: [ ")
    sorting(arr)
    for i := 0; i < len(arr); i++ {
        fmt.Print(arr[i], " ")
    }
    fmt.Print("]\n")
}
