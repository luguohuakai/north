# SrunNorthbound

深澜北向接口 接口文档请参考北向接口文档
> 引入方式: `go get github.com/luguohuakai/SrunNorthbound`

* 首次使用请配置配置文件 `/etc/northbound.conf` 没有的话需要先创建, 样例如下

```ini
protocol = "https"
interface_ip = "127.0.0.1"
port = 8001
```

> 一般情况下只需要修改 `interface_ip`

### 使用方式 仅作参考

```go
if httpRequest, err := srun.OnlineIndexEduroam_(queryMap); err != nil {
return result.ReturnNoData(http.StatusBadRequest, i18n.I18NLoad("Failure", "return"))
} else if httpRequest.Code != 0 {
return result.ReturnNoData(http.StatusBadRequest, httpRequest.Message)
}
```
