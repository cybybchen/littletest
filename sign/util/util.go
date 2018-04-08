package util

import (
	"net/url"
	"bytes"
	"sort"
	"crypto/md5"
	"encoding/hex"
	"strings"
	"fmt"
)

func CheckAuth(values url.Values, key string) bool {
	sign := values.Get("verify")
	fmt.Println(sign)
	values.Del("verify")
	values.Del("action")
	s := SortAndJoin(values)
	s += key
	//s = "account=1618532711545162&sid=1448420002&time=1519808359&uuzu_op_id=592CejSh4i7kifA1NZn"
	fmt.Println(s)
	h := md5.New()
	h.Write([]byte(s))
	ver := hex.EncodeToString(h.Sum(nil))
	ver = strings.ToLower(ver)
	fmt.Println(ver)
	return ver == sign
}

func SortAndJoin(v url.Values) string {
	if v == nil {
		return ""
	}

	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(v)
		}
	}

	return buf.String()
}
