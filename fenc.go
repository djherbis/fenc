package fenc

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) OpenAndEncode(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return e.Encode(f)
}

func (e *Encoder) Encode(f *os.File) (err error) {
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	return e.CustomEncode(fi, f)
}

func (e *Encoder) CustomEncode(fi os.FileInfo, r io.Reader) (err error) {
	tdata, err := fi.ModTime().MarshalBinary()
	if err != nil {
		return err
	}
	fmt.Fprintf(e.w, "%s\n%d\n%d\n%d\n", fi.Name(), fi.Size(), fi.Mode(), len(tdata))
	_, err = e.w.Write(tdata)
	if err != nil {
		return err
	}
	_, err = io.Copy(e.w, r)
	return err
}

func NewFileInfo(name string, size int64, mode os.FileMode, modtime time.Time) os.FileInfo {
	return &fileInfo{
		name:    name,
		size:    size,
		mode:    mode,
		modtime: modtime,
	}
}

type fileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modtime time.Time
}

func (fi *fileInfo) Name() string       { return fi.name }
func (fi *fileInfo) Size() int64        { return fi.size }
func (fi *fileInfo) Mode() os.FileMode  { return fi.mode }
func (fi *fileInfo) ModTime() time.Time { return fi.modtime }
func (fi *fileInfo) IsDir() bool        { return fi.mode.IsDir() }
func (fi *fileInfo) Sys() interface{}   { return nil }

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (d *Decoder) Decode() (ofi os.FileInfo, r io.Reader, err error) {
	var fi fileInfo
	var tlen int
	if _, err = fmt.Fscanf(d.r, "%s\n%d\n%d\n%d\n", &fi.name, &fi.size, &fi.mode, &tlen); err != nil {
		return nil, nil, err
	}
	tdata := make([]byte, tlen)
	_, err = d.r.Read(tdata)
	if err != nil {
		return nil, nil, err
	}
	if err = fi.modtime.UnmarshalBinary(tdata); err != nil {
		return nil, nil, err
	}
	return &fi, io.LimitReader(d.r, fi.size), nil
}

func Discard(r io.Reader) {
	io.Copy(ioutil.Discard, r)
}
