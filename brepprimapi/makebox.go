package brepprimapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepprimapi.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/gp"
	"github.com/marcuswu/gooccwrapper/topods"
)

type MakeBox struct {
	makeBox C.BRepPrimAPIMakeBox
}

func NewMakeBox(position gp.Ax2, dx, dy, dz float64) MakeBox {
	return MakeBox{C.BRepPrimAPIMakeBox_Init(
		C.gpAx2(position.Ax2),
		C.double(dx),
		C.double(dy),
		C.double(dz),
	)}
}

func (mp MakeBox) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(C.BRepPrimAPIMakeBox_Shape(mp.makeBox)))
}

func (mp MakeBox) Free() {
	C.BRepPrimAPIMakeBox_Free(mp.makeBox)
}
