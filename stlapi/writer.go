package stlapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/stlapi_writer.h>
import "C"
import "github.com/marcuswu/gooccwrapper/topods"

type Writer struct {
	Writer C.StlAPIWriter
}

func NewWriter() Writer {
	return Writer{C.StlAPIWriter_Init()}
}

func (w Writer) Write(res topods.Compound, filename string) bool {
	return bool(C.StlAPIWriter_Write(w.Writer, C.TopoDSCompound(res.Compound), C.CString(filename)))
}

func (w Writer) Free() {
	C.StlAPIWriter_Free(w.Writer)
}
