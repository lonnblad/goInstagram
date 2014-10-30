package structs

type Tag struct {
	MediaCount int    `json:"media_count"`
	Name       string `json:"name"`
}

type Media struct {
	Type    string   `json:"type"`
	Tags    []string `json:"tags"`
	Caption Caption  `json:"caption"`
	Images  Images   `json:"images"`
	Id      string   `json:"id"`
}

type Caption struct {
	Text string `json:"text"`
}
type Images struct {
	LowRes Image `json:"low_resolution"`
	Thumb  Image `json:"thumbnail"`
	Normal Image `json:"standard_resolution"`
}
type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (a Tag) Equal(b Tag) bool {
	if a.Name != b.Name {
		return false
	}
	if a.MediaCount != b.MediaCount {
		return false
	}
	return true
}
