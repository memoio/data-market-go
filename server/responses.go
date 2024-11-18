package server

import "time"

type uploadResponse struct {
	Id string `json:"id"`
}

type fileInfoResponse struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Owner         string    `json:"owner"`
	Data          []byte    `json:"data"`
	Price         int64     `json:"price"`
	Description   string    `json:"description"`
	Size          int64     `json:"size"`
	UpTime        time.Time `json:"upTime"`
	Visitors      int       `json:"visitors"`
	DownloadTimes int       `json:"downloadTimes"`
	Etag          string    `json:"etag"`
}

type ownerProductListResponse struct {
	Products []productInfo `json:"products"`
	ToTal    int           `json:"total"`
}

type productInfo struct {
	Name      string `json:"name"`
	Data      []byte `json:"data"`
	Price     int64  `json:"price"`
	IsCollect bool   `json:"isCollect"`
	IsUp      bool   `json:"isUp"`
}
