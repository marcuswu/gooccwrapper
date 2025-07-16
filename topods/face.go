package topods

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/occ_types.h>
// #include <occwrapper/brep_tool.h>
// #include <occwrapper/topods_shape.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/geom"
)

type TopoDSFace C.TopoDSFace

type Face struct {
	Face C.TopoDSFace
}

func NewFaceFromRef(ref TopoDSFace) Face {
	return Face{Face: C.TopoDSFace(ref)}
}

func (f Face) Surface() geom.Surface {
	return geom.NewSurfaceFromRef(geom.GeomSurface(C.BRepTool_Surface(C.TopoDSFace(f.Face))))
}

func (f Face) Orientation() int {
	return int(C.TopoDSShape_Orientation(C.TopoDSShape(f.Face)))
}
