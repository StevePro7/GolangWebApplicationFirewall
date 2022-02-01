package cgoexample

//#include <stdio.h>
//#include <stdlib.h>
//#include <errno.h>
import "C"

import (
	"fmt"
	"unsafe"
)

var errText = map[int]string{
	C.EINVAL: "Invalid mode specified",
}

//func bob(e error) error {
//
//}

//func error(e error) error {
//	s, ok := errText[int(e.(os.Errno))]
//	if ok {
//		os.Err
//		return os.NewError(s)
//	}
//	return os.NewError(fmt.Sprintf("Unknown error: %d", int(e.(os.Errno))))
//}

type File C.FILE

func Open(path, mode string) (*File, error) {
	cpath, cmode := C.CString(path), C.CString(mode)
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(cmode))

	f, err := C.fopen(cpath, cmode)
	if f == nil {
		return nil, error(err)
	}
	return (*File)(f), nil
}

func (f *File) Close() {
	if f != nil {
		C.fclose((*C.FILE)(f))
	}
}

func Steven() {
	fmt.Println("sgb")
}