package geom

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/geom_curve.h>
// #include <occwrapper/gp_ax2.h>
import "C"
import "github.com/marcuswu/gooccwrapper/gp"

type Curve struct {
	Curve C.GeomCurve
}

func MakeCircle(center gp.Ax2, radius float64) Curve {
	return Curve{C.gcTrimmedCurve_ToGeomCurve(C.gcMakeCircle(C.gpAx2(center.Ax2), C.double(radius)))}
}

func (c Curve) Free() {
	C.gcTrimmedCurve_Free(C.gcTrimmedCurve(c.Curve))
}

func MakeArc(circle gp.Circ, pt1 gp.Pnt, pt2 gp.Pnt, sense bool) Curve {
	return Curve{C.gcTrimmedCurve_ToGeomCurve(C.gcMakeArcOfCircle(C.gpCirc(circle.Circ), C.gpPnt(pt1.Pnt), C.gpPnt(pt2.Pnt), C.bool(sense)))}
}

func MakeSegment(start gp.Pnt, end gp.Pnt) Curve {
	return Curve{C.GeomCurve(C.gcMakeSegment(C.gpPnt(start.Pnt), C.gpPnt(end.Pnt)))}
}
