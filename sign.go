package ut

import (
	"crypto/sha1"
 	"crypto/sha256"
 	"fmt"
	"os"
	"crypto/md5"
	"io"
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

func Md5File(file string) (string,error){
	f, err := os.Open(file)
	if err != nil {
		return "" ,err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "",err
	}
	return fmt.Sprintf("%x", h.Sum(nil)),nil
}

func Md5FileReader(r io.Reader) (string,error){
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		return "",err
	}
	return fmt.Sprintf("%x", h.Sum(nil)),nil
}
