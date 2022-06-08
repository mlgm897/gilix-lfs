package lfs

const (
	LFS_SERVICE_CLASS_CAM = 10
	CAM_SERVICE_OFFSET    = (LFS_SERVICE_CLASS_CAM * 100)
)

/* CAM Info Commands */
const (
	LFS_INF_CAM_STATUS       = (CAM_SERVICE_OFFSET + 1)
	LFS_INF_CAM_CAPABILITIES = (CAM_SERVICE_OFFSET + 2)
)

/* CAM Execute Commands */
const (
	LFS_CMD_CAM_TAKE_PICTURE    = (CAM_SERVICE_OFFSET + 1)
	LFS_CMD_CAM_RESET           = (CAM_SERVICE_OFFSET + 2)
	LFS_CMD_CAM_TAKE_PICTURE_EX = (CAM_SERVICE_OFFSET + 3)
	LFS_CMD_CAM_DISPLAY         = (CAM_SERVICE_OFFSET + 4)
	LFS_CMD_CAM_FACE_COMPARISON = (CAM_SERVICE_OFFSET + 5)
	LFS_CMD_CAM_RECORD          = (CAM_SERVICE_OFFSET + 6)
	LFS_CMD_CAM_DISPLAY_HC      = (CAM_SERVICE_OFFSET + 7)
	LFS_CMD_CAM_IMPORT_KEY      = (CAM_SERVICE_OFFSET + 8)
	LFS_CMD_CAM_GENERATE_KCV    = (CAM_SERVICE_OFFSET + 9)
	LFS_CMD_CAM_GET_SIGNATURE   = (CAM_SERVICE_OFFSET + 10)
)

/* CAM Messages */
const (
	LFS_USRE_CAM_MEDIATHRESHOLD = (CAM_SERVICE_OFFSET + 1)
	LFS_EXEE_CAM_INVALIDDATA    = (CAM_SERVICE_OFFSET + 2)
	LFS_USRE_CAM_STARTSIGNING   = (CAM_SERVICE_OFFSET + 3)
)

/* values of LFSCAMSTATUS.fwDevice */
const (
	LFS_CAM_DEVONLINE    = LFS_STAT_DEVONLINE
	LFS_CAM_DEVOFFLINE   = LFS_STAT_DEVOFFLINE
	LFS_CAM_DEVPOWEROFF  = LFS_STAT_DEVPOWEROFF
	LFS_CAM_DEVNODEVICE  = LFS_STAT_DEVNODEVICE
	LFS_CAM_DEVHWERROR   = LFS_STAT_DEVHWERROR
	LFS_CAM_DEVUSERERROR = LFS_STAT_DEVUSERERROR
	LFS_CAM_DEVBUSY      = LFS_STAT_DEVBUSY
)

/* number of cameras supported/length of LFSCAMSTATUS.wCameras field */
const (
	LFS_CAM_CAMERAS_SIZE = 8
	LFS_CAM_CAMERAS_MAX  = (LFS_CAM_CAMERAS_SIZE - 1)
)

/* indices of LFSCAMSTATUS.fwMedia[...]
LFSCAMSTATUS.wCameras [...]
LFSCAMSTATUS.usPictures[...]
LFSCAMCAPS.wCameras [...]
LFSCAMTAKEPICT.wCamera */
const (
	LFS_CAM_ROOM     = 0
	LFS_CAM_PERSON   = 1
	LFS_CAM_EXITSLOT = 2
	LFS_CAM_EXTRA    = 3
)

/* values of LFSCAMSTATUS.fwMedia */
const (
	LFS_CAM_MEDIAOK      = 0
	LFS_CAM_MEDIAHIGH    = 1
	LFS_CAM_MEDIAFULL    = 2
	LFS_CAM_MEDIAUNKNOWN = 3
	LFS_CAM_MEDIANOTSUPP = 4
)

/* values of LFSCAMSTATUS.wCameras */
const (
	LFS_CAM_CAMNOTSUPP = 0
	LFS_CAM_CAMOK      = 1
	LFS_CAM_CAMINOP    = 2
	LFS_CAM_CAMUNKNOWN = 3
)

/* values of LFSCAMCAPS.fwType */
const (
	LFS_CAM_TYPE_CAM = 1
)

/* values of LFSCAMCAPS.wCameras */
const (
	LFS_CAM_NOT_AVAILABLE = 0
	LFS_CAM_AVAILABLE     = 1
)

/* values of LFSCAMCAPS.wCam_Data */
const (
	LFS_CAM_NOTADD  = 0
	LFS_CAM_AUTOADD = 1
	LFS_CAM_MANADD  = 2
)

/* values of LFSCAMCAPS.wChar_Support */
const (
	LFS_CAM_ASCII   = 0x0001
	LFS_CAM_UNICODE = 0x0002
)

/* values of LFSCAMDISP.wAction */
const (
	LFS_CAM_CREATE      = 0
	LFS_CAM_DESTROY     = 1
	LFS_CAM_PAUSE       = 2
	LFS_CAM_RESUME      = 3
	LFS_CAM_ERASE       = 4
	LFS_CAM_PENASERASER = 5
	LFS_CAM_ERASERASPEN = 6
)

/* values of LFSCAMDISP.wTextPosition */
const (
	LFS_CAM_TEXT_POS_DEFAULT      = 0
	LFS_CAM_TEXT_POS_LEFT_TOP     = 1
	LFS_CAM_TEXT_POS_LEFT_BOTTOM  = 2
	LFS_CAM_TEXT_POS_RIGHT_TOP    = 3
	LFS_CAM_TEXT_POS_RIGHT_BOTTOM = 4
	LFS_CAM_TEXT_POS_TOP          = 5
	LFS_CAM_TEXT_POS_BOTTOM       = 6
	LFS_CAM_TEXT_POS_CENTOR       = 7
)

/* values of LFSCAMDISP.wBkShow*/
const (
	LFS_CAM_BK_SHOW_DEFAULT = 0
	LFS_CAM_BK_SHOW_CENTOR  = 1
	LFS_CAM_BK_SHOW_STRETCH = 2
)

/*values  LFSCAMSIGNDATA.wStatus,
LFSCAMGETSIGNATUREPLAINOUT.wStatus,
LFSCAMGETSIGNATUREENCRYPTOUT.wStatus  */
const (
	LFS_CAM_DATAOK         = 1
	LFS_CAM_DATAMISSING    = 2
	LFS_CAM_DATASRCNOTSUPP = 3
	LFS_CAM_DATASRCMISSING = 4
)

/* CAM Errors */
const (
	LFS_ERR_CAM_CAMNOTSUPP       = (-(CAM_SERVICE_OFFSET + 0))
	LFS_ERR_CAM_MEDIAFULL        = (-(CAM_SERVICE_OFFSET + 1))
	LFS_ERR_CAM_CAMINOP          = (-(CAM_SERVICE_OFFSET + 2))
	LFS_ERR_CHARSETNOTSUPP       = (-(CAM_SERVICE_OFFSET + 3))
	LFS_ERR_CAM_FILEIOERROR      = (-(CAM_SERVICE_OFFSET + 4))
	LFS_ERR_CAM_KEYNOTFOUND      = (-(CAM_SERVICE_OFFSET + 5))
	LFS_ERR_CAM_ACCESSDENIED     = (-(CAM_SERVICE_OFFSET + 6))
	LFS_ERR_CAM_KEYINVALID       = (-(CAM_SERVICE_OFFSET + 7))
	LFS_ERR_CAM_NOKEYRAM         = (-(CAM_SERVICE_OFFSET + 8))
	LFS_ERR_CAM_INVALIDKEYLENGTH = (-(CAM_SERVICE_OFFSET + 9))
	LFS_ERR_CAM_INVALID_HWND     = (-(CAM_SERVICE_OFFSET + 10))
	LFS_ERR_CAM_NOTSIGN_ERROR    = (-(CAM_SERVICE_OFFSET + 11))
)

/* values of LFSCAMTAKEPICTSIZE.wScan_Size */
const (
	LFS_CAM_SIZEALL                        = 0
	LFS_CAM_SIZEAUTO                       = 1
	LFS_CAM_SIZEA4                         = 2
	LFS_CAM_SIZEA5                         = 3
	LFS_CAM_SIZEA6                         = 4
	LFS_CAM_SIZEA7                         = 5
	LFS_CAM_SIZECALLINGCARD                = 6
	LFS_CAM_SIZEID                         = 7
	LFS_CAM_SIZECUSTOM                     = 8
	LFS_CAM_SIZEAJUSTEDTOWINDOW            = 9
	LFS_CAM_SIZEAJUSTEDTOFRAME             = 10
	LFS_CAM_SIZEAJUSTEDTOFRAMEANDCHECKFACE = 11
)

/* values of LFSCAMIMPORT.wAlgorithm,LFSCAMGETSIGNATUREENCRYPT.wAlgorithm */
const (
	LFS_CAM_CRYPTNONE    = 0
	LFS_CAM_CRYPTDESECB  = 1
	LFS_CAM_CRYPTDESCBC  = 2
	LFS_CAM_CRYPT3DESECB = 3
	LFS_CAM_CRYPT3DESCBC = 4
	LFS_CAM_CRYPTSM2     = 5
	LFS_CAM_CRYPTSM3     = 6
	LFS_CAM_CRYPTSM4ECB  = 7
	LFS_CAM_CRYPTSM4CBC  = 8
)

/* values of LFSCAMIMPORT.wClear */
const (
	LFS_CAM_CLEARNONE = 0
	LFS_CAM_CLEARKEY  = 1
	LFS_CAM_CLEARALL  = 2
)

/* values of LFSCAMIMPORT.wKey_Check_Mode */
const (
	LFS_CAM_KCVNONE = 0
	LFS_CAM_KCVSELF = 1
	LFS_CAM_KCVZERO = 2
)

/*=================================================================*/
/* CAM Info Command Structures and variables */
/*=================================================================*/
type LFSCAMSTATUS struct {
	Device   uint
	Media    [LFS_CAM_CAMERAS_SIZE]uint
	Cameras  [LFS_CAM_CAMERAS_SIZE]uint
	Pictures [LFS_CAM_CAMERAS_SIZE]uint
	Extra    []string
}

type LFSCAMCAPS struct {
	Class         uint
	Type          uint
	Cameras       [LFS_CAM_CAMERAS_SIZE]uint
	MaxPictures   uint
	CamData       uint
	MaxDataLength uint
	CharSupport   uint
	Extra         []string
	PictureFile   uint
}

/*=================================================================*/
/* CAM Execute Command Structures */
/*=================================================================*/
type LFSCAMTAKEPICT struct {
	Camera         uint
	CamData        string
	UNICODECamData []byte
}

type LFSXCAMDATA struct {
	Length uint
	Data   []byte
}

type LFSCAMDISP struct {
	Camera          uint
	Action          uint
	X               uint
	Y               uint
	Width           uint
	Height          uint
	Hwnd            string
	Hpixel          uint
	Vpixel          uint
	TextData        string
	UNICODETextData []byte
	FontWidth       uint
	FontHeight      uint
	Face            string
	TextPosition    uint
	BackgroundFile  string
	BkShow          uint
	Key             string
	EncKey          string
	KeyEncKey       LFSXCAMDATA
	Algorithm       uint
	Extra           []string
}

type LFSCAMTAKEPICTEX struct {
	Camera         uint
	CamData        string
	UNICODECamData []byte
	PictureFile    string
	ScanSize       uint
	Extra          []string
}

type LFSCAMFACECOMPARISON struct {
	Camera       uint
	PictureFileA string
	PictureFileB string
	Extra        []string
}

type LFSCAMFACECOMPARISONOUT struct {
	Score uint
	Extra []string
}

type LFSCAMDISPHC struct {
	Camera      uint
	Action      uint
	Width       uint
	Height      uint
	X           uint
	Y           uint
	Hwnd        string
	Hpixel      uint
	Vpixel      uint
	RotateAngle uint
	ColorType   uint
	CropType    uint
	CropSize    uint
	TexData     string
	Extra       []string
}

type LFSCAMRECORD struct {
	Camera      uint
	Action      uint
	Hpixel      uint
	Vpixel      uint
	Fps         uint
	AudioOption uint
	RecordFile  string
	Extra       []string
}

type LFSCAMIMPORT struct {
	Key           string
	EncKey        string
	Value         LFSXCAMDATA
	ControlVector LFSXCAMDATA
	Clear         uint
	Algorithm     uint
	KeyCheckMode  uint
	KeyCheckValue LFSXCAMDATA
	Extra         []string
}

type LFSCAMGETSIGNATURE struct {
	Camera      uint
	PictureFile string
	TrackFile   string
	Extra       []string
}

type LFSCAMGETSIGNATUREOUT struct {
	Status     uint
	DataLength uint
	Data       []byte
	Extra      []string
}

//自定义
//LFS_USRE_CAM_MEDIATHRESHOLD
type DefUsreLFSCAMMEDIATHRESHOLD struct {
	MediaThreshold uint
}
