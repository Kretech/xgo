package crypt

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func Md5(s string) string {
	return fmt.Sprintf(`%x`, md5.Sum([]byte(s)))
}

func Sha1(s string) string {
	return fmt.Sprintf(`%x`, sha1.Sum([]byte(s)))
}
