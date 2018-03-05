package models

type FolkNews struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Time     string `json:"time"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Details  string `json:"details"`
}

type FolkNewsLite struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Time     string `json:"time"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Img      string `json:"img"`
}

type BottomNewsLite struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Time    string `json:"time"`
	Briefly string `json:"briefly"`
	Img     string `json:"img"`
}

type BottomNews struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Time    string `json:"time"`
	Briefly string `json:"briefly"`
	Img     string `json:"img"`
	Content string `json:"content"`
}

type MainPageSlideNews struct{
	ID int `json:"id"`
	Content string `json:"content"`
	Img string `json:"img"`
	Detail string `json:"detail"`
}