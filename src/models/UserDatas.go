package models

type ActivityDivideImage struct {
	Id       int64  `json:"id"`
	Category string `json:"category"`
	Url      string `json:"url"`
}

type ChannelInformaiton struct {
	Id             int64  `json:"id"`
	Time           string `json:"time"`
	Category       string `json:"category"`
	Location       string `json:"location"`
	Apply_location string `json:"apply_location"`
	Content        string `json:"content"`
	Number         string `json:"number"`
	Img            string `json:"img"`
	Title          string `json:"title"`
}

type ChannelForkInformationLite struct {
	Id             int64  `json:"id"`
	Divide         string `json:"divide"`
	Title          string `json:"title"`
	Apply_location string `json:"apply_location"`
	Img            string `json:"img"`
	Category       string `json:"category"`
}
type FolkData struct {
	Id             int64  `json:"id"`
	Time           string `json:"time"`
	Divide         string `json:"divide"`
	Category       string `json:"category"`
	Location       string `json:"location"`
	Apply_location string `json:"apply_location"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	Number         string `json:"number"`
	Img            string `json:"img"`
}
