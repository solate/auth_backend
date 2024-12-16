package xerr

// OK 成功返回
const OK uint32 = 200
const InternalError uint32 = 500

/**(前3位代表业务,后三位代表具体功能)**/

// ServerCommonError 全局错误码
const ServerCommonError uint32 = 200001
const RequestParamError uint32 = 200002
const DbError uint32 = 200005
const DbUpdateAffectedZeroError uint32 = 100006
const NotifyError uint32 = 200007
const GetGlobalConfigError uint32 = 200008
const RecordNotFound = 200009
const FailedToAcquireLock = 200010
const DbEditRecordFailed = 200012
const DbCreateRecordFailed = 200013
const DbRecordExist = 200014
const DbRecordExistsUnderMerchant = 200015
const DbRecordExistsName = 200016
const DbSelectRecordFailed = 200017
const UnmarshalFailed = 200018
const InnerTokenInvalid = 200019 // 内部鉴权错误
const QueryError = 200020
const UpdateError = 200021
const CreateError = 200022

// 用户
const (
	UserNotExists       uint32 = 200101
	UserPwdInvalid      uint32 = 200102
	UserTokenInvalid    uint32 = 200103
	UserLoginExpired    uint32 = 200104
	UserHasNoPermission uint32 = 200105
	UserDisabled        uint32 = 200106
	UserExpired         uint32 = 200107
)

// 代理商
const (
	AgencyNameExists         uint32 = 200201
	AgencyContactPhoneExists uint32 = 200202
	AgencyEnabled            uint32 = 200203
	AgencyDisabled           uint32 = 200204
	AgencyToWmzjNoExist      uint32 = 200205
	WmzjToAgencyNoExist      uint32 = 200206
	CreateAgencyExpertCount  uint32 = 200207
)

// 商家
const (
	StoreNameExists         uint32 = 200301
	StoreContactPhoneExists uint32 = 200302
	StoreEnabled            uint32 = 202303
	StoreDisabled           uint32 = 202304
	StoreMTExists           uint32 = 200305
	StoreMTAuthUrlError     uint32 = 200306
	StoreElmExists          uint32 = 200307
	StoreElmAuthUrlError    uint32 = 200308
)
