package breptools

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/breptools.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/topods"
)

func UVBounds(face topods.Face) (float64, float64, float64, float64) {
	var umin float64
	var umax float64
	var vmin float64
	var vmax float64
	C.BRepTools_UVBounds(C.TopoDSFace(face.Face), (*C.double)(&umin), (*C.double)(&umax), (*C.double)(&vmin), (*C.double)(&vmax))

	return umin, umax, vmin, vmax
}
