package srun

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"time"
)

// UserRight 查询用户是否正确
// @params map[string]string{"user_name":"xxx", "password":"xxx"}
func UserRight(params map[string]string) (err error) {
	var httpResult *HttpResult
	if httpResult, err = Request("/api/v1/user/validate-users", "post", params); err != nil {
		return
	} else {
		if httpResult.Code != 0 {
			return errors.New(httpResult.Message)
		}
	}
	return
}

// UserExists 查询用户是否存在
func UserExists(username string) (err error) {
	var httpResult *HttpResult
	if httpResult, err = Request("/api/v1/user/search", "post", map[string]string{"value": username, "type": "1"}); err != nil {
		return
	} else {
		if httpResult.Code != 0 {
			return errors.New(httpResult.Message)
		}
	}
	return
}

// Sso 调用单点登录接口
// 8082上的微信临时放行key(必须核实)也可在服务器文件srun4kauth.xml中ApiAuthSecret字段获得需修改EnableAPIAuth=1然后重启srun3kauth
// @params action login:登录 logout:登出
func Sso(ssoSecret, ssoUrl, username, ip, acId, action string) (*HttpResultSso, error) {
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	data := map[string]string{
		"action":   action,
		"api_auth": "1",
		"username": username,
		"ip":       ip,
		"type":     "1001",
		"n":        "100",
		"drop":     "0",
		"pop":      "0",
		"time":     timeStamp,
		"password": MD5(strings.Join([]string{ssoSecret, timeStamp, username, ip, timeStamp, ssoSecret}, "")),
		"ac_id":    acId,
	}

	if action == "logout" {
		delete(data, "drop")
		delete(data, "pop")
	}

	return RequestSso(ssoUrl+"/cgi-bin/srun_portal", data)
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
