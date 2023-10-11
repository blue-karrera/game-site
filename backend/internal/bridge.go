package internal

/*
int return_int();
int* return_ints();
#cgo LDFLAGS: -L. -L../ -lclib -lhlib
*/
import "C"
import "fmt"
import "unsafe"

func return_int() int {
    return int(C.return_int())
}

func return_ints() []int {
    //var ret []int
    //var x = C.return_ints()
    //var y = unsafe.Pointer(x)
    //defer C.free(unsafe.Pointer(cmsg))
    //var y = unsafe.Slice(x, 30)
    //var y = unsafe.Pointer(x)
    //var z = C.GoBytes(y, 30)
    //var a = (*[30]int)(unsafe.Pointer(&z[0]))
    //var b = ([]int)(unsafe.Pointer(x))
    //var c = make(z, 30, 30)
    //var c = make(a[:], 30, 30)
    //var d = a[:]
    //var e = (*[30]int)(unsafe.Pointer(x))[:]
    //return e
    return (*[30]int)(unsafe.Pointer(C.return_ints()))[:]
}


func Test() {
    var i = return_int()
    var j = return_ints()
    fmt.Printf("i: %d\n", i)
    fmt.Printf("j's: %d %d\n", j[0], j[29])
}


