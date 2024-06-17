package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"main/common"
	"main/external"
	"main/model"
	"strconv"
)

func main() {
	// https://search.i815.or.kr/openApiData.do?type=2&page=228
	res, err := external.NewCall("https://search.i815.or.kr", "/openApiData.do", map[string]string{
		"type": "2",
	}, nil, nil).Get()
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var result model.Location
	err = xml.Unmarshal(data, &result)
	// err = json.Unmarshal(data, result)
	if err != nil {
		panic(err)
	}

	// 페이지 개수 선언
	pageCount := result.PageCount
	var results []model.AtHistory

	// 페이지 개수 만큼 페이지를 바꿔가면서 api 호출
	for i := 1; i < pageCount; i++ {
		data, err = CallPerson(strconv.Itoa(i))
		if err != nil {
			return
		}
		var sliceData model.Location
		err = xml.Unmarshal(data, &sliceData)
		if err != nil {
			return
		}

		// 한 페이지에 들어있는 아이템(사건) 개수 만큼 반복하며 필요한 데이터 추출
		for _, openData := range sliceData.Item {
			//road : 위도 경도
			road := openData.AddressRoadname
			// 사건에 들어있는 주소를 위도경도 변환 api를 사용하여 위도 경도 데이터 추출
			// https://api.vworld.kr/req/address?service=address&request=getcoord&version=2.0&crs=epsg:4326&address=전라북도 순창군 순창읍 순창7길 40&refine=true&simple=false&format=json&type=road&key=5E69E1CE-B173-3F10-8221-3FD8A149CCC8
			res, err := external.NewCall("https://api.vworld.kr", "/req/address", map[string]string{
				"service": "address",
				"request": "getcoord",
				"version": "2.0",
				"crs":     "epsg:4326",
				"address": road,
				"refine":  "true",
				"simple":  "false",
				"format":  "json",
				"type":    "road",
				"key":     "5E69E1CE-B173-3F10-8221-3FD8A149CCC8",
			}, nil, nil).Get()

			if err != nil {
				return
			}

			locationData, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			// 지오코드에서 받은 데이터 담을 리소스
			var geoCode model.GeoCode
			err = json.Unmarshal(locationData, &geoCode)
			if err != nil {
				panic(err)
			}

			// history에 들어갈 데이터
			loc := model.AtHistory{
				Name:         openData.Subject,
				Tag:          openData.Groupe,
				Place:        openData.Define,
				Situation:    openData.Situation,
				Organization: openData.Organization,
				Person:       openData.Person,
				Content:      openData.Content,
				Appraisal:    openData.Research,
				Reference:    openData.Reference,
				Address:      openData.Address,
				AddressRoad:  openData.AddressRoadname,
				Longitude:    geoCode.Point.X,
				Latitude:     geoCode.Point.Y,
			}
			results = append(results, loc)

			// results history에 넘기기
			AtMap, err := common.SetEvents(results)
			if err != nil {
				return
			}
			// mapData map에 넘기기
			err = common.SetHistories(AtMap)
			if err != nil {
				return
			}
		}
	}

}

func CallPerson(page string) (data []byte, err error) {
	res, err := external.NewCall("https://search.i815.or.kr", "/openApiData.do", map[string]string{
		"type": "2",
		"page": page,
	}, nil, nil).Get()
	if err != nil {
		panic(err)

	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return
}
