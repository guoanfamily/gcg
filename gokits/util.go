package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	//"golang.org/x/text/encoding/simplifiedchinese"
	//"golang.org/x/text/transform"

	// 支持gif,jpeg,png的处理
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var (
	// GZipMinSize 要启用gzip的最小值
	GZipMinSize = 1024
)

// ToGzipJSON 转换成json并gzip
func ToGzipJSON(obj interface{}) (gziped bool, data []byte, err error) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return
	}
	if len(bs) <= GZipMinSize {
		return false, bs, nil
	}
	buf := &bytes.Buffer{}
	gzipWriter := gzip.NewWriter(buf)
	_, err = gzipWriter.Write(bs)
	gzipWriter.Close()
	if err != nil {
		return
	}
	return true, buf.Bytes(), nil
}

// ToGzip gzip
func ToGzip(bs []byte) (gziped bool, data []byte, err error) {
	if len(bs) <= GZipMinSize {
		return false, bs, nil
	}
	buf := &bytes.Buffer{}
	gzipWriter := gzip.NewWriter(buf)
	_, err = gzipWriter.Write(bs)
	gzipWriter.Close()
	if err != nil {
		return
	}
	return true, buf.Bytes(), nil
}

// FromGzip from gzip
func FromGzip(data []byte) (bs []byte, err error) {
	buf := bytes.NewBuffer(data)
	gzipReader, err := gzip.NewReader(buf)
	if err != nil {
		return
	}
	defer gzipReader.Close()
	bs, err = ioutil.ReadAll(gzipReader)
	return
}

// FromGzipToJSON from gzip json
func FromGzipToJSON(data []byte, obj interface{}) (err error) {
	buf := bytes.NewBuffer(data)
	gzipReader, err := gzip.NewReader(buf)
	if err != nil {
		return
	}
	defer gzipReader.Close()
	bs, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return
	}
	return json.Unmarshal(bs, obj)
}

// ToJSONString to json string
func ToJSONString(data interface{}) string {
	bs, err := json.Marshal(data)
	if err != nil {
		return "<can't serialize to json>"
	}
	return string(bs)
}

// JSONStringToMap json string to map
func JSONStringToMap(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	jsonBytes := []byte(jsonStr)
	dec := json.NewDecoder(bytes.NewReader(jsonBytes))
	dec.UseNumber()
	if err := dec.Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// ToGBK to gbk
//func ToGBK(s string) (string, error) {
//	enc := simplifiedchinese.GBK
//	trans := enc.NewEncoder()
//	r, _, err := transform.String(trans, s)
//	return r, err
//}

// FromGBK from gbk
//func FromGBK(s string) (string, error) {
//	enc := simplifiedchinese.GBK
//	trans := enc.NewDecoder()
//	r, _, err := transform.String(trans, s)
//	return r, err
//}

var (
	chars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// RandString rand string
func RandString(l int) string {
	rand.Seed(time.Now().Unix())
	bs := []byte{}
	for i := 0; i < l; i++ {
		bs = append(bs, chars[rand.Intn(len(chars))])
	}
	return string(bs)
}

// IsNumber is number
func IsNumber(s string) bool {
	_, err := strconv.ParseUint(s, 10, 64)
	return err == nil
}

// ToPercent to percent
func ToPercent(f float64) string {
	if f < 0 {
		f = 0.0
	} else if f > 1.0 {
		f = 1.0
	}
	t := decimal.NewFromFloat(f)
	r := t.Mul(decimal.New(100, 0))
	return strconv.Itoa(int(r.IntPart()))
}

// BoolToByte bool to byte
func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

// ByteToBool byte to bool
func ByteToBool(b byte) bool {
	return b == 1
}

func StringToFloat64(str string) (value float64) {
	value, _ = strconv.ParseFloat(str, 64)
	return
}

// InArray 判断指定数组中包含指定项
func IsInArray(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// IsEmpty 判断是否为空
func IsEmpty(val interface{}) bool {
	if val == nil {
		return true
	}
	switch val.(type) {
	case string:
		return strings.TrimSpace(val.(string)) == ""
	case byte:
		return val == byte(0)
	case int:
		return val == int(0)
	case int32:
		return val == int32(0)
	case int64:
		return val == int64(0)
	case uint:
		return val == uint(0)
	case uint32:
		return val == uint32(0)
	case uint64:
		return val == uint64(0)
	case float32:
		return val == float32(0)
	case float64:
		return val == float64(0)

	default:
		return false
	}
}

func QueryEscape(s string) string {
	s = url.QueryEscape(s)
	s = strings.Replace(s, "+", "%20", -1)
	return s
}

func QueryUnescape(s string) (string, error) {
	var err error
	s = strings.Replace(s, "%20", "+", -1)
	s, err = url.QueryUnescape(s)
	return s, err
}

func ParseQueries(url string) map[string]string {
	queries := make(map[string]string)
	ss := strings.Split(url, "?")
	if len(ss) <= 1 {
		return queries
	}
	for _, s := range strings.Split(ss[1], "&") {
		p := strings.Split(s, "=")
		queries[p[0]] = p[1]
	}
	return queries
}

func StringToInt32(val string) (value int32) {
	v, _ := strconv.ParseInt(val, 10, 32)
	value = int32(v)
	return
}

func StringToInt64(val string) (value int64) {
	value, _ = strconv.ParseInt(val, 10, 64)
	return
}

// FormatNumberToShortString 数字格式化为短字符串，如 10000 ＝ 1万
func FormatNumberToShortString(number int64) string {
	// 万
	tenThousand := 10000
	// 亿
	oneHundredMillion := 100000000
	// 万亿
	tenThousandMillion := 1000000000000

	if number < int64(tenThousand) {
		return strconv.FormatInt(number, 10)
	} else {
		n := decimal.NewFromFloat(float64(number))
		var suffix string
		var m decimal.Decimal
		if number < int64(oneHundredMillion) {
			m = decimal.NewFromFloat(float64(tenThousand))
			suffix = "万"
		} else if number < int64(tenThousandMillion) {
			m = decimal.NewFromFloat(float64(oneHundredMillion))
			suffix = "亿"
		} else {
			m = decimal.NewFromFloat(float64(tenThousandMillion))
			suffix = "万亿"
		}
		return n.DivRound(m, 2).String() + suffix
	}
}

// SleepRandomTime 随机 sleep 时间
func SleepRandomTime(maxSleepTime time.Duration) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(r.Int63n(int64(maxSleepTime))))
}

// GenRandomString 生成随机字符串
func GenRandomString(n int, randomRange []string) string {
	var s string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rangeSize := len(randomRange)
	for i := 0; i < n; i++ {
		s += randomRange[r.Intn(rangeSize)]
	}
	return s
}
