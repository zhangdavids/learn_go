package lib

import (
	"os"
	in "puzzlers/lib2/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
