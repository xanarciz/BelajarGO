# Packages

- Semua program GO ditulis dalam sebuah package.
- Import bisa dilakukan untuk menggunakan package lain.
- Program pertama kali berjalan dalam package main.
- Cara penulisan untuk import :
    ```
    import "fmt"
    import "math"
    ```

    ```
    import (
        "fmt"
        "math"
    )
    ```
- Nama yang diekspor diawali oleh kapital. Misal package fmt menggunakan fungsi println, maka ditulis
`fmt.Println`.