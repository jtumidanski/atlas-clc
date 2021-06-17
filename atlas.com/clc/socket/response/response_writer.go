package response

import (
	"bytes"
	"encoding/binary"
	"github.com/sirupsen/logrus"
)

type Writer struct {
	l logrus.FieldLogger
	o *bytes.Buffer
}

func NewWriter(l logrus.FieldLogger) *Writer {
	return &Writer{l, new(bytes.Buffer)}
}

func (w *Writer) WriteInt(val uint32) {
	err := binary.Write(w.o, binary.LittleEndian, val)
	if err != nil {
		w.l.WithError(err).Fatal("Writing int value")
	}
}

func (w *Writer) WriteShort(val uint16) {
	err := binary.Write(w.o, binary.LittleEndian, val)
	if err != nil {
		w.l.WithError(err).Fatal("Writing short value")
	}
}

func (w *Writer) WriteLong(val uint64) {
	err := binary.Write(w.o, binary.LittleEndian, val)
	if err != nil {
		w.l.WithError(err).Fatal("Writing long value")
	}
}

func (w *Writer) WriteByte(val byte) {
	err := binary.Write(w.o, binary.LittleEndian, val)
	if err != nil {
		w.l.WithError(err).Fatal("Writing byte value")
	}
}

func (w *Writer) WriteByteArray(vals []byte) {
	for i := 0; i < len(vals); i++ {
		err := binary.Write(w.o, binary.LittleEndian, vals[i])
		if err != nil {
			w.l.WithError(err).Fatal("Writing byte value")
		}
	}
}

func (w *Writer) WriteBool(val bool) {
	i := 1
	if !val {
		i = 0
	}
	w.WriteByte(byte(i))
}

func (w *Writer) WriteAsciiString(s string) {
	w.WriteShort(uint16(len(s)))
	w.WriteByteArray([]byte(s))
}

func (w *Writer) WriteKeyValue(key byte, value uint32) {
	w.WriteByte(key)
	w.WriteInt(value)
}

func (w *Writer) Bytes() []byte {
	return w.o.Bytes()
}
