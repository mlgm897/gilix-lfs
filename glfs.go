package glfs

import (
	"context"
	"strings"
	"sync"

	"github.com/lindorof/gilix"
	"github.com/lindorof/gilix/acp"
	acptcp "github.com/lindorof/gilix/acp/tcp"
	acpws "github.com/lindorof/gilix/acp/ws"

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
// gilixLazyEval
/* ********************************************************** */

type lazyEval struct {
	hss chan int
	ids chan int

	wg   sync.WaitGroup
	exit bool
}

func NewLazyEval() *lazyEval {
	l := &lazyEval{
		hss: make(chan int),
		ids: make(chan int),
	}

	f := func(c chan int) {
		l.wg.Add(1)

		n := 0
		for !l.exit {
			n--
			c <- n
		}

		l.wg.Done()
	}

	go f(l.hss)
	go f(l.ids)

	return l
}

func (l *lazyEval) evalHS() gilix.HS {
	return gilix.HS(<-l.hss)
}

func (l *lazyEval) evalID() gilix.ID {
	return gilix.ID(<-l.ids)
}

func (l *lazyEval) Fini() {
	l.exit = true
	<-l.hss
	<-l.ids
	l.wg.Wait()
}

/* ********************************************************** */
// gilixMain
/* ********************************************************** */

func LoopSync() {
	le := NewLazyEval()
	lfs.GeneHs = le.evalHS
	lfs.GeneId = le.evalID

	cps := gilix.NewCPS()
	gini := kit.NewGilixIni()
	zapt := kit.CreateZapt("gilix/glfs", "LoopSync")
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
		}, "./cilix_rdc"+kit.ExeExt(), "-init", p, "libDRV_"+p, r.Ip, r.Port)
	}

	syncer.Async(cps.SotLoopSync, cps.SotLoopBreak)
	for n, a := range gini.Acp {
		var acp acp.Acceptor
		switch strings.ToLower(a.Net) {
		case "tcp":
			acp = acptcp.CreateServer(nil, a.Addr)
			cps.SubmitAcp(acp)
		case "ws":
			acp = acpws.CreateServer(nil, a.Addr, "/")
			cps.SubmitAcp(acp)
		}
		zapt.Infof("acp start [%s][%s][%s][%p]", n, a.Net, a.Addr, acp)
	}

	syncer.WaitRelease(util.SyncerWaitModeIdle)
	proct.Wait()
	le.Fini()
}
