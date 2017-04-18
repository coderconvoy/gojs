package gojs

import "strings"

type Req struct {
	Assetable
}

func (r *Req) Asset(s string) ([]byte, error) {
	return r.Assetable.Asset(s) // todo, get with compile
}

func (r *Req) CompileRequires(s string) ([]byte, error) {
	return r.compileRequires(s, make(map[string]bool))
}

func (r *Req) compileRequires(s string, collected *map[string]*ReqNode) error {
	b, err := r.Assetable.Asset(s)
	if err != nil {
		return b, nil
	}
	ss = string(b)
	resNode = &ReqNode{
		Map: collected,
	}

	content = []byte{}
	d := 0
	for {
		od := d
		sh, d := line(ss[d:])
		if strings.HasPrefix(sh, "//") {
			resNode.Contents = append(resNode.Contents, []byte(sh+"\n"))
		} else {
			ra, ok := parseRequire(sh)
			if !ok {
				resNode.Contents = append(resNode.Contents, b[od:])
				return nil
			}
			resNode.Reqs = append(resNode.Reqs, ra)
			r.compileRequires(ra.R)

		}

		if d < len(ss) {
			return errors.new("No Content")
		}

	}

	return []byte{}, nil
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

func parseRequire(s string) (ReqAs, bool) {
	ident := ""
	fname := ""

}
