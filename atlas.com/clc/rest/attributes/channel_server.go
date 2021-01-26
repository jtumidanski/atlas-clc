package attributes

type ChannelServerListDataContainer struct {
	Data []ChannelServerData `json:"data"`
}

type ChannelServerDataContainer struct {
	Data ChannelServerData `json:"data"`
}

type ChannelServerInputDataContainer struct {
	Data ChannelServerData `json:"data"`
}

type ChannelServerData struct {
	Id         string                  `json:"id"`
	Type       string                  `json:"type"`
	Attributes ChannelServerAttributes `json:"attributes"`
}

type ChannelServerAttributes struct {
	WorldId   byte   `json:"worldId"`
	ChannelId byte   `json:"channelId"`
	Capacity  int    `json:"capacity"`
	IpAddress string `json:"ipAddress"`
	Port      uint16 `json:"port"`
}
