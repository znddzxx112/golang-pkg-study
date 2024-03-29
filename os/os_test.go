package ost

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"strings"
	"syscall"
	"testing"
	"time"
)

var currentDir string

func init() {
	currentDir, _ = os.Getwd()
}

func TestFile(t *testing.T) {

	fileName := fmt.Sprintf("%s/%s", currentDir, "file.csv")
	fd, OpenFileErr := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if OpenFileErr != nil {
		t.Fatal(OpenFileErr)
	}
	defer fd.Close()

	writeHeader(fd, []string{"id", "name", "score"})
	writeBody(fd, [][]string{
		{"1", "z", "88"},
		{"2", "y", "89"},
		{"3", "x", "90"},
	})

	t.Log(fd.Fd())
	fileinfo, StatErr := fd.Stat()
	if StatErr != nil {
		t.Log(os.IsExist(StatErr))
	}
	// fileinfo
	t.Log(fileinfo.Name())
	t.Log(fileinfo.Size())
	t.Log(fileinfo.Mode())
	t.Log(fileinfo.ModTime().Format(time.RFC3339))
}

func writeHeader(fd *os.File, header []string) {
	fd.WriteString(strings.Join(header, ",") + "\n")
}

func writeBody(fd *os.File, body [][]string) {
	for _, line := range body {
		fd.WriteString(strings.Join(line, ",") + "\n")
	}
}

func TestFileUD(t *testing.T) {
	fileName := fmt.Sprintf("%s/%s", currentDir, "file")
	newFIlename := fmt.Sprintf("%s/%s", currentDir, "newfile")
	os.Create(fileName)
	if RenameErr := os.Rename(fileName, newFIlename); RenameErr != nil {
		t.Fatal(RenameErr)
	}
	if RemoveErr := os.Remove(newFIlename); RemoveErr != nil {
		t.Fatal(RemoveErr)
	}
}

func TestFileExist(t *testing.T) {
	fileNameNotExist := fmt.Sprintf("%s/%s", currentDir, "notExist")
	fdNot, fileErr := os.OpenFile(fileNameNotExist, os.O_RDONLY, 0440)
	if fileErr != nil {
		t.Log(os.IsNotExist(fileErr))
		if os.IsNotExist(fileErr) {
			t.Log("not exist")
		} else {
			t.Fatal(fileErr)
		}

	}
	defer fdNot.Close()

	var fileExist bool = true
	_, StatErr := os.Stat(fileNameNotExist)
	if StatErr != nil {
		if os.IsNotExist(StatErr) {
			fileExist = false
		}
	}
	t.Log(fileExist)
}

func TestDir(t *testing.T) {
	dir := "/tmp/dirt"
	MkdirAllErr := os.MkdirAll(dir, 0755)
	if MkdirAllErr != nil {
		if os.IsPermission(MkdirAllErr) {
			t.Log("permission")
		}
		t.Fatal(MkdirAllErr)
	}

	if ChdirErr := os.Chdir(dir); ChdirErr != nil {
		t.Fatal(ChdirErr)
	}

	now := time.Now()
	todayZoreTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.FixedZone("UTC", 8*3600))

	if ChtimesErr := os.Chtimes(dir, todayZoreTime, todayZoreTime); ChtimesErr != nil {
		t.Fatal(ChtimesErr)
	}

	if _, CreateErr := os.Create("/tmp/test.txt"); CreateErr != nil {
		t.Fatal(CreateErr)
	}

	t.Log(os.Getwd())

}

func TestProcess(t *testing.T) {
	hostname, _ := os.Hostname()
	t.Log(hostname)
	userhome, _ := os.UserHomeDir()
	t.Log(userhome)
	usercache, _ := os.UserCacheDir()
	t.Log(usercache)
	t.Log(os.Getenv("GOPATH"))
}

func TestOsExec(t *testing.T) {
	os.Chdir("/tmp/dirt")
	cmdPath, _ := exec.LookPath("ls")
	cmd := exec.Command(cmdPath, "-al")
	byte, CombinedOutErr := cmd.CombinedOutput()
	if CombinedOutErr != nil {
		t.Fatal(CombinedOutErr)
	}
	t.Log(string(byte))

	t.Log(cmd.ProcessState.ExitCode())
	t.Log(cmd.ProcessState.String())
	t.Log(cmd.ProcessState.Pid())
	t.Log(cmd.ProcessState.Success())
}

func TestOsSignal(t *testing.T) {
	fmt.Println(os.Getpid())
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan)
	var isStart bool = true
	for isStart {
		select {
		case sig := <-signalChan:
			fmt.Println(sig.String())
			if sig == syscall.SIGQUIT || sig == syscall.SIGSTOP {
				isStart = false
			}
		}
	}
	// kill -USR2 pid
	// kill -QUIT pid
}

func TestUserAndGroup(t *testing.T) {
	t.Log(user.Lookup("znddzxx112"))
	t.Log(user.Current())
}
