package brepalgoapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepalgoapi.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/topods"
	"github.com/marcuswu/gooccwrapper/toptools"
)

type Boolean struct {
	operation C.BRepAlgoAPIBooleanOperation
}

func (b Boolean) SetTools(tools toptools.ListOfShape) {
	C.BRepAlgoAPIBooleanOperation_SetTools(b.operation, C.TopToolsListOfShape(tools.List))
}

func (b Boolean) SetArguments(args toptools.ListOfShape) {
	C.BRepAlgoAPIBooleanOperation_SetArguments(b.operation, C.TopToolsListOfShape(args.List))
}

func (b Boolean) Build() {
	C.BRepAlgoAPIBooleanOperation_Build(b.operation)
}

func (b Boolean) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(C.BRepAlgoAPIBooleanOperation_Shape(b.operation)))
}

func (b Boolean) Free() {
	C.BRepAlgoAPIBooleanOperation_Free(b.operation)
}
