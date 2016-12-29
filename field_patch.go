package manta

type fieldPatch struct {
	minBuild uint32
	maxBuild uint32
	patch    func(f *field)
}

var fieldPatches = []fieldPatch{
	fieldPatch{0, 990, func(f *field) {
		switch f.varName {
		case
			"angExtraLocalAngles",
			"angLocalAngles",
			"m_angInitialAngles",
			"m_angRotation",
			"m_ragAngles",
			"m_vLightDirection":
			if f.serializerName == "CBodyComponentBaseAnimatingOverlay" {
				f.encoder = "qangle_pitch_yaw"
			} else {
				f.encoder = "QAngle"
			}

		case
			"dirPrimary",
			"localSound",
			"m_flElasticity",
			"m_location",
			"m_poolOrigin",
			"m_ragPos",
			"m_vecEndPos",
			"m_vecLadderDir",
			"m_vecPlayerMountPositionBottom",
			"m_vecPlayerMountPositionTop",
			"m_viewtarget",
			"m_WorldMaxs",
			"m_WorldMins",
			"origin",
			"vecLocalOrigin":
			f.encoder = "coord"
		}
	}},
	fieldPatch{1016, 1027, func(f *field) {
		switch f.varName {
		case
			"m_bItemWhiteList",
			"m_bWorldTreeState",
			"m_iPlayerIDsInControl",
			"m_iPlayerSteamID",
			"m_ulTeamBannerLogo",
			"m_ulTeamBaseLogo",
			"m_ulTeamLogo":
			f.encoder = "fixed64"
		}
	}},
	fieldPatch{0, 0, func(f *field) {
		switch f.varName {
		case "m_flSimulationTime", "m_flAnimTime":
			f.encoder = "simtime"
		}
	}},
}

func (p *fieldPatch) shouldApply(build uint32) bool {
	return build >= p.minBuild && build <= p.maxBuild
}
