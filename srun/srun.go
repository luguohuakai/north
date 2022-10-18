// Package srun @program: 深澜软件-北向接口GO版 SDK-0.0.1
//@author: DM
//@create: 2021-03-26 16:00
package srun

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

// HttpResult @title HttpResult
// @description 用于接收北向接口返回值
// @param Data Data为interface类型是因为其返回值可能为map或其他类型的值
// @author DM
// @time 2021/4/2 21:07
type HttpResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Version string      `json:"version"`
	Data    interface{} `json:"data"`
	Meta    Meta        `json:"_meta"`
}

type Meta struct {
	TotalCount int `json:"totalCount"`
}

// @title TokenUrl
// @description 北向接口列表
// @author DM
// @time 2021/4/2 21:06
const (
	TokenUrl = "/api/v1/auth/get-access-token" // 获取token
	/**
	 * 用户管理
	 */
	LoginUrl             = "/api/v1/user/validate-manager"
	UserCreateUrl        = "/api/v1/users"
	UserUpdateUrl        = "/api/v1/user/update"
	UserDeleteUrl        = "/api/v1/user/delete"
	UserRestPassUrl      = "/api/v1/user/super-reset-password"
	UserStatusControl    = "/api/v1/user/user-status-control"
	ResetPasswordManager = "/api/v1/user/reset-password-manager"
	/**
	 * 产品相关接口
	 */
	ControlCreateUrl       = "/api/v1/strategy/control-create"
	BillingCreateUrl       = "/api/v1/strategy/billing-create"
	ProductCreateUrl       = "/api/v1/product/create"
	ProductDeleteUrl       = "/api/v1/product/delete"
	ProductModifyUrl       = "/api/v1/product/update"
	ProductCanSubscribe    = "/api/v1/group/subscribe"
	ProductTransferUrl     = "/api/v1/product/transfer-product"              //产品转移（立即生效）
	UsersPackages          = "/api/v1/package/users-packages"                //查询已订购的产品套餐
	EnableProduct          = "/api/v1/product/enable-product"                //启用产品接口
	DisableProduct         = "/api/v1/product/disable-product"               //禁用产品接口
	ProductSubscribeUrl    = "/api/v1/product/subscribe"                     //订购产品
	ReservationTransferUrl = "/api/v1/product/reservation-transfer-products" //预约转移
	/**
	 * 财务相关
	 */
	FinancialAddUrl    = "/api/v1/financial/create-payment"
	FinancialEditUrl   = "/api/v1/financial/update-payment"
	FinancialDeleteUrl = "/api/v1/financial/delete-payment"
	RechargeWallet     = "/api/v1/financial/recharge-wallet" //电子钱包充值
	/**
	 *学校域名对应关系
	 */
	SearchEduroamSchool = "/api/eduroam/domain/index"  //查询学校域名关系接口
	AddEduroamSchool    = "/api/eduroam/domain/create" //添加学校域名关系接口
	UpdateEduroamSchool = "/api/eduroam/domain/update" //编辑学校域名关系接口
	DelEduroamSchool    = "/api/eduroam/domain/delete" //删除学校域名关系接口
	/**
	 *访客接口
	 */
	SearchEduroamVisitor     = "/api/eduroam/visitor/index"       //访客组成接口
	ListEduroamVisitor       = "/api/eduroam/visitor/list"        //访客列表接口
	TopEduroamVisitor        = "/api/eduroam/visitor/use-top"     //使用量TOP接口
	DomainTopEduroamVisitor  = "/api/eduroam/visitor/domain-top"  //学校流量统计
	DayDetailEduroamVisitor  = "/api/eduroam/visitor/day-detail"  //每日访客流量/时长接口
	InVisitorEduroamVisitor  = "/api/eduroam/visitor/in-visitor"  //每日来访查询接口
	OutVisitorEduroamVisitor = "/api/eduroam/visitor/out-visitor" //出访查询接口
	VisitorEduroamDetail     = "/api/v1/user/view"                //用户详情查询
	/**
	 * 在线相关
	 */
	OnlineIndexEduroam          = "/api/eduroam/online/index"        //查询在线表接口
	DayOnlineIndexEduroamCount  = "/api/eduroam/online/online-count" //每日在线分析/统计接口
	DropOnlineIndexEduroam      = "/api/v1/base/online-drop"         //在线设备下线接口
	EquipmentOnlineIndexEduroam = "/api/v1/base/online-equipment"    //查询在线设备接口
	/**
	 * 用户组相关
	 */
	GroupCreateUrl = "/api/v1/groups"              //添加用户组接口
	MaxOnlineNum   = "/api/v1/user/max-online-num" // 修改最大在线数接口
	/**
	 * idm相关
	 */
	IdmDeviceIndex   = "/api/idm/device/index"    //查询设备列表接口
	IdmDeviceFactory = "/api/idm/device/factory"  //查询厂商接口
	IdmDeviceOs      = "/api/idm/device/os"       //查询设备操作系统接口
	IdmDeviceCreate  = "/api/idm/device/create"   //添加设备接口
	IdmProductIndex  = "/api/idm/products/index"  //查询产品列表
	IdmProductUpdate = "/api/idm/products/update" //修改产品接口，异步操作可以修改产品所绑定的计费控制两个策略
	/**
	 * 错误信息接口
	 */
	AuthErrMsg = "/api/eduroam/settings/auth-err-msg" //上网认证错误消息
	DropReason = "/api/eduroam/settings/drop-reason"  //下线原因
	SysMsg     = "/api/eduroam/settings/sys-msg"      //系统错误信息
	/**
	 * Redis键名
	 */
	HashUsersSecret = "hash:users:secret:"
	HashUsersLogin  = "helper:login:"
	HashUsersInfo   = "hash:users:"
	HashBillingPre  = "hash:billing:"        //计费策略策略hash的key前缀
	HashProductPre  = "hash:products:"       //产品hash的key前缀
	HashControlPre  = "hash:control:"        //控制策略hash的key前缀
	HashHelperBx    = "hash:server:api:info" //北向接口服務器地址
	HashPackagePre  = "hash:package:"        //套餐的key前缀

	ListControlPre        = "list:control"   //控制策略list的key前缀
	ListBillingPre        = "list:billing"   //计费策略list的key前缀
	ListProductPre        = "list:products:" //产品list的key前缀
	ListProducts          = "list:products"
	ListPackagesPre       = "list:package" //套餐list的key前缀
	ListExcelPortSelected = "list:excel:export:"
	ListProductsControl   = "list:products:control:" //产品对应的控制策略list的key前缀

	ListLogOnlineDetail = "key:advanced:online:detail:selected"  //上网明细日志字段高级设置选项
	ListLogAuth         = "key:advanced:auth:selected"           //认证日志字段高级设置
	ListLogSystem       = "key:advanced:system:selected"         //系统日志字段高级设置
	ListOnlineRadius    = "key:advanced:online:radius:selected"  //在线用户字段高级设置
	ListVisitor         = "key:advanced:visitor-list:selected"   //访客字段高级设置
	ListIdmUsersList    = "key:advanced:idm-users-list:selected" //IDM系统用户表[设备表]字段高级设置

	ListUsersProducts = "list:users:products:" //用户绑定的产品的list前缀

	IncrRadAttrId        = "rad_attr_id"     //radius属性自增id
	HashRadAttr          = "hash:rad_attr:"  //radius属性hash的key前缀
	ListRadAttr          = "list:rad_attr"   //radius属性list
	KeyRadAttrId         = "key:rad_attr_id" //radius属性的key
	StringSystemStyleKey = "key:system:style"

	IdmListDhcpPool6        = "list:dhcp:pool6"           //ipv6地址池，ID列表
	IdmListDhcpPool         = "list:dhcp:pool"            //ipv4地址池，ID列表
	IdmListDhcpPool6UsedIps = "list:dhcp:pool6:used_ips:" //已使用ipv6地址池ip列表
	IdmListDhcpPoolUsedIps  = "list:dhcp:pool:used_ips:"  //已使用ipv4地址池ip列表
	IdmHashDhcpInfoIp       = "hash:dhcp:info:ip:"        //已使用ipv4详细信息
	IdmHashDhcpInfoIp6      = "hash:dhcp:info:ip6:"       //已使用ipv6详细信息
	IdmHashDhcpInfo         = "hash:dhcp:pool:"           //ipv4地址池信息
	IdmHashDhcpInfo6        = "hash:dhcp:pool6:"          //ipv6地址池信息

	ListInterface = "list:interface" //interface核心接口list

	ListNasType = "list:nas_type"
	HashNasType = "hash:nas_type:"
	/**
	 * 数据库扩展字段表名
	 */
	ExtendTableName   = "extends_field"
	TmpUsers          = "tmp_users"
	PayList           = "pay_list"
	PayType           = "pay_type"
	UserGroup         = "users_group"
	Users             = "users"
	OnlineRadiusTable = "online_radius" //在线表
)

var ssoError = map[string]string{
	"client_ip_error":                     "客户端IP错误",
	"Nas type not found.":                 "认证设备(Nas)未找到",
	"user_tab_error":                      "认证程序未启动",
	"username_error":                      "用户名输入错误",
	"logout_error":                        "注销时发生错误,或没有帐号在线！",
	"uid_error":                           "您的账号不在线上.",
	"mac_error":                           "您的MAC地址不正确",
	"password_error":                      "用户名或密码错误,请重新输入！",
	"status_error":                        "您已欠费，请尽快充值。",
	"sync_error":                          "您的资料已被修改正在等待同步，请2钟分后再试。\n如果您的帐号允许多个用户上线，请到WEB登录页面注销。",
	"delete_error":                        "您的帐号已经被删除",
	"ip_exist_error":                      "IP已存在，请稍后再试。",
	"usernum_error":                       "在线用户已满，请稍后再试。",
	"online_num_error":                    "正在注销在线账号，请重新连接",
	"proxy_error":                         "你的IP地址和认证地址不附，可能是经过小路由器登录的。",
	"mode_error":                          "系统已禁止客户端登录，请使用WEB方式登录。",
	"flux_error":                          "您的流量已用尽。",
	"minutes_error":                       "您的时长已用尽。",
	"ip_error":                            "您的IP地址不合法，可能是：\n一、与绑的IP地址附；二、IP不允许在当前区域登录；",
	"time_policy_error":                   "当前时段不允许连接。",
	"available_error":                     "抱歉，您的帐号已禁用",
	"Addr table error~login_error":        "计费系统尚未授权，目前还不能使用！",
	"ipv6_error":                          "您的IPv6地址不正确，请重新配置IPv6地址!",
	"E2611":                               "您当前使用的设备非该账号绑定设备 请绑定或使用绑定的设备登入",
	"E2602":                               "您还没有绑定手机号或绑定的非联通手机号码",
	"E2601":                               "您使用的不是专用客户端,IPOE-PPPoE混杂模式请联系管理员重新打包客户端程序",
	"E2532":                               "您的两次认证的间隔太短,请稍候10秒后再重试登录",
	"E2533":                               "密码错误次数超过限制，请5分钟后再重试登录",
	"E2531":                               "帐号不存在或密码错误",
	"E2553":                               "帐号或密码错误",
	"E2606":                               "用户被禁用",
	"E2613":                               "NAS PORT绑定错误",
	"E2614":                               "MAC地址绑定错误",
	"E2615":                               "IP地址绑定错误",
	"E2616":                               "用户已欠费",
	"E2621":                               "已经达到授权人数",
	"E2806":                               "找不到符合条件的产品",
	"E2807":                               "后台系统配置错误,找不到符合条件的计费策略,请联系管理员检查后台计费策略配置",
	"E2808":                               "后台系统配置错误,找不到符合条件的控制策略,请联系管理员检查后台控制策略配置",
	"E2833":                               "IP不在DHCP表中，需要重新拿地址。",
	"E2840":                               "校内地址不允许访问外网。",
	"E2843":                               "IP地址不正确!",
	"auth_resault_timeout_err":            "认证服务无响应",
	"auth_result_timeout_err":             "认证服务无响应",
	"The server is not responding.":       "后台服务器无响应,请联系管理员检查后台服务运行状态",
	"You have been forcibly disconnected": "您已经被服务器强制下线！",
	"sign_error":                          "签名错误,建议检查secret或参数顺序",
	"login_error":                         "认证失败,请联系管理员",
	"INFO failed, BAS respond timeout.":   "BAS响应超时",
	"E2533: The number you are trying to have reached 3 times, please try again after 5 minutes.": "密码错误次数超过限制，请5分钟后再重试登录",
}

var ssoSuccess = map[string]string{
	"ip_already_online_error":            "当前IP已在线,可以直接上网",
	"E0000":                              "认证成功",
	"You are not online.":                "注销成功",
	"IP has been online, please logout.": "您的IP已经在线,可以直接上网,或者先注销再重新认证",
	"E2842":                              "您的IP地址无需认证即可上网",
	"E2620":                              "已经在线了",
	"non_auth_error":                     "您无须认证，可直接上网",
	"logout_ok":                          "注销成功，请等1分钟后登录。",
	"ok":                                 "操作成功",
	"login_ok":                           "登录成功",
	"注销成功":                               "在线账号注销成功.",
	"不在线上":                               "在线账号注销成功.",
}

type HttpResultSso struct {
	ServerFlag             int    `json:"ServerFlag"`
	ServicesIntfServerIP   string `json:"ServicesIntfServerIP"`
	ServicesIntfServerPort string `json:"ServicesIntfServerPort"`
	CheckoutDate           int    `json:"checkout_date,omitempty"`
	ClientIp               string `json:"client_ip,omitempty"`
	Ecode                  int    `json:"ecode,omitempty"`
	Error                  string `json:"error,omitempty"`
	ErrorMsg               string `json:"error_msg,omitempty"`
	OnlineIp               string `json:"online_ip,omitempty"`
	RealName               string `json:"real_name,omitempty"`
	RemainFlux             int    `json:"remain_flux,omitempty"`
	RemainTimes            int    `json:"remain_times,omitempty"`
	Res                    string `json:"res,omitempty"`
	SrunVer                string `json:"srun_ver,omitempty"`
	SucMsg                 string `json:"suc_msg,omitempty"`
	Sysver                 string `json:"sysver,omitempty"`
	Username               string `json:"username,omitempty"`
	WalletBalance          int    `json:"wallet_balance,omitempty"`
	PloyMsg                string `json:"ploy_msg,omitempty"`
}

type SsoResult struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"is_success"`
}

func GetSsoSuccessOrError(sso HttpResultSso) (res SsoResult) {
	var code string

	//if sso.Ecode != 0 {
	//	code = strconv.Itoa(sso.Ecode)
	//	goto next
	//}
	if sso.PloyMsg != "" {
		arr := strings.Split(sso.PloyMsg, ":")
		if len(arr) != 2 {
			code = sso.PloyMsg
			goto next
		} else {
			code = arr[0]
			goto next
		}
	}
	if sso.SucMsg != "" {
		code = sso.SucMsg
		goto next
	}
	if sso.ErrorMsg != "" {
		arr := strings.Split(sso.ErrorMsg, ":")
		if len(arr) != 2 {
			code = sso.ErrorMsg
			goto next
		} else {
			code = arr[0]
			goto next
		}
	}
	if sso.Res != "" {
		code = sso.Res
		goto next
	}
	if sso.Error != "" {
		code = sso.Error
		goto next
	}

next:
	res.Code = code
	if ssoSuccess[code] != "" {
		res.Message = ssoSuccess[code]
		res.IsSuccess = true
	} else if ssoError[code] != "" {
		res.Message = ssoError[code]
		res.IsSuccess = false
	} else {
		res.Message = "unknown code: " + code
	}

	return
}

// Url @title Url
// @description 拼接url
// @author DM
// @time 2021/4/2 20:59
// @param api
// @return string
func Url(api string) string {
	confPath := "/etc/northbound.ini"
	var conf Northbound
	if dataMap, err := ReadFile(confPath); err != nil {
		conf.Protocol = "https"
		conf.InterfaceIp = "127.0.0.1"
		conf.Port = 8001
	} else {
		if val, ok := dataMap["protocol"]; ok {
			conf.Protocol = val.(string)
		} else {
			conf.Protocol = "https"
		}
		if val, ok := dataMap["interface_ip"]; ok {
			conf.InterfaceIp = val.(string)
		} else {
			conf.InterfaceIp = "127.0.0.1"
		}
		if val, ok := dataMap["port"]; ok {
			conf.Port = val.(int)
		} else {
			conf.Port = 8001
		}
	}
	return fmt.Sprintf("%s://%s:%d%s", conf.Protocol, conf.InterfaceIp, conf.Port, api)
}

// MyFormatter @title MyFormatter
// @description 自定义日志格式
// @author DM
// @time 2021/4/20 8:37
type MyFormatter struct{}

// Format @title MyFormatter
// @description 自定义日志格式
// @author DM
// @time 2021/4/20 8:37
// @receiver s
// @param entry
// @return []byte
// @return error
func (s *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("%s [%s] %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

// @title logError
// @description 格式化北向接口错误日志
// @author DM
// @time 2021/4/2 20:59
// @param method string 接口请求方法
// @param url string 接口地址
// @param responseOrErr string 接口返回数据或error信息
// @param params string 请求参数
func logError(method string, url string, responseOrErr string, params string) {
	//log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	log.SetFormatter(new(MyFormatter))
	//fmt.Println(fmt.Sprintf("%s %s %s", strings.ToUpper(method), url, params))
	//fmt.Println(fmt.Sprintf("%s", responseOrErr))
	log.Error(fmt.Sprintf("%s Request=%s", strings.ToUpper(method), url))
	if strings.ToUpper(method) != "GET" {
		log.Error(fmt.Sprintf("%s Params=%s", strings.ToUpper(method), params))
	}
	log.Error(fmt.Sprintf("%s Response=%s", strings.ToUpper(method), responseOrErr))
}

func LogError(method string, url string, responseOrInfo string, params string) {
	logError(method, url, responseOrInfo, params)
}

// @title logInfo
// @description 格式化北向接口一般日志
// @author DM
// @time 2021/4/2 20:58
// @param method string 接口请求方法
// @param url string 接口地址
// @param responseOrInfo string 接口返回数据或info信息
// @param params string 请求参数
func logInfo(method string, url string, responseOrInfo string, params string) {
	//log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	log.SetFormatter(new(MyFormatter))
	//fmt.Println(fmt.Sprintf("%s %s %s", strings.ToUpper(method), url, params))
	//fmt.Println(fmt.Sprintf("%s", responseOrInfo))
	log.Info(fmt.Sprintf("%s Request=%s", strings.ToUpper(method), url))
	if strings.ToUpper(method) != "GET" {
		log.Info(fmt.Sprintf("%s Params=%s", strings.ToUpper(method), params))
	}
	log.Info(fmt.Sprintf("%s Response=%s", strings.ToUpper(method), responseOrInfo))
}

func LogInfo(method string, url string, responseOrInfo string, params string) {
	logInfo(method, url, responseOrInfo, params)
}

// AccessToken @title AccessToken
// @description 获取北向接口access_token
// @author DM
// @time 2021/4/2 21:00
// @return string
func AccessToken() string {
	var (
		rs         []byte
		err        error
		httpResult HttpResult
		Token      string
	)
	// todo: 取缓存Token 压力不大的话可以每次都从接口获取
	if Token != "" && len(Token) > 0 {
		return Token
	}
	// 获取access_token
	tokenUrl := Url(TokenUrl)
	if rs, err = Get(tokenUrl); err != nil {
		logError("get", tokenUrl, err.Error(), "")
		return ""
	}

	// 记录返回值，写入到日志中
	logInfo("get", tokenUrl, fmt.Sprintf("%s", rs), "")

	if err = json.Unmarshal(rs, &httpResult); err != nil {
		logError("get", tokenUrl, fmt.Sprintf("access_token json unmarshal error: %s", err.Error()), "")
		return ""
	}
	if httpResult.Code == 401 {
		logError("get", tokenUrl, "the api get access_token returned code 401", "")
		return ""
	}
	if v, ok := (httpResult.Data).(map[string]interface{}); ok {
		// todo: 缓存2小时 redis缓存/文件缓存都可以
		return v["access_token"].(string)
	}
	logError("get", tokenUrl, "the api access_token is empty", "")
	return ""
}

// Request @title Request
// @description api string, method string, params map[string]string, noAccessToken bool
// @author DM
// @time 2021/4/2 21:00
// @param api 接口
// @params MethodParamsNoAccessToken
// @param method 请求类型 支持 get post put ... 默认get
// @param params 接口请求参数 map[string]string 默认不传 所有接口都支持此参数包括get
// @param noAccessToken false:默认,自动上传北向接口access_token true:不传access_token
// @return httpResult
// @return err
func Request(api string, MethodParamsNoAccessToken ...interface{}) (httpResult *HttpResult, err error) {
	var rs []byte
	var method string
	var params = make(map[string]string, 0)
	var noAccessToken bool
	var ok bool
	reqUrl := Url(api)
	for key, value := range MethodParamsNoAccessToken {
		switch key {
		case 0:
			if method, ok = value.(string); ok != true {
				errMsg := "request param [method] error,eg: get post put delete"
				logError(method, reqUrl, errMsg, mapToJson(params))
				return nil, errors.New(errMsg)
			}
		case 1:
			if params, ok = value.(map[string]string); ok != true {
				errMsg := "request param [params] error,eg: map[string]string"
				logError(method, reqUrl, errMsg, mapToJson(params))
				return nil, errors.New(errMsg)
			}
		case 2:
			if noAccessToken, ok = value.(bool); ok != true {
				errMsg := "request param [noAccessToken] error,eg: true false"
				logError(method, reqUrl, errMsg, mapToJson(params))
				return nil, errors.New(errMsg)
			}
		}
	}
	if method == "" {
		method = "get"
	}
	if noAccessToken == false {
		var accessToken = AccessToken()
		if accessToken == "" {
			errMsg := fmt.Sprintf("access_token get failed: %s", api)
			logError(method, reqUrl, errMsg, mapToJson(params))
			return nil, errors.New(errMsg)
		}
		params["access_token"] = accessToken
	}
	switch strings.ToLower(method) {
	case "get":
		// 参数拼接
		if params != nil {
			uv := url.Values{}
			for k, v := range params {
				uv.Add(k, v)
			}
			p := uv.Encode()
			if strings.Index(reqUrl, "?") != -1 && strings.Index(reqUrl, "=") != -1 {
				reqUrl = reqUrl + "&" + p
			} else {
				reqUrl = reqUrl + "?" + p
			}
		}
		rs, err = Get(reqUrl)
	case "post":
		rs, err = Post(reqUrl, params)
	case "put":
		rs, err = Put(reqUrl, params)
	default:
		rs, err = DeleteOrUpdate(reqUrl, params, method)
	}
	// 无论请求成功或失败都记录请求日志
	logInfo(method, reqUrl, fmt.Sprintf("%s", rs), mapToJson(params))
	if err != nil {
		errMsg := fmt.Sprintf("api request error: %s", err.Error())
		logError(method, reqUrl, errMsg, mapToJson(params))
		return nil, errors.New(errMsg)
	}
	// 解码json串到httpResult结构体
	if e := json.Unmarshal(rs, &httpResult); e != nil {
		errMsg := fmt.Sprintf("json unmarshal error: %s", e.Error())
		logError(method, reqUrl, errMsg, mapToJson(params))
		return nil, errors.New(errMsg)
	}
	if httpResult.Code != 0 {
		return nil, errors.New(httpResult.Message)
	}
	if httpResult.Data == nil {
		httpResult.Data = make([]string, 0)
	}

	return
}

// RequestSso 8082上的微信临时放行key(必须核实)也可在服务器文件srun4kauth.xml中ApiAuthSecret字段获得需修改EnableAPIAuth=1然后重启srun3kauth
func RequestSso(reqUrl string, params map[string]string) (httpResult *HttpResultSso, err error) {
	var rs []byte

	rs, err = PostSso(reqUrl, params)

	// 无论请求成功或失败都记录请求日志
	logInfo("post", reqUrl, fmt.Sprintf("%s", rs), mapToJson(params))
	if err != nil {
		errMsg := fmt.Sprintf("sso api request error: %s, rs: %s", err.Error(), rs)
		logError("post", reqUrl, errMsg, mapToJson(params))
		return nil, errors.New(errMsg)
	}
	// 解码json串到httpResult结构体
	if e := json.Unmarshal(rs, &httpResult); e != nil {
		errMsg := fmt.Sprintf("sso json unmarshal error: %s, rs: %s", e.Error(), rs)
		logError("post", reqUrl, errMsg, mapToJson(params))
		return nil, errors.New(errMsg)
	}

	return
}

// Get @title Get
// @description 封装https 的get方法
// @author DM
// @time 2021/4/2 21:02
// @param requestUrl string  请求的url
// @return body []byte 返回的结果
// @return err
func Get(requestUrl string) (body []byte, err error) {
	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	var response *http.Response
	// 重连2次
	for i := 0; i < 2; i++ {
		response, err = c.Get(requestUrl)
		if err != nil {
			times := i + 1
			logError("get", requestUrl, fmt.Sprintf("connect times: (%d) net connect has error: %s", times, err.Error()), "")
		} else {
			break
		}
	}
	if response != nil {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, errors.New("api request has error: " + err.Error())
		}
		return body, nil
	}
	if err != nil {
		return nil, errors.New("api request has error: " + err.Error())
	}
	return nil, errors.New("api request has error: err = nil")
}

// Post @title Post
// @description 封装https的post方法
// @author DM
// @time 2021/4/2 21:03
// @param requestUrl string  请求的url
// @param params map[string]string 请求的参数
// @return body []byte 返回的结果
// @return err
func Post(requestUrl string, params map[string]string) (body []byte, err error) {
	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	reqContentType := "application/x-www-form-urlencoded"
	uv := make(url.Values)
	for k, v := range params {
		uv.Add(k, v)
	}
	encode := uv.Encode()
	var response *http.Response
	// 重连2次
	for i := 0; i < 2; i++ {
		response, err = c.Post(requestUrl, reqContentType, strings.NewReader(encode))
		if err != nil {
			times := i + 1
			logError("get", requestUrl, fmt.Sprintf("connect times: (%d) net connect has error: %s", times, err.Error()), mapToJson(params))
		} else {
			break
		}
	}
	if response != nil {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, errors.New("api request has error: " + err.Error())
		}
		return body, nil
	}
	if err != nil {
		return nil, errors.New("api request has error: " + err.Error())
	} else {
		return nil, errors.New("api request has error: err = nil")
	}
}
func PostSso(requestUrl string, params map[string]string) (body []byte, err error) {
	uv := make(url.Values)
	for k, v := range params {
		uv.Add(k, v)
	}
	encode := uv.Encode()
	var response *http.Response
	// 重连2次
	for i := 0; i < 2; i++ {

		response, err = PostSso_(requestUrl, strings.NewReader(encode))
		if err != nil {
			times := i + 1
			logError("get", requestUrl, fmt.Sprintf("connect times: (%d) net connect has error: %s", times, err.Error()), mapToJson(params))
		} else {
			break
		}
	}
	if response != nil {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, errors.New("api request has error: " + err.Error())
		}
		return body, nil
	}
	if err != nil {
		return nil, errors.New("api request has error: " + err.Error())
	} else {
		return nil, errors.New("api request has error: err = nil")
	}
}

func PostSso_(url string, body io.Reader) (resp *http.Response, err error) {
	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	return c.Do(req)
}

// DeleteOrUpdate @title DeleteOrUpdate
// @description 封装https请求方法
// @author DM
// @time 2021/4/2 21:04
// @param requestUrl string  请求的url
// @param params
// @param method
// @return []byte 返回的结果
// @return error
func DeleteOrUpdate(requestUrl string, params map[string]string, method string) ([]byte, error) {
	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	uv := make(url.Values)
	for k, v := range params {
		uv.Add(k, v)
	}

	var response *http.Response
	req, err := http.NewRequest(method, requestUrl, strings.NewReader(uv.Encode()))
	// 重连2次
	for i := 0; i < 2; i++ {
		response, err = c.Do(req)
		if err != nil {
			times := i + 1
			logError("get", requestUrl, fmt.Sprintf("connect times: (%d) net connect has error: %s", times, err.Error()), mapToJson(params))
		} else {
			break
		}
	}
	if response != nil {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, errors.New("api request has error: " + err.Error())
		}
		return body, nil
	}
	if err != nil {
		return nil, errors.New("api request has error: " + err.Error())
	} else {
		return nil, errors.New("api request has error: err = nil")
	}
}

// Put @title Put
// @description 封装https的put方法
// @author DM
// @time 2021/4/2 21:05
// @param requestUrl string  请求的url
// @param params map[string]string 请求的参数
// @return []byte 返回的结果
// @return error
func Put(requestUrl string, params map[string]string) ([]byte, error) {
	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	uv := make(url.Values)
	for k, v := range params {
		uv.Add(k, v)
	}
	var response *http.Response
	req, err := http.NewRequest(http.MethodPut, requestUrl, strings.NewReader(uv.Encode()))
	// 重连2次
	for i := 0; i < 2; i++ {
		response, err = c.Do(req)
		if err != nil {
			times := i + 1
			logError("get", requestUrl, fmt.Sprintf("connect times: (%d) net connect has error: %s", times, err.Error()), mapToJson(params))
		} else {
			break
		}
	}
	if response != nil {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, errors.New("api request has error: " + err.Error())
		}
		return body, nil
	}
	if err != nil {
		return nil, errors.New("api request has error: " + err.Error())
	} else {
		return nil, errors.New("api request has error: err = nil")
	}
}

func mapToJson(param map[string]string) string {
	dataType, _ := json.Marshal(param)
	return string(dataType)
}
