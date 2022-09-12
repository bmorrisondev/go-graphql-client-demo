package main

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
)

func main() {
	client := graphql.NewClient("https://api.hashnode.com/graphql")
	queryData(client)
}

func queryData(client *graphql.Client) {
	query := `
	{
		user(username:"brianmmdev") {
			publication {
				posts {
					_id
					title
					dateAdded
				}
			}
		}
	}
	`
	request := graphql.NewRequest(query)
	var response QueryUserResponse
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}

	for _, el := range response.User.Publication.Posts {
		log.Println(el.Title)
	}
}

type QueryUserResponse struct {
	User struct {
		Publication struct {
			Posts []struct {
				Id        string `json:"_id"`
				Title     string
				dateAdded string
			}
		}
	}
}
