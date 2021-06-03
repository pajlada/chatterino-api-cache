package youtube

import (
	"net/url"
	"path"
	"regexp"
	"strings"
)

type ChannelType string
const (
	// InvalidChannel channel isn't of a known type or doesn't exist
	InvalidChannel ChannelType = ""
	// UserChannel channel ID is a username
	UserChannel = "user"
	// IdentifierChannel channel uses the YouTube channel ID format (UC*)
	IdentifierChannel = "channel"
	// CustomChannel channel uses a custom URL and requires a Search call for the ID
	CustomChannel = "custom"
)

type ChannelId struct {
	id string
	channelType ChannelType
}

func getYoutubeVideoIDFromURL(url *url.URL) string {
	if strings.Contains(url.Path, "embed") {
		return path.Base(url.Path)
	}

	return url.Query().Get("v")
}

func getYoutubeVideoIDFromURL2(url *url.URL) string {
	return path.Base(url.Path)
}

func getChannelTypeFromString(channelType string) ChannelType  {
	switch channelType {
		case "c":
			return CustomChannel
		case "user":
			return UserChannel
		case "channel":
			return IdentifierChannel
	}

	return InvalidChannel
}

func constructCacheKeyFromChannelId(id ChannelId) string {
	return string(id.channelType) + ":" + id.id
}

func deconstructChannelIdFromCacheKey(cacheKey string) ChannelId  {
	splitKey := strings.Split(cacheKey, ":")

	if len(splitKey) < 2 {
		return ChannelId{id: "", channelType: InvalidChannel}
	}

	return ChannelId{id: splitKey[1], channelType: getChannelTypeFromString(splitKey[0])}
}

func getYoutubeChannelIdFromURL(url *url.URL) ChannelId {
	pattern, err := regexp.Compile(`(user|c(?:hannel)?)/([\w-]+)`)
	if err != nil {
		return ChannelId{id: "", channelType: InvalidChannel}
	}

	match := pattern.FindStringSubmatch(url.Path)
	if match == nil || len(match) < 3 {
		return ChannelId{id: "", channelType: InvalidChannel}
	}

	return ChannelId{id: match[2], channelType: getChannelTypeFromString(match[1])}
}
