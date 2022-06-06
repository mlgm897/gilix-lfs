package lfs

import "github.com/lindorof/gilix"

//PTR
func decodeLparaPTR(msg *LFSMsg) {
	if msg.Ltype == gilix.TYPE_INF {
		switch msg.Lcode {

		}
	} else if msg.Ltype == gilix.TYPE_CMD {
		switch msg.Lcode {

		}
	}
}

//PIN
func decodeLparaPIN(msg *LFSMsg) {
	if msg.Ltype == gilix.TYPE_INF {
		switch msg.Lcode {

		}
	} else if msg.Ltype == gilix.TYPE_CMD {
		switch msg.Lcode {

		}
	}
}

//IDC
func decodeLparaIDC(msg *LFSMsg) {
	if msg.Ltype == gilix.TYPE_INF {
		switch msg.Lcode {
		case LFS_INF_IDC_QUERY_FORM:
			msg.Lpara.DecodeVal(&DefInLFSIDCQUERYFORM{})
		}
	} else if msg.Ltype == gilix.TYPE_CMD {
		switch msg.Lcode {
		case LFS_CMD_IDC_READ_TRACK:
			msg.Lpara.DecodeVal(&DefInLFSIDCREADTRACK{})
		case LFS_CMD_IDC_WRITE_TRACK:
			msg.Lpara.DecodeVal(&LFSIDCWRITETRACK{})
		case LFS_CMD_IDC_EJECT_CARD:
			msg.Lpara.DecodeVal(&LFSIDCEJECTCARD{})
		case LFS_CMD_IDC_RETAIN_CARD:
			msg.Lpara.DecodeVal(&LFSIDCRETAINCARD{})
		case LFS_CMD_IDC_SETKEY:
			msg.Lpara.DecodeVal(&LFSIDCSETKEY{})
		case LFS_CMD_IDC_READ_RAW_DATA:
			msg.Lpara.DecodeVal(&DefInLFSIDCREADRAWDATA{})
		case LFS_CMD_IDC_WRITE_RAW_DATA:
			msg.Lpara.DecodeVal(&DefInLFSIDCWRITERAWDATA{})
		case LFS_CMD_IDC_CHIP_IO:
			msg.Lpara.DecodeVal(&LFSIDCCHIPIO{})
		case LFS_CMD_IDC_RESET:
			msg.Lpara.DecodeVal(&DefInLFSIDCRESET{})
		case LFS_CMD_IDC_CHIP_POWER:
			msg.Lpara.DecodeVal(&DefInLFSIDCCHIPPOWER{})
		case LFS_CMD_IDC_PARSE_DATA:
			msg.Lpara.DecodeVal(&LFSIDCPARSEDATA{})
		case LFS_CMD_IDC_SET_GUIDANCE_LIGHT:
			msg.Lpara.DecodeVal(&LFSIDCSETGUIDLIGHT{})
		case LFS_CMD_IDC_POWER_SAVE_CONTROL:
			msg.Lpara.DecodeVal(&LFSIDCPOWERSAVECONTROL{})
		case LFS_CMD_IDC_RETRACT_CARD:
			msg.Lpara.DecodeVal(&LFSIDCRETRACTCARD{})
		case LFS_CMD_IDC_HANDBACK_CARD:
			msg.Lpara.DecodeVal(&LFSIDCHANDBACKCARD{})
		case LFS_CMD_IDC_CLEAR_RETRACT_CARD:
			msg.Lpara.DecodeVal(&LFSIDCCLEARRETRACTCARD{})
		case LFS_CMD_IDC_RETAIN_SLOT_START_EX:
			msg.Lpara.DecodeVal(&LFSIDCRETAINSLOTSTARTIN{})
		case LFS_CMD_IDC_RETAIN_SLOT_END:
			msg.Lpara.DecodeVal(&LFSIDCRETAINSLOT{})
		case LFS_CMD_IDC_EJECT_SLOT_START:
			msg.Lpara.DecodeVal(&LFSIDCEJECTSLOT{})
		}
	}

}

func decodeLpara(msg *LFSMsg) {
	if LFSMsgDecodeLparaCb != nil && LFSMsgDecodeLparaCb(msg) {
		return
	}

	switch {
	case msg.Lcode >= PTR_SERVICE_OFFSET && msg.Lcode < PTR_SERVICE_OFFSET+100:
		decodeLparaPTR(msg)
	case msg.Lcode >= IDC_SERVICE_OFFSET && msg.Lcode < IDC_SERVICE_OFFSET+100:
		decodeLparaIDC(msg)
	case msg.Lcode >= PIN_SERVICE_OFFSET && msg.Lcode < PIN_SERVICE_OFFSET+100:
		decodeLparaPIN(msg)
	}
}
