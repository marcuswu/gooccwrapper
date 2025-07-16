package gcpnts

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/gcpnts.h>
// #include <occwrapper/brep_adapter.h>
import "C"
import "github.com/marcuswu/gooccwrapper/brepadapter"

func CurveLength(c brepadapter.Curve) float64 {
	return float64(C.CurveLength(C.BRepAdapterCurve(c.Curve)))
}
