package sunkentea

import "io"

type WrappedWriter struct {
	Writer io.Writer
}

func (w *WrappedWriter) Write(p []byte) (n int, err error) {
	return w.Writer.Write(p)
}
