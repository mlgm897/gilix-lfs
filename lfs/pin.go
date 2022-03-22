package lfs

/* values of LFSPINCAPS.wClass */
const (
	LFS_SERVICE_CLASS_PIN = 4
	PIN_SERVICE_OFFSET    = (LFS_SERVICE_CLASS_PIN * 100)
)

/* PIN Info Commands */
const (
	LFS_INF_PIN_STATUS                   = (PIN_SERVICE_OFFSET + 1)
	LFS_INF_PIN_CAPABILITIES             = (PIN_SERVICE_OFFSET + 2)
	LFS_INF_PIN_KEY_DETAIL               = (PIN_SERVICE_OFFSET + 4)
	LFS_INF_PIN_FUNCKEY_DETAIL           = (PIN_SERVICE_OFFSET + 5)
	LFS_INF_PIN_HSM_TDATA                = (PIN_SERVICE_OFFSET + 6)
	LFS_INF_PIN_KEY_DETAIL_EX            = (PIN_SERVICE_OFFSET + 7)
	LFS_INF_PIN_SECUREKEY_DETAIL         = (PIN_SERVICE_OFFSET + 8)
	LFS_INF_PIN_QUERY_LOGICAL_HSM_DETAIL = (PIN_SERVICE_OFFSET + 9)
)

/* PIN Command Verbs */
const (
	LFS_CMD_PIN_CRYPT                             = (PIN_SERVICE_OFFSET + 1)
	LFS_CMD_PIN_IMPORT_KEY                        = (PIN_SERVICE_OFFSET + 3)
	LFS_CMD_PIN_GET_PIN                           = (PIN_SERVICE_OFFSET + 5)
	LFS_CMD_PIN_GET_PINBLOCK                      = (PIN_SERVICE_OFFSET + 7)
	LFS_CMD_PIN_GET_DATA                          = (PIN_SERVICE_OFFSET + 8)
	LFS_CMD_PIN_INITIALIZATION                    = (PIN_SERVICE_OFFSET + 9)
	LFS_CMD_PIN_LOCAL_DES                         = (PIN_SERVICE_OFFSET + 10)
	LFS_CMD_PIN_LOCAL_EUROCHEQUE                  = (PIN_SERVICE_OFFSET + 11)
	LFS_CMD_PIN_LOCAL_VISA                        = (PIN_SERVICE_OFFSET + 12)
	LFS_CMD_PIN_CREATE_OFFSET                     = (PIN_SERVICE_OFFSET + 13)
	LFS_CMD_PIN_DERIVE_KEY                        = (PIN_SERVICE_OFFSET + 14)
	LFS_CMD_PIN_PRESENT_IDC                       = (PIN_SERVICE_OFFSET + 15)
	LFS_CMD_PIN_LOCAL_BANKSYS                     = (PIN_SERVICE_OFFSET + 16)
	LFS_CMD_PIN_BANKSYS_IO                        = (PIN_SERVICE_OFFSET + 17)
	LFS_CMD_PIN_RESET                             = (PIN_SERVICE_OFFSET + 18)
	LFS_CMD_PIN_HSM_SET_TDATA                     = (PIN_SERVICE_OFFSET + 19)
	LFS_CMD_PIN_SECURE_MSG_SEND                   = (PIN_SERVICE_OFFSET + 20)
	LFS_CMD_PIN_SECURE_MSG_RECEIVE                = (PIN_SERVICE_OFFSET + 21)
	LFS_CMD_PIN_GET_JOURNAL                       = (PIN_SERVICE_OFFSET + 22)
	LFS_CMD_PIN_IMPORT_KEY_EX                     = (PIN_SERVICE_OFFSET + 23)
	LFS_CMD_PIN_ENC_IO                            = (PIN_SERVICE_OFFSET + 24)
	LFS_CMD_PIN_HSM_INIT                          = (PIN_SERVICE_OFFSET + 25)
	LFS_CMD_PIN_IMPORT_RSA_PUBLIC_KEY             = (PIN_SERVICE_OFFSET + 26)
	LFS_CMD_PIN_EXPORT_RSA_ISSUER_SIGNED_ITEM     = (PIN_SERVICE_OFFSET + 27)
	LFS_CMD_PIN_IMPORT_RSA_SIGNED_DES_KEY         = (PIN_SERVICE_OFFSET + 28)
	LFS_CMD_PIN_GENERATE_RSA_KEY_PAIR             = (PIN_SERVICE_OFFSET + 29)
	LFS_CMD_PIN_EXPORT_RSA_EPP_SIGNED_ITEM        = (PIN_SERVICE_OFFSET + 30)
	LFS_CMD_PIN_LOAD_CERTIFICATE                  = (PIN_SERVICE_OFFSET + 31)
	LFS_CMD_PIN_GET_CERTIFICATE                   = (PIN_SERVICE_OFFSET + 32)
	LFS_CMD_PIN_REPLACE_CERTIFICATE               = (PIN_SERVICE_OFFSET + 33)
	LFS_CMD_PIN_START_KEY_EXCHANGE                = (PIN_SERVICE_OFFSET + 34)
	LFS_CMD_PIN_IMPORT_RSA_ENCIPHERED_PKCS7_KEY   = (PIN_SERVICE_OFFSET + 35)
	LFS_CMD_PIN_EMV_IMPORT_PUBLIC_KEY             = (PIN_SERVICE_OFFSET + 36)
	LFS_CMD_PIN_DIGEST                            = (PIN_SERVICE_OFFSET + 37)
	LFS_CMD_PIN_SECUREKEY_ENTRY                   = (PIN_SERVICE_OFFSET + 38)
	LFS_CMD_PIN_GENERATE_KCV                      = (PIN_SERVICE_OFFSET + 39)
	LFS_CMD_PIN_SET_GUIDANCE_LIGHT                = (PIN_SERVICE_OFFSET + 41)
	LFS_CMD_PIN_MAINTAIN_PIN                      = (PIN_SERVICE_OFFSET + 42)
	LFS_CMD_PIN_KEYPRESS_BEEP                     = (PIN_SERVICE_OFFSET + 43)
	LFS_CMD_PIN_SET_PINBLOCK_DATA                 = (PIN_SERVICE_OFFSET + 44)
	LFS_CMD_PIN_SET_LOGICAL_HSM                   = (PIN_SERVICE_OFFSET + 45)
	LFS_CMD_PIN_IMPORT_KEYBLOCK                   = (PIN_SERVICE_OFFSET + 46)
	LFS_CMD_PIN_POWER_SAVE_CONTROL                = (PIN_SERVICE_OFFSET + 47)
	LFS_CMD_PIN_GET_PINBLOCK_EX                   = (PIN_SERVICE_OFFSET + 53)
	LFS_CMD_PIN_CHN_DIGEST                        = (PIN_SERVICE_OFFSET + 60)
	LFS_CMD_PIN_CHN_SET_SM2_PARAM                 = (PIN_SERVICE_OFFSET + 61)
	LFS_CMD_PIN_CHN_IMPORT_SM2_PUBLIC_KEY         = (PIN_SERVICE_OFFSET + 62)
	LFS_CMD_PIN_CHN_SIGN                          = (PIN_SERVICE_OFFSET + 63)
	LFS_CMD_PIN_CHN_VERIFY                        = (PIN_SERVICE_OFFSET + 64)
	LFS_CMD_PIN_CHN_EXPORT_SM2_ISSUER_SIGNED_ITEM = (PIN_SERVICE_OFFSET + 65)
	LFS_CMD_PIN_CHN_GENERATE_SM2_KEY_PAIR         = (PIN_SERVICE_OFFSET + 66)
	LFS_CMD_PIN_CHN_EXPORT_SM2_EPP_SIGNED_ITEM    = (PIN_SERVICE_OFFSET + 67)
	LFS_CMD_PIN_CHN_IMPORT_SM2_SIGNED_SM4_KEY     = (PIN_SERVICE_OFFSET + 68)
)

/* PIN Messages */
const (
	LFS_EXEE_PIN_KEY                = (PIN_SERVICE_OFFSET + 1)
	LFS_SRVE_PIN_INITIALIZED        = (PIN_SERVICE_OFFSET + 2)
	LFS_SRVE_PIN_ILLEGAL_KEY_ACCESS = (PIN_SERVICE_OFFSET + 3)
	LFS_SRVE_PIN_OPT_REQUIRED       = (PIN_SERVICE_OFFSET + 4)
	LFS_SRVE_PIN_HSM_TDATA_CHANGED  = (PIN_SERVICE_OFFSET + 5)
	LFS_SRVE_PIN_CERTIFICATE_CHANGE = (PIN_SERVICE_OFFSET + 6)
	LFS_SRVE_PIN_HSM_CHANGED        = (PIN_SERVICE_OFFSET + 7)
	LFS_EXEE_PIN_ENTERDATA          = (PIN_SERVICE_OFFSET + 8)
	LFS_SRVE_PIN_DEVICEPOSITION     = (PIN_SERVICE_OFFSET + 9)
	LFS_SRVE_PIN_POWER_SAVE_CHANGE  = (PIN_SERVICE_OFFSET + 10)
)

/* values of LFSPINSTATUS.fwDevice */
const (
	LFS_PIN_DEVONLINE       = LFS_STAT_DEVONLINE
	LFS_PIN_DEVOFFLINE      = LFS_STAT_DEVOFFLINE
	LFS_PIN_DEVPOWEROFF     = LFS_STAT_DEVPOWEROFF
	LFS_PIN_DEVNODEVICE     = LFS_STAT_DEVNODEVICE
	LFS_PIN_DEVHWERROR      = LFS_STAT_DEVHWERROR
	LFS_PIN_DEVUSERERROR    = LFS_STAT_DEVUSERERROR
	LFS_PIN_DEVBUSY         = LFS_STAT_DEVBUSY
	LFS_PIN_DEVFRAUDATTEMPT = LFS_STAT_DEVFRAUDATTEMPT
)

/* values of LFSPINSTATUS.fwEnc_Stat */
const (
	LFS_PIN_ENCREADY          = 0
	LFS_PIN_ENCNOTREADY       = 1
	LFS_PIN_ENCNOTINITIALIZED = 2
	LFS_PIN_ENCBUSY           = 3
	LFS_PIN_ENCUNDEFINED      = 4
	LFS_PIN_ENCINITIALIZED    = 5
	LFS_PIN_ENCPINTAMPERED    = 6
)

/* Size and max index of Guid_Lights array */
const (
	LFS_PIN_GUIDLIGHTS_SIZE = 32
	LFS_PIN_GUIDLIGHTS_MAX  = (LFS_PIN_GUIDLIGHTS_SIZE - 1)
)

/* Indices of LFSPINSTATUS.GuidLights [...]
   LFSPINCAPS.GuidLights [...]
*/
const (
	LFS_PIN_GUIDANCE_PINPAD = 0
)

/* Values of LFSPINSTATUS.GuidLights [...]
   LFSPINCAPS.GuidLights [...]
*/
const (
	LFS_PIN_GUIDANCE_NOT_AVAILABLE = 0x00000000
	LFS_PIN_GUIDANCE_OFF           = 0x00000001
	LFS_PIN_GUIDANCE_ON            = 0x00000002
	LFS_PIN_GUIDANCE_SLOW_FLASH    = 0x00000004
	LFS_PIN_GUIDANCE_MEDIUM_FLASH  = 0x00000008
	LFS_PIN_GUIDANCE_QUICK_FLASH   = 0x00000010
	LFS_PIN_GUIDANCE_CONTINUOUS    = 0x00000080
	LFS_PIN_GUIDANCE_RED           = 0x00000100
	LFS_PIN_GUIDANCE_GREEN         = 0x00000200
	LFS_PIN_GUIDANCE_YELLOW        = 0x00000400
	LFS_PIN_GUIDANCE_BLUE          = 0x00000800
	LFS_PIN_GUIDANCE_CYAN          = 0x00001000
	LFS_PIN_GUIDANCE_MAGENTA       = 0x00002000
	LFS_PIN_GUIDANCE_WHITE         = 0x00004000
)

/* values for LFSPINSTATUS.fwAuto_Beep_Mode and
LFS_CMD_PIN_KEYPRESS_BEEP lpwMode parameter */
const (
	LFS_PIN_BEEP_ON_ACTIVE   = 0x0001
	LFS_PIN_BEEP_ON_INACTIVE = 0x0002
)

/* values of LFSPINSTATUS.wDevicePosition
   LFSPINDEVICEPOSITION.wPosition */
const (
	LFS_PIN_DEVICEINPOSITION    = 0
	LFS_PIN_DEVICENOTINPOSITION = 1
	LFS_PIN_DEVICEPOSUNKNOWN    = 2
	LFS_PIN_DEVICEPOSNOTSUPP    = 3
)

/* values of LFSPINCAPS.fwType */
const (
	LFS_PIN_TYPEEPP = 0x0001
	LFS_PIN_TYPEEDM = 0x0002
	LFS_PIN_TYPEHSM = 0x0004
	LFS_PIN_TYPEETS = 0x0008
)

/* values of LFSPINCAPS.fwAlgorithms, LFSPINCRYPT.wAlgorithm */
const (
	LFS_PIN_CRYPTDESECB    = 0x0001
	LFS_PIN_CRYPTDESCBC    = 0x0002
	LFS_PIN_CRYPTDESCFB    = 0x0004
	LFS_PIN_CRYPTRSA       = 0x0008
	LFS_PIN_CRYPTECMA      = 0x0010
	LFS_PIN_CRYPTDESMAC    = 0x0020
	LFS_PIN_CRYPTTRIDESECB = 0x0040
	LFS_PIN_CRYPTTRIDESCBC = 0x0080
	LFS_PIN_CRYPTTRIDESCFB = 0x0100
	LFS_PIN_CRYPTTRIDESMAC = 0x0200
	LFS_PIN_CRYPTMAAMAC    = 0x0400
	LFS_PIN_CRYPTSM4CBC    = 0x0800
	LFS_PIN_CRYPTSM4       = 0x1000
	LFS_PIN_CRYPTSM4MAC    = 0x2000
)

/* values of LFSPINCAPS.fwPin_Formats */
const (
	LFS_PIN_FORM3624      = 0x0001
	LFS_PIN_FORMANSI      = 0x0002
	LFS_PIN_FORMISO0      = 0x0004
	LFS_PIN_FORMISO1      = 0x0008
	LFS_PIN_FORMECI2      = 0x0010
	LFS_PIN_FORMECI3      = 0x0020
	LFS_PIN_FORMVISA      = 0x0040
	LFS_PIN_FORMDIEBOLD   = 0x0080
	LFS_PIN_FORMDIEBOLDCO = 0x0100
	LFS_PIN_FORMVISA3     = 0x0200
	LFS_PIN_FORMBANKSYS   = 0x0400
	LFS_PIN_FORMEMV       = 0x0800
	LFS_PIN_FORMISO3      = 0x2000
)

/* values of LFSPINCAPS.fwDerivation_Algorithms */
const (
	LFS_PIN_CHIP_ZKA = 0x0001
)

/* values of LFSPINCAPS.fwPresentation_Algorithms */
const (
	LFS_PIN_PRESENT_CLEAR = 0x0001
)

/* values of LFSPINCAPS.fwDisplay */
const (
	LFS_PIN_DISPNONE       = 1
	LFS_PIN_DISPLEDTHROUGH = 2
	LFS_PIN_DISPDISPLAY    = 3
)

/* values of LFSPINCAPS.fwIDKey */
const (
	LFS_PIN_IDKEYINITIALIZATION = 0x0001
	LFS_PIN_IDKEYIMPORT         = 0x0002
)

/* values of LFSPINCAPS.fwValidation_Algorithms */
const (
	LFS_PIN_DES        = 0x0001
	LFS_PIN_EUROCHEQUE = 0x0002
	LFS_PIN_VISA       = 0x0004
	LFS_PIN_DES_OFFSET = 0x0008
	LFS_PIN_BANKSYS    = 0x0010
)

/* values of LFSPINCAPS.fwKey_Check_Modes and
   LFSPINIMPORTKEYEX.wKey_Check_Mode */
const (
	LFS_PIN_KCVNONE = 0x0000
	LFS_PIN_KCVSELF = 0x0001
	LFS_PIN_KCVZERO = 0x0002
)

/* values of LFSPINCAPS.fwAuto_Beep */
const (
	LFS_PIN_BEEP_ACTIVE_AVAILABLE    = 0x0001
	LFS_PIN_BEEP_ACTIVE_SELECTABLE   = 0x0002
	LFS_PIN_BEEP_INACTIVE_AVAILABLE  = 0x0004
	LFS_PIN_BEEP_INACTIVE_SELECTABLE = 0x0008
)

/* values of LFSPINCAPS.fwKey_Block_Import_Formats */
const (
	LFS_PIN_ANSTR31KEYBLOCK = 0x0001
)

/* values of LFSPINKEYDETAIL.fwUse and values of LFSPINKEYDETAILEX.Use */
const (
	LFS_PIN_USECRYPT           = 0x0001
	LFS_PIN_USEFUNCTION        = 0x0002
	LFS_PIN_USEMACING          = 0x0004
	LFS_PIN_USEKEYENCKEY       = 0x0020
	LFS_PIN_USENODUPLICATE     = 0x0040
	LFS_PIN_USESVENCKEY        = 0x0080
	LFS_PIN_USECONSTRUCT       = 0x0100
	LFS_PIN_USESECURECONSTRUCT = 0x0200
	LFS_PIN_USEANSTR31MASTER   = 0x0400
)

/* additional values for LFSPINKEYDETAILEX.Use */
const (
	LFS_PIN_USEPINLOCAL        = 0x00010000
	LFS_PIN_USERSAPUBLIC       = 0x00020000
	LFS_PIN_USERSAPRIVATE      = 0x00040000
	LFS_PIN_USECHIPINFO        = 0x00100000
	LFS_PIN_USECHIPPIN         = 0x00200000
	LFS_PIN_USECHIPPS          = 0x00400000
	LFS_PIN_USECHIPMAC         = 0x00800000
	LFS_PIN_USECHIPLT          = 0x01000000
	LFS_PIN_USECHIPMACLZ       = 0x02000000
	LFS_PIN_USECHIPMACAZ       = 0x04000000
	LFS_PIN_USERSAPUBLICVERIFY = 0x08000000
	LFS_PIN_USERSAPRIVATESIGN  = 0x10000000
)

/* values of LFSPINFUNCKEYDETAIL.ulFunc_Mask */
const (
	LFS_PIN_FK_0         = 0x00000001
	LFS_PIN_FK_1         = 0x00000002
	LFS_PIN_FK_2         = 0x00000004
	LFS_PIN_FK_3         = 0x00000008
	LFS_PIN_FK_4         = 0x00000010
	LFS_PIN_FK_5         = 0x00000020
	LFS_PIN_FK_6         = 0x00000040
	LFS_PIN_FK_7         = 0x00000080
	LFS_PIN_FK_8         = 0x00000100
	LFS_PIN_FK_9         = 0x00000200
	LFS_PIN_FK_ENTER     = 0x00000400
	LFS_PIN_FK_CANCEL    = 0x00000800
	LFS_PIN_FK_CLEAR     = 0x00001000
	LFS_PIN_FK_BACKSPACE = 0x00002000
	LFS_PIN_FK_HELP      = 0x00004000
	LFS_PIN_FK_DECPOINT  = 0x00008000
	LFS_PIN_FK_00        = 0x00010000
	LFS_PIN_FK_000       = 0x00020000
	LFS_PIN_FK_RES1      = 0x00040000
	LFS_PIN_FK_RES2      = 0x00080000
	LFS_PIN_FK_RES3      = 0x00100000
	LFS_PIN_FK_RES4      = 0x00200000
	LFS_PIN_FK_RES5      = 0x00400000
	LFS_PIN_FK_RES6      = 0x00800000
	LFS_PIN_FK_RES7      = 0x01000000
	LFS_PIN_FK_RES8      = 0x02000000
	LFS_PIN_FK_OEM1      = 0x04000000
	LFS_PIN_FK_OEM2      = 0x08000000
	LFS_PIN_FK_OEM3      = 0x10000000
	LFS_PIN_FK_OEM4      = 0x20000000
	LFS_PIN_FK_OEM5      = 0x40000000
	LFS_PIN_FK_OEM6      = 0x80000000
)

/* additional values of LFSPINFUNCKEYDETAIL.ulFunc_Mask */
const (
	LFS_PIN_FK_UNUSED = 0x00000000

	LFS_PIN_FK_A     = LFS_PIN_FK_RES1
	LFS_PIN_FK_B     = LFS_PIN_FK_RES2
	LFS_PIN_FK_C     = LFS_PIN_FK_RES3
	LFS_PIN_FK_D     = LFS_PIN_FK_RES4
	LFS_PIN_FK_E     = LFS_PIN_FK_RES5
	LFS_PIN_FK_F     = LFS_PIN_FK_RES6
	LFS_PIN_FK_SHIFT = LFS_PIN_FK_RES7
)

/* values of LFSPINFDK.ulFDK */
const (
	LFS_PIN_FK_FDK01 = 0x00000001
	LFS_PIN_FK_FDK02 = 0x00000002
	LFS_PIN_FK_FDK03 = 0x00000004
	LFS_PIN_FK_FDK04 = 0x00000008
	LFS_PIN_FK_FDK05 = 0x00000010
	LFS_PIN_FK_FDK06 = 0x00000020
	LFS_PIN_FK_FDK07 = 0x00000040
	LFS_PIN_FK_FDK08 = 0x00000080
	LFS_PIN_FK_FDK09 = 0x00000100
	LFS_PIN_FK_FDK10 = 0x00000200
	LFS_PIN_FK_FDK11 = 0x00000400
	LFS_PIN_FK_FDK12 = 0x00000800
	LFS_PIN_FK_FDK13 = 0x00001000
	LFS_PIN_FK_FDK14 = 0x00002000
	LFS_PIN_FK_FDK15 = 0x00004000
	LFS_PIN_FK_FDK16 = 0x00008000
	LFS_PIN_FK_FDK17 = 0x00010000
	LFS_PIN_FK_FDK18 = 0x00020000
	LFS_PIN_FK_FDK19 = 0x00040000
	LFS_PIN_FK_FDK20 = 0x00080000
	LFS_PIN_FK_FDK21 = 0x00100000
	LFS_PIN_FK_FDK22 = 0x00200000
	LFS_PIN_FK_FDK23 = 0x00400000
	LFS_PIN_FK_FDK24 = 0x00800000
	LFS_PIN_FK_FDK25 = 0x01000000
	LFS_PIN_FK_FDK26 = 0x02000000
	LFS_PIN_FK_FDK27 = 0x04000000
	LFS_PIN_FK_FDK28 = 0x08000000
	LFS_PIN_FK_FDK29 = 0x10000000
	LFS_PIN_FK_FDK30 = 0x20000000
	LFS_PIN_FK_FDK31 = 0x40000000
	LFS_PIN_FK_FDK32 = 0x80000000
)

/* values of LFSPINCRYPT.wMode */
const (
	LFS_PIN_MODEENCRYPT = 1
	LFS_PIN_MODEDECRYPT = 2
	LFS_PIN_MODERANDOM  = 3
)

/* values of LFSPINENTRY.wCompletion */
const (
	LFS_PIN_COMPAUTO      = 0
	LFS_PIN_COMPENTER     = 1
	LFS_PIN_COMPCANCEL    = 2
	LFS_PIN_COMPCONTINUE  = 6
	LFS_PIN_COMPCLEAR     = 7
	LFS_PIN_COMPBACKSPACE = 8
	LFS_PIN_COMPFDK       = 9
	LFS_PIN_COMPHELP      = 10
	LFS_PIN_COMPFK        = 11
	LFS_PIN_COMPCONTFDK   = 12
)

/* values of LFSPINSECMSG.wProtocol */
const (
	LFS_PIN_PROTISOAS     = 1
	LFS_PIN_PROTISOLZ     = 2
	LFS_PIN_PROTISOPS     = 3
	LFS_PIN_PROTCHIPZKA   = 4
	LFS_PIN_PROTRAWDATA   = 5
	LFS_PIN_PROTPBM       = 6
	LFS_PIN_PROTHSMLDI    = 7
	LFS_PIN_PROTGENAS     = 8
	LFS_PIN_PROTCHIPINCHG = 9
	LFS_PIN_PROTPINCMP    = 10
	LFS_PIN_PROTISOPINCHG = 11
)

/* values of LFSPINHSMINIT.wInit_Mode. */
const (
	LFS_PIN_INITTEMP         = 1
	LFS_PIN_INITDEFINITE     = 2
	LFS_PIN_INITIRREVERSIBLE = 3
)

/* values of LFSPINENCIO.wProtocol and LFSPINCAPS.fwPINENCIO_Protocols */
const (
	LFS_PIN_ENC_PROT_CH    = 0x0001
	LFS_PIN_ENC_PROT_GIECB = 0x0002
	LFS_PIN_ENC_PROT_LUX   = 0x0004
	LFS_PIN_ENC_PROT_CHN   = 0x0008
)

/* values for LFS_SRVE_PIN_CERTIFICATE_CHANGE and LFSPINSTATUS.Certificate_State */
const (
	LFS_PIN_CERT_SECONDARY = 0x00000002
)

/* values for LFSPINSTATUS.CertificateState*/
const (
	LFS_PIN_CERT_UNKNOWN  = 0x00000000
	LFS_PIN_CERT_PRIMARY  = 0x00000001
	LFS_PIN_CERT_NOTREADY = 0x00000004
)

/* Values for LFSPINCAPS.RSA_Authentication_Scheme and the fast-track Capabilities
lpszExtra parameter, REMOTE_KEY_SCHEME. */
const (
	LFS_PIN_RSA_AUTH_2PARTY_SIG  = 0x00000001
	LFS_PIN_RSA_AUTH_3PARTY_CERT = 0x00000002
)

/* Values for LFSPINCAPS.Signature_Scheme and the fast-track Capabilities lpzExtra parameter, SIGNATURE_CAPABILITIES. */
const (
	LFS_PIN_SIG_GEN_RSA_KEY_PAIR = 0x00000001
	LFS_PIN_SIG_RANDOM_NUMBER    = 0x00000002
	LFS_PIN_SIG_EXPORT_EPP_ID    = 0x00000004
	LFS_PIN_SIG_ENHANCED_RKL     = 0x00000008
)

/* values of LFSPINIMPORTRSAPUBLICKEY.RSA_Signature_Algorithm and
LFSPINCAPS.RSA_Signature_Algorithm */
const (
	LFS_PIN_SIGN_NA                = 0
	LFS_PIN_SIGN_RSASSA_PKCS1_V1_5 = 0x00000001
	LFS_PIN_SIGN_RSASSA_PSS        = 0x00000002
)

/* values of LFSPINIMPORTRSAPUBLICKEYOUTPUT.RSAKey_Check_Mode */
const (
	LFS_PIN_RSA_KCV_NONE   = 0x00000000
	LFS_PIN_RSA_KCV_SHA1   = 0x00000001
	LFS_PIN_RSA_KCV_SHA256 = 0x00000002
)

/* values of LFSPINEXPORTRSAISSUERSIGNEDITEM.wExport_Item_Type and */
/*           LFSPINEXPORTRSAEPPSIGNEDITEM.wExport_Item_Type        */
const (
	LFS_PIN_EXPORT_EPP_ID     = 0x0001
	LFS_PIN_EXPORT_PUBLIC_KEY = 0x0002
)

/* values of LFSPINIMPORTRSASIGNEDDESKEY.RSA_Encipher_Algorithm and
LFSPINCAPS.RSACryptAlgorithm */
const (
	LFS_PIN_CRYPT_RSAES_PKCS1_V1_5 = 0x00000001
	LFS_PIN_CRYPT_RSAES_OAEP       = 0x00000002
)

/* values of LFSPINGENERATERSAKEYPAIR.wExponent_Value */
const (
	LFS_PIN_DEFAULT     = 0
	LFS_PIN_EXPONENT_1  = 1
	LFS_PIN_EXPONENT_4  = 2
	LFS_PIN_EXPONENT_16 = 3
)

/* values of LFSPINCAPS.wDESKey_Length,                        */
/*   LFSPINIMPORTRSASIGNEDDESKEYOUTPUT.wKey_Length and         */
/*   LFSPINIMPORTRSAENCIPHEREDPKCS7KEYOUTPUT.wKey_Length       */
const (
	LFS_PIN_KEYSINGLE = 0x0001
	LFS_PIN_KEYDOUBLE = 0x0002
)

/* values of LFSPINGETCERTIFICATE.wGet_Certificate and
LFSPINCAPS.wCertificate_Types */
const (
	LFS_PIN_PUBLICENCKEY          = 0x0001
	LFS_PIN_PUBLICVERIFICATIONKEY = 0x0002
)

/* values for LFSPINEMVIMPORTPUBLICKEY.wImport_Scheme
and LFSPINCAPS.lpwEMV_Import_Schemes */
const (
	LFS_PIN_EMV_IMPORT_PLAIN_CA    = 1
	LFS_PIN_EMV_IMPORT_CHKSUM_CA   = 2
	LFS_PIN_EMV_IMPORT_EPI_CA      = 3
	LFS_PIN_EMV_IMPORT_ISSUER      = 4
	LFS_PIN_EMV_IMPORT_ICC         = 5
	LFS_PIN_EMV_IMPORT_ICC_PIN     = 6
	LFS_PIN_EMV_IMPORT_PKCSV1_5_CA = 7
)

/* values for LFSPINDIGEST.wHash_Algorithm and LFSPINCAPS.fwEMV_Hash_Algorithm */
const (
	LFS_PIN_HASH_SHA1_DIGEST   = 0x0001
	LFS_PIN_HASH_SHA256_DIGEST = 0x0002
)

/* values of LFSPINSECUREKEYDETAIL.fwKey_Entry_Mode */
const (
	LFS_PIN_SECUREKEY_NOTSUPP      = 0x0000
	LFS_PIN_SECUREKEY_REG_SHIFT    = 0x0001
	LFS_PIN_SECUREKEY_REG_UNIQUE   = 0x0002
	LFS_PIN_SECUREKEY_IRREG_SHIFT  = 0x0004
	LFS_PIN_SECUREKEY_IRREG_UNIQUE = 0x0008
)

/* PIN Errors */
const (
	LFS_ERR_PIN_KEYNOTFOUND          = (-(PIN_SERVICE_OFFSET + 0))
	LFS_ERR_PIN_MODENOTSUPPORTED     = (-(PIN_SERVICE_OFFSET + 1))
	LFS_ERR_PIN_ACCESSDENIED         = (-(PIN_SERVICE_OFFSET + 2))
	LFS_ERR_PIN_INVALIDID            = (-(PIN_SERVICE_OFFSET + 3))
	LFS_ERR_PIN_DUPLICATEKEY         = (-(PIN_SERVICE_OFFSET + 4))
	LFS_ERR_PIN_KEYNOVALUE           = (-(PIN_SERVICE_OFFSET + 6))
	LFS_ERR_PIN_USEVIOLATION         = (-(PIN_SERVICE_OFFSET + 7))
	LFS_ERR_PIN_NOPIN                = (-(PIN_SERVICE_OFFSET + 8))
	LFS_ERR_PIN_INVALIDKEYLENGTH     = (-(PIN_SERVICE_OFFSET + 9))
	LFS_ERR_PIN_KEYINVALID           = (-(PIN_SERVICE_OFFSET + 10))
	LFS_ERR_PIN_KEYNOTSUPPORTED      = (-(PIN_SERVICE_OFFSET + 11))
	LFS_ERR_PIN_NOACTIVEKEYS         = (-(PIN_SERVICE_OFFSET + 12))
	LFS_ERR_PIN_NOTERMINATEKEYS      = (-(PIN_SERVICE_OFFSET + 14))
	LFS_ERR_PIN_MINIMUMLENGTH        = (-(PIN_SERVICE_OFFSET + 15))
	LFS_ERR_PIN_PROTOCOLNOTSUPP      = (-(PIN_SERVICE_OFFSET + 16))
	LFS_ERR_PIN_INVALIDDATA          = (-(PIN_SERVICE_OFFSET + 17))
	LFS_ERR_PIN_NOTALLOWED           = (-(PIN_SERVICE_OFFSET + 18))
	LFS_ERR_PIN_NOKEYRAM             = (-(PIN_SERVICE_OFFSET + 19))
	LFS_ERR_PIN_NOCHIPTRANSACTIVE    = (-(PIN_SERVICE_OFFSET + 20))
	LFS_ERR_PIN_ALGORITHMNOTSUPP     = (-(PIN_SERVICE_OFFSET + 21))
	LFS_ERR_PIN_FORMATNOTSUPP        = (-(PIN_SERVICE_OFFSET + 22))
	LFS_ERR_PIN_HSMSTATEINVALID      = (-(PIN_SERVICE_OFFSET + 23))
	LFS_ERR_PIN_MACINVALID           = (-(PIN_SERVICE_OFFSET + 24))
	LFS_ERR_PIN_PROTINVALID          = (-(PIN_SERVICE_OFFSET + 25))
	LFS_ERR_PIN_FORMATINVALID        = (-(PIN_SERVICE_OFFSET + 26))
	LFS_ERR_PIN_CONTENTINVALID       = (-(PIN_SERVICE_OFFSET + 27))
	LFS_ERR_PIN_SIG_NOT_SUPP         = (-(PIN_SERVICE_OFFSET + 29))
	LFS_ERR_PIN_INVALID_MOD_LEN      = (-(PIN_SERVICE_OFFSET + 31))
	LFS_ERR_PIN_INVALIDCERTSTATE     = (-(PIN_SERVICE_OFFSET + 32))
	LFS_ERR_PIN_KEY_GENERATION_ERROR = (-(PIN_SERVICE_OFFSET + 33))
	LFS_ERR_PIN_EMV_VERIFY_FAILED    = (-(PIN_SERVICE_OFFSET + 34))
	LFS_ERR_PIN_RANDOMINVALID        = (-(PIN_SERVICE_OFFSET + 35))
	LFS_ERR_PIN_SIGNATUREINVALID     = (-(PIN_SERVICE_OFFSET + 36))
	LFS_ERR_PIN_SNSCDINVALID         = (-(PIN_SERVICE_OFFSET + 37))
	LFS_ERR_PIN_NORSAKEYPAIR         = (-(PIN_SERVICE_OFFSET + 38))
	LFS_ERR_PIN_INVALID_PORT         = (-(PIN_SERVICE_OFFSET + 39))
	LFS_ERR_PIN_POWERSAVETOOSHORT    = (-(PIN_SERVICE_OFFSET + 40))
	LFS_ERR_PIN_INVALIDHSM           = (-(PIN_SERVICE_OFFSET + 41))
	LFS_ERR_PIN_TOOMANYFRAMES        = (-(PIN_SERVICE_OFFSET + 42))
	LFS_ERR_PIN_PARTIALFRAME         = (-(PIN_SERVICE_OFFSET + 43))
	LFS_ERR_PIN_MISSINGKEYS          = (-(PIN_SERVICE_OFFSET + 44))
	LFS_ERR_PIN_FRAMECOORD           = (-(PIN_SERVICE_OFFSET + 45))
	LFS_ERR_PIN_KEYCOORD             = (-(PIN_SERVICE_OFFSET + 46))
	LFS_ERR_PIN_FRAMEOVERLAP         = (-(PIN_SERVICE_OFFSET + 47))
	LFS_ERR_PIN_KEYOVERLAP           = (-(PIN_SERVICE_OFFSET + 48))
	LFS_ERR_PIN_TOOMANYKEYS          = (-(PIN_SERVICE_OFFSET + 49))
	LFS_ERR_PIN_KEYALREADYDEFINED    = (-(PIN_SERVICE_OFFSET + 50))
	LFS_ERR_PIN_COMMANDUNSUPP        = (-(PIN_SERVICE_OFFSET + 51))
	LFS_ERR_PIN_SYNCHRONIZEUNSUPP    = (-(PIN_SERVICE_OFFSET + 52))
	LFS_ERR_PIN_SIGNATUREERROR       = (-(PIN_SERVICE_OFFSET + 53))
	LFS_ERR_PIN_NOPRIVATEKEY         = (-(PIN_SERVICE_OFFSET + 54))
	LFS_ERR_PIN_NOSM2KEYPAIR         = (-(PIN_SERVICE_OFFSET + 55))
)

/* values of LFSPINCHNDIGESTIN.wHash_Algorithm */
const (
	LFS_PIN_HASH_SM3_DIGEST = 0x0001
)

/* values for LFSPINCHNIMPORTSM2PUBLICKEYIN.Use */
const (
	LFS_PIN_USESM2PUBLIC       = 0x00000001
	LFS_PIN_USESM2PUBLICVERIFY = 0x00000002
)

/* values of LFSPINCHNIMPORTSM2PUBLICKEYIN.SM2_Signature_Algorithm,
PINCHNEXPORTSM2ISSUERSIGNEDITEMOUT.SM2_Signature_Algorithm,
PINCHNEXPORTSM2EPPSIGNEDITEMIN.Signature_Algorithm and
PINCHNIMPORTSM2SIGNEDSM4KEY.SM2_Signature_Algorithm */
const (
	LFS_PIN_SIGN_SM2_SIGN_NA   = 0
	LFS_PIN_SIGN_SM2_GM_T_2012 = 0x00000001
)

/* values for LFSPINCHNIMPORTSM2PUBLICKEYOUT.SM2_Key_Check_Mode */
const (
	LFS_PIN_SM2_KCV_NONE = 0x00000001
	LFS_PIN_SM2_KCV_SM3  = 0x00000002
)

/* values for LFSPINCHNGENERATESM2KEYOUT.Use */
const (
	LFS_PIN_USESM2PRIVATE     = 0x00000001
	LFS_PIN_USESM2PRIVATESIGN = 0x00000002
)

/*=================================================================*/
/* PIN Info Command Structures and variables */
/*=================================================================*/
type LFSXDATA struct {
	Length uint
	Data   []byte
}

type LFSPINSTATUS struct {
	Device                uint
	EncStat               uint
	Extra                 []byte
	GuidLights            [LFS_PIN_GUIDLIGHTS_SIZE]uint
	AutoBeepMode          uint
	CertificateState      uint
	DevicePosition        uint
	PowerSaveRecoveryTime uint
}
type LFSPINCAPS struct {
	Class                   uint
	Type                    uint
	Compound                int
	Key_Num                 uint
	Algorithms              uint
	PinFormats              uint
	DerivationAlgorithms    uint
	PresentationAlgorithms  uint
	Display                 uint
	IDConnect               int
	IDKey                   uint
	ValidationAlgorithms    uint
	KeyCheckModes           uint
	Extra                   []byte
	GuidLights              [LFS_PIN_GUIDLIGHTS_SIZE]uint
	PINCanPersistAfterUse   int
	AutoBeep                uint
	HSMVendor               string
	HSMJournaling           int
	RSAAuthenticationScheme uint
	RSASignatureAlgorithm   uint
	RSACryptAlgorithm       uint
	RSAKeyCheckMode         uint
	SignatureScheme         uint
	EMVImport_Schemes       []uint
	EMVHashAlgorithm        uint
	KeyImportThroughParts   int
	ENCIOProtocols          uint
	TypeCombined            int
	SetPinblockDataRequired int
	KeyBlockImportFormats   uint
	PowerSaveControl        int
}

//自定义
//LFS_INF_PIN_KEY_DETAIL Out
type DefOutLFSPINKEYDETAIL []LFSPINKEYDETAIL
type LFSPINKEYDETAIL struct {
	KeyName        string
	Use            uint
	Loaded         int
	KeyBlockHeader LFSXDATA
}

type LFSPINFDK struct {
	FDK       uint
	XPosition uint
	YPosition uint
}

type LFSPINFUNCKEYDETAIL struct {
	Func_Mask   uint
	Number_FDKs uint
	FDKs        []LFSPINFDK
}

//自定义
type DefOutLFSPINKEYDETAILEX []LFSPINKEYDETAILEX
type LFSPINKEYDETAILEX struct {
	KeyName        string
	Use            uint
	Generation     uint
	Version        uint
	ActivatingDate [4]uint
	ExpiryDate     [4]uint
	Loaded         int
	KeyBlockHeader LFSXDATA
}

/* LFS_INF_PIN_SECUREKEY_DETAIL command key layout output structure */
type LFSPINHEXKEYS struct {
	XPos    uint
	YPos    uint
	XSize   uint
	YSize   uint
	FK      uint
	ShiftFK uint
}

/* LFS_INF_PIN_SECUREKEY_DETAIL command output structure */
type LFSPINSECUREKEYDETAIL struct {
	KeyEntryMode  uint
	FuncKeyDetail LFSPINFUNCKEYDETAIL
	ClearFDK      uint
	CancelFDK     uint
	BackspaceFDK  uint
	EnterFDK      uint
	Columns       uint
	Rows          uint
	HexKeys       []LFSPINHEXKEYS
}

//自定义
//LFS_INF_PIN_KEY_DETAIL In
type DefInLFSPINKEYDETAIL struct {
	KeyName string
}

//LFS_INF_PIN_KEY_DETAILEX In
type DefInLFSPINKEYDETAILEX struct {
	KeyName string
}

//LFS_INF_PIN_FUNCKEY_DETAIL In
type DefInLFSPINFUNCKEYDETAIL struct {
	FDKMask uint
}

/*=================================================================*/
/* PIN Execute Command Structures */
/*=================================================================*/
type LFSPINCRYPT struct {
	Mode          uint
	Key           string
	KeyEncKey     LFSXDATA
	Algorithm     uint
	StartValueKey string
	StartValue    LFSXDATA
	Padding       uint
	Compression   uint
	CryptData     LFSXDATA
}

type LFSPINIMPORT struct {
	Key    string
	EncKey string
	Ident  LFSXDATA
	Value  LFSXDATA
	Use    uint
}

type LFSPINDERIVE struct {
	DerivationAlgorithm uint
	Key                 string
	KeyGenKey           string
	StartValueKey       string
	StartValue          LFSXDATA
	Padding             uint
	InputData           LFSXDATA
	Ident               LFSXDATA
}

type LFSPINGETPIN struct {
	MinLen        uint
	MaxLen        uint
	AutoEnd       int
	Echo          int
	ActiveFDKs    uint
	ActiveKeys    uint
	TerminateFDKs uint
	TerminateKeys uint
}

type LFSPINENTRY struct {
	Digits     uint
	Completion uint
}

type LFSPINLOCALDES struct {
	ValidationData string
	Offset         string
	Padding        uint
	MaxPIN         uint
	ValDigits      uint
	NoLeadingZero  int
	Key            string
	KeyEncKey      LFSXDATA
	DecTable       string
}

type LFSPINCREATEOFFSET struct {
	ValidationData string
	Padding        uint
	MaxPIN         uint
	ValDigits      uint
	Key            string
	KeyEncKey      LFSXDATA
	DecTable       string
}

type LFSPINLOCALEUROCHEQUE struct {
	EurochequeData string
	PVV            string
	FirstEncDigits uint
	FirstEncOffset uint
	PVVDigits      uint
	PVVOffset      uint
	Key            string
	KeyEncKey      LFSXDATA
	DecTable       string
}

type LFSPINLOCALVISA struct {
	PAN       string
	PVV       string
	PVVDigits uint
	Key       string
	KeyEncKey LFSXDATA
}

type LFSPINPRESENTIDC struct {
	PresentAlgorithm uint
	ChipProtocol     uint
	ChipDataLength   uint
	ChipData         []byte
	AlgorithmData    []byte //LPVOID
}

type LFSPINPRESENTRESULT struct {
	ChipProtocol   uint
	ChipDataLength uint
	ChipData       []byte
}

type LFSPINPRESENTCLEAR struct {
	PINPointer uint
	PINOffset  uint
}

type LFSPINBLOCK struct {
	CustomerData string
	XORData      string
	Padding      uint
	Format       uint
	Key          string
	KeyEncKey    string
}

type LFSPINGETDATA struct {
	MaxLen        uint
	AutoEnd       int
	ActiveFDKs    uint
	ActiveKeys    uint
	TerminateFDKs uint
	TerminateKeys uint
}

type LFSPINKEY struct {
	Completion uint
	Digit      uint
}

type LFSPINDATA struct {
	Keys       uint
	Pin_Keys   []LFSPINKEY
	Completion uint
}

type LFSPININIT struct {
	Ident LFSXDATA
	Key   LFSXDATA
}

type LFSPINLOCALBANKSYS struct {
	ATMVAC LFSXDATA
}

type LFSPINBANKSYSIO struct {
	Length uint
	Data   []byte
}

type LFSPINSECMSG struct {
	Protocol uint
	Length   uint
	Msg      []byte
}

type LFSPINIMPORTKEYEX struct {
	Key           string
	EncKey        string
	Value         LFSXDATA
	ControlVector LFSXDATA
	Use           uint
	KeyCheckMode  uint
	KeyCheckValue LFSXDATA
}

type LFSPINENCIO struct {
	Protocol   uint
	DataLength uint
	Data       []byte //LPVOID
}

/* LFS_CMD_PIN_SECUREKEY_ENTRY command input structure */
type LFSPINSECUREKEYENTRY struct {
	KeyLen           uint
	AutoEnd          int
	ActiveFDKs       uint
	ActiveKeys       uint
	TerminateFDKs    uint
	TerminateKeys    uint
	VerificationType uint
}

/* LFS_CMD_PIN_SECUREKEY_ENTRY command output structure */
type LFSPINSECUREKEYENTRYOUT struct {
	Digits     uint
	Completion uint
	KCV        LFSXDATA
}

/* LFS_CDM_PIN_IMPORT_KEYBLOCK command input structure */
type LFSPINIMPORTKEYBLOCK struct {
	Key      string
	EncKey   string
	KeyBlock LFSXDATA
}

type LFSPINIMPORTRSAPUBLICKEY struct {
	Key                   string
	Value                 LFSXDATA
	Use                   uint
	SigKey                string
	RSASignatureAlgorithm uint
	Signature             LFSXDATA
}

type LFSPINIMPORTRSAPUBLICKEYOUTPUT struct {
	RSAKeyCheckMode uint
	KeyCheckValue   LFSXDATA
}

type LFSPINEXPORTRSAISSUERSIGNEDITEM struct {
	ExportItemType uint
	Name           string
}

type LFSPINEXPORTRSAISSUERSIGNEDITEMOUTPUT struct {
	Value                 LFSXDATA
	RSASignatureAlgorithm uint
	Signature             LFSXDATA
}

type LFSPINIMPORTRSASIGNEDDESKEY struct {
	Key                   string
	DecryptKey            string
	RSAEncipherAlgorithm  uint
	Value                 LFSXDATA
	Use                   uint
	SigKey                string
	RSASignatureAlgorithm uint
	Signature             LFSXDATA
}

type LFSPINIMPORTRSASIGNEDDESKEYOUTPUT struct {
	KeyLength     uint
	KeyCheckMode  uint
	KeyCheckValue LFSXDATA
}

type LFSPINGENERATERSAKEYPAIR struct {
	Key           string
	Use           uint
	ModulusLength uint
	ExponentValue uint
}

type LFSPINEXPORTRSAEPPSIGNEDITEM struct {
	ExportItemType     uint
	Name               string
	SigKey             string
	SignatureAlgorithm uint
}

type LFSPINEXPORTRSAEPPSIGNEDITEMOUTPUT struct {
	Value         LFSXDATA
	SelfSignature LFSXDATA
	Signature     LFSXDATA
}

type LFSPINLOADCERTIFICATE struct {
	LoadCertificate LFSXDATA
}

type LFSPINLOADCERTIFICATEOUTPUT struct {
	CertificateData LFSXDATA
}

type LFSPINGETCERTIFICATE struct {
	GetCertificate uint
}

type LFSPINGETCERTIFICATEOUTPUT struct {
	Certificate LFSXDATA
}

type LFSPINREPLACECERTIFICATE struct {
	ReplaceCertificate LFSXDATA
}

type LFSPINREPLACECERTIFICATEOUTPUT struct {
	NewCertificateData LFSXDATA
}

type LFSPINSTARTKEYEXCHANGE struct {
	Random_Item LFSXDATA
}

type LFSPINIMPORTRSAENCIPHEREDPKCS7KEY struct {
	ImportRSAKeyIn LFSXDATA
	Key            string
	Use            uint
}

type LFSPINIMPORTRSAENCIPHEREDPKCS7KEYOUTPUT struct {
	KeyLength uint
	RSAData   LFSXDATA
}

type LFSPINEMVIMPORTPUBLICKEY struct {
	Key          string
	Use          uint
	ImportScheme uint
	ImportData   LFSXDATA
	SigKey       string
}

type LFSPINEMVIMPORTPUBLICKEYOUTPUT struct {
	ExpiryDate string
}

type LFSPINDIGEST struct {
	HashAlgorithm uint
	DigestInput   LFSXDATA
}

type LFSPINDIGESTOUTPUT struct {
	DigestOutput LFSXDATA
}

type LFSPINHSMINIT struct {
	InitMode   uint
	OnlineTime LFSXDATA
}

type LFSPINGENERATEKCV struct {
	Key          string
	KeyCheckMode uint
}

type LFSPINKCV struct {
	KCV LFSXDATA
}

type LFSPINSETGUIDLIGHT struct {
	GuidLight uint
	Command   uint
}

type LFSPINMAINTAINPIN struct {
	MaintainPIN int
}

type LFSPINHSMINFO struct {
	HSMSerialNumber uint
	ZKAID           string
}

type LFSPINHSMDETAIL struct {
	ActiveLogicalHSM uint
	HSMInfo          []LFSPINHSMINFO
}

type LFSPINHSMIDENTIFIER struct {
	HSMSerialNumber uint
}

type LFSPINPOWERSAVECONTROL struct {
	MaxPowerSaveRecoveryTime uint
}

type LFSPINBLOCKEX struct {
	CustomerData string
	XORData      string
	Padding      uint
	Format       uint
	Key          string
	KeyEncKey    string
	Algorithm    uint
}

type LFSPINLOADCERTIFICATEEX struct {
	LoadOption      uint
	Signer          uint
	CertificateData LFSXDATA
}

type LFSPINLOADCERTIFICATEEXOUTPUT struct {
	RSAKeyCheckMode uint
	RSAData         LFSXDATA
}

type LFSPINIMPORTRSAENCIPHEREDPKCS7KEYEX struct {
	ImportRSAKeyIn LFSXDATA
	Key            string
	Use            uint
	CRKLLoadOption uint
}

type LFSPINIMPORTRSAENCIPHEREDPKCS7KEYEXOUTPUT struct {
	KeyLength       uint
	RSAKeyCheckMode uint
	RSAData         LFSXDATA
	KeyCheckMode    uint
	KeyCheckValue   LFSXDATA
}

type LFSPINCHNDIGEST struct {
	HashAlgorithm uint
	DigestInput   LFSXDATA
}

type LFSPINCHNDIGESTOUTPUT struct {
	DigestOutput LFSXDATA
}

type LFSPINCHNSM2ALGORITHMPARAMIN struct {
	P  LFSXDATA
	A  LFSXDATA
	B  LFSXDATA
	N  LFSXDATA
	Xg LFSXDATA
	Yg LFSXDATA
}

type LFSPINCHNIMPORTSM2PUBLICKEY struct {
	Key                   string
	Value                 LFSXDATA
	Use                   uint
	SigKey                string
	SM2SignatureAlgorithm uint
	Signature             LFSXDATA
}

type LFSPINCHNIMPORTSM2PUBLICKEYOUTPUT struct {
	SM2KeyCheckMode uint
	KeyCheckValue   LFSXDATA
}

type LFSPINCHNSIGN struct {
	Key           string
	SignerID      string
	PlaintextData LFSXDATA
}

type LFSPINCHNSIGNOUTPUT struct {
	Sign_Data LFSXDATA
}

type LFSPINCHNVERIFY struct {
	Key            string
	Plaintext_Data LFSXDATA
	Sign_Data      LFSXDATA
}

type LFSPINCHNEXPORTSM2ISSUERSIGNEDITEM struct {
	ExportItemType uint
	Name           string
}

type LFSPINCHNEXPORTSM2ISSUERSIGNEDITEMOUTPUT struct {
	Value                 LFSXDATA
	SM2SignatureAlgorithm uint
	Signature             LFSXDATA
}

type LFSPINCHNGENERATESM2KEYPAIR struct {
	Key string
	Use uint
}

type LFSPINCHNEXPORTSM2EPPSIGNEDITEM struct {
	ExportItemType     uint
	Name               string
	SigKey             string
	SignatureAlgorithm uint
}

type LFSPINCHNEXPORTSM2EPPSIGNEDITEMOUTPUT struct {
	Value         LFSXDATA
	SelfSignature LFSXDATA
	Signature     LFSXDATA
}

type LFSPINCHNIMPORTSM2SIGNEDSM4KEY struct {
	Key                   string
	DecryptKey            string
	SM2EncipherAlgorithm  uint
	Value                 LFSXDATA
	Use                   uint
	SigKey                string
	SM2SignatureAlgorithm uint
	Signature             LFSXDATA
}

type LFSPINCHNIMPORTSM2SIGNEDSM4KEYOUTPUT struct {
	KeyCheckMode  uint
	KeyCheckValue LFSXDATA
}

/*=================================================================*/
/* PIN Message Structures */
/*=================================================================*/
type LFSPINACCESS struct {
	KeyName   string
	ErrorCode int
}

type LFSPINDEVICEPOSITION struct {
	Position uint
}

type LFSPINPOWERSAVECHANGE struct {
	PowerSaveRecoveryTime uint
}
