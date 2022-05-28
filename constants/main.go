package main

import "fmt"

const (
    _ = 1 << (10 * iota)
    KiB // 1024
    MiB // 1048576
    GiB // 1073741824
    TiB // 1099511627776              (exceeds 1 << 32)
    PiB // 1125899906842624
    EiB // 1152921504606846976
    ZiB // 1180591620717411303424     (exceeds 1 << 64)
    YiB // 1208925819614629174706176
)

func main() {
	fmt.Println(KiB)
	fmt.Println(MiB)
	fmt.Println(GiB)
	fmt.Println(TiB)
	fmt.Println(PiB)
	fmt.Println(EiB)
	// fmt.Println(ZiB)
	// fmt.Println(YiB)

	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%c\n", '\000') // "int32" (rune)
}