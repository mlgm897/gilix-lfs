package lfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/lindorof/gilix"
)

func TestTypeAssertion(t *testing.T) {
	cases := []struct {
		v interface{}
		a interface{}
		e bool
	}{
		{LFSMsg{Lhs: 1}, LFSMsg{Lhs: 1}, true},
		{LFSMsg{Lhs: 8}, LFSMsg{Lhs: 9}, true},
		{LFSMsg{Lhs: 9}, &LFSMsg{Lhs: 9}, false},
		{&LFSMsg{Lhs: 2}, &LFSMsg{Lhs: 2}, true},
		{&LFSMsg{Lhs: 5}, &LFSMsg{Lhs: 6}, true},
		{[3]int{1}, [3]int{2}, true},
		{[3]int{1}, &[3]int{2}, false},
		{[3]int{}, [3]string{}, false},
		{[3]int{}, [4]int{}, false},
		{[]int{}, []int{}, true},
		{[]int{1}, []int{1}, true},
		{[]int{2}, []int{2, 3}, true},
		{[]int{4}, [1]int{4}, false},
		{&[]int{5}, []int{5}, false},
		{&[]int{5}, &[]int{5}, true},
		{&[]int{6}, &[]int{6, 7}, true},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v_$_%v", c.v, c.a), func(t *testing.T) {
			b := reflect.ValueOf(c.v).Type() == reflect.ValueOf(c.a).Type()
			if b != c.e {
				t.Error("type assertion failed")
			}
		})
	}
}

func TestTypeReflect(t *testing.T) {
	cases := []struct {
		v interface{}
		e reflect.Kind
	}{
		{LFSMsg{}, reflect.Struct},
		{&LFSMsg{}, reflect.Ptr},
		{"str", reflect.String},
		{1, reflect.Int},
		{int32(7), reflect.Int32},
		{int64(8), reflect.Int64},
		{16.3, reflect.Float64},
		{float32(16.6), reflect.Float32},
		{true, reflect.Bool},
		{[3]int{}, reflect.Array},
		{[4]int{1, 2}, reflect.Array},
		{[]int{1, 2, 3}, reflect.Slice},
		{[]LFSMsg{}, reflect.Slice},
		{&[]LFSMsg{}, reflect.Ptr},
		{[]*LFSMsg{}, reflect.Slice},
		{&[]*LFSMsg{}, reflect.Ptr},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.v), func(t *testing.T) {
			v := reflect.ValueOf(c.v)
			t.Logf("v [%v] [%v]\n", v.Type(), v.Kind())
			if v.Kind() == reflect.Ptr {
				ve := reflect.ValueOf(v.Elem().Interface())
				t.Logf("ve [%v] [%v]\n", ve.Type(), ve.Kind())
			}
			if v.Kind() != c.e {
				t.Error("kind fail")
			}
		})
	}
}

func TestJson(t *testing.T) {
	type stru struct {
		Fint   int
		Fuint  int
		Fbytes []byte
		Fstr   string
	}

	cases := []struct {
		pm interface{}
		pu interface{}
	}{
		{&stru{1, 2, []byte{0x03, 0x04}, "str1"}, &stru{}},
		{stru{3, 4, []byte{0x00, 0x05}, "str3"}, &stru{}},
		{nil, &stru{}},
		{[]*stru{{Fint: 8}, {Fint: 9}, {Fint: 10}}, &[]*stru{}},
		{[]*stru{{Fint: 11}, {Fint: 12}, {Fint: 13}}, &[]stru{}},
		{[]*stru{}, &[]*stru{}},
		{[]*stru{}, &[]stru{}},
		{nil, &[]*stru{}},
		{nil, &[]stru{}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.pm), func(t *testing.T) {
			d, err := json.MarshalIndent(c.pm, "", "  ")
			if err != nil {
				t.Error("!!! marsharl error : ", err)
				return
			}
			err = json.Unmarshal(d, c.pu)
			if err != nil {
				t.Error("!!! unmarsharl error : ", err)
				return
			}
			var vv interface{} = nil
			if len(d) > 0 && !bytes.EqualFold(d[:1], []byte("n")) {
				vv = reflect.ValueOf(c.pu).Elem().Interface()
			}
			t.Logf("### %#v", vv)
		})
	}
}

func TestLFS(t *testing.T) {
	type stru struct {
		Fint   int
		Fuint  int
		Fbytes []byte
		Fstr   string
	}

	cases := []struct {
		tp gilix.TYPE
		cd gilix.CODE
		pm interface{}
		pu interface{}
	}{
		{3, 11, nil, &stru{}},
		{4, 12, &stru{1, 2, []byte{0x01}, "str1"}, &stru{}},
		{5, 13, stru{3, 4, []byte{0x00}, "str2"}, &stru{}},
		{6, 14, &[]stru{{}, {Fint: 1}, {Fuint: 2}, {Fbytes: []byte{0x03}}, {Fstr: "s4"}}, &[]stru{}},
		{7, 15, &[]*stru{{}, {Fint: 2}, {Fuint: 3}, {Fbytes: []byte{0x04}}, {Fstr: "s5"}}, &[]stru{}},
		{8, 16, &[]stru{}, &[]stru{}},
		{9, 17, &[]*stru{}, &[]stru{}},
		{10, 18, nil, &[]stru{}},
		{10, 18, nil, &stru{}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.pm), func(t *testing.T) {
			msg := &LFSMsg{}
			msg.Lhs = 1
			msg.Lid = 2
			msg.Ltype = c.tp
			msg.Lcode = c.cd
			msg.Ltimeout = -1
			msg.Ltimestamp.Stamp()
			msg.Laux.Logic = "CardReader"
			msg.Laux.Hwnd = "awind"
			msg.Lpara.Val = c.pm
			msg.Lret = 0

			m1 := LFSMsgEncode(msg)

			u1 := LFSMsgDecode(m1).(*LFSMsg)
			if u1.Ltype == c.tp && u1.Lcode == c.cd {
				u1.Lpara.DecodeVal(c.pu)
			}

			m2 := LFSMsgEncode(u1)

			if !bytes.Equal(m1, m2) {
				t.Errorf("error : m1 != m2")
				t.Logf("pm : %+v", msg.Lpara.Val)
				t.Logf("pu : %+v", u1.Lpara.Val)
			}
		})
	}
}