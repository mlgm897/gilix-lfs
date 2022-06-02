package lfs

import "github.com/lindorof/gilix"

//PTR
func lfsParaPTR(msg *LFSMsg) {
	switch msg.Lcode {
	case LFS_INF_PTR_QUERY_FORM:

		break
	}
}

//IDC
func lfsParaIDC(msg *LFSMsg) {

	if msg.Ltype == gilix.TYPE_INF {
		switch msg.Lcode {
		case LFS_INF_IDC_QUERY_FORM:
			msg.Lpara.DecodeVal(&DefInLFSIDCQUERYFORM{})
			break
		}
	} else if msg.Ltype == gilix.TYPE_CMD {
		switch msg.Lcode {
		case LFS_CMD_IDC_READ_TRACK:
			msg.Lpara.DecodeVal(&DefInLFSIDCREADTRACK{})
			break
		case LFS_CMD_IDC_WRITE_TRACK:
			msg.Lpara.DecodeVal(&LFSIDCWRITETRACK{})
			break
		case LFS_CMD_IDC_EJECT_CARD:
			msg.Lpara.DecodeVal(&LFSIDCEJECTCARD{})
			break
		case LFS_CMD_IDC_RETAIN_CARD:
			msg.Lpara.DecodeVal(&LFSIDCRETAINCARD{})
			break
		case LFS_CMD_IDC_SETKEY:
			msg.Lpara.DecodeVal(&LFSIDCSETKEY{})
			break
		case LFS_CMD_IDC_READ_RAW_DATA:
			msg.Lpara.DecodeVal(&DefInLFSIDCREADRAWDATA{})
			break
		case LFS_CMD_IDC_WRITE_RAW_DATA:
			msg.Lpara.DecodeVal(&DefInLFSIDCWRITERAWDATA{})
			break
		case LFS_CMD_IDC_CHIP_IO:
			msg.Lpara.DecodeVal(&LFSIDCCHIPIO{})
			break
		case LFS_CMD_IDC_RESET:
			msg.Lpara.DecodeVal(&DefInLFSIDCRESET{})
			break
		case LFS_CMD_IDC_CHIP_POWER:
			msg.Lpara.DecodeVal(&DefInLFSIDCCHIPPOWER{})
			break
		case LFS_CMD_IDC_PARSE_DATA:
			msg.Lpara.DecodeVal(&LFSIDCPARSEDATA{})
			break
		case LFS_CMD_IDC_SET_GUIDANCE_LIGHT:
			msg.Lpara.DecodeVal(&LFSIDCSETGUIDLIGHT{})
			break
		case LFS_CMD_IDC_POWER_SAVE_CONTROL:
			msg.Lpara.DecodeVal(&LFSIDCPOWERSAVECONTROL{})
			break
		case LFS_CMD_IDC_RETRACT_CARD:
			msg.Lpara.DecodeVal(&LFSIDCRETRACTCARD{})
			break
		case LFS_CMD_IDC_HANDBACK_CARD:
			msg.Lpara.DecodeVal(&LFSIDCHANDBACKCARD{})
			break
		case LFS_CMD_IDC_CLEAR_RETRACT_CARD:
			msg.Lpara.DecodeVal(&LFSIDCCLEARRETRACTCARD{})
			break
		case LFS_CMD_IDC_RETAIN_SLOT_START_EX:
			msg.Lpara.DecodeVal(&LFSIDCRETAINSLOTSTARTIN{})
			break
		case LFS_CMD_IDC_RETAIN_SLOT_END:
			msg.Lpara.DecodeVal(&LFSIDCRETAINSLOT{})
			break
		case LFS_CMD_IDC_EJECT_SLOT_START:
			msg.Lpara.DecodeVal(&LFSIDCEJECTSLOT{})
			break
		}
	}

}

//PIN
func lfsParaPIN(msg *LFSMsg) {
	switch msg.Lcode {
	case LFS_CMD_PIN_LOCAL_DES:
		break
	}
}

func decodeLpara(msg *LFSMsg) {
	if msg.Lcode >= PTR_SERVICE_OFFSET && msg.Lcode < PTR_SERVICE_OFFSET+100 {
		lfsParaPTR(msg)
	} else if msg.Lcode >= IDC_SERVICE_OFFSET && msg.Lcode < IDC_SERVICE_OFFSET+100 {
		lfsParaIDC(msg)
	} else if msg.Lcode >= PIN_SERVICE_OFFSET && msg.Lcode < PIN_SERVICE_OFFSET+100 {
		lfsParaPIN(msg)
	}
}
