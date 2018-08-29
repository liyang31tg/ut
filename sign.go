package ut

import (
	"crypto/sha1"
 	"crypto/sha256"
 	"fmt"
 	)

func SignSha1(str string) string {
	s := sha1.New()
	s.Write([]byte(str))
	return fmt.Sprintf("%x",s.Sum(nil))
}

func SignSha256(str string) string {
	s := sha256.New()
	s.Write([]byte(str))
	return fmt.Sprintf("%x",s.Sum(nil))
}
