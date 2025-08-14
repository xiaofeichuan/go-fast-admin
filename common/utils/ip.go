package utils

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

//
//type GeoIP2Util struct{}
//
//var GeoIP2 = new(GeoIP2Util)

var (
	db *geoip2.Reader
)

// InitGeoIP2 初始化
func InitGeoIP2() {
	var err error
	db, err = geoip2.Open("resource/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
}

// GetIPInfo 获取Ip信息
func GetIPInfo(s string) string {
	if db == nil {
		InitGeoIP2()
	}
	ip := net.ParseIP(s)
	city, err := db.City(ip)
	if err == nil {
		return city.Country.Names["en"] + " " + city.City.Names["en"]
	}
	return ""
}
