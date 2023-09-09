package brepgprop

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepgprop.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/gprop"
	"github.com/marcuswu/gooccwrapper/topods"
)

func SurfaceProperties(shape topods.Shape, props gprop.GProps, skipShared bool, useTriangulation bool) {
	C.BRepGProp_SurfaceProperties(
		C.TopoDSShape(shape.Shape),
		C.GPropGProps(props.Props),
		C.bool(skipShared),
		C.bool(useTriangulation),
	)
}
