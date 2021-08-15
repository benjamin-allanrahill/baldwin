package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"brdbth/graph/generated"
	"brdbth/graph/model"
	"context"
)

func (r *mutationResolver) CreateTweet(ctx context.Context, input *model.NewTweet) (*model.Tweet, error) {

	var tweet model.Tweet
	tweet.AuthorID = input.AuthorID
	tweet.Text = input.Text
	tweet.ID = "!#@@##"
	return &tweet, nil
}

func (r *queryResolver) Tweets(ctx context.Context) ([]*model.Tweet, error) {
	var tweets []*model.Tweet

	dummyTweet := model.Tweet{ID: "sfjahfdshalksdh", Text: "this is my dummy tweet", AuthorID: "@beeejar"}
	tweets = append(tweets, &dummyTweet)
	return tweets, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
