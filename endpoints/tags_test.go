package endpoints

import (
	"instagram/network"
	"instagram/structs"
	"mock"
	"testing"
)

var networkMock mock.Mock
var endpoint *Endpoint

const clientId = "sleepy"
const apiURL = "https://api.instagram.com/v1/"

func init() {
	m := network.NetworkMock{}
	networkMock = &m
	endpoint = &Endpoint{&m, apiURL, clientId}
}

func Test_GetTag_EmptyTag(t *testing.T) {
	networkMock.InitMock()
	tag := endpoint.GetTag("")
	networkMock.ExpectedCallsCalled(t)
	result := structs.Tag{}
	if !tag.Equal(result) {
		t.Errorf("mismatch:\n\t[%+v]\n\t[%+v]", tag, result)
	}
}

func Test_GetTag(t *testing.T) {
	networkMock.InitMock()
	networkMock.ExpectCall("GetRequest",
		apiURL+"tags/alten?client_id="+clientId,
	)
	endpoint.GetTag("alten")
	networkMock.ExpectedCallsCalled(t)
}

func Test_GetTaggedMedia_EmptyTag(t *testing.T) {
	networkMock.InitMock()
	endpoint.GetTaggedMedia("", 0, "", "")
	networkMock.ExpectedCallsCalled(t)
}

func Test_GetTaggedMedia_Count_1(t *testing.T) {
	networkMock.InitMock()
	networkMock.ExpectCall("GetRequest",
		apiURL+"tags/alten/media/recent?count=1&client_id="+clientId,
	)
	endpoint.GetTaggedMedia("alten", 1, "", "")
	networkMock.ExpectedCallsCalled(t)
}

func Test_GetTaggedMedia_Min_TAG_ID_1(t *testing.T) {
	networkMock.InitMock()
	networkMock.ExpectCall("GetRequest",
		apiURL+"tags/alten/media/recent?min_tag_id=1&client_id="+clientId,
	)
	endpoint.GetTaggedMedia("alten", 0, "1", "")
	networkMock.ExpectedCallsCalled(t)
}

func Test_GetTaggedMedia_Max_TAG_ID_1(t *testing.T) {
	networkMock.InitMock()
	networkMock.ExpectCall("GetRequest",
		apiURL+"tags/alten/media/recent?max_tag_id=1&client_id="+clientId,
	)
	endpoint.GetTaggedMedia("alten", 0, "", "1")
	networkMock.ExpectedCallsCalled(t)
}

func Test_SearchTags_EmptyTag(t *testing.T) {
	networkMock.InitMock()
	endpoint.SearchTag("")
	networkMock.ExpectedCallsCalled(t)
}

func Test_SearchTags(t *testing.T) {
	networkMock.InitMock()
	networkMock.ExpectCall("GetRequest",
		apiURL+"tags/search?q=alten&client_id="+clientId,
	)
	endpoint.SearchTag("alten")
	networkMock.ExpectedCallsCalled(t)
}
