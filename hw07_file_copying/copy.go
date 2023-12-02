package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Place your code here.
	fileFrom, err := os.Open(fromPath)
	if err != nil {
		if os.IsNotExist(err) {
			// файл не найден
			return err
		}
	}
	defer fileFrom.Close()
	fi, err := fileFrom.Stat()
	if err != nil {
		return err
	}
	if fi.Size() < limit+offset {
		limit = 0
	}
	if offset > fi.Size() {
		return err
	}
	if _, err := os.Stat(toPath); err == nil {
		// path/to/whatever exists
		os.Remove(toPath)
	}
	fileTo, err := os.Create(toPath)
	if err != nil {
		return err
	}

	defer fileTo.Close()
	fileFrom.Seek(offset, io.SeekStart)

	if limit == 0 {
		b, _ := io.ReadAll(fileFrom)
		fileTo.Write(b)
	} else {
		b := make([]byte, limit)
		fileFrom.Read(b)
		fileTo.Write(b)
	}
	return nil
}
