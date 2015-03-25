fenc 
======

[![GoDoc](https://godoc.org/github.com/djherbis/fenc?status.svg)](https://godoc.org/github.com/djherbis/fenc)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.txt)

Usage
-----

```go
import(
  "bytes"
  "fmt"
  "github.com/djherbis/fenc"
)

// Buffer as a network stand-in
buf := bytes.NewBuffer(nil)

// Stream encode your file onto your "network". 
// You can call this multiple times to encode multiple files in the same stream
enc := fenc.NewEncoder(buf)
if err := enc.OpenAndEncode("fenc.go"); err != nil {
  fmt.Println(err.Error())
  return
}
// enc.Encode(os.File), and enc.CustomEncode(os.FileInfo, io.Reader) are also available

// Stream decode your file on the other end of your "network".
// err == io.EOF when there are no more files.
dec := fenc.NewDecoder(buf)
fi, r, err := dec.Decode()
if err != nil {
  fmt.Println(err.Error())
  return
}
// fi is a os.FileInfo for your file, and r is a streaming reader for the file

```

Installation
------------
```sh
go get github.com/djherbis/fenc
```
