package glfs

import (
	"bytes"
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/lindorof/gilix"
	acptcp "github.com/lindorof/gilix/acp/tcp"
	"github.com/lindorof/gilix/rdc"
	"github.com/lindorof/gilix/util"

	"github.com/lindorof/gilix-lfs/kit"
)

func init() {
	gilix.CBS = &gilixCBS{}
}

/* ********************************************************** */
// gilixCBS
/* ********************************************************** */

type gilixCBS struct{}

func (g *gilixCBS) UsrSnap(m gilix.Msg) gilix.Usr {
	return LFSUsrSnap(m)
}
func (g *gilixCBS) MsgEncode(m gilix.Msg) []byte {
	return LFSMsgEncode(m)
}
func (g *gilixCBS) MsgDecode(b []byte) gilix.Msg {
	return LFSMsgDecode(b)
}
func (g *gilixCBS) MsgByRsp(req gilix.Msg, ret gilix.RET, rsp gilix.Rsp) gilix.Msg {
	return LFSMsgByRsp(req, ret, rsp)
}
func (g *gilixCBS) MsgByEvt(hi gilix.HsId, usr gilix.Usr, evt gilix.Evt) []gilix.Msg {
	return LFSMsgByEvt(hi, usr, evt)
}
func (g *gilixCBS) DevInit(phy string, dc gilix.DevCp) gilix.Dev {
	return LFSDevInit(phy, dc)
}
func (g *gilixCBS) DevFini(d gilix.Dev) {
	LFSDevFini(d)
}
func (g *gilixCBS) ZaptCfg() (path string, mode string, purge int, lelvel string) {
	return kit.ZaptCfg()
}

/* ********************************************************** */
// gilixMain
/* ********************************************************** */

func LoopSync() {
	gini := kit.NewGilixIni()
	zapt := kit.CreateZapt("gilix", "main")
	syncer := util.CreateSyncerWithSig(context.Background())

	zapt.Infof("entry")
	defer zapt.Infof("exit")

	proct := util.NewProct(syncer.Ctx())
	for p, r := range gini.Rdc {
		zapt.Infof("rdc start [%s][%s:%s][dial:%d]", p, r.Ip, r.Port, r.Dial)
		proct.AddCmd(func(cmd *util.ProctCmd) {
			zapt.Infof("rdc exit [%s][pid:%d][%v]", p, cmd.C.Process.Pid, cmd.E)
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

/* ********************************************************** */
// lfs-api
/* ********************************************************** */

// 错误码
const (
	LFS_SUCCESS                    = 0
	LFS_ERR_ALREADY_STARTED        = -1
	LFS_ERR_API_VER_TOO_HIGH       = -2
	LFS_ERR_API_VER_TOO_LOW        = -3
	LFS_ERR_CANCELED               = -4
	LFS_ERR_DEV_NOT_READY          = -13
	LFS_ERR_HARDWARE_ERROR         = -14
	LFS_ERR_INTERNAL_ERROR         = -15
	LFS_ERR_INVALID_ADDRESS        = -16
	LFS_ERR_INVALID_APP_HANDLE     = -17
	LFS_ERR_INVALID_BUFFER         = -18
	LFS_ERR_INVALID_CATEGORY       = -19
	LFS_ERR_INVALID_COMMAND        = -20
	LFS_ERR_INVALID_EVENT_CLASS    = -21
	LFS_ERR_INVALID_HSERVICE       = -22
	LFS_ERR_INVALID_HPROVIDER      = -23
	LFS_ERR_INVALID_MSG_OBJECT     = -24
	LFS_ERR_INVALID_MSG_REG_OBJECT = -25
	LFS_ERR_INVALID_POINTER        = -26
	LFS_ERR_INVALID_REQ_ID         = -27
	LFS_ERR_INVALID_RESULT         = -28
	LFS_ERR_INVALID_SERVPROV       = -29
	LFS_ERR_INVALID_TIMER          = -30
	LFS_ERR_INVALID_TRACELEVEL     = -31
	LFS_ERR_LOCKED                 = -32
	LFS_ERR_NO_BLOCKING_CALL       = -33
	LFS_ERR_NO_SERVPROV            = -34
	LFS_ERR_NO_SUCH_THREAD         = -35
	LFS_ERR_NO_TIMER               = -36
	LFS_ERR_NOT_LOCKED             = -37
	LFS_ERR_NOT_OK_TO_UNLOAD       = -38
	LFS_ERR_NOT_STARTED            = -39
	LFS_ERR_NOT_REGISTERED         = -40
	LFS_ERR_OP_IN_PROGRESS         = -41
	LFS_ERR_OUT_OF_MEMORY          = -42
	LFS_ERR_SERVICE_NOT_FOUND      = -43
	LFS_ERR_SPI_VER_TOO_HIGH       = -44
	LFS_ERR_SPI_VER_TOO_LOW        = -45
	LFS_ERR_SRVC_VER_TOO_HIGH      = -46
	LFS_ERR_SRVC_VER_TOO_LOW       = -47
	LFS_ERR_TIMEOUT                = -48
	LFS_ERR_UNSUPP_CATEGORY        = -49
	LFS_ERR_UNSUPP_COMMAND         = -50
	LFS_ERR_VERSION_ERROR_IN_SRVC  = -51
	LFS_ERR_INVALID_DATA           = -52
	LFS_ERR_SOFTWARE_ERROR         = -53
	LFS_ERR_CONNECTION_LOST        = -54
	LFS_ERR_USER_ERROR             = -55
	LFS_ERR_UNSUPP_DATA            = -56
)

// 事件类型
const (
	SERVICE_EVENTS = 1
	USER_EVENTS    = 2
	SYSTEM_EVENTS  = 4
	EXECUTE_EVENTS = 8
)

// 系统事件ID
const (
	LFS_SYSE_UNDELIVERABLE_MSG = 1
	LFS_SYSE_HARDWARE_ERROR    = 2
	LFS_SYSE_VERSION_ERROR     = 3
	LFS_SYSE_DEVICE_STATUS     = 4
	LFS_SYSE_APP_DISCONNECT    = 5
	LFS_SYSE_SOFTWARE_ERROR    = 6
	LFS_SYSE_USER_ERROR        = 7
	LFS_SYSE_LOCK_REQUESTED    = 8
)

/* ********************************************************** */
// util
/* ********************************************************** */

func marshal(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

func unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func ret2lfs(gret gilix.RET) gilix.RET {
	switch gret {
	case gilix.RET_SUCCESS:
		return LFS_SUCCESS
	case gilix.RET_TIMEOUT:
		return LFS_ERR_TIMEOUT
	case gilix.RET_CANCELLED:
		return LFS_ERR_CANCELED
	case gilix.RET_LOCKED:
		return LFS_ERR_LOCKED
	case gilix.RET_ALREADY_LOCKED:
		return LFS_ERR_LOCKED
	case gilix.RET_NOT_LOCKED_YET:
		return LFS_ERR_NOT_LOCKED
	case gilix.RET_UNSUPP_CATEGORY:
		return LFS_ERR_UNSUPP_CATEGORY
	case gilix.RET_UNSUPP_COMMAND:
		return LFS_ERR_UNSUPP_COMMAND
	default:
		return LFS_ERR_INVALID_RESULT
	}
}

func evt2c(evt gilix.TYPE, c uint) bool {
	if evt == gilix.TYPE_EVT_USR && c&USER_EVENTS != 0 {
		return true
	}
	if evt == gilix.TYPE_EVT_SRV && c&SERVICE_EVENTS != 0 {
		return true
	}
	if evt == gilix.TYPE_EVT_EXE && c&EXECUTE_EVENTS != 0 {
		return true
	}
	if evt == gilix.TYPE_EVT_SYS && c&SYSTEM_EVENTS != 0 {
		return true
	}
	return false
}

/* ********************************************************** */
// LFSVal
/* ********************************************************** */

type LFSVal struct {
	Val gilix.PARA
	raw json.RawMessage
}

func (t *LFSVal) MarshalJSON() ([]byte, error) {
	return marshal(t.Val)
}
func (t *LFSVal) UnmarshalJSON(data []byte) error {
	t.raw = data
	return nil
}
func (t *LFSVal) DecodeVal(v gilix.PARA) error {
	if err := unmarshal(t.raw, v); err != nil {
		return err
	}
	if len(t.raw) > 0 && !bytes.EqualFold(t.raw[:1], []byte("n")) {
		t.Val = reflect.ValueOf(v).Elem().Interface()
	}
	return nil
}

/* ********************************************************** */
// LFSUsr
/* ********************************************************** */

type LFSUsr struct {
	Logic     string
	LogicType string
	Regm      map[string]uint
}

// CBS
func LFSUsrSnap(req gilix.Msg) gilix.Usr {
	rm := req.(*LFSMsg)

	return &LFSUsr{
		Logic:     rm.Laux.Logic,
		LogicType: kit.Logic2Type(rm.Laux.Logic),
		Regm:      make(map[string]uint, 256),
	}
}

func (t *LFSUsr) Regx(msg *LFSMsg) {
	switch msg.Ltype {
	case gilix.TYPE_REG:
		c := t.Regm[msg.Laux.RegHwnd]
		c |= msg.Laux.RegEvt
		t.Regm[msg.Laux.RegHwnd] = c
	case gilix.TYPE_DEREG:
		c := t.Regm[msg.Laux.RegHwnd]
		c -= (c & msg.Laux.RegEvt)
		if c == 0 || msg.Laux.RegEvt == 0 {
			delete(t.Regm, msg.Laux.RegHwnd)
		} else {
			t.Regm[msg.Laux.RegHwnd] = c
		}
	}
}

/* ********************************************************** */
// LFSTime
/* ********************************************************** */

type LFSTime struct {
	Year        int          `json:"YEAR"`
	Month       time.Month   `json:"MONTH"`
	DayOfWeek   time.Weekday `json:"DAYOFWEEK"`
	Day         int          `json:"DAY"`
	Hour        int          `json:"HOUR"`
	Minute      int          `json:"MINUTE"`
	Second      int          `json:"SECOND"`
	MilliSecond int          `json:"MILLISECOND"`
}

func (t *LFSTime) Stamp() {
	tm := time.Now()

	t.Year = tm.Year()
	t.Month = tm.Month()
	t.DayOfWeek = tm.Weekday()
	t.Day = tm.Day()
	t.Hour = tm.Hour()
	t.Minute = tm.Minute()
	t.Second = tm.Second()
	t.MilliSecond = tm.Nanosecond() / int(time.Millisecond)
}

/* ********************************************************** */
// LFSAux
/* ********************************************************** */

type LFSAux struct {
	Logic   string `json:"LOGIC"`
	Hwnd    string `json:"HWND"`
	RegHwnd string `json:"REGHWND"`
	RegEvt  uint   `json:"REGEVT"`
}

/* ********************************************************** */
// LFSMsg
/* ********************************************************** */

type LFSMsg struct {
	Lhs        gilix.HS      `json:"HS"`
	Lid        gilix.ID      `json:"ID"`
	Ltype      gilix.TYPE    `json:"TYPE"`
	Lcode      gilix.CODE    `json:"CODE"`
	Ltimeout   gilix.TIMEOUT `json:"TIMEOUT"`
	Ltimestamp LFSTime       `json:"TIMESTAMP"`
	Laux       LFSAux        `json:"AUX"`
	Lpara      LFSVal        `json:"PARA"`
	Lret       gilix.RET     `json:"RET"`
}

func (t *LFSMsg) Hs() gilix.HS {
	return t.Lhs
}
func (t *LFSMsg) Id() gilix.ID {
	return t.Lid
}
func (t *LFSMsg) Type() gilix.TYPE {
	return t.Ltype
}
func (t *LFSMsg) Code() gilix.CODE {
	return t.Lcode
}
func (t *LFSMsg) Timeout() gilix.TIMEOUT {
	return t.Ltimeout
}
func (t *LFSMsg) Aux() gilix.PARA {
	return &t.Laux
}
func (t *LFSMsg) Para() gilix.PARA {
	return t.Lpara.Val
}
func (t *LFSMsg) Ret() gilix.RET {
	return t.Lret
}
func (t *LFSMsg) Phyname() string {
	return kit.Logic2Phy(t.Laux.Logic)
}

// CBS
func LFSMsgEncode(m gilix.Msg) []byte {
	d, e := marshal(m)
	if e == nil {
		return d
	}
	return nil
}

// CBS
func LFSMsgDecode(data []byte) gilix.Msg {
	m := &LFSMsg{}
	e := unmarshal(data, m)
	if e == nil {
		return m
	}
	return nil
}

// CBS
func LFSMsgByRsp(req gilix.Msg, ret gilix.RET, rsp gilix.Rsp) gilix.Msg {
	rm := req.(*LFSMsg)
	msg := &LFSMsg{}

	msg.Lhs = rm.Lhs
	msg.Lid = rm.Lid
	msg.Ltype = rm.Ltype
	msg.Lcode = rm.Lcode
	msg.Ltimeout = rm.Ltimeout
	msg.Ltimestamp.Stamp()

	msg.Laux = rm.Laux

	if rsp != nil {
		msg.Lpara.Val = rsp.Para()
	}

	msg.Lret = ret2lfs(ret)
	if rsp != nil {
		msg.Lret = rsp.Ret()
	}

	return msg
}

// CBS
func LFSMsgByEvt(hi gilix.HsId, usr gilix.Usr, evt gilix.Evt) []gilix.Msg {
	lu := usr.(*LFSUsr)
	ms := make([]gilix.Msg, 0, len(lu.Regm))

	for h, c := range lu.Regm {
		if !evt2c(evt.Type(), c) {
			continue
		}

		msg := &LFSMsg{}
		ms = append(ms, msg)

		msg.Lhs = hi.Hs()
		msg.Lid = hi.Id()
		msg.Ltype = evt.Type()
		msg.Lcode = evt.Code()
		msg.Ltimeout = 0
		msg.Ltimestamp.Stamp()

		msg.Laux.Logic = lu.Logic
		msg.Laux.Hwnd = h

		msg.Lpara.Val = evt.Para()

		msg.Lret = LFS_SUCCESS
	}

	return ms
}

/* ********************************************************** */
// LFSRsp
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

/* ********************************************************** */
// LFSEvt
/* ********************************************************** */

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
	r := req.(*LFSMsg)
	u := usr.(*LFSUsr)

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
	case LFS_SYSE_DEVICE_STATUS:
		return gilix.ERCV_ALL, gilix.EHSU_ALL
	case LFS_SYSE_HARDWARE_ERROR:
		return gilix.ERCV_ALL, gilix.EHSU_ALL
	case LFS_SYSE_SOFTWARE_ERROR:
		return gilix.ERCV_ALL, gilix.EHSU_ALL
	case LFS_SYSE_USER_ERROR:
		return gilix.ERCV_ALL, gilix.EHSU_ALL
	case LFS_SYSE_LOCK_REQUESTED:
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
		ECode: LFS_SYSE_LOCK_REQUESTED,
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
	OnInf(*LFSMsg, *LFSUsr) (qut gilix.QUEUET, cee gilix.Callee, pci int)
	OnCmd(*LFSMsg, *LFSUsr) (qut gilix.QUEUET, cee gilix.Callee, chk bool, chg bool)
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

type DecodeLpara func(*LFSMsg)

var NewDevpXCb NewDevpX = nil
var DecodeLparaCb DecodeLpara = nil
