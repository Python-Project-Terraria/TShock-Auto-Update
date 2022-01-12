package ziper

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unpack(file string, out string) error {
	err := Unzip(file, out)
	if err != nil {
		return err
	}
	os.Remove(file)
	return nil
}

func Unzip(file string, out string) error {
	zf, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer zf.Close()

	for _, f := range zf.File {
		path := filepath.Join(out, f.Name)

		if f.FileInfo().IsDir() {
			os.Mkdir(path, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer file.Close()

		arcive, err := f.Open()
		if err != nil {
			return err
		}
		defer arcive.Close()

		if _, err := io.Copy(file, arcive); err != nil {
			return err
		}
	}

	return nil
}
