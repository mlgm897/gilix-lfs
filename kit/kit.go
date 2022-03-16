package kit

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lindorof/gilix/rdc"
	rdctcp "github.com/lindorof/gilix/rdc/tcp"
	"github.com/lindorof/gilix/util"
)

// CBS
func ZaptCfg() (path string, mode string, purge int, lelvel string) {
	gini := NewGilixIni()

	return gini.Trace.Path, gini.Trace.Mode, gini.Trace.Purge, gini.Trace.Lelvel
}

func CreateZapt(mod string, file string) *util.Zapt {
	gini := NewGilixIni()

	return util.CreateZapt(gini.Trace.Path, mod, file, gini.Trace.Mode, gini.Trace.Purge, gini.Trace.Lelvel)
}

func CreateRdc(phy string) rdc.Rdc {
	gini := NewGilixIni()

	addr := ""
	dial := 0
	if r, o := gini.Rdc[phy]; o {
		addr = r.Ip + ":" + r.Port
		dial = r.Dial
	}

	return rdctcp.CreateCaller(addr, time.Duration(dial)*time.Millisecond)
}

func IniChildRealName(n string, p string) string {
	return n[len(p+"."):]
}

func ExeNameNoExt() string {
	ex, er := os.Executable()
	if er != nil {
		return ""
	}

	dot := strings.LastIndex(ex, ".")
	sep := strings.LastIndex(ex, string(os.PathSeparator))
	if sep < 0 {
		return ""
	}
	if dot >= 0 && dot <= sep {
		return ""
	}

	if dot < 0 {
		dot = len(ex)
	}

	return ex[sep+1 : dot]
}

func ExeExt() string {
	ex, er := os.Executable()
	if er != nil {
		return ""
	}

	return filepath.Ext(ex)
}
