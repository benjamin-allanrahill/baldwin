package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"brdbth/auth"
	"brdbth/graph/generated"
	"brdbth/graph/model"
	"context"

	api "github.com/dghubble/go-twitter/twitter"
)

func (r *mutationResolver) CreateTweet(ctx context.Context, input *model.NewTweet) (*api.Tweet, error) {
	var tweet model.Tweet
	tweet.CreatedAt = "1232423"
	tweet.Text = input.Text
	return &tweet, nil
}

func (r *queryResolver) AllTweets(ctx context.Context) ([]*api.Tweet, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}
	tweets, _, err := client.Timelines.HomeTimeline(&api.HomeTimelineParams{})

	if err != nil {
		return nil, err
	}
	resp := make([]*api.Tweet, len(tweets))

	for i := range tweets {
		resp[i] = &tweets[i]
	}

	return resp, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
