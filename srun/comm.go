package srun

// UserRight 查询用户是否正确
func UserRight(params map[string]string) (flag bool) {
	httpResult, err := Request("/api/v1/user/validate-users", "post", params)
	if err != nil {
		return
	} else {
		if httpResult.Code == 0 {
			flag = true
		}
	}
	return
}
