package lfs

/****** 错误码 ************************************************************/

const (
	LFS_SUCCESS                    = 0
	LFS_ERR_ALREADY_STARTED        = -1
	LFS_ERR_API_VER_TOO_HIGH       = -2
	LFS_ERR_API_VER_TOO_LOW        = -3
	LFS_ERR_CANCELED               = -4
	LFS_ERR_DEV_NOT_READY          = -13
	LFS_ERR_HARDWARE_ERROR         = -14
	LFS_ERR_INTERNAL_ERROR         = -15
	LFS_ERR_INVALID_ADDRESS        = -16
	LFS_ERR_INVALID_APP_HANDLE     = -17
	LFS_ERR_INVALID_BUFFER         = -18
	LFS_ERR_INVALID_CATEGORY       = -19
	LFS_ERR_INVALID_COMMAND        = -20
	LFS_ERR_INVALID_EVENT_CLASS    = -21
	LFS_ERR_INVALID_HSERVICE       = -22
	LFS_ERR_INVALID_HPROVIDER      = -23
	LFS_ERR_INVALID_MSG_OBJECT     = -24
	LFS_ERR_INVALID_MSG_REG_OBJECT = -25
	LFS_ERR_INVALID_POINTER        = -26
	LFS_ERR_INVALID_REQ_ID         = -27
	LFS_ERR_INVALID_RESULT         = -28
	LFS_ERR_INVALID_SERVPROV       = -29
	LFS_ERR_INVALID_TIMER          = -30
	LFS_ERR_INVALID_TRACELEVEL     = -31
	LFS_ERR_LOCKED                 = -32
	LFS_ERR_NO_BLOCKING_CALL       = -33
	LFS_ERR_NO_SERVPROV            = -34
	LFS_ERR_NO_SUCH_THREAD         = -35
	LFS_ERR_NO_TIMER               = -36
	LFS_ERR_NOT_LOCKED             = -37
	LFS_ERR_NOT_OK_TO_UNLOAD       = -38
	LFS_ERR_NOT_STARTED            = -39
	LFS_ERR_NOT_REGISTERED         = -40
	LFS_ERR_OP_IN_PROGRESS         = -41
	LFS_ERR_OUT_OF_MEMORY          = -42
	LFS_ERR_SERVICE_NOT_FOUND      = -43
	LFS_ERR_SPI_VER_TOO_HIGH       = -44
	LFS_ERR_SPI_VER_TOO_LOW        = -45
	LFS_ERR_SRVC_VER_TOO_HIGH      = -46
	LFS_ERR_SRVC_VER_TOO_LOW       = -47
	LFS_ERR_TIMEOUT                = -48
	LFS_ERR_UNSUPP_CATEGORY        = -49
	LFS_ERR_UNSUPP_COMMAND         = -50
	LFS_ERR_VERSION_ERROR_IN_SRVC  = -51
	LFS_ERR_INVALID_DATA           = -52
	LFS_ERR_SOFTWARE_ERROR         = -53
	LFS_ERR_CONNECTION_LOST        = -54
	LFS_ERR_USER_ERROR             = -55
	LFS_ERR_UNSUPP_DATA            = -56
)

/****** 事件类型 **********************************************************/

const (
	SERVICE_EVENTS = 1
	USER_EVENTS    = 2
	SYSTEM_EVENTS  = 4
	EXECUTE_EVENTS = 8
)

/****** 系统事件ID ********************************************************/

const (
	LFS_SYSE_UNDELIVERABLE_MSG = 1
	LFS_SYSE_HARDWARE_ERROR    = 2
	LFS_SYSE_VERSION_ERROR     = 3
	LFS_SYSE_DEVICE_STATUS     = 4
	LFS_SYSE_APP_DISCONNECT    = 5
	LFS_SYSE_SOFTWARE_ERROR    = 6
	LFS_SYSE_USER_ERROR        = 7
	LFS_SYSE_LOCK_REQUESTED    = 8
)
