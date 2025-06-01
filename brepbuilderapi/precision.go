package brepbuilderapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brep_builder.h>
import "C"

func SetPrecision(precision float64) {
	C.BRepBuilderAPI_SetPrecision(C.double(precision))
}
