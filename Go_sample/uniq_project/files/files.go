package files

import (
	. "awesomeProject/uniq_project/ans_struct"
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

type MyFileR struct {
	FileName string
}
type MyFileW struct {
	FileName string
}

func (f MyFileR) OpenFile() (*os.File, error) {
	flS, err := os.OpenFile(f.FileName, os.O_RDONLY, os.ModePerm)
	if f.FileName == "" {
		return os.Stdin, nil
	}
	if err == nil {
		return flS, nil
	} else {
		return flS, err
	}
}

func (f MyFileW) OpenFile() (*os.File, error) {
	flS, err := os.OpenFile(f.FileName, os.O_WRONLY, os.ModePerm)
	if f.FileName == "" {
		return os.Stdout, nil
	}
	if err == nil {
		return flS, nil
	} else {
		return flS, err
	}
}

type Reader interface {
	OpenFile() (*os.File, error)
}

func Scan(r Reader) []string {
	f, _ := r.OpenFile()
	scanner := bufio.NewScanner(f)
	arr := make([]string, 0)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	return arr
}

type Writer interface {
	OpenFile() (*os.File, error)
}

func Clear(w Writer) error {
	f, err := w.OpenFile()
	if err != nil {
		return err
	}
	f.Seek(0, io.SeekStart)
	f.Truncate(0)
	return nil
}

func Write(w Writer, ans []AnsStruct, flag_c bool) error {
	f, err := w.OpenFile()
	if err != nil {
		return err
	}
	for i := 0; i < len(ans); i++ {
		var str string
		if flag_c {
			str = strconv.Itoa(ans[i].Count) + " " + ans[i].Str + "\n"
		} else {
			str = ans[i].Str + "\n"
		}
		_, err = f.WriteString(str)
		if err != nil {
			log.Fatal("Can't write to file")
			return err
		}
	}
	return nil
}
