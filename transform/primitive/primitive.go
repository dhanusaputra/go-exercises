package primitive

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Mode ...
type Mode int

// All mode ...
const (
	Combo Mode = iota
	Triangle
	Rect
	Ellipse
	Cicle
	Rotatedrect
	Beziers
	Rotatedellipse
	Polygon
)

// WithMode ...
func WithMode(mode Mode) func() []string {
	return func() []string {
		return []string{"-n", fmt.Sprintf("%d", mode)}
	}
}

// Transform ...
func Transform(image io.Reader, ext string, numShapes int, opts ...func() []string) (io.Reader, error) {
	var args []string
	for _, opt := range opts {
		args = append(args, opt()...)
	}
	in, err := tempfile("in_", ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	out, err := tempfile("in_", ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(out.Name())

	_, err = io.Copy(in, image)
	if err != nil {
		return nil, err
	}

	std, err := primitive(in.Name(), out.Name(), numShapes, args...)
	if err != nil {
		return nil, err
	}
	fmt.Println(std)

	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, out)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func primitive(inputFile, outputFile string, numShapes int, args ...string) (string, error) {
	argStr := fmt.Sprintf("-i %s -o %s -n %d -m %d", inputFile, outputFile, numShapes)
	args = append(strings.Fields(argStr), args...)
	cmd := exec.Command("primitive", args...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}

func tempfile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("", prefix)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
