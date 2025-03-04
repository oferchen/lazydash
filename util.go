//MIT License

//Copyright (c) 2019 Jason Witting

//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:

//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.

//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//LoadFromFile loads prometheus metrics fromfile
func LoadFromFile(path string) ([]byte, error) {
	promFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	b, err := ioutil.ReadAll(promFile)

	if err != nil {
		log.Fatalf("could not read from file: %v", err)
	}

	return b, nil
}

//LoadFromStdin loads prometheus metrics from STDIN
func LoadFromStdin() []byte {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	//is a pipe?
	if info.Mode()&os.ModeNamedPipe == 0 {
		//is a redirect?
		if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
			return nil
		}
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return ([]byte(string(output)))
}
