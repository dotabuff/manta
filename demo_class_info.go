package manta

import "github.com/dotabuff/manta/dota"

func (p *Parser) onCDemoClassInfo(classInfo *dota.CDemoClassInfo) {
	for _, class := range classInfo.GetClasses() {
		p.classInfo[int(class.GetClassId())] = class.GetNetworkName()
		if class.TableName != nil {
			PP(class)
		}
	}
}
