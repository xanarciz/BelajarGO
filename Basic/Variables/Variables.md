# Variables

- Variable didefinisikan dengan statment `var`.
- Variable bisa di dalam package atau di dalam fungsi.
- Inisialisasi variable bisa dilakukan di satu statemen var.
- Inisialisasi variable tidak perlu ditulis type datanya.
- Operator `:=` digunakan untuk memperpendek cara mendeklarasikan variable.
- Operator `:=` hanya bekerja didalam fungsi.
- Type data di GO
    ```
    bool

    string

    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr

    byte // alias for uint8

    rune // alias for int32
        // represents a Unicode code point

    float32 float64

    complex64 complex128

    ```
- Variable yang tidak mempunyai nilai akan menghasilkan nilai kosong.
    - Pada integer akan menghasilkan 0.
    - Pada string akan menghasilkan "".
    - Pada boolean akan menghasilkan false.
- Inisialisasi konstanta menggunakan statment `const`