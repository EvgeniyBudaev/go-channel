package entity

type HubContent struct {
	Message      string `json:"message"`
	Type         string `json:"type"`
	UserId       uint64 `json:"userId"`
	UserImageUrl string `json:"userImageUrl"`
	Username     string `json:"username"`
}

type Hub struct {
	Broadcast chan *HubContent
}

func NewHub() *Hub {
	return &Hub{
		Broadcast: make(chan *HubContent, 5),
	}
}
