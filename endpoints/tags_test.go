package endpoints

import (
	"goInstagram/network"
	"goInstagram/structs"
	"testing"
)

var networkMock network.NetworkMockIf
var endpoint *Endpoint

const clientId = "sleepy"
const apiURL = "https://api.instagram.com/v1/"

func init() {
	m := network.NetworkMock{}
	networkMock = &m
	endpoint = &Endpoint{&m, apiURL, clientId}
}

func Test_GetTag_EmptyTag(t *testing.T) {
	networkMock.InitMock(t)
	tag := endpoint.GetTag("")
	networkMock.ExpectedCallsCalled()
	result := structs.Tag{}
	if !tag.Equal(result) {
		t.Errorf("mismatch:\n\t[%+v]\n\t[%+v]", tag, result)
	}
}

func Test_GetTag(t *testing.T) {
	networkMock.InitMock(t)
	arg := apiURL + "tags/alten?client_id=" + clientId
	ret := []byte{}
	networkMock.ExpectGetRequestToCall(arg, ret)
	endpoint.GetTag("alten")
	networkMock.ExpectedCallsCalled()
}

func Test_GetTaggedMedia_EmptyTag(t *testing.T) {
	networkMock.InitMock(t)
	endpoint.GetTaggedMedia("", 0, "", "")
	networkMock.ExpectedCallsCalled()
}

func Test_GetTaggedMedia_Count_1(t *testing.T) {
	networkMock.InitMock(t)
	arg := apiURL + "tags/alten/media/recent?count=1&client_id=" + clientId
	ret := []byte{}
	networkMock.ExpectGetRequestToCall(arg, ret)
	endpoint.GetTaggedMedia("alten", 1, "", "")
	networkMock.ExpectedCallsCalled()
}

func Test_GetTaggedMedia_Min_TAG_ID_1(t *testing.T) {
	networkMock.InitMock(t)
	arg := apiURL + "tags/alten/media/recent?min_tag_id=1&client_id=" + clientId
	ret := []byte{}
	networkMock.ExpectGetRequestToCall(arg, ret)
	endpoint.GetTaggedMedia("alten", 0, "1", "")
	networkMock.ExpectedCallsCalled()
}

func Test_GetTaggedMedia_Max_TAG_ID_1(t *testing.T) {
	networkMock.InitMock(t)
	arg := apiURL + "tags/alten/media/recent?max_tag_id=1&client_id=" + clientId
	ret := []byte{}
	networkMock.ExpectGetRequestToCall(arg, ret)
	endpoint.GetTaggedMedia("alten", 0, "", "1")
	networkMock.ExpectedCallsCalled()
}

func Test_SearchTags_EmptyTag(t *testing.T) {
	networkMock.InitMock(t)
	endpoint.SearchTag("")
	networkMock.ExpectedCallsCalled()
}

func Test_SearchTags(t *testing.T) {
	networkMock.InitMock(t)
	arg := apiURL + "tags/search?q=alten&client_id=" + clientId
	ret := []byte{}
	networkMock.ExpectGetRequestToCall(arg, ret)
	endpoint.SearchTag("alten")
	networkMock.ExpectedCallsCalled()
}
