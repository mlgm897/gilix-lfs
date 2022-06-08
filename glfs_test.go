package glfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/lindorof/gilix"
	"github.com/lindorof/gilix-lfs/lfs"
)

func TestTypeAssertion(t *testing.T) {
	cases := []struct {
		v interface{}
		a interface{}
		e bool
	}{
		{lfs.LFSMsg{Lhs: 1}, lfs.LFSMsg{Lhs: 1}, true},
		{lfs.LFSMsg{Lhs: 8}, lfs.LFSMsg{Lhs: 9}, true},
		{lfs.LFSMsg{Lhs: 9}, &lfs.LFSMsg{Lhs: 9}, false},
		{&lfs.LFSMsg{Lhs: 2}, &lfs.LFSMsg{Lhs: 2}, true},
		{&lfs.LFSMsg{Lhs: 5}, &lfs.LFSMsg{Lhs: 6}, true},
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
		{lfs.LFSMsg{}, reflect.Struct},
		{&lfs.LFSMsg{}, reflect.Ptr},
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
		{[]lfs.LFSMsg{}, reflect.Slice},
		{&[]lfs.LFSMsg{}, reflect.Ptr},
		{[]*lfs.LFSMsg{}, reflect.Slice},
		{&[]*lfs.LFSMsg{}, reflect.Ptr},
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
		Fint    int
		Fuint   int
		Fbytes  []byte
		Fbytess [][]byte
		Fstr    string
	}

	cases := []struct {
		pm interface{}
		pu interface{}
	}{
		{&stru{1, 2, []byte{0x03, 0x04}, [][]byte{{0x01, 0x02}, {0x10, 0x11}}, "str1"}, &stru{}},
		{stru{3, 4, []byte{0x00, 0x05}, [][]byte{{0xa1, 0xa2}, {0x1a, 0x1b}}, "str3"}, &stru{}},
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
		Fint    int
		Fuint   int
		Fbytes  []byte
		Fbytess [][]byte
		Fstr    string
	}

	cases := []struct {
		tp gilix.TYPE
		cd gilix.CODE
		pm interface{}
		pu interface{}
	}{
		{1, 11, nil, &stru{}},
		{2, 12, &stru{1, 2, []byte{0x01}, [][]byte{{0x01, 0x02}, {0x10, 0x11}}, "str1"}, &stru{}},
		{3, 13, stru{3, 4, []byte{0x00}, [][]byte{{0xa1, 0xa2}, {0x1a, 0x1b}}, "str2"}, &stru{}},
		{4, 14, &[]stru{{}, {Fint: 1}, {Fuint: 2}, {Fbytes: []byte{0x03}}, {Fstr: "s4"}}, &[]stru{}},
		{5, 15, &[]*stru{{}, {Fint: 2}, {Fuint: 3}, {Fbytes: []byte{0x04}}, {Fstr: "s5"}}, &[]stru{}},
		{6, 16, &[]stru{}, &[]stru{}},
		{7, 17, &[]*stru{}, &[]stru{}},
		{8, 18, []stru{{}, {Fint: 1}, {Fuint: 2}, {Fbytes: []byte{0x03}}, {Fstr: "s4"}}, &[]stru{}},
		{9, 19, []*stru{{}, {Fint: 2}, {Fuint: 3}, {Fbytes: []byte{0x04}}, {Fstr: "s5"}}, &[]stru{}},
		{10, 20, nil, &[]stru{}},
		{11, 21, nil, &stru{}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.pm), func(t *testing.T) {
			msg := &lfs.LFSMsg{}
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

			m1 := lfs.LFSMsgEncode(msg)

			u1 := lfs.LFSMsgDecode(m1).(*lfs.LFSMsg)
			if u1.Ltype == c.tp && u1.Lcode == c.cd {
				u1.Lpara.DecodeVal(c.pu)
			}

			m2 := lfs.LFSMsgEncode(u1)

			if !bytes.Equal(m1, m2) {
				t.Errorf("error : m1 != m2")
				t.Logf("pm : %+v", msg.Lpara.Val)
				t.Logf("pu : %+v", u1.Lpara.Val)
			}
		})
	}
}
