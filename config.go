package main

import "encoding/json"

type planet struct {
	PlanetID             string      `json:"PlanetIdentifier"`
	TypeFlag             int         `json:"TypeFlag"`
	PlanetaryMassJpt     json.Number `json:"PlanetaryMassJpt"`
	RadiusJpt            json.Number `json:"RadiusJpt"`
	PeriodDays           json.Number `json:"PeriodDays"`
	SemiMajorAxisAU      json.Number `json:"SemiMajorAxisAU"`
	Eccentricity         json.Number `json:"Eccentricity"`
	PeriastronDeg        json.Number `json:"PeriastronDeg"`
	LongitudeDeg         json.Number `json:"LongitudeDeg"`
	AscendingNodeDeg     json.Number `json:"AscendingNodeDeg"`
	InclinationDeg       json.Number `json:"InclinationDeg"`
	SurfaceTempK         json.Number `json:"SurfaceTempK"`
	AgeGyr               json.Number `json:"AgeGyr"`
	DiscoveryMethod      json.Number `json:"DiscoveryMethod"`
	DiscoveryYear        json.Number `json:"DiscoveryYear"`
	LastUpdated          json.Number `json:"LastUpdated"`
	RightAscension       json.Number `json:"RightAscension"`
	Declination          json.Number `json:"Declination"`
	DistFromSunParsec    json.Number `json:"DistFromSunParsec"`
	HostStarMassSlrMass  json.Number `json:"HostStarMassSlrMass"`
	HostStarRadiusSlrRad json.Number `json:"HostStarRadiusSlrRad"`
	HostStarMetallicity  json.Number `json:"HostStarMetallicity"`
	HostStarTempK        json.Number `json:"HostStarTempK"`
	HostStarAgeGyr       json.Number `json:"HostStarAgeGyr"`
	ListsPlanetIsOn      json.Number `json:"ListsPlanetIsOn"`
}

var Planets []planet
