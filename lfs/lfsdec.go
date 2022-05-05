package lfs

//ptr
func lfsPara_PTR(msg *LFSMsg) {
	switch msg.Lcode {
	case LFS_INF_PTR_QUERY_FORM:

		break
	}
}

//idc
func lfsPara_IDC(msg *LFSMsg) {
	switch msg.Lcode {
	case LFS_CMD_IDC_READ_RAW_DATA:
		msg.Lpara.DecodeVal(&DefInLFSIDCREADRAWDATA{})
		break
	case LFS_CMD_IDC_CHIP_IO:
		msg.Lpara.DecodeVal(&LFSIDCCHIPIO{})
		break
	case LFS_CMD_IDC_RESET:
		msg.Lpara.DecodeVal(&DefInLFSIDCRESET{})
		break
	case LFS_CMD_IDC_WRITE_RAW_DATA:
		msg.Lpara.DecodeVal(&DefInLFSIDCWRITERAWDATA{})
		break
	}
}

//pin
func lfsPara_PIN(msg *LFSMsg) {
	switch msg.Lcode {
	case LFS_CMD_PIN_LOCAL_DES:
		break
	}
}

func decodeLpara(msg *LFSMsg) {
	if msg.Lcode >= PTR_SERVICE_OFFSET && msg.Lcode < PTR_SERVICE_OFFSET+100 {
		lfsPara_PTR(msg)
	} else if msg.Lcode >= IDC_SERVICE_OFFSET && msg.Lcode < IDC_SERVICE_OFFSET+100 {
		lfsPara_IDC(msg)
	} else if msg.Lcode >= PIN_SERVICE_OFFSET && msg.Lcode < PIN_SERVICE_OFFSET+100 {
		lfsPara_PIN(msg)
	}
}
