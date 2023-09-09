package breptool

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brep_tool.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/geom"
	"github.com/marcuswu/gooccwrapper/gp"
	"github.com/marcuswu/gooccwrapper/topods"
)

func Surface(ref topods.Face) geom.Surface {
	return geom.NewSurfaceFromRef(geom.GeomSurface(C.BRepTool_Surface(C.TopoDSFace(ref.Face))))
}

func Pnt(ref topods.Vertex) gp.Pnt {
	return gp.NewPntFromRef(gp.GPPnt(C.BRepTool_Pnt(C.TopoDSVertex(ref.Vertex))))
}
