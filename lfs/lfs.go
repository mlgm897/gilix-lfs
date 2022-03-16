package lfs

import (
	"bytes"
	"encoding/json"
	"reflect"
	"time"

	"github.com/lindorof/gilix-lfs/kit"

	"github.com/lindorof/gilix"
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
