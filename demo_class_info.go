package manta

import (
	"github.com/dotabuff/manta/dota"
)

func (p *Parser) onCDemoClassInfo(classInfo *dota.CDemoClassInfo) error {
	for _, class := range classInfo.GetClasses() {
		p.classInfo[class.GetClassId()] = class.GetNetworkName()
	}
	return nil
}
