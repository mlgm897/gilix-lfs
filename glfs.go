package glfs

import (
	"context"
	"strings"

	"github.com/lindorof/gilix"
	acptcp "github.com/lindorof/gilix/acp/tcp"

	_ "github.com/lindorof/gilix/sot"
	"github.com/lindorof/gilix/util"

	"github.com/lindorof/gilix-lfs/kit"
	"github.com/lindorof/gilix-lfs/lfs"
)

func init() {
	gilix.CBS = &gilixCBS{}
}

/* ********************************************************** */
// gilixCBS
/* ********************************************************** */

type gilixCBS struct{}

func (g *gilixCBS) UsrSnap(m gilix.Msg) gilix.Usr {
	return lfs.LFSUsrSnap(m)
}
func (g *gilixCBS) MsgEncode(m gilix.Msg) []byte {
	return lfs.LFSMsgEncode(m)
}
func (g *gilixCBS) MsgDecode(b []byte) gilix.Msg {
	return lfs.LFSMsgDecode(b)
}
func (g *gilixCBS) MsgByRsp(req gilix.Msg, ret gilix.RET, rsp gilix.Rsp) gilix.Msg {
	return lfs.LFSMsgByRsp(req, ret, rsp)
}
func (g *gilixCBS) MsgByEvt(hi gilix.HsId, usr gilix.Usr, evt gilix.Evt) []gilix.Msg {
	return lfs.LFSMsgByEvt(hi, usr, evt)
}
func (g *gilixCBS) DevInit(phy string, dc gilix.DevCp) gilix.Dev {
	return lfs.LFSDevInit(phy, dc)
}
func (g *gilixCBS) DevFini(d gilix.Dev) {
	lfs.LFSDevFini(d)
}
func (g *gilixCBS) ZaptCfg() (path string, mode string, purge int, lelvel string) {
	return kit.ZaptCfg()
}

/* ********************************************************** */
// gilixMain
/* ********************************************************** */

func LoopSync() {
	gini := kit.NewGilixIni()
	zapt := kit.CreateZapt("gilix/glfs", "glfs")
	syncer := util.CreateSyncerWithSig(context.Background())

	zapt.Infof("entry")
	defer zapt.Infof("exit")
	defer func() {
		if err := recover(); err != nil {
			zapt.DPanicf("[!!! panic recover] %v", err)
		}
	}()

	proct := util.NewProct(syncer.Ctx())
	for p, r := range gini.Rdc {
		zapt.Infof("rdc start [%s][%s:%s][dial:%d]", p, r.Ip, r.Port, r.Dial)
		proct.AddCmd(p, func(cmd *util.ProctCmd) {
			zapt.Infof("rdc exit [%s][pid:%d][%v]", cmd.N, cmd.I, cmd.E)
		}, "./cilix_rdc"+kit.ExeExt(), "-init", p, p+"_DRV", r.Ip, r.Port)
	}

	syncer.Async(gilix.CPS.SotLoopSync, gilix.CPS.SotLoopBreak)
	for n, a := range gini.Acp {
		switch strings.ToLower(a.Net) {
		case "tcp":
			zapt.Infof("acp start [%s][%s][net:%s]", n, a.Addr, a.Net)
			gilix.CPS.SubmitAcp(acptcp.CreateServer(nil, a.Addr))
		}
	}

	syncer.WaitRelease(util.SyncerWaitModeIdle)
	proct.Wait()
}
