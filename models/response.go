package models

type Response struct {
	Kind          string   `json:"kind"`
	Etag          string   `json:"etag"`
	NextPageToken string   `json:"nextPageToken"`
	RegionCode    string   `json:"regionCode"`
	PageInfo      PageInfo `json:"pageInfo"`
	Items         []Items  `json:"items"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type Items struct {
	Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	Id      Id      `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type Id struct {
	Kind    string `json:"kind"`
	VideoId string `json:"videoId"`
}

type Snippet struct {
	PublishedAt          string     `json:"publishedAt"`
	ChannelId            string     `json:"channelId"`
	Title                string     `json:"title"`
	Description          string     `json:"description"`
	Thumbnails           Thumbnails `json:"thumbnails"`
	ChannelTitle         string     `json:"channelTitle"`
	LiveBroadcastContent string     `json:"liveBroadcastContent"`
}

type Thumbnails struct {
	Default ThumbnailSize `json:"default"`
	Medium  ThumbnailSize `json:"medium"`
	High    ThumbnailSize `json:"high"`
}

type ThumbnailSize struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
