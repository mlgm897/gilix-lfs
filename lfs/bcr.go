package lfs

const (
	LFS_SERVICE_CLASS_BCR = 15
	BCR_SERVICE_OFFSET    = (LFS_SERVICE_CLASS_BCR * 100)
)

/* BCR Info Commands */

const (
	LFS_INF_BCR_STATUS       = (BCR_SERVICE_OFFSET + 1)
	LFS_INF_BCR_CAPABILITIES = (BCR_SERVICE_OFFSET + 2)
)

/* BCR Execute Commands */

const (
	LFS_CMD_BCR_READ               = (BCR_SERVICE_OFFSET + 1)
	LFS_CMD_BCR_RESET              = (BCR_SERVICE_OFFSET + 2)
	LFS_CMD_BCR_SET_GUIDANCE_LIGHT = (BCR_SERVICE_OFFSET + 3)
	LFS_CMD_BCR_POWER_SAVE_CONTROL = (BCR_SERVICE_OFFSET + 4)
)

/* BCR Messages */

const (
	LFS_SRVE_BCR_DEVICEPOSITION    = (BCR_SERVICE_OFFSET + 1)
	LFS_SRVE_BCR_POWER_SAVE_CHANGE = (BCR_SERVICE_OFFSET + 3)
)

/* values of LFSBCRSTATUS.fwDevice */

const (
	LFS_BCR_DEVONLINE       = LFS_STAT_DEVONLINE
	LFS_BCR_DEVOFFLINE      = LFS_STAT_DEVOFFLINE
	LFS_BCR_DEVPOWEROFF     = LFS_STAT_DEVPOWEROFF
	LFS_BCR_DEVNODEVICE     = LFS_STAT_DEVNODEVICE
	LFS_BCR_DEVHWERROR      = LFS_STAT_DEVHWERROR
	LFS_BCR_DEVUSERERROR    = LFS_STAT_DEVUSERERROR
	LFS_BCR_DEVBUSY         = LFS_STAT_DEVBUSY
	LFS_BCR_DEVFRAUDATTEMPT = LFS_STAT_DEVFRAUDATTEMPT
)

/* values of LFSBCRSTATUS.fwBCRScanner */

const (
	LFS_BCR_SCANNERON      = 1
	LFS_BCR_SCANNEROFF     = 2
	LFS_BCR_SCANNERINOP    = 3
	LFS_BCR_SCANNERUNKNOWN = 4
)

/* values of LFSBCRSTATUS.wDevicePosition
LFSBCRDEVICEPOSITION.wPosition */
const (
	LFS_BCR_DEVICEINPOSITION    = 1
	LFS_BCR_DEVICENOTINPOSITION = 2
	LFS_BCR_DEVICEPOSUNKNOWN    = 3
	LFS_BCR_DEVICEPOSNOTSUPP    = 4
)

/* values of LFSBCRCAPS.lpwSymbologies
LFSBCRREADINPUT.lpwSymbologies
LFSBCRREADOUTPUT.wSymbology */
const (
	LFS_BCR_SYM_UNKNOWN          = 0
	LFS_BCR_SYM_EAN128           = 1
	LFS_BCR_SYM_EAN8             = 2
	LFS_BCR_SYM_EAN8_2           = 3
	LFS_BCR_SYM_EAN8_5           = 4
	LFS_BCR_SYM_EAN13            = 5
	LFS_BCR_SYM_EAN13_2          = 6
	LFS_BCR_SYM_EAN13_5          = 7
	LFS_BCR_SYM_JAN13            = 8
	LFS_BCR_SYM_UPCA             = 9
	LFS_BCR_SYM_UPCE0            = 10
	LFS_BCR_SYM_UPCE0_2          = 11
	LFS_BCR_SYM_UPCE0_5          = 12
	LFS_BCR_SYM_UPCE1            = 13
	LFS_BCR_SYM_UPCE1_2          = 14
	LFS_BCR_SYM_UPCE1_5          = 15
	LFS_BCR_SYM_UPCA_2           = 16
	LFS_BCR_SYM_UPCA_5           = 17
	LFS_BCR_SYM_CODABAR          = 18
	LFS_BCR_SYM_ITF              = 19
	LFS_BCR_SYM_11               = 20
	LFS_BCR_SYM_39               = 21
	LFS_BCR_SYM_49               = 22
	LFS_BCR_SYM_93               = 23
	LFS_BCR_SYM_128              = 24
	LFS_BCR_SYM_MSI              = 25
	LFS_BCR_SYM_PLESSEY          = 26
	LFS_BCR_SYM_STD2OF5          = 27
	LFS_BCR_SYM_STD2OF5_IATA     = 28
	LFS_BCR_SYM_PDF_417          = 29
	LFS_BCR_SYM_MICROPDF_417     = 30
	LFS_BCR_SYM_DATAMATRIX       = 31
	LFS_BCR_SYM_MAXICODE         = 32
	LFS_BCR_SYM_CODEONE          = 33
	LFS_BCR_SYM_CHANNELCODE      = 34
	LFS_BCR_SYM_TELEPEN_ORIGINAL = 35
	LFS_BCR_SYM_TELEPEN_AIM      = 36
	LFS_BCR_SYM_RSS              = 37
	LFS_BCR_SYM_RSS_EXPANDED     = 38
	LFS_BCR_SYM_RSS_RESTRICTED   = 39
	LFS_BCR_SYM_COMPOSITE_CODE_A = 40
	LFS_BCR_SYM_COMPOSITE_CODE_B = 41
	LFS_BCR_SYM_COMPOSITE_CODE_C = 42
	LFS_BCR_SYM_POSICODE_A       = 43
	LFS_BCR_SYM_POSICODE_B       = 44
	LFS_BCR_SYM_TRIOPTIC_CODE_39 = 45
	LFS_BCR_SYM_CODABLOCK_F      = 46
	LFS_BCR_SYM_CODE_16K         = 47
	LFS_BCR_SYM_QRCODE           = 48
	LFS_BCR_SYM_AZTEC            = 49
	LFS_BCR_SYM_UKPOST           = 50
	LFS_BCR_SYM_PLANET           = 51
	LFS_BCR_SYM_POSTNET          = 52
	LFS_BCR_SYM_CANADIANPOST     = 53
	LFS_BCR_SYM_NETHERLANDSPOST  = 54
	LFS_BCR_SYM_AUSTRALIANPOST   = 55
	LFS_BCR_SYM_JAPANESEPOST     = 56
	LFS_BCR_SYM_CHINESEPOST      = 57
	LFS_BCR_SYM_KOREANPOST       = 58
)

/* Size and max index of dwGuidLights array */
const (
	LFS_BCR_GUIDLIGHTS_SIZE = 32
	LFS_BCR_GUIDLIGHTS_MAX  = (LFS_BCR_GUIDLIGHTS_SIZE - 1)
)

/* Indices of LFSBCRSTATUS.dwGuidLights [...]
LFSBCRCAPS.dwGuidLights [...]
*/
const (
	LFS_BCR_GUIDANCE_BCR = 0
)

/* Values of LFSBCRSTATUS.dwGuidLights [...]
LFSBCRCAPS.dwGuidLights [...],
LFSBCRSETGUIDLIGHT.wGuidLight */
const (
	LFS_BCR_GUIDANCE_NOT_AVAILABLE = 0x00000000
	LFS_BCR_GUIDANCE_OFF           = 0x00000001
	LFS_BCR_GUIDANCE_ON            = 0x00000002
	LFS_BCR_GUIDANCE_SLOW_FLASH    = 0x00000004
	LFS_BCR_GUIDANCE_MEDIUM_FLASH  = 0x00000008
	LFS_BCR_GUIDANCE_QUICK_FLASH   = 0x00000010
	LFS_BCR_GUIDANCE_CONTINUOUS    = 0x00000080
	LFS_BCR_GUIDANCE_RED           = 0x00000100
	LFS_BCR_GUIDANCE_GREEN         = 0x00000200
	LFS_BCR_GUIDANCE_YELLOW        = 0x00000400
	LFS_BCR_GUIDANCE_BLUE          = 0x00000800
	LFS_BCR_GUIDANCE_CYAN          = 0x00001000
	LFS_BCR_GUIDANCE_MAGENTA       = 0x00002000
	LFS_BCR_GUIDANCE_WHITE         = 0x00004000
)

/* LFSBCR Errors */
const (
	LFS_ERR_BCR_BARCODEINVALID    = (-(BCR_SERVICE_OFFSET + 0))
	LFS_ERR_BCR_INVALID_PORT      = (-(BCR_SERVICE_OFFSET + 1))
	LFS_ERR_BCR_POWERSAVETOOSHORT = (-(BCR_SERVICE_OFFSET + 2))
)

/*=================================================================*/
/* BCR Info Command Structures */
/*=================================================================*/

type LFSBCRSTATUS struct {
	Device                uint
	BcrScanner            uint
	GuidLights            [LFS_BCR_GUIDLIGHTS_SIZE]uint
	Extra                 []string
	DevicePosition        uint
	PowerSaveRecoveryTime uint
}

type LFSBCRCAPS struct {
	Class                uint
	Compound             uint
	CanFilterSymbologies uint
	Symbologies          []uint
	GuidLights           [LFS_BCR_GUIDLIGHTS_SIZE]uint
	Extra                []string
	PowerSaveControl     uint
}

/*=================================================================*/
/* BCR Execute Command Structures */
/*=================================================================*/
type LFSBCRXDATA struct {
	Length uint
	Data   []byte
}

type LFSBCRREADINPUT struct {
	Symbologies []uint
}
type DefOutLFSBCRREADOUTPUT []LFSBCRREADOUTPUT
type LFSBCRREADOUTPUT struct {
	Symbology     uint
	BarcodeData   LFSBCRXDATA
	SymbologyName string
}

type LFSBCRSETGUIDLIGHT struct {
	GuidLight uint
	Command   uint
}

type LFSBCRPOWERSAVECONTROL struct {
	MaxPowerSaveRecoveryTime uint
}

/*=================================================================*/
/* BCR Message Structures */
/*=================================================================*/
type LFSBCRDEVICEPOSITION struct {
	Position uint
}

type LFSBCRPOWERSAVECHANGE struct {
	PowerSaveRecoveryTime uint
}
