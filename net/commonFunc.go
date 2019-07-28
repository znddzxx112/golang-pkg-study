package net

import (
	"bytes"
	"encoding/binary"
)

func int32Tobytes(len int32) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, len); err == nil {
		return buf, nil
	} else {
		return nil, err
	}
}

func bytesToint32(buf []byte) (int32, error) {
	b_buf := bytes.NewBuffer(buf)
	var i32 int32
	if err := binary.Read(b_buf, binary.BigEndian, &i32); err == nil {
		return i32, nil
	} else {
		return 0, err
	}
}