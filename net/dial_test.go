package net

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math/rand"
	"net"
	"strconv"
	"testing"
)

func TestDial(t *testing.T) {
	conn, err := net.Dial("tcp", ":8088")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	transCoder := NewTransCoder(conn)


	sbyte := []byte("helloword" + strconv.Itoa(rand.Intn(100)))
	if sendErr := transCoder.Send(&sbyte); sendErr != nil {
		t.Fatal(sendErr)
	}

	rbyte, _ := transCoder.Receive()
	t.Log(string(rbyte))
	//}
}


type TransCoder struct {
	rw io.ReadWriter
}

func NewTransCoder(rw io.ReadWriter) *TransCoder {
	t := new(TransCoder)
	t.rw = rw
	return t
}

func (t *TransCoder) Send(body *[]byte) error {
	header, berr := cint32Tobytes(int32(len(*body)))
	if berr != nil {
		return berr
	}
	t.rw.Write(header)
	t.rw.Write(*body)
	return nil
}

func (t *TransCoder) Receive() ([]byte, error) {
	headbuf := make([]byte, 4)
	rlen, err := io.ReadFull(t.rw, headbuf)
	if err != nil {
		return nil, err
	}
	if rlen < 4 {
		return nil, errors.New("headbuf length less 4")
	}

	bodyLen, cbytesToint32err := cbytesToint32(headbuf)
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

func cint32Tobytes(len int32) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, len); err == nil {
		return buf.Bytes(), nil
	} else {
		return nil, err
	}
}

func cbytesToint32(buf []byte) (int32, error) {
	b_buf := bytes.NewBuffer(buf)
	var i32 int32
	if err := binary.Read(b_buf, binary.BigEndian, &i32); err == nil {
		return i32, nil
	} else {
		return 0, err
	}
}


func TestCint32Tobytes(t *testing.T) {
	var len int32 = 500
	buf, _ := cint32Tobytes(len)
	t.Log(buf)
	ilen, _ := cbytesToint32(buf)
	t.Log(ilen)
}

func TestTransCoder(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})
	rd := bufio.NewReader(buffer)
	wt := bufio.NewWriter(buffer)

	rw := bufio.NewReadWriter(rd, wt)
	transCoder := NewTransCoder(rw)

	sbyte := []byte("helloword")
	if sendErr := transCoder.Send(&sbyte); sendErr != nil {
		t.Fatal(sendErr)
	}
	t.Log(buffer.Bytes())

	rbyte, _ := transCoder.Receive()
	t.Log(rbyte)
}
