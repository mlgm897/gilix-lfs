package kit

import (
	"sync"

	"github.com/go-ini/ini"
)

func LOGICAL_SERVICES(logic string) string {
	return "/etc/LFS/logical_services/" + logic + ".ini"
}

func SERVICE_PROVIDERS(provider string) string {
	return "/etc/LFS/service_providers/" + provider + ".ini"
}

func PHYSICAL_SERVICES(phy string) string {
	return "./" + phy + ".ini"
}

/* ********************************************************** */
// LfsIni
/* ********************************************************** */

func Logic2Type(logic string) string {
	l, e := ini.Load(LOGICAL_SERVICES(logic))
	if e != nil {
		return ""
	}

	return l.Section("default").Key("type").String()
}

func Logic2Phy(logic string) string {
	l, e := ini.Load(LOGICAL_SERVICES(logic))
	if e != nil {
		return ""
	}

	p := l.Section("default").Key("provider").String()

	s, e := ini.Load(SERVICE_PROVIDERS(p))
	if e != nil {
		return ""
	}

	return s.Section("default").Key("physic_name").String()
}

/* ********************************************************** */
// PhyIni
/* ********************************************************** */

type PhyIni struct {
	PHY string
	INI *ini.File

	Class string
	Type  string
	Polli int
}

func NewPhyIni(phy string) *PhyIni {
	i, e := ini.Load(PHYSICAL_SERVICES(phy))
	phyini := &PhyIni{
		PHY: phy,
		INI: i,
	}

	if e == nil {
		phyini.Class = i.Section("").Key("Class").String()
		phyini.Type = i.Section("").Key("Type").String()
		phyini.Polli = i.Section("").Key("Polli").MustInt()
	}

	return phyini
}

/* ********************************************************** */
// GilixIni
/* ********************************************************** */

var (
	gilixIni *GilixIni
	once     sync.Once
)

type GilixIni struct {
	INI *ini.File

	Trace *GilixTrace
	Rdc   map[string]*GilixRdc
	Acp   map[string]*GilixAcp
}

type GilixTrace struct {
	Path   string
	Mode   string
	Purge  int
	Lelvel string
}

type GilixRdc struct {
	Ip   string
	Port string
	Dial int
}

type GilixAcp struct {
	Net  string
	Addr string
}

func NewGilixIni() *GilixIni {
	once.Do(func() {
		i, e := ini.Load("./" + ExeNameNoExt() + ".ini")
		gilixIni = &GilixIni{
			INI:   i,
			Trace: &GilixTrace{},
			Rdc:   make(map[string]*GilixRdc),
			Acp:   make(map[string]*GilixAcp),
		}

		if e != nil {
			return
		}

		st := "trace"
		gilixIni.Trace = &GilixTrace{
			Path:   i.Section(st).Key("path").String(),
			Mode:   i.Section(st).Key("mode").String(),
			Purge:  i.Section(st).Key("purge").MustInt(),
			Lelvel: i.Section(st).Key("level").String(),
		}

		sr := "rdc"
		for _, s := range i.ChildSections(sr) {
			gilixIni.Rdc[IniChildRealName(s.Name(), sr)] = &GilixRdc{
				Ip:   s.Key("ip").String(),
				Port: s.Key("port").String(),
				Dial: s.Key("dial").MustInt(),
			}
		}

		sa := "acp"
		for _, s := range i.ChildSections(sa) {
			gilixIni.Acp[IniChildRealName(s.Name(), sa)] = &GilixAcp{
				Net:  s.Key("net").String(),
				Addr: s.Key("addr").String(),
			}
		}
	})

	return gilixIni
}
