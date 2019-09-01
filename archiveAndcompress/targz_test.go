package archiveAndcompress

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestGzip(t *testing.T) {
	ChdirErr := os.Chdir("./targzdir")
	if ChdirErr != nil {
		t.Fatal(ChdirErr)
	}

	file, OpenErr := os.Open("privoxy-3.0.28-stable-src.tar.gz")
	if OpenErr != nil {
		t.Fatal(OpenErr)
	}
	defer file.Close()

	gp, NewReaderErr := gzip.NewReader(file)
	if NewReaderErr != nil {
		t.Fatal(NewReaderErr)
	}
	defer gp.Close()

	createFile, CreateErr := os.Create(gp.Name)
	if CreateErr != nil {
		t.Fatal(CreateErr)
	}
	defer createFile.Close()

	io.Copy(createFile, gp)

}

func TestTar(t *testing.T) {
	ChdirErr := os.Chdir("./targzdir")
	if ChdirErr != nil {
		t.Fatal(ChdirErr)
	}

	file, OpenErr := os.Open("privoxy-3.0.28-stable-src.tar")
	if OpenErr != nil {
		t.Fatal(OpenErr)
	}
	defer file.Close()

	tr := tar.NewReader(file)

	for true {
		trHeader, NextErr := tr.Next()
		if NextErr != nil {
			if NextErr == io.EOF {
				break
			} else {
				t.Fatal(NextErr)
				break
			}
		}

		nameBytes := []byte(trHeader.Name)
		dir := nameBytes[:bytes.LastIndex(nameBytes, []byte("/"))]

		if MkdirAllErr := os.MkdirAll(string(dir), 0755); MkdirAllErr != nil {
			t.Fatal(MkdirAllErr)
			break
		}

		createFile, CreateErr := os.Create(trHeader.Name)
		if CreateErr != nil {
			t.Fatal(CreateErr)
			break
		}

		io.Copy(createFile, tr)

		createFile.Close()
	}

}

func TestTarEncode(t *testing.T) {
	ChdirErr := os.Chdir("./targzdir")
	if ChdirErr != nil {
		t.Fatal(ChdirErr)
	}

	tarfile, CreateErr := os.Create("privoxy-3.0.28-stable-new.tar")
	if CreateErr != nil {
		t.Fatal(CreateErr)
	}
	defer tarfile.Close()

	tr := tar.NewWriter(tarfile)
	defer tr.Close()

	fileInfos, ReadDirErr := ioutil.ReadDir("privoxy-3.0.28-stable")
	if ReadDirErr != nil {
		t.Fatal(ReadDirErr)
	}

	for _, fileinfo := range fileInfos {
		tarEncode(fileinfo, "privoxy-3.0.28-stable", tr)
	}
}

func tarEncode(fileinfo os.FileInfo, prefix string, tr *tar.Writer) {
	if fileinfo.IsDir() {
		fileInfos, ReadDirErr := ioutil.ReadDir(prefix + "/" + fileinfo.Name())
		if ReadDirErr != nil {
			return
		}
		for _, fi := range fileInfos {
			tarEncode(fi, prefix+"/"+fileinfo.Name(), tr)
		}
	} else {
		f, OpenErr := os.Open(prefix + "/" + fileinfo.Name())
		if OpenErr != nil {
			return
		}

		header, FileInfoHeaderErr := tar.FileInfoHeader(fileinfo, "")
		if FileInfoHeaderErr != nil {
			return
		}
		header.Name = prefix + "/" + fileinfo.Name()
		WriteHeaderErr := tr.WriteHeader(header)
		if WriteHeaderErr != nil {
			return
		}

		io.Copy(tr, f)
		f.Close()
	}
}

func TestGzipCompress(t *testing.T) {
	ChdirErr := os.Chdir("./targzdir")
	if ChdirErr != nil {
		t.Fatal(ChdirErr)
	}

	tarfile, OpenErr := os.Open("privoxy-3.0.28-stable-new.tar")
	if OpenErr != nil {
		t.Fatal(OpenErr)
	}
	defer tarfile.Close()

	gzfile, CreateErr := os.Create("privoxy-3.0.28-stable-new.tar.gz")
	if CreateErr != nil {
		t.Fatal(CreateErr)
	}
	defer gzfile.Close()

	gzWrite := gzip.NewWriter(gzfile)
	defer gzWrite.Close()

	fileinfo, StatErr := tarfile.Stat()
	if StatErr != nil {
		t.Fatal(StatErr)
	}

	gzWrite.Name = fileinfo.Name()
	gzWrite.ModTime = fileinfo.ModTime()

	io.Copy(gzWrite, tarfile)

}
