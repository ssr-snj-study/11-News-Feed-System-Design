package service

import (
	"context"
	"github.com/machinebox/graphql"
	"log"
	"webserver/internal/feed/entity"
	"webserver/internal/feed/repository"
)

type Follower struct {
	FollowerID         int         `json:"follower_id"`
	UserTbByFollowerID UserDetails `json:"userTbByFollowerId"`
	UserID             int         `json:"user_id"`
}

type UserDetails struct {
	Name       string      `json:"name"`
	DeviceByID interface{} `json:"deviceById"` // null 값을 받을 수 있도록 interface{} 사용
}

type FeedService struct {
	Repo repository.FeedRepository
}

func (s *FeedService) PostingFeed(req *entity.Req) (int, error) {
	postNo, err := s.Repo.PostFeed(req)
	if err != nil {
		return 0, err
	}
	return postNo, nil

}

func (s *FeedService) SendFeedToFollower(req *entity.Req) (int, error) {
	postNo, err := s.Repo.PostFeed(req)
	if err != nil {
		return 0, err
	}
	return postNo, nil
}

func getFollower(userId int) {
	client := graphql.NewClient("http://127.0.0.1/v1/graphql")
	req := graphql.NewRequest(`
query GetFollowersWithPagination {
  followers(where: {user_id: {_eq: $user_id}}) {
    follower_id
    userTbByFollowerId {
      name
      deviceById {
        user_id
        device_token
      }
    }
    user_id
  }
}
	`)

	// 변수 설정
	req.Var("user_id", 1)

	// 요청 실행
	var resp struct {
		Followers []Follower `json:"followers"`
	}

	err := client.Run(context.Background(), req, &resp)
	if err != nil {
		log.Fatalf("Failed to execute GraphQL query: %v", err)
	}

	log.Printf("Inserted Follower: %+v\n", resp.Followers)

}

func (s *FeedService) GetFeed(req *entity.Req) (*entity.Posting, error) {
	return s.Repo.GetFeed(req)
}
