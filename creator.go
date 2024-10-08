package cli_cleancode_manager

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type file struct {
	content    string
	tmplPath   string
	targetPath string
	filename   string
	params     any
}

func create(f file) error {
	err := os.MkdirAll(f.targetPath, os.ModeDir)
	if err != nil {
		return err
	}

	t, err := template.New("tmpl").Parse(f.content)
	if err != nil {
		return err
	}
	ff, err := os.Create(filepath.Join(f.targetPath, f.filename))
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ff)

	err = t.Execute(ff, f.params)
	if err != nil {
		return err
	}

	return nil
}
