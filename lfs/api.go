package lfs

/****** LFSDEVSTATUS.state 的取值 *****************************************/

const (
	LFS_STAT_DEVONLINE         = 0
	LFS_STAT_DEVOFFLINE        = 1
	LFS_STAT_DEVPOWEROFF       = 2
	LFS_STAT_DEVNODEVICE       = 3
	LFS_STAT_DEVHWERROR        = 4
	LFS_STAT_DEVUSERERROR      = 5
	LFS_STAT_DEVBUSY           = 6
	LFS_STAT_DEVFRAUDATTEMPT   = 7
	LFS_STAT_DEVPOTENTIALFRAUD = 8
)

/****** 消息结构体定义 ***************************************************/

type LFSDEVSTATUS struct {
	PhysicalName    string
	WorkstationName string
	State           uint
}

type LFSAPPDISC struct {
	LogicalName     string
	WorkstationName string
	AppID           string
}

type LFSHWERROR struct {
	LogicalName     string
	PhysicalName    string
	WorkstationName string
	AppID           string
	Action          uint
	Size            uint
	Description     []byte
}

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

/****** LFS 错误时采取的动作 *********************************************/

const (
	LFS_ERR_ACT_NOACTION = 0x0000
	LFS_ERR_ACT_RESET    = 0x0001
	LFS_ERR_ACT_SWERROR  = 0x0002
	LFS_ERR_ACT_CONFIG   = 0x0004
	LFS_ERR_ACT_HWCLEAR  = 0x0008
	LFS_ERR_ACT_HWMAINT  = 0x0010
	LFS_ERR_ACT_SUSPEND  = 0x0020
)
