package gp

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/gp_vec.h>
import "C"

func Resolution() float64 {
	return float64(C.gpResolution())
}
