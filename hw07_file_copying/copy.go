package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
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
		return fmt.Errorf("смещение за границами файла")
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
		count := int(fi.Size())
		bar := pb.StartNew(count)
		b, _ := io.ReadAll(fileFrom)
		fileTo.Write(b)
		bar.Add(count)
		bar.Finish()
	} else {
		count := int(limit)
		bar := pb.StartNew(count)
		io.CopyN(fileTo, fileFrom, limit)
		bar.Add(count)
		bar.Finish()
	}

	return nil
}
