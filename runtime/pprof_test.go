package runtimet

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"sync"
	"testing"
	"time"
)

func TestProfiles(t *testing.T) {
	profiles := pprof.Profiles()
	for _, profile := range profiles {
		t.Log(profile.Name())
	}

	profile := pprof.Lookup("heap")
	t.Log(profile.Name())
}

func counter() {
	list := []int{1}
	c := 1
	for i := 0; i < 15; i++ {
		httpGet()
		c = i + 1 + 2 + 3 + 4 - 5
		list = append(list, c)
	}
	fmt.Println(c)
	fmt.Println(list[0])
}

func work(wg *sync.WaitGroup) {

	counter()
	time.Sleep(1 * time.Second)

	wg.Done()
}

func httpGet() int {
	queue := []string{"start..."}
	resp, err := http.Get("http://www.163.com")
	if err != nil {
		// handle error
	}

	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	queue = append(queue, string(body))
	return len(queue)
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func TestCpuAndMemProfile(t *testing.T) {

	f, err := os.Create("cpuprofile")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg)
	}

	wg.Wait()
	time.Sleep(3 * time.Second)

}
