//Package golangexamples has examples of golang
package golangexamples

import (
	"github.com/ehteshamz/greetings"
)

//ConcatSlice is a public function to concatenate byte array
func ConcatSlice(sliceToConcat []byte) string {

	str := ""
	for i, _ := range sliceToConcat {
		str += string(sliceToConcat[i])
		if i != len(sliceToConcat)-1 {
			str += "-"
		}
	}

	return str
}

//Encrypt is a public function
func Encrypt(sliceToConcat []byte, ceaserCount int) {

	for i, _ := range sliceToConcat {
		sliceToConcat[i] = sliceToConcat[i] + byte(ceaserCount)
	}
}

//EZGreetings is a public function
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
