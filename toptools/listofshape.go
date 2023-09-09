package toptools

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/toptools_listofshape.h>
import "C"
import "github.com/marcuswu/gooccwrapper/topods"

type ListOfShape struct {
	List C.TopToolsListOfShape
}

type ShapeIterator struct {
	Iterator C.ShapeIterator
}

func NewListOfShape() ListOfShape {
	return ListOfShape{C.TopToolsListOfShape_Init()}
}

func (l ListOfShape) Free() {
	C.TopToolsListOfShape_Free(l.List)
}

func (l ListOfShape) Append(shape topods.Shape) topods.Shape {
	C.TopToolsListOfShape_Append(l.List, C.TopoDSShape(shape.Shape))
	return shape
}

func (l ListOfShape) Begin() ShapeIterator {
	return ShapeIterator{C.TopToolsListOfShape_Begin(l.List)}
}

func (i ShapeIterator) Next() ShapeIterator {
	return ShapeIterator{C.ShapeIterator_Next(i.Iterator)}
}

func (i ShapeIterator) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(C.ShapeIterator_Shape(i.Iterator)))
}
