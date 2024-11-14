package main

/*
extern int _main(int argc, char** argv);
*/
import "C"
import (
	"os"
	"unsafe"
)

func main() {
	argc := len(os.Args)
	argv := make([]*C.char, argc)
	for i, arg := range os.Args {
		argv[i] = C.CString(arg)
	}
	C._main(C.int(argc), (**C.char)(unsafe.Pointer(&argv[0])))
}
