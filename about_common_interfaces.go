package go_koans

import (
	"bytes"
	"fmt"
	"io"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")
		//defer in.Close()

		out := new(bytes.Buffer)
		//defer out.Close()
		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		bytes, err := io.Copy(out, in)

		if err != nil {
			fmt.Println("error ", err.Error())
		}

		fmt.Println("ke copy ", bytes)

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		io.CopyN(out, in,5 )

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
