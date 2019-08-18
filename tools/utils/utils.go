package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"log"
	"regexp"
)

const (
	str = `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
)

func CheckError(err error, v interface{}) bool {
	if err != nil {
		log.Printf("err is %s,%s", err, v)
		return false
	}
	return true
}

func EncodeMd5(value string) string {

	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func SqlRegex(param string) bool {

	re, err := regexp.Compile(str)
	if err != nil {
		log.Printf("error param error is %s", err.Error())
		return true
	}

	return re.MatchString(param)
}

func StringJoin(a ...string) string {
	var buf bytes.Buffer
	for _, k := range a {
		buf.WriteString(k)
	}
	return buf.String()
}
