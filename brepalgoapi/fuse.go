package brepalgoapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepalgoapi.h>
import "C"
import "github.com/marcuswu/gooccwrapper/topods"

type Fuse struct {
	fuse C.BRepAlgoAPIFuse
}

func NewFuse() Fuse {
	return Fuse{C.BRepAlgoAPIFuse_Init()}
}

func (f Fuse) Free() {
	C.BRepAlgoAPIFuse_Free(f.fuse)
}

func (f Fuse) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(
		C.BRepAlgoAPIBooleanOperation_Shape(C.BRepAlgoAPIBooleanOperation(f.fuse)),
	))
}

func (f Fuse) ToBooleanOperation() Boolean {
	return Boolean{C.BRepAlgoAPIFuse_ToBooleanOperation(f.fuse)}
}
