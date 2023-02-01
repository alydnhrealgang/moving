package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const EmptyString = ""

var EmptyStrings = make([]string, 0, 0)

func EmptyOrWhiteSpace(text interface{}) bool {
	switch text.(type) {
	case string:
		return strings.TrimSpace(text.(string)) == EmptyString
	case *string:
		value := text.(*string)
		return nil == value || EmptyOrWhiteSpace(*value)
	}

	return nil == text
}

func Compare(left, right *string, ignoreCase bool) int {
	if left == right {
		return 0
	}

	if nil == left && nil != right {
		return -1
	}

	if nil != left && nil == right {
		return 1
	}

	if ignoreCase {
		return strings.Compare(strings.ToUpper(*left), strings.ToUpper(*right))
	}

	return strings.Compare(*left, *right)
}

func Ptr2Str(text *string) (string, bool) {
	if nil == text {
		return EmptyString, false
	}

	return *text, true
}

func StrPtr2Ptr(getter func() (string, bool)) *string {
	if value, ok := getter(); ok {
		return &value
	}

	return nil

}

func ParseValue(value interface{}, stringValue string) (err error) {
	switch v := value.(type) {
	case *int:
		*v, err = strconv.Atoi(stringValue)
	case *int32:
		var iv int64
		if iv, err = strconv.ParseInt(stringValue, 10, 32); nil == err {
			*v = int32(iv)
		}
	case *int64:
		*v, err = strconv.ParseInt(stringValue, 10, 64)
	case *uint32:
		var uv uint64
		if uv, err = strconv.ParseUint(stringValue, 10, 32); nil == err {
			*v = uint32(uv)
		}
	case *uint64:
		*v, err = strconv.ParseUint(stringValue, 10, 64)
	case *string:
		*v = stringValue
	case *bool:
		*v, err = strconv.ParseBool(stringValue)
	case *float32:
		var fv float64
		if fv, err = strconv.ParseFloat(stringValue, 32); nil == err {
			*v = float32(fv)
		}
	case *float64:
		*v, err = strconv.ParseFloat(stringValue, 64)
	default:
		err = fmt.Errorf("INVALID_TYPE")
	}

	return
}

func FormatValue(value interface{}) string {
	switch v := value.(type) {
	case bool:
		return strconv.FormatBool(v)
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', 3, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', 3, 64)
	default:
		return fmt.Sprint(v)
	}
}

func RandomID() (string, error) {
	bytes := make([]byte, 64)
	_, err := rand.Read(bytes)
	if nil != err {
		return EmptyString, err
	}
	text := base64.URLEncoding.EncodeToString(bytes)
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func EncodeURIComponent(uri string) string {
	r := url.QueryEscape(uri)
	return strings.Replace(r, "+", "%20", -1)
}
