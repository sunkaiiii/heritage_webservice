package models

type UserInfo struct {
	ID                     int    `json:"id"`
	UserName               string `json:"userName"`
	FocusNumber            int    `json:"focusNumber"`
	FansNumber             int    `json:"fansNumber"`
	Permission             int    `json:"permission"`
	FocusAndFansPermission int    `json:"focusAndFansPermission"`
}

type FollowInformation struct {
	FocusFansID  int    `json:"focusFansID"`
	FocusFocusID int    `json:"focusFocusID"`
	UserName     string `json:"userName"`
	Checked      bool   `json:"checked"`
}

type SearchInfo struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
}

type CollectionInfo struct {
	CollectionType string `json:"collectionType"`
	TypeID         int    `json:"typeID"`
}
