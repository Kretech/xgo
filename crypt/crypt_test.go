package crypt

import (
	"testing"

	"github.com/Kretech/common.go/test"
)

func TestMd5(t *testing.T) {
	test.AssertEqual(t, Md5(`hello`), `5d41402abc4b2a76b9719d911017c592`)
}

func TestSha1(t *testing.T) {
	test.AssertEqual(t, Sha1(`hello`), `aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d`)
}
