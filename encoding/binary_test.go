package encodingt

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	bs := []byte{0x01, 0x02, 0x00, 0x01}
	t.Log(string(bs))
	t.Log([]rune(string(bs)))

	bsReader1 := bytes.NewReader(bs[:2])
	var i1 int16
	binary.Read(bsReader1, binary.LittleEndian, &i1)
	t.Log(i1)

	bsReader2 := bytes.NewReader(bs[:2])
	var i2 int16
	binary.Read(bsReader2, binary.BigEndian, &i2)
	t.Log(i2)

	buf := make([]byte, 2)
	bsReader3 := bytes.NewReader(bs)
	t.Log(io.ReadFull(bsReader3, buf))
	bsReader4 := bytes.NewReader(buf)
	var i3 int16
	binary.Read(bsReader4, binary.BigEndian, &i3)
	t.Log(i3)

	t.Log(io.ReadFull(bsReader3, buf))
	bsReader5 := bytes.NewReader(buf)
	binary.Read(bsReader5, binary.BigEndian, &i3)
	t.Log(i3)

	// 0 EOF
	t.Log(io.ReadFull(bsReader3, buf))
}

func TestWriter(t *testing.T) {

	buf := bytes.NewBuffer([]byte{})
	var c byte = 'L'
	buf.WriteByte(c)

	var i1 int32 = 10
	binary.Write(buf, binary.BigEndian, i1)

	var ds string = "iamfoobara"
	buf.WriteString(ds)

	t.Log(len(buf.Bytes()))

	bs := make([]byte, 5)
	io.ReadFull(buf, bs)

	if bs[0] == 'L' {
		t.Logf("%c", bs[0])
	}

	ir := bytes.NewReader(bs[1:5])
	var i2 int32
	binary.Read(ir, binary.BigEndian, &i2)
	t.Log(i2)

	dsbs := make([]byte, 100)
	io.ReadFull(buf, dsbs)

	t.Log(string(dsbs))

}

type Stu struct {
	Id    int32
	Name  []byte
	Sex   int8
	Score float32
}

type StudentFormater struct {
	rd io.Reader
	wt io.Writer
}

func NewStudentEncoder(wt io.Writer) *StudentFormater {
	s := new(StudentFormater)
	s.wt = wt
	return s
}

func NewStudentDecoder(rd io.Reader) *StudentFormater {
	s := new(StudentFormater)
	s.rd = rd
	return s
}

// 4byte Id
// 1byte Sex
// 4byte Score
// 4byte NameLength
// NameLength byte Name
func (s *StudentFormater) Encode(student *Stu) error {
	if err := binary.Write(s.wt, binary.BigEndian, student.Id); err != nil {
		return err
	}
	if err := binary.Write(s.wt, binary.BigEndian, student.Sex); err != nil {
		return err
	}
	if err := binary.Write(s.wt, binary.BigEndian, student.Score); err != nil {
		return err
	}
	var nameLength int32
	nameLength = int32(len(student.Name))
	if err := binary.Write(s.wt, binary.BigEndian, nameLength); err != nil {
		return err
	}
	if _, err := s.wt.Write(student.Name); err != nil {
		return err
	}
	return nil
}

func (s *StudentFormater) Decode() (*Stu, error) {
	student := &Stu{}
	header := make([]byte, 13)
	if _, err := io.ReadFull(s.rd, header); err != nil {
		return nil, err
	}
	idReader := bytes.NewReader(header[:3])
	if err := binary.Read(idReader, binary.BigEndian, &student.Id); err != nil && err != io.ErrUnexpectedEOF {
		return nil, err
	}
	sexReader := bytes.NewReader(header[4:5])
	if err := binary.Read(sexReader, binary.BigEndian, &student.Sex); err != nil && err != io.ErrUnexpectedEOF {
		return nil, err
	}
	scoreReader := bytes.NewReader(header[5:9])
	if err := binary.Read(scoreReader, binary.BigEndian, &student.Score); err != nil && err != io.ErrUnexpectedEOF {
		return nil, err
	}
	var nameLength int32
	nameLengthReader := bytes.NewReader(header[9:13])
	if err := binary.Read(nameLengthReader, binary.BigEndian, &nameLength); err != nil && err != io.ErrUnexpectedEOF {
		return nil, err
	}

	namebuf := make([]byte, nameLength)
	if _, ReadFullErr := io.ReadFull(s.rd, namebuf); ReadFullErr != nil {
		return nil, ReadFullErr
	}
	student.Name = namebuf[:]
	return student, nil
}

func TestStudentEncode(t *testing.T) {
	student := &Stu{
		456,
		[]byte("miller linkon"),
		1,
		73.6,
	}

	buf := bytes.NewBuffer([]byte{})
	formater := NewStudentEncoder(buf)
	if err := formater.Encode(student); err != nil {
		t.Fatal(err)
	}
	t.Log(len(buf.Bytes()))
	t.Log(buf.Bytes())
	// output:
	// [0 0 1 200 1 66 147 51 51 0 0 0 13 109 105 108 108 101 114 32 108 105 110 107 111 110]
}

func TestStudentDecode(t *testing.T) {
	student := &Stu{
		456,
		[]byte("miller linkon"),
		1,
		73.6,
	}
	buf := bytes.NewBuffer([]byte{})
	formater := NewStudentEncoder(buf)
	if err := formater.Encode(student); err != nil {
		t.Fatal(err)
	}

	bs := buf.Bytes()
	bsReader := bytes.NewReader(bs)
	formater2 := NewStudentDecoder(bsReader)
	stu, DecodeErr := formater2.Decode()
	if DecodeErr != nil {
		t.Fatal(DecodeErr)
	}
	t.Log(stu.Id)
	t.Log(string(stu.Name))
	t.Log(stu.Score)
	t.Log(stu.Sex)
}
