package topexp

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/topexp_explorer.h>
import "C"
import "github.com/marcuswu/gooccwrapper/topods"

type ShapeEnum int

// This list echos opencascade/STEPControl_StepModelType.hxx and must be updated if opencascade is updated
const (
	Compound ShapeEnum = iota
	CompSolid
	Solid
	Shell
	Face
	Wire
	Edge
	Vertex
	Shape
)

type Explorer struct {
	Explorer C.TopExpExplorer
}

func NewExplorer(shape topods.Shape, toFind ShapeEnum) Explorer {
	return Explorer{C.TopExpExplorer_Init(C.TopoDSShape(shape.Shape), C.int(toFind))}
}

func (e Explorer) Free() {
	C.TopExpExplorer_Free(e.Explorer)
}

func (e Explorer) More() bool {
	return bool(C.TopExpExplorer_More(e.Explorer))
}

func (e Explorer) Next() {
	C.TopExpExplorer_Next(e.Explorer)
}

func (e Explorer) Current() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(C.TopExpExplorer_Current(e.Explorer)))
}

func (e Explorer) Depth() int {
	return int(C.TopExpExplorer_Depth(e.Explorer))
}
