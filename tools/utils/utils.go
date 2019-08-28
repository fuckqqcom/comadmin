package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"time"
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

func Time2Str(format string) string {
	now := time.Now()
	//formatNow := now.Format("2006-01-02 15:04:05")
	formatNow := now.Format(format)
	fmt.Println(formatNow)
	return formatNow
}

func NowTime() time.Time {
	timeUnix := time.Now().Format("2006-01-02")
	location, _ := time.ParseInLocation("2006-01-02", timeUnix, time.Local)
	return location
}

func Str2Time(format, value string) time.Time {
	local, _ := time.LoadLocation("Local")
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-06-20 18:16:15", local)
	t, _ := time.ParseInLocation(format, value, local)

	fmt.Println(t)
	return t
}

/**
打印es sql
*/
func PrintQuery(src interface{}) {
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	log.Printf("es sql--->%s", string(data))
}
