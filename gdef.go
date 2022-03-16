package glfs

import (
	"github.com/lindorof/gilix"
	"github.com/lindorof/gilix/rdc"

	"github.com/lindorof/gilix-lfs/kit"
	"github.com/lindorof/gilix-lfs/lfs"
)

/* ********************************************************** */
// Rsp / Evt
/* ********************************************************** */

type LFSRsp struct {
	Out    gilix.PARA
	Result gilix.RET
}

func (r *LFSRsp) Para() gilix.PARA {
	return r.Out
}
func (r *LFSRsp) Ret() gilix.RET {
	return r.Result
}

type LFSEvt struct {
	EType gilix.TYPE
	ECode gilix.CODE
	EAux  gilix.PARA
	EData gilix.PARA
}

func (e *LFSEvt) Type() gilix.TYPE {
	return e.EType
}
func (e *LFSEvt) Code() gilix.CODE {
	return e.ECode
}
func (e *LFSEvt) Aux() gilix.PARA {
	return e.EAux
}
func (e *LFSEvt) Para() gilix.PARA {
	return e.EData
}

/* ********************************************************** */
// LFSDevCp
/* ********************************************************** */

type LFSDevCp interface {
	Rdc() rdc.Rdc
}

/* ********************************************************** */
// LFSDevpX
/* ********************************************************** */

// NewXXX , e.g. NewIDCCRD
type NewDevpX func(phyini *kit.PhyIni, dc gilix.DevCp, ldc LFSDevCp) (LFSDevpX, LFSDevX)

type LFSDevpX interface {
	PollFuncs() []gilix.Callee
	OnInf(*lfs.LFSMsg, *lfs.LFSUsr) (qut gilix.QUEUET, cee gilix.Callee, pci int)
	OnCmd(*lfs.LFSMsg, *lfs.LFSUsr) (qut gilix.QUEUET, cee gilix.Callee, chk bool, chg bool)
	OnEvt(cur gilix.PollCache, e *LFSEvt) (gilix.ERCV, gilix.EHSU)
	PollChange(old gilix.PollCache, new gilix.PollCache)
}

/* ********************************************************** */
// LFSDevX
/* ********************************************************** */

// NewXXX , e.g. NewMTCR
type NewDevX func(phyini *kit.PhyIni, dc gilix.DevCp, ldc LFSDevCp) LFSDevX

type LFSDevX interface {
	Init()
	Fini()

	// Inf & Cmd with type of gilix.Callee
}

/* ********************************************************** */
// CBS
/* ********************************************************** */

type DecodeLpara func(*lfs.LFSMsg)

var NewDevpXCb NewDevpX = nil
var DecodeLparaCb DecodeLpara = nil
