package marshalfs

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"testing/fstest"
)

func TestMarshalFS(t *testing.T) {
	const f0 = "your/file"
	const f1 = "my/file"
	const glob2 = "their/*"
	const f2 = "their/file"
	m := &FS{
		Marshal: json.Marshal,
		Files: map[string]*File{
			f0: {
				value: struct {
					Thingy []byte
					Number int
				}{Thingy: []byte("hello, world\n"), Number: 10},
			},
			f1: {value: struct{ Info string }{"Some interesting info.\n"}},
			//glob2: {Value: struct{ Info string }{"Some globbed info.\n"}},
		},
		Patterns: map[string]Generator{
			glob2: func(name string) (*File, error) {
				return &File{value: struct{ Info string }{"Some globbed info.\n"}}, nil
			},
		},
	}
	for _, fn := range []string{f0, f1, f2} {
		t.Run(fn, func(t *testing.T) {
			f, err := m.Open(fn)
			if err != nil {
				t.Fatal(err)
			}
			b, err := ioutil.ReadAll(f)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("file: %s", fn)
			t.Logf("file contents: %s", string(b))
		})
	}

	t.Run("TestFS", func(t *testing.T) {
		if err := fstest.TestFS(m, f0, f1); err != nil {
			t.Fatal(err)
		}
	})

}
