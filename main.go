package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

var (
	url = flag.String("url", "https://gist.githubusercontent.com/joelbirchler/66cf8045fcbb6515557347c05d789b4a/raw/9a196385b44d4288431eef74896c0512bad3defe/exoplanets", "Using url to get data from exoplanet data")
)

func main() {
	flag.Parse()
	b, err := getData(*url)
	if err != nil {
		log.Fatalln(err)
	}

	err = getJson(b)
	if err != nil {
		log.Fatalln(err)
	}

	// First part of the assignment
	np := getNumberOfPlanet(Planets)
	log.Println("Number of orphan planets (no star) : ", np)

	// Second part of the assignment
	p, err := getPlanetInHotest(Planets)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("The planet Identifier of the planet orbiting the hottest star is '%+v'\n", p)
	}

	// Third part of the assignment
	gp, err := groupingPlanets(Planets)
	if err != nil {
		log.Fatalln(err)
	}
	gpSort := make([]int, 0, len(gp))
	for k := range gp {
		gpSort = append(gpSort, k)
	}
	sort.Ints(gpSort)
	for _, v := range gpSort {
		log.Printf("In %d we discovered %d small planets, %d medium planets, and %d large planets. ", v, gp[v][0], gp[v][1], gp[v][2])
	}
}

// groupingPlanets is for getting the number of planets discovered per year grouped by size
func groupingPlanets(p []planet) (map[int][]int, error) {
	groupPlanet := map[int][]int{}
	for _, v := range p {
		if v.RadiusJpt.String() == "" {
			continue
		}
		r, err := v.RadiusJpt.Float64()
		if err != nil {
			return map[int][]int{}, err
		}
		if r == 0 {
			continue
		}
		if v.DiscoveryYear.String() == "" {
			continue
		}
		y, err := v.DiscoveryYear.Float64()
		if err != nil {
			return map[int][]int{}, err
		}
		if y == 0 {
			continue
		}

		if r < 1 {
			if len(groupPlanet[int(y)]) == 0 {
				groupPlanet[int(y)] = []int{1, 0, 0}
			} else {
				groupPlanet[int(y)][0]++
			}
		} else if r < 2 {
			if len(groupPlanet[int(y)]) == 0 {
				groupPlanet[int(y)] = []int{0, 1, 0}
			} else {
				groupPlanet[int(y)][1]++
			}
		} else {
			if len(groupPlanet[int(y)]) == 0 {
				groupPlanet[int(y)] = []int{0, 0, 1}
			} else {
				groupPlanet[int(y)][2]++
			}
		}
	}
	return groupPlanet, nil
}

// getPlanetInHotest is for getting of the planet orbiting the hottest star
func getPlanetInHotest(p []planet) (string, error) {
	maxTempStar := []int64{0, 0}
	for k, v := range p {
		if v.HostStarTempK.String() == "" {
			continue
		}
		t, err := v.HostStarTempK.Float64()
		if err != nil {
			return "", fmt.Errorf("Error %v - %v", v.HostStarTempK, err)
		}
		if int64(t) > maxTempStar[1] {
			maxTempStar = []int64{int64(k), int64(t)}
		}
	}
	return p[int(maxTempStar[0])].PlanetID, nil
}

// getNumbetOrfetnPlanet is getting the total number of orphan planets (no star)
func getNumberOfPlanet(p []planet) (counter int) {
	for _, v := range p {
		if v.TypeFlag == 3 {
			counter++
		}
	}
	return counter
}

// getData is for getting bytes from the URL with Get request
func getData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// getJson is for input the json file from bytes into the planets
func getJson(b []byte) error {
	if !json.Valid(b) {
		return fmt.Errorf("The data is not valid %v", b)
	}
	if err := json.Unmarshal(b, &Planets); err != nil {
		return err
	}
	return nil
}
