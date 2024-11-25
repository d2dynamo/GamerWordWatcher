package guild

import (
	"GamerWordWatcher/internal/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Guild struct {
	id        primitive.ObjectID `bson:"_id,omitempty"`
	discordId string             `bson:"discordId,omitempty""`
	name      string             `bson:"name,omitempty""`
	createdAt time.Time          `bson:"createdAt,omitempty""`
	updatedAt time.Time          `bson:"updatedAt,omitempty""`
}

type DiscordGuild interface {
	GetID() primitive.ObjectID
	GetDiscordID() string
}

type GuildRepository interface {
	FindAll() ([]*Guild, error)
	FindByOID(id primitive.ObjectID) (*Guild, error)
	FindByDID(discordId string) (*Guild, error)
}

func (g *Guild) GetID() primitive.ObjectID {
	return g.id
}

func (g *Guild) GetDiscordID() string {
	return g.discordId
}

func (g *Guild) FindAll() ([]*Guild, error) {
	guildColl, ctx, err := database.ConnectCollection("guilds")

	if err != nil {
		return nil, err
	}

	var guilds []*Guild

	cursor, err := guildColl.Find(*ctx, nil)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(*ctx)

	for cursor.Next(*ctx) {
		var guild Guild
		err := cursor.Decode(&guild)

		if err != nil {
			return nil, err
		}

		guilds = append(guilds, &guild)
	}

	return guilds, nil
}

func (g *Guild) FindByOID(id primitive.ObjectID) (*Guild, error) {
	guildColl, ctx, err := database.ConnectCollection("guilds")

	if err != nil {
		return nil, err
	}

	var guild Guild

	filter := bson.D{{"_id", id}}

	err = guildColl.FindOne(*ctx, filter).Decode(&guild)

	if err != nil {
		return nil, err
	}

	return &guild, nil
}

func (g *Guild) FindByDID(discordId string) (*Guild, error) {
	guildColl, ctx, err := database.ConnectCollection("guilds")
	if err != nil {
		return nil, err
	}

	var guild Guild

	filter := bson.D{{"discordId", discordId}}

	err = guildColl.FindOne(*ctx, filter).Decode(&guild)

	if err != nil {
		return nil, err
	}

	return &guild, nil
}
