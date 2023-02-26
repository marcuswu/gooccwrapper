package brepalgoapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepalgoapi.h>
import "C"
import "github.com/marcuswu/gooccwrapper/pkg/topods"

type Cut struct {
	cut C.BRepAlgoAPICut
}

func NewCut() Cut {
	return Cut{C.BRepAlgoAPICut_Init()}
}

func (c Cut) Free() {
	C.BRepAlgoAPICut_Free(c.cut)
}

func (c Cut) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(
		C.BRepAlgoAPIBooleanOperation_Shape(C.BRepAlgoAPIBooleanOperation(c.cut)),
	))
}
