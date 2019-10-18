package main

import (
	"encoding/json"
	"gopkg.in/h2non/gock.v1"
	"reflect"
	"testing"
	"time"
)

var (
	responseMsg = `[
	{
		"PlanetIdentifier": "KOI-1843.03",
		"TypeFlag": 0,
		"PlanetaryMassJpt": 0.0014,
		"RadiusJpt": 0.054,
		"PeriodDays": 0.1768913,
		"SemiMajorAxisAU": 0.0048,
		"Eccentricity": "",
		"PeriastronDeg": "",
		"LongitudeDeg": "",
		"AscendingNodeDeg": "",
		"InclinationDeg": 72,
		"SurfaceTempK": "",
		"AgeGyr": "",
		"DiscoveryMethod": "transit",
		"DiscoveryYear": 2012,
		"LastUpdated": "13/07/15",
		"RightAscension": "19 00 03.14",
		"Declination": "+40 13 14.7",
		"DistFromSunParsec": "",
		"HostStarMassSlrMass": 0.46,
		"HostStarRadiusSlrRad": 0.45,
		"HostStarMetallicity": 0,
		"HostStarTempK": 3584,
		"HostStarAgeGyr": ""
	},
	{
		"PlanetIdentifier": "KOI-1843.01",
		"TypeFlag": 0,
		"PlanetaryMassJpt": "",
		"RadiusJpt": 0.114,
		"PeriodDays": 4.194525,
		"SemiMajorAxisAU": 0.039,
		"Eccentricity": "",
		"PeriastronDeg": "",
		"LongitudeDeg": "",
		"AscendingNodeDeg": "",
		"InclinationDeg": 89.38,
		"SurfaceTempK": "",
		"AgeGyr": "",
		"DiscoveryMethod": "transit",
		"DiscoveryYear": "",
		"LastUpdated": "",
		"RightAscension": "19 00 03.14",
		"Declination": "+40 13 14.7",
		"DistFromSunParsec": "",
		"HostStarMassSlrMass": 0.46,
		"HostStarRadiusSlrRad": 0.45,
		"HostStarMetallicity": 0,
		"HostStarTempK": 3584,
		"HostStarAgeGyr": ""
	}]`
)

func Test_getPlanetInHotest(t *testing.T) {
	tests := []struct {
		name     string
		p        []planet
		expected string
		wantErr  bool
	}{
		{
			name: "Positive Test",
			p: []planet{
				{
					PlanetID:      "AAA 123",
					HostStarTempK: json.Number("20"),
				},
				{
					PlanetID:      "AAB 124",
					HostStarTempK: json.Number("50"),
				},
			},
			expected: "AAB 124",
			wantErr:  false,
		},
		{
			name: "Negative Test One - Wrong data type",
			p: []planet{
				{
					PlanetID:      "AAA 123",
					HostStarTempK: json.Number(20),
				},
				{
					PlanetID:      "AAB 124",
					HostStarTempK: json.Number(50),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := getPlanetInHotest(tt.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPlanetInHotest() error = %v, wantErr %v", err, tt.wantErr)
			}
			if p != tt.expected {
				t.Fatalf("Expected = %v | Got = %v", tt.expected, p)
			}
		})
	}
}

func Test_getNumbetOrfetnPlanet(t *testing.T) {
	tests := []struct {
		name        string
		p           []planet
		wantCounter int
	}{
		{
			name: "Positive Test for one",
			p: []planet{
				{
					PlanetID: "A1",
					TypeFlag: 1,
				},
				{
					PlanetID: "A2",
					TypeFlag: 2,
				},
				{
					PlanetID: "A3",
					TypeFlag: 3,
				},
			},
			wantCounter: 1,
		},
		{
			name: "Positive Test for two",
			p: []planet{
				{
					PlanetID: "A1",
					TypeFlag: 1,
				},
				{
					PlanetID: "A2",
					TypeFlag: 2,
				},
				{
					PlanetID: "A3",
					TypeFlag: 3,
				},
				{
					PlanetID: "A4",
					TypeFlag: 1,
				},
				{
					PlanetID: "A5",
					TypeFlag: 2,
				},
				{
					PlanetID: "A6",
					TypeFlag: 3,
				},
			},
			wantCounter: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCounter := getNumbetOrfetnPlanet(tt.p); gotCounter != tt.wantCounter {
				t.Errorf("getNumbetOrfetnPlanet() = %v, want %v", gotCounter, tt.wantCounter)
			}
		})
	}
}

func Test_getData(t *testing.T) {
	defer gock.Off()
	gock.New("http://localhost:8000").
		Get("/test").
		Reply(200).
		JSON(responseMsg)
	time.Sleep(time.Second)
	tests := []struct {
		name    string
		url     string
		want    []byte
		wantErr bool
	}{
		{
			name:    "Positive Test",
			url:     "http://localhost:8000/test",
			want:    []byte(responseMsg),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getData(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getJson(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		wantErr bool
	}{
		{
			name:    "positive test",
			b:       []byte(responseMsg),
			wantErr: false,
		},
		{
			name:    "negative test",
			b:       []byte(responseMsg + "Wrong"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getJson(tt.b); (err != nil) != tt.wantErr {
				t.Errorf("getJson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
