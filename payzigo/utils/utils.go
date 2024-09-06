package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

func GenerateAuthorizationHeader(apiKey, hashedPki string) string {
	return fmt.Sprintf("IYZWS %s:%s", apiKey, hashedPki)
}

func HashSha1(pki string) string {
	hasher := sha1.New()
	hasher.Write([]byte(pki))
	sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func GenerateRandomString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// PKI String (apiKey + x-iyzi-rnd + secretKey + requestString)
func GeneratePKIString(apiKey string, randomString string, secretKey string, requestString string) string {
	return apiKey + randomString + secretKey + requestString
}

func GenerateRequestString(v any) (res string) {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Struct {
		return fmt.Sprintf("%v", v)
	}

	num := rv.NumField()
	for i := 0; i < num; i++ {
		fv := rv.Field(i)
		st := rv.Type().Field(i)
		jsonName, err := st.Tag.Lookup("json")
		if strings.Contains(jsonName, ",omitempty") {
			continue
		}
		if err != true {
			res += st.Name + "="
		} else {
			res += fmt.Sprintf("%s=", jsonName)
		}
		switch fv.Kind() {
		case reflect.String:
			res += fmt.Sprintf("%s", fv.String())
		case reflect.Int:
			res += fmt.Sprint(fv.Int())
		case reflect.Struct:
			res += GenerateRequestString(fv.Interface())
		case reflect.Slice:
			res += "["
			for k := 0; k < fv.Len(); k++ {
				e := fv.Index(k)
				res += GenerateRequestString(e.Interface())
				if k != fv.Len()-1 {
					res += ", "
				}
			}
			res += "]"
		}
		if i != num-1 {
			res += ","
		}
	}

	if res[len(res)-1:] == "," {
		res = res[0 : len(res)-1]
	}
	res = "[" + res + "]"

	return res
}
