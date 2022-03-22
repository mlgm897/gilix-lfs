package lfs

const (
	LFS_SERVICE_CLASS_CRD = 14
	CRD_SERVICE_OFFSET    = (LFS_SERVICE_CLASS_CRD * 100)
)
const (
	LFS_CRD_SLOT_SIZE = 2000
)

/* CRD Info Commands */
const (
	LFS_INF_CRD_STATUS         = (CRD_SERVICE_OFFSET + 1)
	LFS_INF_CRD_CAPABILITIES   = (CRD_SERVICE_OFFSET + 2)
	LFS_INF_CRD_CARD_UNIT_INFO = (CRD_SERVICE_OFFSET + 3)
	LFS_INF_CRD_SLOT_INFO      = (CRD_SERVICE_OFFSET + 4)
)

/* CRD Execute Commands */
const (
	LFS_CMD_CRD_DISPENSE_CARD      = (CRD_SERVICE_OFFSET + 1)
	LFS_CMD_CRD_EJECT_CARD         = (CRD_SERVICE_OFFSET + 2)
	LFS_CMD_CRD_RETAIN_CARD        = (CRD_SERVICE_OFFSET + 3)
	LFS_CMD_CRD_RESET              = (CRD_SERVICE_OFFSET + 4)
	LFS_CMD_CRD_SET_GUIDANCE_LIGHT = (CRD_SERVICE_OFFSET + 5)
	LFS_CMD_CRD_SET_CARD_UNIT_INFO = (CRD_SERVICE_OFFSET + 6)
	LFS_CMD_CRD_POWER_SAVE_CONTROL = (CRD_SERVICE_OFFSET + 7)
	LFS_CMD_CRD_SLOT_DEPOSIT       = (CRD_SERVICE_OFFSET + 8)
	LFS_CMD_CRD_SLOT_DISPENSE      = (CRD_SERVICE_OFFSET + 9)
	LFS_CMD_CRD_MOVE_CARD          = (CRD_SERVICE_OFFSET + 50)
)

/* CRD Events  */
const (
	LFS_SRVE_CRD_MEDIAREMOVED        = (CRD_SERVICE_OFFSET + 1)
	LFS_SRVE_CRD_CARDUNITINFOCHANGED = (CRD_SERVICE_OFFSET + 2)
	LFS_SRVE_CRD_MEDIADETECTED       = (CRD_SERVICE_OFFSET + 3)
	LFS_USRE_CRD_CARDUNITTHRESHOLD   = (CRD_SERVICE_OFFSET + 4)
	LFS_EXEE_CRD_CARDUNITERROR       = (CRD_SERVICE_OFFSET + 5)
	LFS_SRVE_CRD_DEVICEPOSITION      = (CRD_SERVICE_OFFSET + 6)
	LFS_SRVE_CRD_POWER_SAVE_CHANGE   = (CRD_SERVICE_OFFSET + 7)
)

/* values of LFSCRDSTATUS.wDevice */
const (
	LFS_CRD_DEVONLINE       = LFS_STAT_DEVONLINE
	LFS_CRD_DEVOFFLINE      = LFS_STAT_DEVOFFLINE
	LFS_CRD_DEVPOWEROFF     = LFS_STAT_DEVPOWEROFF
	LFS_CRD_DEVNODEVICE     = LFS_STAT_DEVNODEVICE
	LFS_CRD_DEVHWERROR      = LFS_STAT_DEVHWERROR
	LFS_CRD_DEVUSERERROR    = LFS_STAT_DEVUSERERROR
	LFS_CRD_DEVBUSY         = LFS_STAT_DEVBUSY
	LFS_CRD_DEVFRAUDATTEMPT = LFS_STAT_DEVFRAUDATTEMPT
)

/* values of LFSCRDSTATUS.wDispenser */
const (
	LFS_CRD_DISPCUOK      = 0
	LFS_CRD_DISPCUSTATE   = 1
	LFS_CRD_DISPCUSTOP    = 2
	LFS_CRD_DISPCUUNKNOWN = 3
)

/* values of LFSCRDSTATUS.wMedia,
   LFSCRDRETAINCARD.wPosition, and
   LFSCRDMEDIADETECTED.wPosition */
const (
	LFS_CRD_MEDIAPRESENT    = 1
	LFS_CRD_MEDIANOTPRESENT = 2
	LFS_CRD_MEDIAJAMMED     = 3
	LFS_CRD_MEDIANOTSUPP    = 4
	LFS_CRD_MEDIAUNKNOWN    = 5
	LFS_CRD_MEDIAEXITING    = 6
	LFS_CRD_MEDIARETAINED   = 7
)

/* values of LFSCRDSTATUS.wTransport */
const (
	LFS_CRD_TPOK           = 0
	LFS_CRD_TPINOP         = 1
	LFS_CRD_TPUNKNOWN      = 2
	LFS_CRD_TPNOTSUPPORTED = 3
)

/* Size and max index of dwGuid_Lights array */
const (
	LFS_CRD_GUIDLIGHTS_SIZE = 32
	LFS_CRD_GUIDLIGHTS_MAX  = (LFS_CRD_GUIDLIGHTS_SIZE - 1)
)

/* Indices of LFSCRDSTATUS.dwGuid_Lights [...]
   LFSCRDCAPS.dwGuid_Lights [...] */
const (
	LFS_CRD_GUIDANCE_CARDDISP = 0
)

/* Values of LFSCRDSTATUS.dwGuid_Lights [...]
   LFSCRDCAPS.dwGuid_Lights [...] */
const (
	LFS_CRD_GUIDANCE_NOT_AVAILABLE = 0x00000000
	LFS_CRD_GUIDANCE_OFF           = 0x00000001
	LFS_CRD_GUIDANCE_SLOW_FLASH    = 0x00000004
	LFS_CRD_GUIDANCE_MEDIUM_FLASH  = 0x00000008
	LFS_CRD_GUIDANCE_QUICK_FLASH   = 0x00000010
	LFS_CRD_GUIDANCE_CONTINUOUS    = 0x00000080
	LFS_CRD_GUIDANCE_RED           = 0x00000100
	LFS_CRD_GUIDANCE_GREEN         = 0x00000200
	LFS_CRD_GUIDANCE_YELLOW        = 0x00000400
	LFS_CRD_GUIDANCE_BLUE          = 0x00000800
	LFS_CRD_GUIDANCE_CYAN          = 0x00001000
	LFS_CRD_GUIDANCE_MAGENTA       = 0x00002000
	LFS_CRD_GUIDANCE_WHITE         = 0x00004000
)

/* values of LFSCRDSTATUS.wDevice_Position
   LFSCRDDEVICEPOSITION.wPosition */
const (
	LFS_CRD_DEVICEINPOSITION    = 0
	LFS_CRD_DEVICENOTINPOSITION = 1
	LFS_CRD_DEVICEPOSUNKNOWN    = 2
	LFS_CRD_DEVICEPOSNOTSUPP    = 3
)

/*values of LFSCRDCAPS.wDispense_To */
const (
	LFS_CRD_DISPTO_CONSUMER  = 0x0001
	LFS_CRD_DISPTO_TRANSPORT = 0x0002
)

/*values of LFSCRDCARDUNIT.usStatus */
const (
	LFS_CRD_STATCUOK      = 0
	LFS_CRD_STATCULOW     = 1
	LFS_CRD_STATCUEMPTY   = 2
	LFS_CRD_STATCUINOP    = 3
	LFS_CRD_STATCUMISSING = 4
	LFS_CRD_STATCUHIGH    = 5
	LFS_CRD_STATCUFULL    = 6
	LFS_CRD_STATCUUNKNOWN = 7
)

/*values of LFSCRDCARDUNIT.usType */
const (
	LFS_CRD_SUPPLYBIN = 1
	LFS_CRD_RETAINBIN = 2
)

/* values of LFSCRDSTATUS.wShutter */
const (
	LFS_CRD_SHTCLOSED       = 0
	LFS_CRD_SHTOPEN         = 1
	LFS_CRD_SHTJAMMED       = 2
	LFS_CRD_SHTUNKNOWN      = 3
	LFS_CRD_SHTNOTSUPPORTED = 4
)

/* values of LFSCRDCAPS.wPower_On_Option,
   LFSCRDCAPS.wPower_Off_Option,
   LFSCRDRESET.usAction  */
const (
	LFS_CRD_NOACTION        = 1
	LFS_CRD_EJECT           = 2
	LFS_CRD_RETAIN          = 3
	LFS_CRD_EJECTTHENRETAIN = 4
)

/*values of LFSCRDCUERROR.wFailure */
const (
	LFS_CRD_CARDUNITEMPTY   = 1
	LFS_CRD_CARDUNITERROR   = 2
	LFS_CRD_CARDUNITINVALID = 3
)

/* values of LFSCRDSLOTUNIT.wSlot_Status */
const (
	LFS_CRD_SLOT_MEDIAPRESENT    = 1
	LFS_CRD_SLOT_MEDIANOTPRESENT = 2
	LFS_CRD_SLOT_MEDIAUNKNOW     = 3
)

/* LFS CRD Errors */
const (
	LFS_ERR_CRD_MEDIAJAM              = (-(CRD_SERVICE_OFFSET + 0))
	LFS_ERR_CRD_NOMEDIA               = (-(CRD_SERVICE_OFFSET + 1))
	LFS_ERR_CRD_MEDIARETAINED         = (-(CRD_SERVICE_OFFSET + 2))
	LFS_ERR_CRD_RETAINBINFULL         = (-(CRD_SERVICE_OFFSET + 3))
	LFS_ERR_CRD_SHUTTERFAIL           = (-(CRD_SERVICE_OFFSET + 4))
	LFS_ERR_CRD_DEVICE_OCCUPIED       = (-(CRD_SERVICE_OFFSET + 5))
	LFS_ERR_CRD_CARDUNITERROR         = (-(CRD_SERVICE_OFFSET + 6))
	LFS_ERR_CRD_INVALIDCARDUNIT       = (-(CRD_SERVICE_OFFSET + 7))
	LFS_ERR_CRD_INVALID_PORT          = (-(CRD_SERVICE_OFFSET + 8))
	LFS_ERR_CRD_INVALIDRETAINBIN      = (-(CRD_SERVICE_OFFSET + 9))
	LFS_ERR_CRD_POWERSAVETOOSHORT     = (-(CRD_SERVICE_OFFSET + 10))
	LFS_ERR_CRD_POWERSAVEMEDIAPRESENT = (-(CRD_SERVICE_OFFSET + 11))
	LFS_ERR_CRD_MEDIA_EXIST           = (-(CRD_SERVICE_OFFSET + 12))
	LFS_ERR_CRD_MEDIA_NOTEXIST        = (-(CRD_SERVICE_OFFSET + 13))
)

/*=================================================================*/
/* CRD Info Command Structures and variables */
/*=================================================================*/
type LFSCRDSTATUS struct {
	Device                uint
	Dispenser             uint
	Transport             uint
	Media                 uint
	Shutter               uint
	Extra                 []byte
	GuidLights            [LFS_CRD_GUIDLIGHTS_SIZE]uint
	DevicePosition        uint
	PowerSaveRecoveryTime uint
}
type LFSCRDCAPS struct {
	Class            uint
	Compound         int
	PowerOnOption    uint
	PowerOffOption   uint
	CardTakenSensor  int
	DispenseTo       uint
	Extra            []byte
	GuidLights       [LFS_CRD_GUIDLIGHTS_SIZE]uint
	PowerSaveControl int
}

type LFSCRDCARDUNIT struct {
	Number         uint
	CardName       string
	CardType       uint
	InitialCount   uint
	Count          uint
	RetainCount    uint
	Threshold      uint
	Status         uint
	HardwareSensor int
}

type LFSCRDCUINFO struct {
	Count uint
	List  []LFSCRDCARDUNIT
}

type LFSCRDSLOTINFO struct {
	RealTotalSlotCount uint
	SlotStatus         [LFS_CRD_SLOT_SIZE]uint
}

/*=================================================================*/
/* CRD Execute Command Structures */
/*=================================================================*/
type LFSCRDDISPENSE struct {
	usNumber uint
	bPresent int
}

type LFSCRDRETAINCARD struct {
	usNumber uint
}

type LFSCRDRESET struct {
	usAction uint
}

type LFSCRDSETGUIDLIGHT struct {
	GuidLight uint
	Command   uint
}

type LFSCRDPOWERSAVECONTROL struct {
	MaxPowerSaveRecoveryTime uint
}

type LFSCRDSLOTDEPOSIT struct {
	SlotNo uint
}

type LFSCRDSLOTDISPENSE struct {
	SlotNo  uint
	Present int
}

//自定义
//LFS_CMD_CRD_MOVE_CARD
type DefInLFSCRDMOVECARD struct {
	MoveOptionIn uint
}

/*=================================================================*/
/* CRD Message Structures */
/*=================================================================*/
type LFSCRDMEDIADETECTED struct {
	Position uint
	Number   uint
}

type LFSCRDCUERROR struct {
	Failure  uint
	CardUnit LFSCRDCARDUNIT
}

type LFSCRDDEVICEPOSITION struct {
	Position uint
}

type LFSCRDPOWERSAVECHANGE struct {
	PowerSaveRecoveryTime uint
}
