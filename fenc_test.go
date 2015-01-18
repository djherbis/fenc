package fenc

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestSelf(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	e := NewEncoder(buf)
	for i := 0; i < 3; i++ {
		if err := e.OpenAndEncode("fenc.go"); err != nil {
			t.Error(err)
		}
	}

	d := NewDecoder(buf)
	for {
		fi, r, er := d.Decode()

		if er != nil {
			if er != io.EOF {
				t.Error(er)
			}
			return
		}

		fmt.Println(fi.Name(), fi.Size(), fi.Mode(), fi.ModTime(), fi.IsDir())
		Discard(r)
	}
}
