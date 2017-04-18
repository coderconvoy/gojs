package gojs

import "errors"

//Assetable (asset-able)the core interface that allows anything module to be added with relative ease.
type Assetable interface {
	Asset(s string) ([]byte, error)
	AssetDir(s string) ([]string, error)
}

//The Global object
var Single *AssetGroup = &AssetGroup{
	G: []Assetable{
		able{
			ass:    Asset,
			assdir: AssetDir,
		},
	},
}

//Asset Group, a collection of Assetables to allow multiple libraries access at the same time. Asset Group implements Assetable so they can be treed
type AssetGroup struct {
	G []Assetable
}

//Asset , loop assets until finding a match, error on no match
func (ag *AssetGroup) Asset(s string) ([]byte, error) {
	for _, ab := range ag.G {
		res, err := ab.Asset(s)
		if err == nil {
			return res, nil
		}
	}
	return []byte{}, errors.New("Asset not found:" + s)
}

//List all files in all internal Assetable directorys with a match
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

//able simple struct, holds 2 functions, normally clojures, but not necesarily
type able struct {
	ass    func(string) ([]byte, error)
	assdir func(string) ([]string, error)
}

func (a able) Asset(s string) ([]byte, error) {
	return a.ass(s)
}

func (a able) AssetDir(s string) ([]string, error) {
	return a.assdir(s)
}

//Add an Assetable to the group
func (ag *AssetGroup) Add(a Assetable) {
	ag.G = append(ag.G, a)
}

//AddFuncs, take two funcs, one an assetgrabber, and add them to the asset group as a new Assetable
func (ag *AssetGroup) AddFuncs(ass func(string) ([]byte, error), assdir func(string) ([]string, error)) {
	ag.G = append(ag.G, able{ass, assdir})
}
