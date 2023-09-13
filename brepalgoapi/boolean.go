package brepalgoapi

import "github.com/marcuswu/gooccwrapper/topods"

type Boolean struct {
	operation C.BRepAlgoAPIBoolean
}

func (b Boolean) SetTools(toptools.ListOfShape tools) {
	C.BRepAlgoAPIBooleanOperation_SetTools(b.operation, tools.List)
}

func (b Boolean) SetArguments(toptools.ListOfShape args) {
	C.BRepAlgoAPIBooleanOperation_SetArguments(b.operation, tools.List)
}

func (b Boolean) Build() {
	C.BRepAlgoAPIBooleanOperation_Build(b.operation)
}

func (b Boolean) Shape() topods.Shape {
	topods.NewShapeFromRef(C.BRepAlgoAPIBooleanOperation_Shape(b.operation))
}

func (b Boolean) Free() {
	C.BRepAlgoAPIBooleanOperation_Free(b.operation)
}
