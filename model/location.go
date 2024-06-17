package model

import "encoding/xml"

type Location struct {
	XMLName      xml.Name `xml:"root"`
	TotalCount   string   `xml:"total_count"`
	PageCount    int      `xml:"page_count"`
	Page         string   `xml:"page"`
	ArticleCount string   `xml:"article_count"`
	Item         []struct {
		Subject         string `xml:"subject"`
		Category        string `xml:"category"`
		Groupe          string `xml:"groupe"`
		Sort            string `xml:"sort"`
		Situation       string `xml:"situation"`
		Define          string `xml:"define"`
		Historic        string `xml:"historic"`
		Organization    string `xml:"organization"`
		Person          string `xml:"person"`
		Content         string `xml:"content"`
		Reference       string `xml:"reference"`
		AddressThen     string `xml:"addressThen"`
		Address         string `xml:"address"`
		AddressRoadname string `xml:"addressRoadname"`
		Research        string `xml:"research"`
	} `xml:"item"`
}

// 지오코드 위도경도 데이터
type GeoCode struct {
	Point struct {
		X float64 `xml:"x"`
		Y float64 `xml:"y"`
	} `xml:"point"`
}

// History에 들어갈 데이터
type AtHistory struct {
	Name         string
	Tag          string
	Place        string
	Situation    string
	Organization string
	Person       string
	Content      string
	Reference    string
	Appraisal    string
	LocationName string
	Address      string
	AddressRoad  string
	Latitude     float64
	Longitude    float64
}

// Map에 들어갈 데이터
type AtMap struct {
	EventId      uint
	LocationName string
	Address      string
	AddressRoad  string
	Latitude     float64
	Longitude    float64
}
