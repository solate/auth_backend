package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[RequestParamError] = "参数错误"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[DbUpdateAffectedZeroError] = "更新数据影响行数为0"
	message[NotifyError] = "!!! 请联系开发人员确认操作 !!!"
	message[GetGlobalConfigError] = "获取全局配置失败"
	message[RecordNotFound] = "未找到该记录"
	message[DbSelectRecordFailed] = "查询记录失败"
	message[FailedToAcquireLock] = "获取锁失败"
	message[InternalError] = "内部错误"
	message[DbEditRecordFailed] = "修改数据库记录失败"
	message[DbCreateRecordFailed] = "创建数据库记录失败"
	message[DbRecordExist] = "数据库记录已存在"
	message[DbRecordExistsUnderMerchant] = "该poi门店已在所修改的门店下存在"
	message[DbRecordExistsName] = "名称已存在"
	message[UnmarshalFailed] = "数据序列化错误"
	message[QueryError] = "查询错误"
	message[UpdateError] = "更新错误"
	message[CreateError] = "保存错误"

	message[UserNotExists] = "用户不存在"
	message[UserPwdInvalid] = "用户密码错误"
	message[UserTokenInvalid] = "无效的token"
	message[UserLoginExpired] = "用户登录已过期"
	message[UserHasNoPermission] = "用户无权限"
	message[InnerTokenInvalid] = "内部鉴权错误"

	// 商户映射
	message[StoreMTExists] = "该商户美团商铺已存在"
	message[StoreMTAuthUrlError] = "获取美团认证链接失败"
	message[StoreElmExists] = "该商户饿了么商铺已存在"
	message[StoreElmAuthUrlError] = "获取饿了么认证链接失败"

	// 代理商映射
	message[AgencyToWmzjNoExist] = "代理商映射外卖专家账号未找到"
	message[WmzjToAgencyNoExist] = "外卖专家映射代理商账号未找到"
	message[CreateAgencyExpertCount] = "服务商剩余专家版个数少于输入专家版数量"
	message[UserDisabled] = "用户已被禁用"
	message[UserExpired] = "用户已过期"

	message[AgencyNameExists] = "代理商名称已存在"
	message[AgencyContactPhoneExists] = "代理商联系人电话已存在"

	message[StoreNameExists] = "商家名称已存在"
	message[StoreContactPhoneExists] = "商家联系人电话已存在"
}

func MapErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errCode uint32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}
