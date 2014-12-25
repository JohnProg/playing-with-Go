package data

import "encoding/xml"

type Event struct {
	Id int64
	Name string
}

type User struct {
	Id int64
	UserName string
	UserEmail string
	UserPassword string
}

type Config struct {
	Domain   string
	BindAddr string
	Database struct {
		User string
		Host string
		Server string
		Name  string
		Password string
	}
}

type ExampleXml struct {
    XMLName xml.Name `xml:"example"`
    One     string   `xml:"one,attr"`
    Two     string   `xml:"two,attr"`
}