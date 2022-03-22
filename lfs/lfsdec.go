package lfs

func decodeLpara(msg *LFSMsg) {
	switch msg.Lcode {
	case LFS_CMD_IDC_READ_RAW_DATA:
		msg.Lpara.DecodeVal(&DefInLFSIDCREADRAWDATA{})
	case LFS_CMD_IDC_CHIP_IO:
		msg.Lpara.DecodeVal(&LFSIDCCHIPIO{})
	}
}
