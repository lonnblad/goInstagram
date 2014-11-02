package endpoints

import (
	"encoding/json"
	"fmt"
	"goInstagram/structs"
	"log"
)

type ApiTag struct {
	Data structs.Tag `json:"data"`
}
type ApiTags struct {
	Data []structs.Tag `json:"data"`
}
type ApiMedia struct {
	Data []structs.Media `json:"data"`
}

func (e *Endpoint) GetTag(tagName string) (tag structs.Tag) {
	if tagName == "" {
		return structs.Tag{}
	}
	data := e.net.GetRequest(e.makeGetTagAddress(tagName))
	var apiTag ApiTag
	err := json.Unmarshal(data, &apiTag)
	if err != nil {
		log.Printf("%+v\n", err)
		return structs.Tag{}
	}
	return apiTag.Data
}

func (e *Endpoint) makeGetTagAddress(tag string) string {
	return e.apiURL + "tags/" + tag + "?client_id=" + e.clientId
}

func (e *Endpoint) GetTaggedMedia(
	tag string,
	count int,
	min_tag_id string,
	max_tag_id string,
) []structs.Media {
	if tag == "" {
		return []structs.Media{}
	}
	data := e.net.GetRequest(e.makeGetRecentMediaAddress(tag, count, min_tag_id, max_tag_id))
	var apiMedia ApiMedia
	err := json.Unmarshal(data, &apiMedia)
	if err != nil {
		log.Printf("%+v\n", err)
		return []structs.Media{}
	}
	return apiMedia.Data
}

func (e *Endpoint) makeGetRecentMediaAddress(
	tag string,
	count int,
	min_tag_id string,
	max_tag_id string,
) string {
	var c, min, max string
	tag = "tags/" + tag + "/media/recent?"

	if count > 0 {
		c = fmt.Sprintf("count=%d&", count)
	}

	if min_tag_id != "" {
		min = "min_tag_id=" + min_tag_id + "&"
	}

	if max_tag_id != "" {
		min = "max_tag_id=" + max_tag_id + "&"
	}

	return e.apiURL + tag + c + min + max + "client_id=" + e.clientId
}

func (e *Endpoint) SearchTag(query string) []structs.Tag {
	if query == "" {
		return []structs.Tag{}
	}
	var apiTags ApiTags

	data := e.net.GetRequest(e.makeSearchTagAddress(query))
	err := json.Unmarshal(data, &apiTags)
	if err != nil {
		log.Printf("%+v\n", err)
		return []structs.Tag{}
	}

	return apiTags.Data
}

func (e *Endpoint) makeSearchTagAddress(query string) string {
	return e.apiURL + "tags/search?q=" + query + "&client_id=" + e.clientId
}
