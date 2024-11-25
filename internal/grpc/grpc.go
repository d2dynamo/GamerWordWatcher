package grpc

import (
	"context"
	"fmt"

	"GamerWordWatcher/internal/database"
	"GamerWordWatcher/internal/guild"
	"GamerWordWatcher/internal/message"
	messageRPC "GamerWordWatcher/proto/message"

	"google.golang.org/protobuf/proto"
)


type WatcherService struct {
	messageRPC.UnimplementedGamerWordWatcherServer
}

func (s *WatcherService) MessageCreated(ctx context.Context, in *messageRPC.TextMessage) (*messageRPC.WatcherResponse, error) {
	fmt.Printf("Received message: %v\n", in.Content)
	
	content := in.Content
	userId := in.UserId
	guildId := in.GuildId
	channelId := in.ChannelId

	msg := message.Message{
		content: in.content,
		userId: in.userId,
		guildId: in.guildId,
		channelId: in.channelId,
	}
	
	if proto.Equal()
}