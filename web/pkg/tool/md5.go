package tool

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5ByString(str string) (string, error) {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		return "", err
	}

	arr := m.Sum(nil)

	return fmt.Sprintf("%x", arr), nil
}
