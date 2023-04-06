package protocol

import (
	"fmt"
	"io"
)

type Writer struct{
	writer io.Writer
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		writer:writer,
	}
}

func (w *Writer) writeString(msg string) error {
	_,err := w.writer.Writer([]byte(msg))
	return err
}

func (w *Writer) Write(command interface{}) error {
	var err error

	switch v := command.(type) {
	case SendCmd:
		err = w.writeString(fmt.Sprintf("SEND %v\n"),v.Message)
	case Message:
		err = w.writeString(fmt.Sprintf("MESSAGE %v %v\n",v.Name,v.Message))
	case NameCmd:
		err = w.writeString(fmt.Sprintf("NAME %v\n",v.Name))
	default:
		err = UnknowCommand
	}
	return err
}