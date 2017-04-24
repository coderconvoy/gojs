package gojs

import (
	"strings"

	"github.com/pkg/errors"
)

type Req struct {
	Assetable
}

func (r *Req) Asset(s string) ([]byte, error) {
	return r.Assetable.Asset(s) // todo, get with compile
}

func (r *Req) CompileRequires(s string) error {
	m := make(map[string]*ReqNode)
	return r.compileRequires(s, &m)
}

func (r *Req) compileRequires(s string, collected *map[string]*ReqNode) error {
	b, err := r.Assetable.Asset(s)
	if err != nil {
		return errors.Wrap(err, "Could not load Asset")
	}
	ss := string(b)
	resNode := &ReqNode{
		Map: collected,
	}

	d := 0
	for {
		od := d
		sh, d := line(ss[d:])
		if strings.HasPrefix(sh, "//") {
			resNode.Contents = append(resNode.Contents, []byte(sh+"\n")...)
		} else {
			ra, ok := parseRequire(sh)
			if !ok {
				resNode.Contents = append(resNode.Contents, b[od:]...)
				return nil
			}
			resNode.Reqs = append(resNode.Reqs, ra)
			r.compileRequires(ra.R, collected)

		}

		if d < len(ss) {
			return errors.New("No Content")
		}

	}

	return nil
}

func line(s string) (string, int) {
	rfnd := false
	for k, v := range s {
		if v == '\n' {
			if rfnd {
				return s[:k-1], k + 1
			}
			return s[:k], k + 1
		}
		if rfnd {
			return s[:k-1], k
		}
		if v == '\r' {
			rfnd = true
		}
	}
	return s, len(s)
}

//TODO - Check if contains require keyword, etc
func parseRequire(s string) (ReqAs, bool) {
	ident, fname := "", ""
	_, _ = ident, fname
	return ReqAs{}, false
}
