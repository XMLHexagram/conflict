package pkg

import (
	"bytes"
	"embed"
	"io/ioutil"
	"os"
	"path"
)

//go:embed template
var Template embed.FS
var ignore = []string{".DS_Store", "template.go", "template_test.go"}
var normalFileMode = os.FileMode(0644)
var normalDirMode = os.FileMode(0755)

type replace struct {
	before string
	after  string
}

var replaces = []replace{
	{
		before: "conflict",
		after:  "test",
	},
}

func GenerateDir(src, dst string) error {
	if err := os.MkdirAll(dst, normalDirMode); err != nil {
		return err
	}

	dirs, err := Template.ReadDir(src)
	if err != nil {
		return err
	}

	for _, v := range dirs {
		if hasSets(v.Name(), ignore) {
			continue
		}

		srcfp := path.Join(src, v.Name())
		dstfp := path.Join(dst, v.Name())

		if v.IsDir() {
			err = GenerateDir(srcfp, dstfp)
			if err != nil {
				return err
			}
		} else {
			err := GenerateFile(srcfp, dstfp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GenerateFile(src, dst string) error {
	buf, err := Template.ReadFile(src)
	if err != nil {
		return err
	}
	for _, v := range replaces {
		buf = bytes.ReplaceAll(buf, []byte(v.before), []byte(v.after))
	}
	err = ioutil.WriteFile(dst, buf, normalFileMode)
	if err != nil {
		return err
	}
	return nil
}

func hasSets(name string, sets []string) bool {
	for _, ig := range sets {
		if ig == name {
			return true
		}
	}
	return false
}
