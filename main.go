package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"os"
)

//query to fetch PR's that don't hae a CLA signed
var query struct {
	Repository struct {
		PullRequests struct {
			Edges []struct {
				Node struct {
					Id     githubv4.ID
					Author struct {
						Login githubv4.String
					}
				}
			}
		} `graphql:"pullRequests(labels:[\"cncf-cla: no\"], first: 100, states:[OPEN])"`
	} `graphql:"repository(owner: \"kubernetes\", name: \"website\")"`
}

//publish the comment
var mutation struct {
	AddComment struct {
		CommentEdge struct {
			Node struct {
				Id githubv4.ID
			}
		}
		Subject struct {
			Id githubv4.ID
		}
	} `graphql:"addComment(input: $input)"`
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	bdy := "Hey there! @%s, looks like you haven't signed the CLA yet. " +
		"Could I please have you do that? https://github.com/kubernetes/community/blob/master/CLA.md"
	for _, v := range query.Repository.PullRequests.Edges {
		input := githubv4.AddCommentInput{
			SubjectID: v.Node.Id,
			Body:      githubv4.String(fmt.Sprintf(bdy, v.Node.Author.Login)),
		}
		err := client.Mutate(context.Background(), &mutation, input, nil)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
