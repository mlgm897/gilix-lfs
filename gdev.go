package glfs

import (
	"sync"

	"github.com/lindorof/gilix-lfs/kit"
	"github.com/lindorof/gilix-lfs/lfs"

	"github.com/lindorof/gilix"
	"github.com/lindorof/gilix/rdc"
)

/* ********************************************************** */
// LFSDevp
/* ********************************************************** */

type LFSDevp struct {
	phyini *kit.PhyIni
	dc     gilix.DevCp

	rdc  rdc.Rdc
	once sync.Once

	devpx LFSDevpX
	devx  LFSDevX
}

func (d *LFSDevp) Rdc() rdc.Rdc {
	d.once.Do(func() {
		d.rdc = kit.CreateRdc(d.phyini.PHY)
	})
	return d.rdc
}

// CBS
func LFSDevInit(phy string, dc gilix.DevCp) gilix.Dev {
	d := &LFSDevp{
		phyini: kit.NewPhyIni(phy),
		dc:     dc,
	}

	d.devpx, d.devx = NewDevpXCb(d.phyini, d.dc, d)
	if d.devpx == nil {
		return d
	}
	if d.devx != nil {
		d.devx.Init()
	}

	return d
}

// CBS
func LFSDevFini(pd gilix.Dev) {
	d := pd.(*LFSDevp)
	if d.devx != nil {
		d.devx.Fini()
	}
	if d.rdc != nil {
		d.rdc.Fini()
	}
}

func (d *LFSDevp) PollInterval() int {
	return d.phyini.Polli
}

func (d *LFSDevp) PollFuncs() []gilix.Callee {
	if d.devpx == nil {
		return nil
	}
	return d.devpx.PollFuncs()
}

func (d *LFSDevp) OnReq(req gilix.Msg, usr gilix.Usr) (qut gilix.QUEUET, cee gilix.Callee, pci int, chk bool, chg bool) {
	r := req.(*lfs.LFSMsg)
	u := usr.(*lfs.LFSUsr)

	u.Regx(r)

	if d.devpx == nil {
		return
	}

	DecodeLparaCb(r)
	switch r.Ltype {
	case gilix.TYPE_INF:
		qut, cee, pci = d.devpx.OnInf(r, u)
	case gilix.TYPE_CMD:
		qut, cee, chk, chg = d.devpx.OnCmd(r, u)
	}

	return
}

func (d *LFSDevp) OnEvt(cur gilix.PollCache, e gilix.Evt) (gilix.ERCV, gilix.EHSU) {
	evt := e.(*LFSEvt)

	switch evt.ECode {
	case lfs.LFS_SYSE_DEVICE_STATUS:
		return gilix.ERCV_ALL, gilix.EHSU_ALL
	case lfs.LFS_SYSE_HARDWARE_ERROR:
		return gilix.ERCV_ALL, gilix.EHSU_ALL
	case lfs.LFS_SYSE_SOFTWARE_ERROR:
		return gilix.ERCV_ALL, gilix.EHSU_ALL
	case lfs.LFS_SYSE_USER_ERROR:
		return gilix.ERCV_ALL, gilix.EHSU_ALL
	case lfs.LFS_SYSE_LOCK_REQUESTED:
		return gilix.ERCV_LOCKER, gilix.EHSU_CURRENT
	}

	if d.devpx == nil {
		return gilix.ERCV_CURRENT, gilix.ERCV_CURRENT
	}

	return d.devpx.OnEvt(cur, evt)
}

func (d *LFSDevp) OnLockTry() {
	d.dc.PostEvt(&LFSEvt{
		EType: gilix.TYPE_EVT_SYS,
		ECode: lfs.LFS_SYSE_LOCK_REQUESTED,
		EAux:  nil,
		EData: nil,
	})
}

func (d *LFSDevp) PollChange(old gilix.PollCache, new gilix.PollCache) {
	if d.devpx == nil {
		return
	}
	d.devpx.PollChange(old, new)
}
