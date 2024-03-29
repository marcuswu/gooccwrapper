package brepprimapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepprimapi.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/gp"
	"github.com/marcuswu/gooccwrapper/topods"
)

type MakePrism struct {
	makePrism C.BRepPrimAPIMakePrism
}

func NewMakePrism(face topods.Face, vec gp.Vec) MakePrism {
	return MakePrism{C.BRepPrimAPIMakePrism_Init(
		C.TopoDSFace(face.Face),
		C.gpVec(vec.Vec),
	)}
}

func (mp MakePrism) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(C.BRepPrimAPIMakePrism_Shape(mp.makePrism)))
}

func (mp MakePrism) Free() {
	C.BRepPrimAPIMakePrism_Free(mp.makePrism)
}
