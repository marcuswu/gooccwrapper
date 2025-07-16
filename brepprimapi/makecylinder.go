package brepprimapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepprimapi.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/gp"
	"github.com/marcuswu/gooccwrapper/topods"
)

type MakeCylinder struct {
	makeCylinder C.BRepPrimAPIMakeCylinder
}

func NewMakeCylinder(position gp.Ax2, radius, height float64) MakeCylinder {
	return MakeCylinder{C.BRepPrimAPIMakeCylinder_Init(
		C.gpAx2(position.Ax2),
		C.double(radius),
		C.double(height),
	)}
}

func (mp MakeCylinder) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(C.BRepPrimAPIMakeCylinder_Shape(mp.makeCylinder)))
}

func (mp MakeCylinder) Free() {
	C.BRepPrimAPIMakeCylinder_Free(mp.makeCylinder)
}
