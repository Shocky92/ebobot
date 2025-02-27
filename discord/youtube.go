package discord

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func newYoutubeService(ctx context.Context, apiKey string) (*youtube.Service, error) {
	return youtube.NewService(ctx, option.WithAPIKey(apiKey))
}

func searchVideo(service *youtube.Service, query string) (*youtube.SearchListResponse, error) {
	call := service.Search.List([]string{"id", "snippet"}).
		Q(query).
		MaxResults(1)

	return call.Do()
}

func createURL(item *youtube.SearchResult) string {
	switch item.Id.Kind {
	case "youtube#channel":
		return "https://www.youtube.com/channel/" + item.Id.ChannelId
	default:
		return "https://www.youtube.com/watch?v=" + item.Id.VideoId
	}
}
