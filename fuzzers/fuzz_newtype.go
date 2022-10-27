package myfuzz

import "github.com/h2non/filetype"

var fooType = filetype.NewType("foo", "foo/foo")

func fooMatcher(buf []byte) bool {
  return len(buf) > 1 && buf[0] == 0x01 && buf[1] == 0x02
}

func Fuzz(data[] byte) int {
  // Register the new matcher and its type
  filetype.AddMatcher(fooType, fooMatcher)

  // Try to match the file
  //fooFile := []byte{0x01, 0x02}
  kind, _ := filetype.Match(data)
  if kind == filetype.Unknown {
    //fmt.Println("Unknown file type")
	  return 1
  }
  return 0
}
