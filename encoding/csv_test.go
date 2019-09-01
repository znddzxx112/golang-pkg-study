package encodingt

import (
	"encoding/csv"
	"io"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

type Student struct {
	Id    int
	Score int
}

func TestCsv(t *testing.T) {
	os.Chdir("encodingTemp")

	students := make([]Student, 20)
	for i := 0; i < 20; i++ {
		s := Student{
			i + 1,
			rand.Intn(100),
		}
		students[i] = s
	}

	file, CreateErr := os.Create("dump.csv")
	if CreateErr != nil {
		t.Fatal(CreateErr)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{"id", "socre"})
	for _, s := range students {
		record := []string{
			strconv.Itoa(s.Id),
			strconv.Itoa(s.Score),
		}
		csvWriter.Write(record)
	}
	csvWriter.Flush()
}

func TestCsvDump(t *testing.T) {

	os.Chdir("encodingTemp")

	file, OpenErr := os.Open("dump.csv")
	if OpenErr != nil {
		t.Fatal(OpenErr)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	for {
		record, ReadErr := csvReader.Read()
		if ReadErr != nil {
			if ReadErr == io.EOF {
				break
			}
			t.Fatal(ReadErr)
		}
		t.Log(record)
	}
}
