package lfs

const (
	LFS_SERVICE_CLASS_IDC = 2
	IDC_SERVICE_OFFSET    = (LFS_SERVICE_CLASS_IDC * 100)
)

/* IDC Info Commands */

const (
	LFS_INF_IDC_STATUS               = (IDC_SERVICE_OFFSET + 1)
	LFS_INF_IDC_CAPABILITIES         = (IDC_SERVICE_OFFSET + 2)
	LFS_INF_IDC_FORM_LIST            = (IDC_SERVICE_OFFSET + 3)
	LFS_INF_IDC_QUERY_FORM           = (IDC_SERVICE_OFFSET + 4)
	LFS_INF_IDC_QUERY_IFM_IDENTIFIER = (IDC_SERVICE_OFFSET + 5)
	LFS_INF_IDC_MCR_CAPABILITIES     = (IDC_SERVICE_OFFSET + 51)
	LFS_INF_IDC_MCR_STATUS           = (IDC_SERVICE_OFFSET + 52)
	LFS_INF_IDC_SLOTS_INFO           = (IDC_SERVICE_OFFSET + 53)
)

/* IDC Execute Commands */

const (
	LFS_CMD_IDC_READ_TRACK           = (IDC_SERVICE_OFFSET + 1)
	LFS_CMD_IDC_WRITE_TRACK          = (IDC_SERVICE_OFFSET + 2)
	LFS_CMD_IDC_EJECT_CARD           = (IDC_SERVICE_OFFSET + 3)
	LFS_CMD_IDC_RETAIN_CARD          = (IDC_SERVICE_OFFSET + 4)
	LFS_CMD_IDC_RESET_COUNT          = (IDC_SERVICE_OFFSET + 5)
	LFS_CMD_IDC_SETKEY               = (IDC_SERVICE_OFFSET + 6)
	LFS_CMD_IDC_READ_RAW_DATA        = (IDC_SERVICE_OFFSET + 7)
	LFS_CMD_IDC_WRITE_RAW_DATA       = (IDC_SERVICE_OFFSET + 8)
	LFS_CMD_IDC_CHIP_IO              = (IDC_SERVICE_OFFSET + 9)
	LFS_CMD_IDC_RESET                = (IDC_SERVICE_OFFSET + 10)
	LFS_CMD_IDC_CHIP_POWER           = (IDC_SERVICE_OFFSET + 11)
	LFS_CMD_IDC_PARSE_DATA           = (IDC_SERVICE_OFFSET + 12)
	LFS_CMD_IDC_SET_GUIDANCE_LIGHT   = (IDC_SERVICE_OFFSET + 13)
	LFS_CMD_IDC_POWER_SAVE_CONTROL   = (IDC_SERVICE_OFFSET + 14)
	LFS_CMD_IDC_RETRACT_CARD         = (IDC_SERVICE_OFFSET + 91)
	LFS_CMD_IDC_HANDBACK_CARD        = (IDC_SERVICE_OFFSET + 92)
	LFS_CMD_IDC_CLEAR_RETRACT_CARD   = (IDC_SERVICE_OFFSET + 93)
	LFS_CMD_IDC_RETAIN_SLOT_START    = (IDC_SERVICE_OFFSET + 51)
	LFS_CMD_IDC_RETAIN_SLOT_END      = (IDC_SERVICE_OFFSET + 52)
	LFS_CMD_IDC_EJECT_SLOT_START     = (IDC_SERVICE_OFFSET + 53)
	LFS_CMD_IDC_EJECT_SLOT_END       = (IDC_SERVICE_OFFSET + 54)
	LFS_CMD_IDC_MCR_RESET            = (IDC_SERVICE_OFFSET + 58)
	LFS_CMD_IDC_RETAIN_SLOT_START_EX = (IDC_SERVICE_OFFSET + 96)
	LFS_CMD_IDC_MCR_RETAIN_CARD      = (IDC_SERVICE_OFFSET + 97)
	LFS_CMD_IDC_OPEN_SHUTTER         = (IDC_SERVICE_OFFSET + 98)
	LFS_CMD_IDC_CLOSE_SHUTTER        = (IDC_SERVICE_OFFSET + 99)
)

/* IDC Messages */

const (
	LFS_EXEE_IDC_INVALIDTRACKDATA   = (IDC_SERVICE_OFFSET + 1)
	LFS_EXEE_IDC_MEDIAINSERTED      = (IDC_SERVICE_OFFSET + 3)
	LFS_SRVE_IDC_MEDIAREMOVED       = (IDC_SERVICE_OFFSET + 4)
	LFS_SRVE_IDC_CARDACTION         = (IDC_SERVICE_OFFSET + 5)
	LFS_USRE_IDC_RETAINBINTHRESHOLD = (IDC_SERVICE_OFFSET + 6)
	LFS_EXEE_IDC_INVALIDMEDIA       = (IDC_SERVICE_OFFSET + 7)
	LFS_EXEE_IDC_MEDIARETAINED      = (IDC_SERVICE_OFFSET + 8)
	LFS_SRVE_IDC_MEDIADETECTED      = (IDC_SERVICE_OFFSET + 9)
	LFS_SRVE_IDC_RETAINBININSERTED  = (IDC_SERVICE_OFFSET + 10)
	LFS_SRVE_IDC_RETAINBINREMOVED   = (IDC_SERVICE_OFFSET + 11)
	LFS_EXEE_IDC_INSERTCARD         = (IDC_SERVICE_OFFSET + 12)
	LFS_SRVE_IDC_DEVICEPOSITION     = (IDC_SERVICE_OFFSET + 13)
	LFS_SRVE_IDC_POWER_SAVE_CHANGE  = (IDC_SERVICE_OFFSET + 14)
)

/* values of LFSIDCSTATUS.wDevice,LFSIDCMCRSTATUS.wDevice */

const (
	LFS_IDC_DEVONLINE       = LFS_STAT_DEVONLINE
	LFS_IDC_DEVOFFLINE      = LFS_STAT_DEVOFFLINE
	LFS_IDC_DEVPOWEROFF     = LFS_STAT_DEVPOWEROFF
	LFS_IDC_DEVNODEVICE     = LFS_STAT_DEVNODEVICE
	LFS_IDC_DEVHWERROR      = LFS_STAT_DEVHWERROR
	LFS_IDC_DEVUSERERROR    = LFS_STAT_DEVUSERERROR
	LFS_IDC_DEVBUSY         = LFS_STAT_DEVBUSY
	LFS_IDC_DEVFRAUDATTEMPT = LFS_STAT_DEVFRAUDATTEMPT
)

/* values of LFSIDCSTATUS.wMedia, LFSIDCRETAINCARD.wPosition,  */
/* LFSIDCCARDACT.wPosition  */

const (
	LFS_IDC_MEDIAPRESENT    = 1
	LFS_IDC_MEDIANOTPRESENT = 2
	LFS_IDC_MEDIAJAMMED     = 3
	LFS_IDC_MEDIANOTSUPP    = 4
	LFS_IDC_MEDIAUNKNOWN    = 5
	LFS_IDC_MEDIAENTERING   = 6
	LFS_IDC_MEDIALATCHED    = 7
)

/* values of LFSIDCSTATUS.wRetainBin */

const (
	LFS_IDC_RETAINBINOK      = 1
	LFS_IDC_RETAINNOTSUPP    = 2
	LFS_IDC_RETAINBINFULL    = 3
	LFS_IDC_RETAINBINHIGH    = 4
	LFS_IDC_RETAINBINMISSING = 5
)

/* values of LFSIDCSTATUS.wSecurity */

const (
	LFS_IDC_SECNOTSUPP  = 1
	LFS_IDC_SECNOTREADY = 2
	LFS_IDC_SECOPEN     = 3
)

/* values of LFSIDCSTATUS.wChipPower */

const (
	LFS_IDC_CHIPONLINE     = 0
	LFS_IDC_CHIPPOWEREDOFF = 1
	LFS_IDC_CHIPBUSY       = 2
	LFS_IDC_CHIPNODEVICE   = 3
	LFS_IDC_CHIPHWERROR    = 4
	LFS_IDC_CHIPNOCARD     = 5
	LFS_IDC_CHIPNOTSUPP    = 6
	LFS_IDC_CHIPUNKNOWN    = 7
)

/* Size and max index of dwGuidLights array */

const (
	LFS_IDC_GUIDLIGHTS_SIZE = 32
	LFS_IDC_GUIDLIGHTS_MAX  = (LFS_IDC_GUIDLIGHTS_SIZE - 1)
)

/* Indices of LFSIDCSTATUS.dwGuidLights [...]
   LFSIDCCAPS.dwGuidLights [...]
*/

const (
	LFS_IDC_GUIDANCE_CARDUNIT = 0
)

/* Values of LFSIDCSTATUS.dwGuidLights [...]
   LFSIDCCAPS.dwGuidLights [...]
*/

const (
	LFS_IDC_GUIDANCE_NOT_AVAILABLE = 0x00000000
	LFS_IDC_GUIDANCE_OFF           = 0x00000001
	LFS_IDC_GUIDANCE_ON            = 0x00000002
	LFS_IDC_GUIDANCE_SLOW_FLASH    = 0x00000004
	LFS_IDC_GUIDANCE_MEDIUM_FLASH  = 0x00000008
	LFS_IDC_GUIDANCE_QUICK_FLASH   = 0x00000010
	LFS_IDC_GUIDANCE_CONTINUOUS    = 0x00000080
	LFS_IDC_GUIDANCE_RED           = 0x00000100
	LFS_IDC_GUIDANCE_GREEN         = 0x00000200
	LFS_IDC_GUIDANCE_YELLOW        = 0x00000400
	LFS_IDC_GUIDANCE_BLUE          = 0x00000800
	LFS_IDC_GUIDANCE_CYAN          = 0x00001000
	LFS_IDC_GUIDANCE_MAGENTA       = 0x00002000
	LFS_IDC_GUIDANCE_WHITE         = 0x00004000
)

/* values of LFSIDCSTATUS.wChipModule */

const (
	LFS_IDC_CHIPMODOK      = 1
	LFS_IDC_CHIPMODINOP    = 2
	LFS_IDC_CHIPMODUNKNOWN = 3
	LFS_IDC_CHIPMODNOTSUPP = 4
)

/* values of LFSIDCSTATUS.wMagReadModule and
   LFSIDCSTATUS.wMagWriteModule  */

const (
	LFS_IDC_MAGMODOK      = 1
	LFS_IDC_MAGMODINOP    = 2
	LFS_IDC_MAGMODUNKNOWN = 3
	LFS_IDC_MAGMODNOTSUPP = 4
)

/* values of LFSIDCSTATUS.wFrontImageModule and
   LFSIDCSTATUS.wBackImageModule */

const (
	LFS_IDC_IMGMODOK      = 1
	LFS_IDC_IMGMODINOP    = 2
	LFS_IDC_IMGMODUNKNOWN = 3
	LFS_IDC_IMGMODNOTSUPP = 4
)

/* values of LFSIDCSTATUS.wDevicePosition
   LFSIDCDEVICEPOSITION.wPosition */

const (
	LFS_IDC_DEVICEINPOSITION    = 0
	LFS_IDC_DEVICENOTINPOSITION = 1
	LFS_IDC_DEVICEPOSUNKNOWN    = 2
	LFS_IDC_DEVICEPOSNOTSUPP    = 3
)

/* values of LFSIDCCAPS.wType */

const (
	LFS_IDC_TYPEMOTOR       = 1
	LFS_IDC_TYPESWIPE       = 2
	LFS_IDC_TYPEDIP         = 3
	LFS_IDC_TYPECONTACTLESS = 4
	LFS_IDC_TYPELATCHEDDIP  = 5
	LFS_IDC_TYPEPERMANENT   = 6
)

/* values of LFSIDCCAPS.wReadTracks,
   LFSIDCCAPS.wWriteTracks,
   LFSIDCCARDDATA.wDataSource,
   LFSIDCCAPS.wChipProtocols,
   LFSIDCCAPS.wWriteMode,
   LFSIDCCAPS.wChipPower */

const (
	LFS_IDC_NOTSUPP = 0x0000
)

/* values of LFSIDCCAPS.wReadTracks, LFSIDCCAPS.wWriteTracks,
   LFSIDCCARDDATA.wDataSource,
   LFS_CMD_IDC_READ_RAW_DATA */

const (
	LFS_IDC_TRACK1        = 0x0001
	LFS_IDC_TRACK2        = 0x0002
	LFS_IDC_TRACK3        = 0x0004
	LFS_IDC_FRONT_TRACK_1 = 0x0080
)

/* further values of LFSIDCCARDDATA.wDataSource (except
   LFS_IDC_FLUXINACTIVE), LFS_CMD_IDC_READ_RAW_DATA */

const (
	LFS_IDC_CHIP         = 0x0008
	LFS_IDC_SECURITY     = 0x0010
	LFS_IDC_FLUXINACTIVE = 0x0020
	LFS_IDC_TRACK_WM     = 0x8000
	LFS_IDC_MEMORY_CHIP  = 0x0040
	LFS_IDC_FRONTIMAGE   = 0x0100
	LFS_IDC_BACKIMAGE    = 0x0200
)

/* values of LFSIDCCAPS.wChipProtocols */

const (
	LFS_IDC_CHIPT0                     = 0x0001
	LFS_IDC_CHIPT1                     = 0x0002
	LFS_IDC_CHIP_PROTOCOL_NOT_REQUIRED = 0x0004
)

/* values of LFSIDCCAPS.wSecType */

const (
	//LFS_IDC_SECNOTSUPP = 1  LFSIDCSTATUS.wSecurity 重复
	LFS_IDC_SECMMBOX = 2
	LFS_IDC_SECCIM86 = 3
)

/* values of LFSIDCCAPS.wEjectModuleEmptyType*/

const (
	LFS_EM_EMPTYSTORAGE    = 1
	LFS_EM_EMPTYCAPTUREBIN = 2
	LFS_EM_EMPTYALL        = 3
)

/* values of LFSIDCCAPS.wPowerOnOption, LFSIDCCAPS.wPowerOffOption*/

const (
	LFS_IDC_NOACTION        = 1
	LFS_IDC_EJECT           = 2
	LFS_IDC_RETAIN          = 3
	LFS_IDC_EJECTTHENRETAIN = 4
	LFS_IDC_READPOSITION    = 5
)

/* values of LFSIDCCAPS.wWriteMode; LFSIDCWRITETRACK.wWriteMethod,
LFSIDCCARDDATA.wWriteMethod */

/* Note: LFS_IDC_UNKNOWN was removed as it was an invalid value */

const (
	LFS_IDC_LOCO = 0x0002
	LFS_IDC_HICO = 0x0004
	LFS_IDC_AUTO = 0x0008
)

/* values of LFSIDCCAPS.wChipPower */

const (
	LFS_IDC_CHIPPOWERCOLD = 0x0002
	LFS_IDC_CHIPPOWERWARM = 0x0004
	LFS_IDC_CHIPPOWEROFF  = 0x0008

	LFS_IDC_SAM1 = 0x8000
	LFS_IDC_SAM2 = 0x4000
	LFS_IDC_SAM3 = 0x2000
	LFS_IDC_SAM4 = 0x1000
	LFS_IDC_SAM5 = 0x0800
	LFS_IDC_SAM6 = 0x0400
	LFS_IDC_SAM7 = 0x0200
	LFS_IDC_SAM8 = 0x0100
)

/* values of LFSIDCCAPS.wDIPMode */

const (
	LFS_IDC_DIP_UNKNOWN    = 0x0001
	LFS_IDC_DIP_EXIT       = 0x0002
	LFS_IDC_DIP_ENTRY      = 0x0004
	LFS_IDC_DIP_ENTRY_EXIT = 0x0008
)

/* values of LFSIDCCAPS. lpwMemoryChipProtocols */

const (
	LFS_IDC_MEM_SIEMENS4442 = 0x0001
	LFS_IDC_MEM_GPM896      = 0x0002
	LFS_IDC_MEM_1608        = 0x0004
	LFS_IDC_MEM_102         = 0x0008
)

/* values of LFSIDCFORM.wAction */

const (
	LFS_IDC_ACTIONREAD  = 0x0001
	LFS_IDC_ACTIONWRITE = 0x0002
)

/* values of LFSIDCTRACKEVENT.wStatus, LFSIDCCARDDATA.wStatus */

const (
	LFS_IDC_DATAOK         = 0
	LFS_IDC_DATAMISSING    = 1
	LFS_IDC_DATAINVALID    = 2
	LFS_IDC_DATATOOLONG    = 3
	LFS_IDC_DATATOOSHORT   = 4
	LFS_IDC_DATASRCNOTSUPP = 5
	LFS_IDC_DATASRCMISSING = 6
)

/* values LFSIDCCARDACT.wAction */

const (
	LFS_IDC_CARDRETAINED     = 1
	LFS_IDC_CARDEJECTED      = 2
	LFS_IDC_CARDREADPOSITION = 3
	LFS_IDC_CARDJAMMED       = 4
)

/* values of LFSIDCCARDDATA.lpbData if security is read */

const (
	LFS_IDC_SEC_READLEVEL1   = '1'
	LFS_IDC_SEC_READLEVEL2   = '2'
	LFS_IDC_SEC_READLEVEL3   = '3'
	LFS_IDC_SEC_READLEVEL4   = '4'
	LFS_IDC_SEC_READLEVEL5   = '5'
	LFS_IDC_SEC_BADREADLEVEL = '6'
	LFS_IDC_SEC_NODATA       = '7'
	LFS_IDC_SEC_DATAINVAL    = '8'
	LFS_IDC_SEC_HWERROR      = '9'
	LFS_IDC_SEC_NOINIT       = 'A'
)

/* values of LFSIDCIFMIDENTIFIER.wIFMAuthority */

const (
	LFS_IDC_IFMEMV     = 1
	LFS_IDC_IFMEUROPAY = 2
	LFS_IDC_IFMVISA    = 3
	LFS_IDC_IFMGIECB   = 4
)

/* values LFSIDCCAPS.wEjectPosition, LFSIDCEJECTCARD.wEjectPosition */
const (
	LFS_IDC_EXITPOSITION      = 0x0001
	LFS_IDC_TRANSPORTPOSITION = 0x0002
)

/* values if LPLFSIDCCLEARRETRACTCARD.wNum */

const (
	LFS_IDC_CLEARALL = 0
	LFS_IDC_CLEARONE = 1
)

/* value of LFSIDCMCRCAPS.wType */

const (
	LFS_IDC_TYPE = 10
)

/* values of LFSIDCMCRSTATUS.wSlots */

const (
	LFS_IDC_SLOTESMPTY = 0
	LFS_IDC_SLOTSOK    = 1
	LFS_IDC_SLOTSFULL  = 2
)

/*values of LFSIDCSLOTUNITINFO.wStatus */
const (
	LFS_IDC_SLOTUNIT_NOTPRESENT = 0
	LFS_IDC_SLOTUNIT_PRESENT    = 1
)

/* values of LFSIDCEJECTSLOTOUT.wType */

const (
	LFS_IDC_TYPE_SLOTID   = 0
	LFS_IDC_TYPE_CARDDATA = 1
)

/* IDC Errors */

const (
	LFS_ERR_IDC_MEDIAJAM              = (-(IDC_SERVICE_OFFSET + 0))
	LFS_ERR_IDC_NOMEDIA               = (-(IDC_SERVICE_OFFSET + 1))
	LFS_ERR_IDC_MEDIARETAINED         = (-(IDC_SERVICE_OFFSET + 2))
	LFS_ERR_IDC_RETAINBINFULL         = (-(IDC_SERVICE_OFFSET + 3))
	LFS_ERR_IDC_INVALIDDATA           = (-(IDC_SERVICE_OFFSET + 4))
	LFS_ERR_IDC_INVALIDMEDIA          = (-(IDC_SERVICE_OFFSET + 5))
	LFS_ERR_IDC_FORMNOTFOUND          = (-(IDC_SERVICE_OFFSET + 6))
	LFS_ERR_IDC_FORMINVALID           = (-(IDC_SERVICE_OFFSET + 7))
	LFS_ERR_IDC_DATASYNTAX            = (-(IDC_SERVICE_OFFSET + 8))
	LFS_ERR_IDC_SHUTTERFAIL           = (-(IDC_SERVICE_OFFSET + 9))
	LFS_ERR_IDC_SECURITYFAIL          = (-(IDC_SERVICE_OFFSET + 10))
	LFS_ERR_IDC_PROTOCOLNOTSUPP       = (-(IDC_SERVICE_OFFSET + 11))
	LFS_ERR_IDC_ATRNOTOBTAINED        = (-(IDC_SERVICE_OFFSET + 12))
	LFS_ERR_IDC_INVALIDKEY            = (-(IDC_SERVICE_OFFSET + 13))
	LFS_ERR_IDC_WRITE_METHOD          = (-(IDC_SERVICE_OFFSET + 14))
	LFS_ERR_IDC_CHIPPOWERNOTSUPP      = (-(IDC_SERVICE_OFFSET + 15))
	LFS_ERR_IDC_CARDTOOSHORT          = (-(IDC_SERVICE_OFFSET + 16))
	LFS_ERR_IDC_CARDTOOLONG           = (-(IDC_SERVICE_OFFSET + 17))
	LFS_ERR_IDC_INVALID_PORT          = (-(IDC_SERVICE_OFFSET + 18))
	LFS_ERR_IDC_POWERSAVETOOSHORT     = (-(IDC_SERVICE_OFFSET + 19))
	LFS_ERR_IDC_POWERSAVEMEDIAPRESENT = (-(IDC_SERVICE_OFFSET + 20))
	LFS_ERR_IDC_CARDPRESENT           = (-(IDC_SERVICE_OFFSET + 21))
	LFS_ERR_IDC_POSITIONINVALID       = (-(IDC_SERVICE_OFFSET + 22))
	LFS_ERR_IDC_NORETRACTBIN          = (-(IDC_SERVICE_OFFSET + 23))
	LFS_ERR_IDC_RETRACTBINFULL        = (-(IDC_SERVICE_OFFSET + 24))
	LFS_ERR_IDC_CARD_NOTFOUND         = (-(IDC_SERVICE_OFFSET + 25))
	LFS_ERR_IDC_MEDIAEXIST            = (-(IDC_SERVICE_OFFSET + 52))
	LFS_ERR_IDC_MEDIARETRACTJAM       = (-(IDC_SERVICE_OFFSET + 53))
	LFS_ERR_IDC_SLOTFULL              = (-(IDC_SERVICE_OFFSET + 54))
	LFS_ERR_IDC_SLOTNOTEXIST          = (-(IDC_SERVICE_OFFSET + 55))
	LFS_ERR_IDC_SHUTTERBLOCKED        = (-(IDC_SERVICE_OFFSET + 56))
	LFS_ERR_IDC_SHUTTERCLOSEFAILED    = (-(IDC_SERVICE_OFFSET + 57))
	LFS_ERR_IDC_SHUTTEROPENFAILED     = (-(IDC_SERVICE_OFFSET + 58))
	LFS_ERR_IDC_CODENOTMATCH          = (-(IDC_SERVICE_OFFSET + 59))
	LFS_ERR_IDC_CAPTUREIMPURITY       = (-(IDC_SERVICE_OFFSET + 60))
)

/*=================================================================*/
/* IDC Info Command Structures and variables */
/*=================================================================*/
type LFSIDCSTATUS struct {
	Device                uint
	Media                 uint
	RetainBin             uint
	Security              uint
	UsCards               uint
	ChipPower             uint
	Extra                 []string
	GuidLights            [LFS_IDC_GUIDLIGHTS_SIZE]uint
	ChipModule            uint
	MagReadModule         uint
	MagWriteModule        uint
	FrontImageModule      uint
	BackImageModule       uint
	DevicePosition        uint
	PowerSaveRecoveryTime uint
}

type LFSIDCCAPS struct {
	Class                         uint
	Type                          uint
	Compound                      uint
	ReadTracks                    uint
	WriteTracks                   uint
	ChipProtocols                 uint
	UsCards                       uint
	SecType                       uint
	PowerOnOption                 uint
	PowerOffOption                uint
	FluxSensorProgrammable        uint
	ReadWriteAccessFollowingEject uint
	WriteMode                     uint
	ChipPower                     uint
	Extra                         []string
	DIPMode                       uint
	MemoryChipProtocols           []uint
	GuidLights                    [LFS_IDC_GUIDLIGHTS_SIZE]uint
	EjectPosition                 uint
	PowerSaveControl              uint
}

type LFSIDCFORM struct {
	FormName             string
	FieldSeparatorTrack1 byte
	FieldSeparatorTrack2 byte
	FieldSeparatorTrack3 byte
	Action               uint
	Tracks               string
	Secure               uint
	Track1Fields         []string
	Track2Fields         []string
	Track3Fields         []string
}

//自定义 LFS_INF_IDC_QUERY_IFM_IDENTIFIER Out
type DefOutLFSIDCQUERYIFMIDENTIFIER []LFSIDCIFMIDENTIFIER

type LFSIDCIFMIDENTIFIER struct {
	IFMAuthority  uint
	IFMIdentifier string
}

//自定义 LFS_INF_IDC_FORM_LIST Out
type DefOutLFSIDCFORMLIST []string

//自定义 LFS_INF_IDC_QUERY_FORM In
type DefInLFSIDCQUERYFORM struct {
	FormName string
}

/*=================================================================*/
/* IDC Execute Command Structures */
/*=================================================================*/

//LFS_CMD_IDC_READ_TRACK In
type DefInLFSIDCREADTRACK struct {
	FormName string
}

//LFS_CMD_IDC_READ_TRACK Out
type DefOutLFSIDCREADTRACk struct {
	TrackData [][]string
}

type LFSIDCWRITETRACK struct {
	FormName    string
	TrackData   []string
	WriteMethod uint
}

type LFSIDCRETAINCARD struct {
	UsCount  uint
	Position uint
}

type LFSIDCSETKEY struct {
	UsKeyLen uint
	KeyValue []byte
}

//自定义LFS_CMD_IDC_READ_RAW_DATA In
type DefInLFSIDCREADRAWDATA struct {
	ReadData uint
}

//自定义LFS_CMD_IDC_READ_RAW_DATA Out
type DefOutLFSIDCREADRAWDATA []LFSIDCCARDDATA

type LFSIDCCARDDATA struct {
	DataSource  uint
	Status      uint
	DataLength  uint
	Data        []byte
	WriteMethod uint
}

//自定义LFS_CMD_IDC_WRITE_RAW_DATA In
type DefInLFSIDCWRITERAWDATA []LFSIDCCARDDATA

type LFSIDCCHIPIO struct {
	ChipProtocol   uint
	ChipDataLength uint
	ChipData       []byte
}

//自定义LFS_CMD_IDC_RESET In
type DefInLFSIDCRESET struct {
	ResetIn uint
}

//自定义LFS_CMD_IDC_CHIP_POWER In
type DefInLFSIDCCHIPPOWER struct {
	ChipPower uint
}

type LFSIDCCHIPPOWEROUT struct {
	ChipDataLength uint
	ChipData       []byte
}

type LFSIDCPARSEDATA struct {
	FormName string
	CardData []LFSIDCCARDDATA
}

//自定义LFS_CMD_IDC_PARSE_DATA Out
type DefOutLFSIDCPARSEDATA struct {
	TrackData [][]string
}

type LFSIDCSETGUIDLIGHT struct {
	GuidLight uint
	Command   uint
}

type LFSIDCEJECTCARD struct {
	EjectPosition uint
}

type LFSIDCPOWERSAVECONTROL struct {
	MaxPowerSaveRecoveryTime uint
}

/*=================================================================*/
/* IDC Message Structures */
/*=================================================================*/

type LFSIDCTRACKEVENT struct {
	Status uint
	Track  string
	Data   []byte
}

type LFSIDCCARDACT struct {
	Action   uint
	Position uint
}

type LFSIDCDEVICEPOSITION struct {
	Position uint
}

type LFSIDCPOWERSAVECHANGE struct {
	PowerSaveRecoveryTime uint
}

//自定义LFS_SRVE_IDC_MEDIADETECTED
type DefSrveLFSIDCMEDIADETECTED struct {
	ResetOut uint
}

//自定义LFS_USRE_IDC_RETAINBINTHRESHOLD
type DefUsreLFSIDCRETAINBINTHRESHOLD struct {
	RetainBin uint
}

type LFSIDCRETRACTCARD struct {
	CardNo      string
	RetractTime string
}

type LFSIDCHANDBACKCARD struct {
	CardNo string
}

type LFSIDCCLEARRETRACTCARD struct {
	Num uint
}

type LFSIDCMCRSTATUS struct {
	Device    uint
	Slots     uint
	RetainBin uint
	UseSlots  uint
	Shutter   uint
	Extra     []string
}

type LFSIDCMCRCAPS struct {
	Class       uint
	Type        uint
	BigSlots    uint
	NormalSlots uint
	Extra       []string
}

type LFSIDCSLOTUNITINFO struct {
	Number         uint
	Status         uint
	IsRetained     uint
	CardInfo       string
	ReatinCardInfo string
	RestoreTime    string
	RetainTime     string
	Position       string
	SlotSize       uint
}
type LFSIDCSLOTSINFO struct {
	UsCount uint
	List    []LFSIDCSLOTUNITINFO
}

type LFSIDCRETAINSLOTSTARTIN struct {
	Data string
}

type LFSIDCRETAINSLOTSTARTOUT struct {
	Data string
}

type LFSIDCRETAINSLOT struct {
	CardInfo   string
	RetainTime string
}

type LFSIDCEJECTSLOT struct {
	MCRType uint
	CmdData string
}

type LFSIDCEJECTSLOTOUT struct {
	MCRType uint
	CmdData string
}

type LFSIDCRETAINSLOTOUT struct {
	CardInfo   string
	RetainTime string
	Position   string
	Status     uint
}
