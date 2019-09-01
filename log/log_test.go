package logt

import (
	"log"
	"testing"
)

// prefix flag
// println panic fatal
// log.std = New(os.Stderr, "", LstdFlags)
func TestLog(t *testing.T) {
	log.SetPrefix("bar:")
	log.SetFlags(log.Llongfile | log.LstdFlags)
	t.Log(log.Flags())
	log.Println("hh")
	//log.Panic("help")
	//log.Fatal("fatal")
}
