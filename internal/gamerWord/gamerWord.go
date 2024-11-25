package gamerWord

import (
	"strings"
	"time"

	"GamerWordWatcher/internal/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GamerWord struct {
	id        primitive.ObjectID `bson:"_id"`
	word      string             `bson:"word"`
	phrases   []string           `bson:"phrases"`
	createdAt time.Time          `bson:"createdAt"`
	updatedAt time.Time          `bson:"updatedAt"`
}

type GamerWordRepository interface {
	FindAll() ([]*GamerWord, error)
	FindByOID(id primitive.ObjectID) (*GamerWord, error)
}

func (g *GamerWord) GetID() primitive.ObjectID {
	return g.id
}

func (g *GamerWord) GetWord() string {
	return g.word
}

func (g *GamerWord) GetPhrases() []string {
	return g.phrases
}

func (g *GamerWord) FindAll() ([]*GamerWord, error) {
	gamerWordColl, ctx, err := database.ConnectCollection("gamerWords")
	if err != nil {
		return nil, err
	}

	var gamerWords []*GamerWord

	cursor, err := gamerWordColl.Find(*ctx, nil)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(*ctx)

	for cursor.Next(*ctx) {
		var gamerWord GamerWord
		err := cursor.Decode(&gamerWord)

		if err != nil {
			return nil, err
		}

		gamerWords = append(gamerWords, &gamerWord)
	}

	return gamerWords, nil
}

func (g *GamerWord) FindByOID(id primitive.ObjectID) (*GamerWord, error) {
	gamerWordColl, ctx, err := database.ConnectCollection("gamerWords")
	if err != nil {
		return nil, err
	}

	filter := map[string]interface{}{
		"_id": id,
	}

	var gamerWord GamerWord

	err = gamerWordColl.FindOne(*ctx, filter).Decode(&gamerWord)
	if err != nil {
		return nil, err
	}

	return &gamerWord, nil
}

func (g *GamerWord) CheckMatch(content *string) bool {
	for _, phrase := range g.phrases {
		var contentWords []string

		var wordBuffer string
		for _, r := range *content {
			if r == ' ' {
				contentWords = append(contentWords, wordBuffer)
				wordBuffer = ""
			} else {
				wordBuffer += string(r)
			}
		}

		for _, word := range contentWords {
			if strings.EqualFold(word, phrase) {
				return true
			}
		}
	}

	return false
}
