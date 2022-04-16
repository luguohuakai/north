package srun

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

// 操作文件相关
var (
	bt = sync.Pool{
		New: func() interface{} {
			b := make([]byte, 1024)
			return &b
		},
	}
)

// ReadFile 读取配置文件
func ReadFile(fileName string) (m map[string]interface{}, err error) {
	var (
		b *[]byte
		f *os.File
		r *bufio.Reader
	)
	if _, err = os.Stat(fileName); err != nil && os.IsNotExist(err) {
		err = nil
		return
	}
	m = make(map[string]interface{})
	f, err = os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		log.Error(err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Error(err)
		}
	}(f)
	r = bufio.NewReader(f)
	for {
		b = bt.Get().(*[]byte)
		*b, err = r.ReadBytes('\n')
		if r := doStringToArray(string(*b)); len(r) > 0 {
			m[r[0]] = r[1]
		}
		*b = (*b)[:0]
		bt.Put(b)
		if err != nil {
			if err == io.EOF {
				return m, nil
			}
			log.Error(err)
		}

	}
}

// string reserve to Array struct
func doStringToArray(r string) (s []string) {
	//先判断字符串是否以#开头，注释
	if b := strings.HasPrefix(r, "#"); b {
		return nil
	}
	// 分号也是注释，用于企业微信配置文件
	if b := strings.HasPrefix(r, ";"); b {
		return nil
	}
	result := strings.Split(r, "=")
	if len(result) < 2 {
		return nil
	}
	ss := make([]string, 2)
	re3, _ := regexp.Compile("\"|(\r\n)|\n")

	// 避免出现值中出现=的情况
	var key, value string
	for k, v := range result {
		if k == 0 {
			key = re3.ReplaceAllString(v, "")
			key = strings.TrimSpace(key)
		} else if value == "" {
			value = re3.ReplaceAllString(v, "")
			value = strings.TrimSpace(value)
		} else {
			valueNew := re3.ReplaceAllString(v, "")
			valueNew = strings.TrimSpace(valueNew)
			value = value + "=" + valueNew
		}
	}

	ss[0] = key
	ss[1] = value
	return ss
}
