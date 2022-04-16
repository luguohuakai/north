// Package srun @program: 深澜软件-北向接口GO版 SDK-0.0.1
//@author: DM
//@create: 2021-03-26 16:01
package srun

import "net/http"

// LoginUrl_ @title LoginUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:09
// @param params
// @return httpResult
// @return err
func LoginUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(LoginUrl, "post", params)
	return
}

// UserCreateUrl_ @title UserCreateUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:09
// @param params
// @return httpResult
// @return err
func UserCreateUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(UserCreateUrl, "post", params)
	return
}

// UserUpdateUrl_ @title UserUpdateUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:10
// @param params
// @return httpResult
// @return err
func UserUpdateUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(UserUpdateUrl, "post", params)
	return
}

// UserDeleteUrl_ @title UserDeleteUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:10
// @param params
// @return httpResult
// @return err
func UserDeleteUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(UserDeleteUrl, http.MethodDelete, params)
	return
}

// UserRestPassUrl_ @title UserRestPassUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func UserRestPassUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(UserRestPassUrl, "post", params)
	return
}

// UserStatusControl_ @title UserStatusControl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func UserStatusControl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(UserStatusControl, "post", params)
	return
}

// ResetPasswordManager_ @title ResetPasswordManager_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func ResetPasswordManager_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(ResetPasswordManager, "post", params)
	return
}

// ControlCreateUrl_ @title ControlCreateUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func ControlCreateUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(ControlCreateUrl, "post", params)
	return
}

// BillingCreateUrl_ @title BillingCreateUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func BillingCreateUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(BillingCreateUrl, "post", params)
	return
}

// ProductCreateUrl_ @title ProductCreateUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func ProductCreateUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(ProductCreateUrl, "post", params)
	return
}

// ProductDeleteUrl_ @title ProductDeleteUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func ProductDeleteUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(ProductDeleteUrl, "DELETE", params)
	return
}

// ProductModifyUrl_ @title ProductModifyUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func ProductModifyUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(ProductModifyUrl, "put", params)
	return
}

// ProductCanSubscribe_ @title ProductCanSubscribe_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func ProductCanSubscribe_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(ProductCanSubscribe, "post", params)
	return
}

// ProductTransferUrl_ @title ProductTransferUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func ProductTransferUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(ProductTransferUrl, "post", params)
	return
}

// ReservationTransferUrl_ @title ReservationTransferUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func ReservationTransferUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(ReservationTransferUrl, "post", params)
	return
}

// UsersPackages_ @title UsersPackages_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func UsersPackages_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(UsersPackages, "get", params)
	return
}

// EnableProduct_ @title EnableProduct_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func EnableProduct_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(EnableProduct, "post", params)
	return
}

// DisableProduct_ @title DisableProduct_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func DisableProduct_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(DisableProduct, "post", params)
	return
}

// ProductSubscribeUrl_ @title ProductSubscribeUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func ProductSubscribeUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(ProductSubscribeUrl, "post", params)
	return
}

// FinancialAddUrl_ @title FinancialAddUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func FinancialAddUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(FinancialAddUrl, "post", params)
	return
}

// FinancialEditUrl_ @title FinancialEditUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func FinancialEditUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(FinancialEditUrl, "post", params)
	return
}

// FinancialDeleteUrl_ @title FinancialDeleteUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func FinancialDeleteUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(FinancialDeleteUrl, "post", params)
	return
}

// RechargeWallet_ @title RechargeWallet_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func RechargeWallet_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(RechargeWallet, "post", params)
	return
}

// SearchEduroamSchool_ @title SearchEduroamSchool_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func SearchEduroamSchool_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(SearchEduroamSchool, "get", params)
	return
}

// AddEduroamSchool_ @title AddEduroamSchool_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func AddEduroamSchool_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(AddEduroamSchool, "post", params)
	return
}

// UpdateEduroamSchool_ @title UpdateEduroamSchool_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func UpdateEduroamSchool_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(UpdateEduroamSchool, "post", params)
	return
}

// DelEduroamSchool_ @title DelEduroamSchool_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func DelEduroamSchool_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(DelEduroamSchool, "post", params)
	return
}

// SearchEduroamVisitor_ @title SearchEduroamVisitor_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func SearchEduroamVisitor_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(SearchEduroamVisitor, "get", params)
	return
}

// ListEduroamVisitor_ @title ListEduroamVisitor_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func ListEduroamVisitor_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(ListEduroamVisitor, "get", params)
	return
}

// TopEduroamVisitor_ @title TopEduroamVisitor_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func TopEduroamVisitor_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(TopEduroamVisitor, "get", params)
	return
}

// DomainTopEduroamVisitor_ @title DomainTopEduroamVisitor_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func DomainTopEduroamVisitor_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(DomainTopEduroamVisitor, "get", params)
	return
}

// DayDetailEduroamVisitor_ @title DayDetailEduroamVisitor_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func DayDetailEduroamVisitor_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(DayDetailEduroamVisitor, "get", params)
	return
}

// InVisitorEduroamVisitor_ @title InVisitorEduroamVisitor_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func InVisitorEduroamVisitor_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(InVisitorEduroamVisitor, "get", params)
	return
}

// OutVisitorEduroamVisitor_ @title OutVisitorEduroamVisitor_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func OutVisitorEduroamVisitor_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(OutVisitorEduroamVisitor, "get", params)
	return
}

// VisitorEduroamDetail_ @title VisitorEduroamDetail_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func VisitorEduroamDetail_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(VisitorEduroamDetail, "get", params)
	return
}

// OnlineIndexEduroam_ @title OnlineIndexEduroam_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func OnlineIndexEduroam_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(OnlineIndexEduroam, "get", params)
	return
}

// DayOnlineIndexEduroamCount_ @title DayOnlineIndexEduroamCount_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func DayOnlineIndexEduroamCount_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(DayOnlineIndexEduroamCount, "get", params)
	return
}

// DropOnlineIndexEduroam_ @title DropOnlineIndexEduroam_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func DropOnlineIndexEduroam_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(DropOnlineIndexEduroam, "post", params)
	return
}

// EquipmentOnlineIndexEduroam_ @title EquipmentOnlineIndexEduroam_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func EquipmentOnlineIndexEduroam_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(EquipmentOnlineIndexEduroam, "post", params)
	return
}

// GroupCreateUrl_ @title GroupCreateUrl_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func GroupCreateUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(GroupCreateUrl, "post", params)
	return
}

// MaxOnlineNumUrl_ @title MaxOnlineNumUrl_
// @description omit
// @auth zhanglianfeng 2021-04-07
// @param params map[string]string
// @return httpResult *HttpResult, err error
func MaxOnlineNumUrl_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(MaxOnlineNum, "post", params)
	return
}

// IdmDeviceIndex_ @title IdmDeviceIndex_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func IdmDeviceIndex_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(IdmDeviceIndex, "post", params)
	return
}

// IdmDeviceFactory_ @title IdmDeviceFactory_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func IdmDeviceFactory_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(IdmDeviceFactory, "post", params)
	return
}

// IdmDeviceOs_ @title IdmDeviceOs_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func IdmDeviceOs_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(IdmDeviceOs, "post", params)
	return
}

// IdmDeviceCreate_ @title IdmDeviceCreate_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func IdmDeviceCreate_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(IdmDeviceCreate, "post", params)
	return
}

// IdmProductIndex_ @title IdmProductIndex_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func IdmProductIndex_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(IdmProductIndex, "post", params)
	return
}

// IdmProductUpdate_ @title IdmProductUpdate_
// @description omit
// @author DM
// @time 2021/4/2 21:11
// @param params
// @return httpResult
// @return err
func IdmProductUpdate_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(IdmProductUpdate, "post", params)
	return
}

// AuthErrMsg_ @title AuthErrMsg_
// @description omit
// @author FangWenQiang
// @time 2021/4/18 14:11
// @param params
// @return httpResult
// @return err
func AuthErrMsg_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(AuthErrMsg, "get", params)
	return
}

// DropReason_ @title DropReason_
// @description omit
// @author FangWenQiang
// @time 2021/4/18 14:11
// @param params
// @return httpResult
// @return err
func DropReason_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(DropReason, "get", params)
	return
}

// SysMsg_ @title SysMsg_
// @description omit
// @author FangWenQiang
// @time 2021/4/18 14:11
// @param params
// @return httpResult
// @return err
func SysMsg_(params map[string]string) (httpResult *HttpResult, err error) {
	httpResult, err = Request(SysMsg, "get", params)
	return
}
