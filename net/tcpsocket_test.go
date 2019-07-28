package net

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"net"
	"testing"
)

func TestListen(t *testing.T) {
	l, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// 等待下一个连接,如果没有连接,l.Accept会阻塞
		conn, err := l.Accept()
		if err != nil {
			t.Fatal(err)
		}
		// 将新连接放入一个goroute里,然后再等下一个新连接.
		go func(c net.Conn) {
			sTransCoder := NewSTransCoder(conn)
			bs,err := sTransCoder.Receive()
			if err != nil {
				c.Close()
				return
			}
			t.Log(string(bs))
			sTransCoder.Send(&bs)
		}(conn)
	}
}

type STransCoder struct {
	rw io.ReadWriter
}

func NewSTransCoder(rw io.ReadWriter) *STransCoder {
	t := new(STransCoder)
	t.rw = rw
	return t
}

func (t *STransCoder) Send(body *[]byte) error {
	header, berr := sint32Tobytes(int32(len(*body)))
	if berr != nil {
		return berr
	}
	t.rw.Write(header)
	t.rw.Write(*body)
	return nil
}

func (t *STransCoder) Receive() ([]byte, error) {
	headbuf := make([]byte, 4)
	rlen, err := io.ReadFull(t.rw, headbuf)
	if err != nil {
		return nil, err
	}
	if rlen < 4 {
		return nil, errors.New("headbuf length less 4")
	}

	bodyLen, cbytesToint32err := sbytesToint32(headbuf)
	if cbytesToint32err != nil {
		return nil, cbytesToint32err
	}

	body := make([]byte, bodyLen)
	_, readFullerr := io.ReadFull(t.rw, body)
	if readFullerr != nil {
		return nil, readFullerr
	}
	return body, nil
}

func sint32Tobytes(len int32) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, len); err == nil {
		return buf.Bytes(), nil
	} else {
		return nil, err
	}
}

func sbytesToint32(buf []byte) (int32, error) {
	b_buf := bytes.NewBuffer(buf)
	var i32 int32
	if err := binary.Read(b_buf, binary.BigEndian, &i32); err == nil {
		return i32, nil
	} else {
		return 0, err
	}
}