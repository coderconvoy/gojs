package gojs

import "errors"

type Assetable interface {
	Asset(s string) ([]byte, error)
	AssetDir(s string) ([]string, error)
}

var Single *AssetGroup = &AssetGroup{
	G: []Assetable{
		able{
			ass:    Asset,
			assdir: AssetDir,
		},
	},
}

type AssetGroup struct {
	G []Assetable
}

func (ag *AssetGroup) Asset(s string) ([]byte, error) {
	for _, ab := range ag.G {
		res, err := ab.Asset(s)
		if err == nil {
			return res, nil
		}
	}
	return []byte{}, errors.New("Asset not found:" + s)
}

func (ag *AssetGroup) AssetDir(s string) ([]string, error) {
	res := []string{}
	for _, ab := range ag.G {
		plus, err := ab.AssetDir(s)
		if err == nil {
			res = append(res, plus...)
		}
	}
	return res, nil
}

type able struct {
	ass    func(string) ([]byte, error)
	assdir func(string) ([]string, error)
}

func (a able) Asset(s string) ([]byte, error) {
	return a.ass(s)
}

func (a able) AssetDir(s string) ([]strin
	ag.G = append(ag.G, able{ass, assdir})
}
