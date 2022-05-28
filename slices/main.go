package main

import "fmt"

func appendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        // There is room to grow.  Extend the slice.
        z = x[:zlen]
    } else {
        // There is insufficient space.  Allocate a new array.
        // Grow by doubling, for amortized linear complexity.
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2 * len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z, x) // a built-in function; see text
    }
    z[len(x)] = y
    return z
}
func reverse(int *[6]int) {
	for i, j := 0, len(*int)-1; i < j; i, j = i+1, j-1 {
		(*int)[i], (*int)[j] = (*int)[j], (*int)[i]
	}
}
// func reverse(s *[]int) {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }

func main() {

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)

    // var x, y []int
    // for i := 0; i < 10; i++ {
    //     y = appendInt(x, i)
    //     fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
    //     x = y
    // }

	// months := [...]string{1: "January",2: "February",3: "March",4: "April",5: "May",6: "June",7: "July",8: "August",9: "September",10: "October",11:"November",12: "December"}
	// Q2 := months[4:7]
	// summer := months[6:9]
	// fmt.Println(Q2)     // ["April" "May" "June"]
	// fmt.Println(summer) // ["June" "July" "August"]

	// for _, s := range summer {
    // for _, q := range Q2 {
    //     if s == q {
    //         fmt.Printf("%s appears in both\n", s)
    //     }
    // }
	// }

	// // fmt.Println(summer[:20]) // panic: out of range

	// endlessSummer := summer[:5] // extend a slice (within capacity)
	// fmt.Println(endlessSummer)  // "[June July August September October]"

	// array := [...]int{0, 1, 2, 3, 4, 5}

	// slice := array[:2]
	// change := array[:1]
	// change[0] = 10
	
	// fmt.Println(&array[0])
	// fmt.Println(change)
	// fmt.Println(slice)

	// array[0] = 100

	// fmt.Println(&array[0])
	// fmt.Println(change)
	// fmt.Println(slice)
}