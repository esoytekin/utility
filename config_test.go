package utility

import (
	"os"
	"testing"

	"gotest.tools/assert"
)

type Config struct {
	Name string `json:"name"`
}

const appName = "dovetail"

func TestReadConfig(t *testing.T) {

	c := new(Config)

	open = func(p string) (*os.File, error) {
		f, _ := os.Create("text.json")
		f.WriteString("{\"name\":\"x\"}")
		err := f.Close()

		f, _ = os.Open("text.json")

		return f, err
	}
	err := ReadConfiguration(appName, c)

	os.Remove("text.json")
	assert.NilError(t, err, "not nil error")

	assert.Assert(t, c.Name == "x", "expected %q, got %q", "x", c.Name)

}
