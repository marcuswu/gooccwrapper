package topods

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/topods_shape.h>
import "C"
import "github.com/marcuswu/gooccwrapper/toploc"

type TopoDSShape C.TopoDSShape

type Shape struct {
	Shape C.TopoDSShape
}

func NewShapeFromRef(ref TopoDSShape) Shape {
	return Shape{C.TopoDSShape(ref)}
}

func (s Shape) IsEqual(other TopoDSShape) bool {
	return bool(C.TopoDSShape_IsEqual(s.Shape, C.TopoDSShape(other)))
}

func (s Shape) Location() toploc.Location {
	return toploc.NewLocationFromRef(toploc.TopLocLocation(C.TopoDSShape_Location(s.Shape)))
}

func (s Shape) Free() {
	C.TopoDSShape_Free(s.Shape)
}
