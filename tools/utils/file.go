package utils

import "os"

func FileMode(p string) (fm os.FileMode, err error) {
	s, err := FileInfo(p)
	if err != nil {
		return fm, err
	}
	fm = s.Mode()
	return fm, nil
}

func FileInfo(p string) (fi os.FileInfo, err error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Stat()
}
