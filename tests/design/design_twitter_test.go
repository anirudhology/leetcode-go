package design_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/design"
	"github.com/stretchr/testify/assert"
)

func TestTwitter_PostTweet(t *testing.T) {
	twitter := design.TwitterConstructor()
	twitter.PostTweet(1, 101)
	feed := twitter.GetNewsFeed(1)
	assert.Equal(t, []int{101}, feed)
}

func TestTwitter_GetNewsFeed(t *testing.T) {
	twitter := design.TwitterConstructor()
	twitter.PostTweet(1, 101)
	twitter.PostTweet(2, 102)
	twitter.Follow(1, 2)
	feed := twitter.GetNewsFeed(1)
	assert.Equal(t, []int{102, 101}, feed)
}

func TestTwitter_FollowUnfollow(t *testing.T) {
	twitter := design.TwitterConstructor()
	twitter.PostTweet(1, 101)
	twitter.PostTweet(2, 102)
	twitter.Follow(1, 2)
	feed := twitter.GetNewsFeed(1)
	assert.Equal(t, []int{102, 101}, feed)
	twitter.Unfollow(1, 2)
	feed = twitter.GetNewsFeed(1)
	assert.Equal(t, []int{101}, feed)
}
