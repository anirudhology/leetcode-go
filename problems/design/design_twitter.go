package design

import (
	"container/heap"
)

// Tweet represents a tweet with its ID and timestamp.
type Tweet struct {
	ID        int
	Timestamp int
}

// TwitterMaxHeap is a priority queue (max heap) for storing tweets by timestamp.
type TwitterMaxHeap []Tweet

func (h TwitterMaxHeap) Len() int            { return len(h) }
func (h TwitterMaxHeap) Less(i, j int) bool  { return h[i].Timestamp > h[j].Timestamp }
func (h TwitterMaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *TwitterMaxHeap) Push(x interface{}) { *h = append(*h, x.(Tweet)) }
func (h *TwitterMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	tweet := old[n-1]
	*h = old[0 : n-1]
	return tweet
}

// Twitter represents the Twitter service.
type Twitter struct {
	userTweets    map[int][]int            // userId -> list of tweetIds
	userFollowing map[int]map[int]struct{} // userId -> set of following userIds
	tweets        map[int]int              // tweetId -> timestamp
	timestamp     int
}

// Constructor initializes a new Twitter instance.
func TwitterConstructor() Twitter {
	return Twitter{
		userTweets:    make(map[int][]int),
		userFollowing: make(map[int]map[int]struct{}),
		tweets:        make(map[int]int),
		timestamp:     0,
	}
}

// PostTweet adds a tweet to the user's tweet list and records its timestamp.
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.userTweets[userId] = append(this.userTweets[userId], tweetId)
	this.timestamp++
	this.tweets[tweetId] = this.timestamp
}

// GetNewsFeed returns the 10 most recent tweetIds in the user's news feed.
func (this *Twitter) GetNewsFeed(userId int) []int {
	// Collect all the users the current user is following, including the user themselves.
	users := make(map[int]struct{})
	if followees, ok := this.userFollowing[userId]; ok {
		for followee := range followees {
			users[followee] = struct{}{}
		}
	}
	users[userId] = struct{}{}

	// Max heap to store the most recent tweets.
	h := &TwitterMaxHeap{}
	heap.Init(h)

	// Gather tweets from all the users in the news feed.
	for user := range users {
		if tweets, ok := this.userTweets[user]; ok {
			for i := len(tweets) - 1; i >= 0; i-- {
				heap.Push(h, Tweet{ID: tweets[i], Timestamp: this.tweets[tweets[i]]})
			}
		}
	}

	// Retrieve the 10 most recent tweets.
	var newsFeed []int
	for h.Len() > 0 && len(newsFeed) < 10 {
		newsFeed = append(newsFeed, heap.Pop(h).(Tweet).ID)
	}
	return newsFeed
}

// Follow adds followeeId to the followerId's following list.
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.userFollowing[followerId]; !ok {
		this.userFollowing[followerId] = make(map[int]struct{})
	}
	this.userFollowing[followerId][followeeId] = struct{}{}
}

// Unfollow removes followeeId from the followerId's following list.
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followees, ok := this.userFollowing[followerId]; ok {
		delete(followees, followeeId)
	}
}
