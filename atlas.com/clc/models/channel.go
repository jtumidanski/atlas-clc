package models

type Channel struct {
	worldId   byte
	channelId byte
	capacity  int
	ipAddress string
	port      uint16
}

func NewChannel(worldId byte, channelId byte, capacity int, ipAddress string, port uint16) *Channel {
	return &Channel{worldId, channelId, capacity, ipAddress, port}
}

func (c *Channel) WorldId() byte {
	return c.worldId
}

func (c *Channel) ChannelId() byte {
	return c.channelId
}

func (c *Channel) Capacity() int {
	return c.capacity
}

func (c *Channel) IpAddress() string {
	return c.ipAddress
}

func (c *Channel) Port() uint16 {
	return c.port
}

type ChannelLoad struct {
	channelId byte
	capacity  int
}

func NewChannelLoad(channelId byte, capacity int) *ChannelLoad {
	return &ChannelLoad{channelId, capacity}
}

func (cl *ChannelLoad) ChannelId() byte {
	return cl.channelId
}

func (cl *ChannelLoad) Capacity() int {
	return cl.capacity
}
