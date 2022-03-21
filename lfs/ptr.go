package lfs

/*LFSPTRCAPS.wClass 取值 */
const (
	LFS_SERVICE_CLASS_PTR = 1
	PTR_SERVICE_OFFSET    = (LFS_SERVICE_CLASS_PTR * 100)
)

/*信息指令 */
const (
	LFS_INF_PTR_STATUS           = (PTR_SERVICE_OFFSET + 1)
	LFS_INF_PTR_CAPABILITIES     = (PTR_SERVICE_OFFSET + 2)
	LFS_INF_PTR_FORM_LIST        = (PTR_SERVICE_OFFSET + 3)
	LFS_INF_PTR_MEDIA_LIST       = (PTR_SERVICE_OFFSET + 4)
	LFS_INF_PTR_QUERY_FORM       = (PTR_SERVICE_OFFSET + 5)
	LFS_INF_PTR_QUERY_MEDIA      = (PTR_SERVICE_OFFSET + 6)
	LFS_INF_PTR_QUERY_FIELD      = (PTR_SERVICE_OFFSET + 7)
	LFS_INF_PTR_CODELINE_MAPPTRG = (PTR_SERVICE_OFFSET + 8)
)

/*执行指令 */
const (
	LFS_CMD_PTR_CONTROL_MEDIA      = (PTR_SERVICE_OFFSET + 1)
	LFS_CMD_PTR_PRINT_FORM         = (PTR_SERVICE_OFFSET + 2)
	LFS_CMD_PTR_READ_FORM          = (PTR_SERVICE_OFFSET + 3)
	LFS_CMD_PTR_RAW_DATA           = (PTR_SERVICE_OFFSET + 4)
	LFS_CMD_PTR_MEDIA_EXTENTS      = (PTR_SERVICE_OFFSET + 5)
	LFS_CMD_PTR_RESET_COUNT        = (PTR_SERVICE_OFFSET + 6)
	LFS_CMD_PTR_READ_IMAGE         = (PTR_SERVICE_OFFSET + 7)
	LFS_CMD_PTR_RESET              = (PTR_SERVICE_OFFSET + 8)
	LFS_CMD_PTR_RETRACT_MEDIA      = (PTR_SERVICE_OFFSET + 9)
	LFS_CMD_PTR_DISPENSE_PAPER     = (PTR_SERVICE_OFFSET + 10)
	LFS_CMD_PTR_SET_GUIDANCE_LIGHT = (PTR_SERVICE_OFFSET + 11)
	LFS_CMD_PTR_PRINT_RAW_FILE     = (PTR_SERVICE_OFFSET + 12)
	LFS_CMD_PTR_LOAD_DEFINITION    = (PTR_SERVICE_OFFSET + 13)
	LFS_CMD_PTR_SUPPLY_REPLENISH   = (PTR_SERVICE_OFFSET + 14)
	LFS_CMD_PTR_POWER_SAVE_CONTROL = (PTR_SERVICE_OFFSET + 15)
	LFS_CMD_PTR_READ_IMAGE_EX      = (PTR_SERVICE_OFFSET + 16)
	LFS_CMD_PTR_IMPORT_KEY         = (PTR_SERVICE_OFFSET + 17)
)

/*事件 */
const (
	LFS_EXEE_PTR_NOMEDIA             = (PTR_SERVICE_OFFSET + 1)
	LFS_EXEE_PTR_MEDIAINSERTED       = (PTR_SERVICE_OFFSET + 2)
	LFS_EXEE_PTR_FIELDERROR          = (PTR_SERVICE_OFFSET + 3)
	LFS_EXEE_PTR_FIELDWARNING        = (PTR_SERVICE_OFFSET + 4)
	LFS_USRE_PTR_RETRACTBINTHRESHOLD = (PTR_SERVICE_OFFSET + 5)
	LFS_SRVE_PTR_MEDIATAKEN          = (PTR_SERVICE_OFFSET + 6)
	LFS_USRE_PTR_PAPERTHRESHOLD      = (PTR_SERVICE_OFFSET + 7)
	LFS_USRE_PTR_TONERTHRESHOLD      = (PTR_SERVICE_OFFSET + 8)
	LFS_SRVE_PTR_MEDIAINSERTED       = (PTR_SERVICE_OFFSET + 9)
	LFS_USRE_PTR_LAMPTHRESHOLD       = (PTR_SERVICE_OFFSET + 10)
	LFS_USRE_PTR_INKTHRESHOLD        = (PTR_SERVICE_OFFSET + 11)
	LFS_SRVE_PTR_MEDIADETECTED       = (PTR_SERVICE_OFFSET + 12)
	LFS_SRVE_PTR_RETRACTBINSTATUS    = (PTR_SERVICE_OFFSET + 13)
	LFS_EXEE_PTR_MEDIAPRESENTED      = (PTR_SERVICE_OFFSET + 14)
	LFS_SRVE_PTR_DEFINITIONLOADED    = (PTR_SERVICE_OFFSET + 15)
	LFS_EXEE_PTR_MEDIAREJECTED       = (PTR_SERVICE_OFFSET + 16)
	LFS_SRVE_PTR_MEDIAPRESENTED      = (PTR_SERVICE_OFFSET + 17)
	LFS_SRVE_PTR_MEDIAAUTORETRACTED  = (PTR_SERVICE_OFFSET + 18)
	LFS_SRVE_PTR_DEVICEPOSITION      = (PTR_SERVICE_OFFSET + 19)
	LFS_SRVE_PTR_POWER_SAVE_CHANGE   = (PTR_SERVICE_OFFSET + 20)
)

/*LFSPTRSTATUS.wDevice取值 */
const (
	LFS_PTR_DEVONLINE       = LFS_STAT_DEVONLINE
	LFS_PTR_DEVOFFLINE      = LFS_STAT_DEVOFFLINE
	LFS_PTR_DEVPOWEROFF     = LFS_STAT_DEVPOWEROFF
	LFS_PTR_DEVNODEVICE     = LFS_STAT_DEVNODEVICE
	LFS_PTR_DEVHWERROR      = LFS_STAT_DEVHWERROR
	LFS_PTR_DEVUSERERROR    = LFS_STAT_DEVUSERERROR
	LFS_PTR_DEVBUSY         = LFS_STAT_DEVBUSY
	LFS_PTR_DEVFRAUDATTEMPT = LFS_STAT_DEVFRAUDATTEMPT
)

/*LFSPTRSTATUS.wMedia and
LFSPTRMEDIADETECTED.wPosition
取值 */
const (
	LFS_PTR_MEDIAPRESENT    = 0
	LFS_PTR_MEDIANOTPRESENT = 1
	LFS_PTR_MEDIAJAMMED     = 2
	LFS_PTR_MEDIANOTSUPP    = 3
	LFS_PTR_MEDIAUNKNOWN    = 4
	LFS_PTR_MEDIAENTERING   = 5
	LFS_PTR_MEDIARETRACTED  = 6
)

/*wPaper数组的大小和下标上限  */
const (
	LFS_PTR_SUPPLYSIZE = 16
	LFS_PTR_SUPPLYMAX  = (LFS_PTR_SUPPLYSIZE - 1)
)

/*LFSPTRSTATUS.wPaper[...] 数组的下标 */
const (
	LFS_PTR_SUPPLYUPPER    = 0
	LFS_PTR_SUPPLYLOWER    = 1
	LFS_PTR_SUPPLYEXTERNAL = 2
	LFS_PTR_SUPPLYAUX      = 3
	LFS_PTR_SUPPLYAUX2     = 4
	LFS_PTR_SUPPLYPARK     = 5
)

/*LFSPTRSTATUS.wPaper  and
LFSPTRPAPERTHRESHOLD.wPaper_Source
取值 */
const (
	LFS_PTR_PAPERFULL    = 0
	LFS_PTR_PAPERLOW     = 1
	LFS_PTR_PAPEROUT     = 2
	LFS_PTR_PAPERNOTSUPP = 3
	LFS_PTR_PAPERUNKNOWN = 4
	LFS_PTR_PAPERJAMMED  = 5
)

/*LFSPTRSTATUS.wToner取值 */
const (
	LFS_PTR_TONERFULL    = 0
	LFS_PTR_TONERLOW     = 1
	LFS_PTR_TONEROUT     = 2
	LFS_PTR_TONERNOTSUPP = 3
	LFS_PTR_TONERUNKNOWN = 4
)

/*LFSPTRSTATUS.wInk取值 */
const (
	LFS_PTR_INKFULL    = 0
	LFS_PTR_INKLOW     = 1
	LFS_PTR_INKOUT     = 2
	LFS_PTR_INKNOTSUPP = 3
	LFS_PTR_INKUNKNOWN = 4
)

/*LFSPTRSTATUS.wLamp取值 */
const (
	LFS_PTR_LAMPOK      = 0
	LFS_PTR_LAMPFADING  = 1
	LFS_PTR_LAMPINOP    = 2
	LFS_PTR_LAMPNOTSUPP = 3
	LFS_PTR_LAMPUNKNOWN = 4
)

/*LFSPTRSTATUS.lppRetract_Bins and
LFSPTRBINTHRESHOLD.wRetract_Bin
		取值 */
const (
	LFS_PTR_RETRACTBINOK   = 0
	LFS_PTR_RETRACTBINFULL = 1
	LFS_PTR_RETRACTNOTSUPP = 2
	LFS_PTR_RETRACTUNKNOWN = 3
	LFS_PTR_RETRACTBINHIGH = 4

	/* additional values of LFSPTRSTATUS.lppRetract_Bins*/
	LFS_PTR_RETRACTBINMISSING = 5
)

/* Size and max index of dwGuid_Lights array */
const (
	LFS_PTR_GUIDLIGHTS_SIZE = 32
	LFS_PTR_GUIDLIGHTS_MAX  = (LFS_PTR_GUIDLIGHTS_SIZE - 1)
)

/* Indices of LFSPTRSTATUS.dwGuid_Lights [...]
LFSPTRCAPS.dwGuid_Lights [...] */
const (
	LFS_PTR_GUIDANCE_PRINTER = 0
)

/* Values of LFSPTRSTATUS.dwGuid_Lights [...]
LFSPTRCAPS.dwGuid_Lights [...] */
const (
	LFS_PTR_GUIDANCE_NOT_AVAILABLE = 0x00000000
	LFS_PTR_GUIDANCE_OFF           = 0x00000001
	LFS_PTR_GUIDANCE_SLOW_FLASH    = 0x00000004
	LFS_PTR_GUIDANCE_MEDIUM_FLASH  = 0x00000008
	LFS_PTR_GUIDANCE_QUICK_FLASH   = 0x00000010
	LFS_PTR_GUIDANCE_CONTINUOUS    = 0x00000080
	LFS_PTR_GUIDANCE_RED           = 0x00000100
	LFS_PTR_GUIDANCE_GREEN         = 0x00000200
	LFS_PTR_GUIDANCE_YELLOW        = 0x00000400
	LFS_PTR_GUIDANCE_BLUE          = 0x00000800
	LFS_PTR_GUIDANCE_CYAN          = 0x00001000
	LFS_PTR_GUIDANCE_MAGENTA       = 0x00002000
	LFS_PTR_GUIDANCE_WHITE         = 0x00004000
)

/* values of LFSPTRSTATUS.wDevice_Position
LFSPTRDEVICEPOSITION.wPosition */
const (
	LFS_PTR_DEVICEINPOSITION    = 0
	LFS_PTR_DEVICENOTINPOSITION = 1
	LFS_PTR_DEVICEPOSUNKNOWN    = 2
	LFS_PTR_DEVICEPOSNOTSUPP    = 3
)

/*LFSPTRCAPS.wType取值 */
const (
	LFS_PTR_TYPERECEIPT  = 0x0001
	LFS_PTR_TYPEPASSBOOK = 0x0002
	LFS_PTR_TYPEJOURNAL  = 0x0004
	LFS_PTR_TYPEDOCUMENT = 0x0008
	LFS_PTR_TYPESCANNER  = 0x0010
)

/*LFSPTRCAPS.wDev_Resolution, LFSPTRPRINTFORM.wResolution取值 */
const (
	LFS_PTR_RESLOW      = 0x0001
	LFS_PTR_RESMED      = 0x0002
	LFS_PTR_RESHIGH     = 0x0004
	LFS_PTR_RESVERYHIGH = 0x0008
)

/*LFSPTRCAPS.wRead_form 取值 */
const (
	LFS_PTR_READOCR       = 0x0001
	LFS_PTR_READMICR      = 0x0002
	LFS_PTR_READMSF       = 0x0004
	LFS_PTR_READBARCODE   = 0x0008
	LFS_PTR_READPAGEMARK  = 0x0010
	LFS_PTR_READIMAGE     = 0x0020
	LFS_PTR_READEMPTYLINE = 0x0040
)

/*LFSPTRCAPS.wWrite_form 取值 */
const (
	LFS_PTR_WRITETEXT     = 0x0001
	LFS_PTR_WRITEGRAPHICS = 0x0002
	LFS_PTR_WRITEOCR      = 0x0004
	LFS_PTR_WRITEMICR     = 0x0008
	LFS_PTR_WRITEMSF      = 0x0010
	LFS_PTR_WRITEBARCODE  = 0x0020
	LFS_PTR_WRITESTAMP    = 0x0040
)

/*LFSPTRCAPS.wDev_Extents 取值 */
const (
	LFS_PTR_EXTHORIZONTAL = 0x0001
	LFS_PTR_EXTVERTICAL   = 0x0002
)

/*LFSPTRCAPS.wDev_Control, lpdwMedia_Control 取值 */
const (
	LFS_PTR_CTRLEJECT            = 0x0001
	LFS_PTR_CTRLPERFORATE        = 0x0002
	LFS_PTR_CTRLCUT              = 0x0004
	LFS_PTR_CTRLSKIP             = 0x0008
	LFS_PTR_CTRLFLUSH            = 0x0010
	LFS_PTR_CTRLRETRACT          = 0x0020
	LFS_PTR_CTRLSTACK            = 0x0040
	LFS_PTR_CTRLPARTIALCUT       = 0x0080
	LFS_PTR_CTRLALARM            = 0x0100
	LFS_PTR_CTRLATPFORWARD       = 0x0200
	LFS_PTR_CTRLATPBACKWARD      = 0x0400
	LFS_PTR_CTRLTURNMEDIA        = 0x0800
	LFS_PTR_CTRLSTAMP            = 0x1000
	LFS_PTR_CTRLPARK             = 0x2000
	LFS_PTR_CTRLEXPEL            = 0x4000
	LFS_PTR_CTRLEJECTTOTRANSPORT = 0x8000
)

/*LFSPTRCAPS.wPaper_Sources,
LFSFRMMEDIA.wPaper_Sources,
LFSPTRPRINTFORM.wPaper_Sources and
LFSPTRPAPERTHRESHOLD.wPaper_Sources
取值 */
const (
	LFS_PTR_PAPERANY      = 0x0001
	LFS_PTR_PAPERUPPER    = 0x0002
	LFS_PTR_PAPERLOWER    = 0x0004
	LFS_PTR_PAPEREXTERNAL = 0x0008
	LFS_PTR_PAPERAUX      = 0x0010
	LFS_PTR_PAPERAUX2     = 0x0020
	LFS_PTR_PAPERPARK     = 0x0040
)

/*LFSPTRCAPS.wImage_Type,
LFSPTRIMAGEREQUEST.wFront_Image_Type and
LFSPTRIMAGEREQUEST.wBack_Image_Type
取值 */
const (
	LFS_PTR_IMAGETIF = 0x0001
	LFS_PTR_IMAGEWMF = 0x0002
	LFS_PTR_IMAGEBMP = 0x0004
	LFS_PTR_IMAGEJPG = 0x0008
)

/*LFSPTRCAPS.wFront_Image_Color_Format,
LFSPTRCAPS.wBack_Image_Color_Format,
LFSPTRIMAGEREQUEST.wFront_Image_Color_Format and
LFSPTRIMAGEREQUEST.wBack_Image_Color_Format
取值 */
const (
	LFS_PTR_IMAGECOLORBINARY    = 0x0001
	LFS_PTR_IMAGECOLORGRAYSCALE = 0x0002
	LFS_PTR_IMAGECOLORFULL      = 0x0004
)

/*LFSPTRCAPS.wCodeline_Format and
LFSPTRIMAGEREQUEST.wCodeline_Format
取值 */
const (
	LFS_PTR_CODELINECMC7 = 0x0001
	LFS_PTR_CODELINEE13B = 0x0002
	LFS_PTR_CODELINEOCR  = 0x0004
)

/*LFSPTRCAPS.wImage_Source,
LFSPTRIMAGEREQUEST.wImage_Source and
LFSPTRIMAGE.wImage_Source
取值 */
const (
	LFS_PTR_IMAGEFRONT = 0x0001
	LFS_PTR_IMAGEBACK  = 0x0002
	LFS_PTR_CODELINE   = 0x0004

	LFS_PTR_IMAGETYPE0 = 0x8000
	LFS_PTR_IMAGETYPE1 = 0x4000
	LFS_PTR_IMAGETYPE2 = 0x2000
)

/*LFSPTRCAPS.wChar_Support, LFSFRMHEADER.wChar_Support 取值 */
const (
	LFS_PTR_UTF8    = 0x0001
	LFS_PTR_UNICODE = 0x0002
)

/*LFSFRMHEADER.wBase, LFSFRMMEDIA.wBase, LFSPTRMEDIAUNIT.wBase取值 */
const (
	LFS_FRM_INCH      = 0
	LFS_FRM_MM        = 1
	LFS_FRM_ROWCOLUMN = 2
)

/*LFSFRMHEADER.wAlignment取值 */
const (
	LFS_FRM_TOPLEFT     = 0
	LFS_FRM_TOPRIGHT    = 1
	LFS_FRM_BOTTOMLEFT  = 2
	LFS_FRM_BOTTOMRIGHT = 3
)

/*LFSFRMHEADER.wOrientation取值 */
const (
	LFS_FRM_PORTRAIT  = 0
	LFS_FRM_LANDSCAPE = 1
)

/*LFSFRMMEDIA.wMedia_Type 取值 */
const (
	LFS_FRM_MEDIAGENERIC   = 0
	LFS_FRM_MEDIAPASSBOOK  = 1
	LFS_FRM_MEDIAMULTIPART = 2
)

/*LFSFRMMEDIA.wFold_Type 取值 */
const (
	LFS_FRM_FOLDNONE       = 0
	LFS_FRM_FOLDHORIZONTAL = 1
	LFS_FRM_FOLDVERTICAL   = 2
)

/*LFSFRMFIELD.wType取值 */
const (
	LFS_FRM_FIELDTEXT     = 0
	LFS_FRM_FIELDMICR     = 1
	LFS_FRM_FIELDOCR      = 2
	LFS_FRM_FIELDMSF      = 3
	LFS_FRM_FIELDBARCODE  = 4
	LFS_FRM_FIELDGRAPHIC  = 5
	LFS_FRM_FIELDPAGEMARK = 6
)

/*LFSFRMFIELD.wField_Class取值 */
const (
	LFS_FRM_CLASSSTATIC   = 0
	LFS_FRM_CLASSOPTIONAL = 1
	LFS_FRM_CLASSREQUIRED = 2
)

/*LFSFRMFIELD.wAccess取值 */
const (
	LFS_FRM_ACCESSREAD  = 0x0001
	LFS_FRM_ACCESSWRITE = 0x0002
)

/*LFSFRMFIELD.wOverflow取值 */
const (
	LFS_FRM_OVFTERMINATE = 0
	LFS_FRM_OVFTRUNCATE  = 1
	LFS_FRM_OVFBESTFIT   = 2
	LFS_FRM_OVFOVERWRITE = 3
	LFS_FRM_OVFWORDWRAP  = 4
)

/*LFSPTRFIELDFAIL.wFailure取值 */
const (
	LFS_PTR_FIELDREQUIRED         = 0
	LFS_PTR_FIELDSTATICOVWR       = 1
	LFS_PTR_FIELDOVERFLOW         = 2
	LFS_PTR_FIELDNOTFOUND         = 3
	LFS_PTR_FIELDNOTREAD          = 4
	LFS_PTR_FIELDNOTWRITE         = 5
	LFS_PTR_FIELDHWERROR          = 6
	LFS_PTR_FIELDTYPENOTSUPPORTED = 7
	LFS_PTR_FIELDGRAPHIC          = 8
	LFS_PTR_CHARSETFORM           = 9
)

/*LFSPTRPRINTFORM.wAlignment取值 */
const (
	LFS_PTR_ALNUSEFORMDEFN = 0
	LFS_PTR_ALNTOPLEFT     = 1
	LFS_PTR_ALNTOPRIGHT    = 2
	LFS_PTR_ALNBOTTOMLEFT  = 3
	LFS_PTR_ALNBOTTOMRIGHT = 4
)

/*LFSPTRPRINTFORM.wOffset_X and LFSPTRPRINTFORM.wOffset_Y 取值*/
const (
	LFS_PTR_OFFSETUSEFORMDEFN = 0xffff
)

/*LFSPTRRAWDATA.wInput_Data 取值 */
const (
	LFS_PTR_NOINPUTDATA = 0
	LFS_PTR_INPUTDATA   = 1
)

/*LFSPTRIMAGE.wStatus取值 */
const (
	LFS_PTR_DATAOK         = 0
	LFS_PTR_DATASRCNOTSUPP = 1
	LFS_PTR_DATASRCMISSING = 2
)

/*LFSPTRBINSTATUS.wRetract_Bin 取值*/
const (
	LFS_PTR_RETRACTBININSERTED = 1
	LFS_PTR_RETRACTBINREMOVED  = 2
)

/*LFSPTRDEFINITIONLOADED.dwDefinition_Type取值*/
const (
	LFS_PTR_FORMLOADED  = 0x00000001
	LFS_PTR_MEDIALOADED = 0x00000002
)

/*LFSPTRSUPPLYREPLEN.fwSupply_Replen取值*/
const (
	LFS_PTR_REPLEN_PAPERUPPER = 0x0001
	LFS_PTR_REPLEN_PAPERLOWER = 0x0002
	LFS_PTR_REPLEN_PAPERAUX   = 0x0004
	LFS_PTR_REPLEN_PAPERAUX2  = 0x0008
	LFS_PTR_REPLEN_TONER      = 0x0010
	LFS_PTR_REPLEN_INK        = 0x0020
	LFS_PTR_REPLEN_LAMP       = 0x0040
)

/*LFSPTRMEDIAREJECTED.wMedia_Rejected取值*/
const (
	LFS_PTR_REJECT_SHORT       = 0
	LFS_PTR_REJECT_LONG        = 1
	LFS_PTR_REJECT_MULTIPLE    = 2
	LFS_PTR_REJECT_ALIGN       = 3
	LFS_PTR_REJECT_MOVETOALIGN = 4
	LFS_PTR_REJECT_SHUTTER     = 5
	LFS_PTR_REJECT_ESCROW      = 6
	LFS_PTR_REJECT_THICK       = 7
	LFS_PTR_REJECT_OTHER       = 8
)

/*LFSPTRMEDIARETRACTED.wRetract_Result取值*/
const (
	LFS_PTR_AUTO_RETRACT_OK          = 0
	LFS_PTR_AUTO_RETRACT_MEDIAJAMMED = 1
)

/*LFSPTRIMAGEREQUESTEX.wEncrypt_Mode取值*/
const (
	LFS_PTR_NOT_ENCRYPT = 0
	LFS_PTR_MODESM4     = 1
)

/*LFSPTRIMPORTKEYEX.dwUse取值*/
const (
	LFS_PTR_USECRYPT           = 0x0001
	LFS_PTR_USEFUNCTION        = 0x0002
	LFS_PTR_USEMACING          = 0x0004
	LFS_PTR_USEKEYENCKEY       = 0x0020
	LFS_PTR_USENODUPLICATE     = 0x0040
	LFS_PTR_USESVENCKEY        = 0x0080
	LFS_PTR_USECONSTRUCT       = 0x0100
	LFS_PTR_USESECURECONSTRUCT = 0x0200
	LFS_PTR_USEANSTR31MASTER   = 0x0400
)

/* LFSPTRIMPORTKEYEX.wKey_Check_Mode */
const (
	LFS_PTR_KCVNONE = 0x0000
	LFS_PTR_KCVSELF = 0x0001
	LFS_PTR_KCVZERO = 0x0002
)

/*错误码 */
const (
	LFS_ERR_PTR_FORMNOTFOUND       = (-(PTR_SERVICE_OFFSET + 0))
	LFS_ERR_PTR_FIELDNOTFOUND      = (-(PTR_SERVICE_OFFSET + 1))
	LFS_ERR_PTR_NOMEDIAPRESENT     = (-(PTR_SERVICE_OFFSET + 2))
	LFS_ERR_PTR_READNOTSUPPORTED   = (-(PTR_SERVICE_OFFSET + 3))
	LFS_ERR_PTR_FLUSHFAIL          = (-(PTR_SERVICE_OFFSET + 4))
	LFS_ERR_PTR_MEDIAOVERFLOW      = (-(PTR_SERVICE_OFFSET + 5))
	LFS_ERR_PTR_FIELDSPECFAILURE   = (-(PTR_SERVICE_OFFSET + 6))
	LFS_ERR_PTR_FIELDERROR         = (-(PTR_SERVICE_OFFSET + 7))
	LFS_ERR_PTR_MEDIANOTFOUND      = (-(PTR_SERVICE_OFFSET + 8))
	LFS_ERR_PTR_EXTENTNOTSUPPORTED = (-(PTR_SERVICE_OFFSET + 9))
	LFS_ERR_PTR_MEDIAINVALID       = (-(PTR_SERVICE_OFFSET + 10))
	LFS_ERR_PTR_FORMINVALID        = (-(PTR_SERVICE_OFFSET + 11))
	LFS_ERR_PTR_FIELDINVALID       = (-(PTR_SERVICE_OFFSET + 12))
	LFS_ERR_PTR_MEDIASKEWED        = (-(PTR_SERVICE_OFFSET + 13))
	LFS_ERR_PTR_RETRACTBINFULL     = (-(PTR_SERVICE_OFFSET + 14))
	LFS_ERR_PTR_STACKERFULL        = (-(PTR_SERVICE_OFFSET + 15))
	LFS_ERR_PTR_PAGETURNFAIL       = (-(PTR_SERVICE_OFFSET + 16))
	LFS_ERR_PTR_MEDIATURNFAIL      = (-(PTR_SERVICE_OFFSET + 17))
	LFS_ERR_PTR_SHUTTERFAIL        = (-(PTR_SERVICE_OFFSET + 18))
	LFS_ERR_PTR_MEDIAJAMMED        = (-(PTR_SERVICE_OFFSET + 19))
	LFS_ERR_PTR_FILE_IO_ERROR      = (-(PTR_SERVICE_OFFSET + 20))
	LFS_ERR_PTR_CHARSETDATA        = (-(PTR_SERVICE_OFFSET + 21))
	LFS_ERR_PTR_PAPERJAMMED        = (-(PTR_SERVICE_OFFSET + 22))
	LFS_ERR_PTR_PAPEROUT           = (-(PTR_SERVICE_OFFSET + 23))
	LFS_ERR_PTR_INKOUT             = (-(PTR_SERVICE_OFFSET + 24))
	LFS_ERR_PTR_TONEROUT           = (-(PTR_SERVICE_OFFSET + 25))
	LFS_ERR_PTR_LAMPINOP           = (-(PTR_SERVICE_OFFSET + 26))
	LFS_ERR_PTR_SOURCEINVALID      = (-(PTR_SERVICE_OFFSET + 27))
	LFS_ERR_PTR_SEQUENCEINVALID    = (-(PTR_SERVICE_OFFSET + 28))
	LFS_ERR_PTR_MEDIASIZE          = (-(PTR_SERVICE_OFFSET + 29))
	LFS_ERR_PTR_INVALID_PORT       = (-(PTR_SERVICE_OFFSET + 30))
	LFS_ERR_PTR_MEDIARETAINED      = (-(PTR_SERVICE_OFFSET + 31))
	LFS_ERR_PTR_BLACKMARK          = (-(PTR_SERVICE_OFFSET + 32))
	LFS_ERR_PTR_DEFINITIONEXISTS   = (-(PTR_SERVICE_OFFSET + 33))
	LFS_ERR_PTR_MEDIAREJECTED      = (-(PTR_SERVICE_OFFSET + 34))
	LFS_ERR_PTR_MEDIARETRACTED     = (-(PTR_SERVICE_OFFSET + 35))
	LFS_ERR_PTR_MSFERROR           = (-(PTR_SERVICE_OFFSET + 36))
	LFS_ERR_PTR_NOMSF              = (-(PTR_SERVICE_OFFSET + 37))
	LFS_ERR_PTR_FILENOTFOUND       = (-(PTR_SERVICE_OFFSET + 38))
)

/*=================================================================*/
/*信息指令结构体 */
/*=================================================================*/

type LFSPTRRETRACTBINS struct {
	RetractBin   uint
	RetractCount uint
}

type LFSPTRSTATUS struct {
	Device                uint
	Media                 uint
	Paper                 [LFS_PTR_SUPPLYSIZE]uint
	Toner                 uint
	Ink                   uint
	Lamp                  uint
	RetractBins           []LFSPTRRETRACTBINS
	MediaOnStacker        uint
	Extra                 []byte
	GuidLights            [LFS_PTR_GUIDLIGHTS_SIZE]uint
	DevicePosition        uint
	PowerSaveRecoveryTime uint
}

type LFSPTRCAPS struct {
	Class                 uint
	Type                  uint
	Compound              int
	DevResolution         uint
	Readform              uint
	Writeform             uint
	DevExtents            uint
	DevControl            uint
	MaxMediaOnStacker     uint
	AcceptMedia           int
	MultiPage             int
	PaperSources          uint
	MediaTaken            int
	RetractBins           uint
	MaxRetract            []uint
	ImageType             uint
	FrontImageColorFormat uint
	BackImageColorFormat  uint
	CodelineFormat        uint
	ImageSource           uint
	CharSupport           uint
	DispensePaper         int
	Extra                 []byte
	GuidLights            [LFS_PTR_GUIDLIGHTS_SIZE]uint
	Printer               string
	MediaPresented        int
	AutoRetractPeriod     uint
	RetractToTransport    int
	PowerSaveControl      int
}

type LFSFRMHEADER struct {
	FormName     string
	Base         uint
	UnitX        uint
	UnitY        uint
	Width        uint
	Height       uint
	Alignment    uint
	Orientation  uint
	OffsetX      uint
	OffsetY      uint
	VersionMajor uint
	VersionMinor uint
	UserPrompt   []byte
	CharSupport  uint
	Fields       []string
	LanguageID   uint
}

type LFSFRMMEDIA struct {
	MediaType            uint
	Base                 uint
	UnitX                uint
	UnitY                uint
	SizeWidth            uint
	SizeHeight           uint
	PageCount            uint
	LineCount            uint
	PrintAreaX           uint
	PrintAreaY           uint
	PrintAreaWidth       uint
	PrintAreaHeight      uint
	RestrictedAreaX      uint
	RestrictedAreaY      uint
	RestrictedAreaWidth  uint
	RestrictedAreaHeight uint
	Stagger              uint
	FoldType             uint
	PaperSources         uint
}

type LFSPTRQUERYFIELD struct {
	Form_Name  string
	Field_Name string
}

//LFS_INF_PTR_QUERY_FIELD Out
type DefOutLFSFRMFIELD []LFSFRMFIELD
type LFSFRMFIELD struct {
	FieldName           string
	IndexCount          uint
	Type                uint
	FieldClass          uint
	Access              uint
	Overflow            uint
	InitialValue        []byte
	UnicodeInitialValue []byte
	Format              []byte
	UnicodeFormat       []byte
	LanguageID          uint
}

type LFSPTRXDATA struct {
	Length uint
	Data   []byte
}

type LFSPTRCODELINEMAPPTRG struct {
	wCodeline_Format uint
}

//自定义
//LFS_INF_PTR_QUERY_FORM In
type DefInQUERYFORM struct {
	FormName string
}

//LFS_INF_PTR_QUERY_MEDIA In
type DefInQUERYMEDIA struct {
	MediaName string
}

//LFS_INF_PTR_FORM_LIST Out
type DefOutFORMLIST struct {
	FormList []byte
}

//LFS_INF_PTR_MEDIA_LIST Out
type DefOutMEDIALIST struct {
	MediaList []byte
}

//LFS_CMD_PTR_RETRACT_MEDIA Out
type DefOutRETRACTMEDIA struct {
	BinNumber uint
}

/*=================================================================*/
/*执行指令结构体 */
/*=================================================================*/

type LFSPTRPRINTFORM struct {
	FormName      string
	MediaName     string
	Alignment     uint
	OffsetX       uint
	OffsetY       uint
	Resolution    uint
	MediaControl  uint
	Fields        [][]byte
	UnicodeFields [][]byte
	PaperSource   uint
}

type LFSPTRREADFORM struct {
	FormName     string
	FieldNames   []string
	MediaName    string
	MediaControl uint
}

type LFSPTRREADFORMOUT struct {
	Fields        [][]byte
	UnicodeFields [][]byte //LPWSTR
}

type LFSPTRRAWDATA struct {
	InputData uint
	Size      uint
	Data      []byte
}

type LFSPTRRAWDATAIN struct {
	Size uint
	Data []byte
}

type LFSPTRMEDIAUNIT struct {
	Base  uint
	UnitX uint
	UnitY uint
}

type LFSPTRIMAGEREQUEST struct {
	FrontImageType        uint
	BackImageType         uint
	FrontImageColorFormat uint
	BackImageColorFormat  uint
	CodelineFormat        uint
	ImageSource           uint
	FrontImageFile        string
	BackImageFile         string
}

//自定义LFS_CMD_PTR_CONTROL_MEDIA In
type DefInCONTROLMEDIA struct {
	MediaControl uint
}

//LFS_CMD_PTR_RESET_COUNT In
type DefInRESETCOUNT struct {
	BinNumber uint
}

//LFS_CMD_PTR_RETRACT_MEDIA In
type DefInRETRACTMEDIA struct {
	BinNumber uint
}

//LFS_CMD_PTR_DISPENSE_PAPER
type DefInDISPENSEPAPER struct {
	PaperSource uint
}

//自定义LFS_CMD_PTR_READ_IMAGE Out
type DefOutLFSPTRIMAGE []LFSPTRIMAGE
type LFSPTRIMAGE struct {
	Image_Source uint
	Status       uint
	Data_Length  uint
	Data         []byte
}

type LFSPTRRESET struct {
	MediaControl     uint
	RetractBinNumber uint
}

type LFSPTRSETGUIDLIGHT struct {
	GuidLight uint
	Command   uint
}

type LFSPTRPRINTRAWFILE struct {
	FileName     string
	MediaControl uint
	PaperSource  uint
}

type LFSPTRLOADDEFINITION struct {
	FileName  string
	Overwrite int
}

type LFSPTRSUPPLYREPLEN struct {
	SupplyReplen uint
}

type LFSPTRPOWERSAVECONTROL struct {
	MaxPowerSaveRecoveryTime uint
}

/*=================================================================*/
/*事件结构体 */
/*=================================================================*/

type LFSPTRFIELDFAIL struct {
	FormName  string
	FieldName string
	Failure   uint
}

type LFSPTRBINTHRESHOLD struct {
	BinNumber  uint
	RetractBin uint
}

type LFSPTRPAPERTHRESHOLD struct {
	PaperSource    uint
	PaperThreshold uint
}

type LFSPTRMEDIADETECTED struct {
	Position         uint
	RetractBinNumber uint
}

type LFSPTRBINSTATUS struct {
	BinNumber  uint
	RetractBin uint
}

type LFSPTRMEDIAPRESENTED struct {
	WadIndex  uint
	TotalWads uint
}

type LFSPTRDEFINITIONLOADED struct {
	DefinitionName string
	DefinitionType uint
}

type LFSPTRMEDIAREJECTED struct {
	MediaRejected uint
}

type LFSPTRMEDIARETRACTED struct {
	RetractResult uint
	BinNumber     uint
}

type LFSPTRDEVICEPOSITION struct {
	Position uint
}

type LFSPTRPOWERSAVECHANGE struct {
	PowerSaveRecoveryTime uint
}

type LFSPTRIMAGEREQUESTEX struct {
	FrontImageType         uint
	BackImageType          uint
	FrontImageColor_Format uint
	BackImageColorFormat   uint
	CodelineFormat         uint
	ImageSource            uint
	FrontImageFile         string
	BackImageFile          string
	EncryptMode            uint
	Key                    string
	Feature                string
	Extra                  string
}

type LFSPTRIMAGEEX struct {
	ImageSource uint
	Status      uint
	DataLength  uint
	Data        []byte
	ImageData   string
	Code        uint
}

type LFSPTRCODELINEMAPPTRGOUT struct {
	CodelineFormat uint
	CharMapPTRg    LFSPTRXDATA
}

type LFSPTRIMPORTKEY struct {
	Key           string
	Enc_Key       string
	Value         LFSPTRXDATA
	ControlVector LFSPTRXDATA
	Use           uint
	KeyCheckMode  uint
	KeyCheckValue LFSPTRXDATA
	Extra         []byte
}

type LFSPTRMEDIAEXT struct {
	SizeX uint
	SizeY uint
}

//自定义
//LFS_EXEE_PTR_NOMEDIA
type DefEventEXEENOMEDIA struct {
	UserPrompt uint
}

//LFS_USRE_PTR_TONERTHRESHOLD
type DefEventUSRETONERTHRESHOLD struct {
	TonerThreshold uint
}

//LFS_USRE_PTR_LAMPTHRESHOLD
type DefEventUSRELAMPTHRESHOLD struct {
	LampThreshold uint
}

//LFS_USRE_PTR_INKTHRESHOLD
type DefEventUSREINKTHRESHOLD struct {
	InkThreshold uint
}
