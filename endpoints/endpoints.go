package endpoints

import (
	"instagram/network"
	"instagram/structs"
)

type Instagram interface {
	GetTag(string) structs.Tag
	GetTaggedMedia(string, int, string, string) []structs.Media
	SearchTag(string) []structs.Tag
}

type Endpoint struct {
	net      network.NetworkInterface
	apiURL   string
	clientId string
}
