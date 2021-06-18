package youtube

type youtubeVideoTooltipData struct {
	Title        string
	ChannelTitle string
	Duration     string
	PublishDate  string
	Views        string
	LikeCount    string
	DislikeCount string
}

type youtubeChannelTooltipData struct {
	Title        string
	PublishDate  string
	Subscribers  string
	Views        string
}

type ChannelId struct {
	id string
	channelType ChannelType
}
