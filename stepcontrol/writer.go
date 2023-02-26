package stepcontrol

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/stepcontrol_writer.h>
import "C"
import "github.com/marcuswu/gooccwrapper/pkg/topods"

type StepModelType int

// This list echos opencascade/STEPControl_StepModelType.hxx and must be updated if opencascade is updated
const (
	AsIs StepModelType = iota
	ManifoldSolidBrep
	BrepWithVoids
	FacetedBrep
	FacetedBrepAndBrepWithVoids
	ShellBasedSurfaceModel
	GeometricCurveSet
	Hybrid
)

type Writer struct {
	Writer C.STEPControlWriter
}

func NewWriter() Writer {
	return Writer{C.STEPControlWriter_Init()}
}

func (w Writer) Transfer(shape topods.Shape, modelType StepModelType) int {
	return int(C.STEPControlWriter_Transfer(w.Writer, C.TopoDSShape(shape.Shape), C.int(modelType)))
}

func (w Writer) Write(filename string) int {
	return int(C.STEPControlWriter_Write(w.Writer, C.CString(filename)))
}

func (w Writer) Free() {
	C.STEPControlWriter_Free(w.Writer)
}
