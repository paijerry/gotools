//Package wtfhash - key/value format & MD5
package wtfhash

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"sort"
	"strconv"
)

//JSONtoMD5 - json request generate hash
func JSONtoMD5(j []byte, sKey string) string {

	var m map[string]interface{}

	json.Unmarshal(j, &m)

	result := MaptoMD5(m, sKey)

	return result
}

//MaptoMD5 - map data generate hash
func MaptoMD5(m map[string]interface{}, sKey string) string {
	mk := make([]string, len(m))
	i := 0
	for k := range m {
		mk[i] = k
		i++
	}
	sort.Strings(mk)

	var str string
	s := "&"
	for i, key := range mk {
		if i >= len(mk)-1 {
			s = sKey
		}

		str += (key + "=")
		switch m[key].(type) {
		case string:
			str += m[key].(string)
		case int:
			str += strconv.Itoa(m[key].(int))
		case bool:
			str += strconv.FormatBool(m[key].(bool))
		case float64:
			str += strconv.FormatFloat(m[key].(float64), 'f', -1, 64)
		}
		str += s
	}
	//fmt.Println(str)
	//MD5 encode
	h := md5.New()
	io.WriteString(h, str)
	hByte := h.Sum(nil)

	//convert [Size]byte to string
	result := hex.EncodeToString(hByte)

	return result
}
