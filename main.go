package main

import (
	"fmt"
	"os"

	"github.com/nicklaw5/helix"
)

func main() {
	clientId := os.Getenv("TWITCH_CLIENT_ID")
	accessToken := os.Getenv("TWITCH_ACCESS_TOKEN")
	client, err := helix.NewClient(&helix.Options{
		ClientID: clientId,
		UserAccessToken: accessToken,
	})


	if err != nil {
		// handle error
	}

	usersResponse, err := client.GetUsers(&helix.UsersParams{
		Logins: []string{"ShizukaSakuya"},
	})
	if err != nil {
		// handle error
	}

	fmt.Printf("%+v\n", usersResponse)
	id := usersResponse.Data.Users[0].ID

	fmt.Printf("%+v\n", id)

	rewardResponse, err := client.CreateCustomReward(&helix.ChannelCustomRewardsParams{
		BroadcasterID: id,
		Title:         "game analysis 1v1",
		Cost:          50000,
	})
	if err != nil {
		// handle error
	}

	fmt.Printf("%+v\n", rewardResponse)

}
