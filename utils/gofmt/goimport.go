package gofmt

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/tools/imports"
)

var (
	// main operation modes
	list   = new(bool)
	write  = new(bool)
	doDiff = new(bool)
	srcdir = new(string)

	options = &imports.Options{
		TabWidth:  8,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
	}
)

// argumentType is which mode goimports was invoked as.
type argumentType int

const (
	// fromStdin means the user is piping their source into goimports.
	fromStdin argumentType = iota

	// singleArg is the faas_common case from editors, when goimports is run on
	// a single file.
	singleArg

	// multipleArg is when the user ran "goimports file1.go file2.go"
	// or ran goimports on a directory tree.
	multipleArg
)

func processFile(filename string, in io.Reader, out io.Writer, argType argumentType) error { // cbt_skip
	opt := options
	if argType == fromStdin {
		nopt := *options
		nopt.Fragment = true
		opt = &nopt
	}

	if in == nil {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		in = f
	}

	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	target := filename
	if *srcdir != "" {
		// Determine whether the provided -srcdirc is a directory or file
		// and then use it to override the target.
		//
		// See https://github.com/dominikh/go-mode.el/issues/146
		if isFile(*srcdir) {
			if argType == multipleArg {
				return errors.New("-srcdir value can't be a file when passing multiple arguments or when walking directories")
			}
			target = *srcdir
		} else if argType == singleArg && strings.HasSuffix(*srcdir, ".go") && !isDir(*srcdir) {
			// For a file which doesn't exist on disk yet, but might shortly.
			// e.g. user in editor opens $DIR/newfile.go and newfile.go doesn't yet exist on disk.
			// The goimports on-save hook writes the buffer to a temp file
			// first and runs goimports before the actual save to newfile.go.
			// The editor's buffer is named "newfile.go" so that is passed to goimports as:
			//      goimports -srcdir=/gopath/src/pkg/newfile.go /tmp/gofmtXXXXXXXX.go
			// and then the editor reloads the result from the tmp file and writes
			// it to newfile.go.
			target = *srcdir
		} else {
			// Pretend that file is from *srcdir in order to decide
			// visible imports correctly.
			target = filepath.Join(*srcdir, filepath.Base(filename))
		}
	}

	res, err := imports.Process(target, src, opt)
	if err != nil {
		return err
	}

	if !bytes.Equal(src, res) {
		// formatting has changed
		if *list {
			fmt.Fprintln(out, filename)
		}
		if *write {
			if argType == fromStdin {
				// filename is "<standard input>"
				return errors.New("can't use -w on stdin")
			}
			err = ioutil.WriteFile(filename, res, 0)
			if err != nil {
				return err
			}
		}
		if *doDiff {
			if argType == fromStdin {
				filename = "stdin.go" // because <standard input>.orig looks silly
			}
			data, err := diff(src, res, filename)
			if err != nil {
				return fmt.Errorf("computing diff: %s", err)
			}
			fmt.Printf("diff -u %s %s\n", filepath.ToSlash(filename+".orig"), filepath.ToSlash(filename))
			out.Write(data)
		}
	}

	if !*list && !*write && !*doDiff {
		_, err = out.Write(res)
	}

	return err
}

func GoFmtInMem(fileSource string) string {

	buf := bytes.NewBufferString("")
	in := bytes.NewBufferString(fileSource)
	if err := processFile("", in, buf, singleArg); err != nil {
		panic(err)
	}
	return buf.String()
}

func GoFmtInMemForBytes(fileSource []byte) []byte {

	buf := bytes.NewBufferString("")
	in := bytes.NewBuffer(fileSource)
	if err := processFile("", in, buf, singleArg); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func GofmtMain(filepath string) string {
	buf := bytes.NewBufferString("")
	if err := processFile(filepath, nil, buf, singleArg); err != nil {
		panic(err)
	}
	return buf.String()
}

func writeTempFile(dir, prefix string, data []byte) (string, error) {
	file, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		return "", err
	}
	_, err = file.Write(data)
	if err1 := file.Close(); err == nil {
		err = err1
	}
	if err != nil {
		os.Remove(file.Name())
		return "", err
	}
	return file.Name(), nil
}

func diff(b1, b2 []byte, filename string) (data []byte, err error) {
	f1, err := writeTempFile("", "gofmt", b1)
	if err != nil {
		return
	}
	defer os.Remove(f1)

	f2, err := writeTempFile("", "gofmt", b2)
	if err != nil {
		return
	}
	defer os.Remove(f2)

	cmd := "diff"
	if runtime.GOOS == "plan9" {
		cmd = "/bin/ape/diff"
	}

	data, err = exec.Command(cmd, "-u", f1, f2).CombinedOutput()
	if len(data) > 0 {
		// diff exits with a non-zero status when the files don't match.
		// Ignore that failure as long as we get output.
		return replaceTempFilename(data, filename)
	}
	return
}

// replaceTempFilename replaces temporary filenames in diff with actual one.
//
// --- /tmp/gofmt316145376	2017-02-03 19:13:00.280468375 -0500
// +++ /tmp/gofmt617882815	2017-02-03 19:13:00.280468375 -0500
// ...
// ->
// --- path/to/file.go.orig	2017-02-03 19:13:00.280468375 -0500
// +++ path/to/file.go	2017-02-03 19:13:00.280468375 -0500
// ...
func replaceTempFilename(diff []byte, filename string) ([]byte, error) {
	bs := bytes.SplitN(diff, []byte{'\n'}, 3)
	if len(bs) < 3 {
		return nil, fmt.Errorf("got unexpected diff for %s", filename)
	}
	// Preserve timestamps.
	var t0, t1 []byte
	if i := bytes.LastIndexByte(bs[0], '\t'); i != -1 {
		t0 = bs[0][i:]
	}
	if i := bytes.LastIndexByte(bs[1], '\t'); i != -1 {
		t1 = bs[1][i:]
	}
	// Always print filepath with slash separator.
	f := filepath.ToSlash(filename)
	bs[0] = []byte(fmt.Sprintf("--- %s%s", f+".orig", t0))
	bs[1] = []byte(fmt.Sprintf("+++ %s%s", f, t1))
	return bytes.Join(bs, []byte{'\n'}), nil
}

// isFile reports whether name is a file.
func isFile(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.Mode().IsRegular()
}

// isDir reports whether name is a directory.
func isDir(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.IsDir()
}

func PluralToSingular(word string) string {
	// 以 "s" 结尾的单词，直接去掉 "s"（适用于大部分简单复数情况，如 books -> book）
	if strings.HasSuffix(word, "s") && len(word) > 1 {
		return word[:len(word)-1]
	}
	// 以 "es" 结尾的单词，去掉 "es"（如 boxes -> box），但需排除一些特殊情况，这里简单判断单词长度大于3
	if strings.HasSuffix(word, "es") && len(word) > 3 {
		return word[:len(word)-2]
	}
	// 以 "ies" 结尾的单词，将 "ies" 变为 "y"（如 cities -> city），同样简单判断单词长度大于4
	if strings.HasSuffix(word, "ies") && len(word) > 4 {
		return word[:len(word)-3] + "y"
	}
	return word
}
