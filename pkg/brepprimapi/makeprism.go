package brepprimapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepprimapi.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/pkg/gp"
	"github.com/marcuswu/gooccwrapper/pkg/topods"
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

func (mr MakePrism) Free() {
	C.BRepPrimAPIMakePrism_Free(mr.makePrism)
}
