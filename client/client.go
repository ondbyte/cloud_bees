package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ondbyte/cloud_bees/blog"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Start(port uint) error {
	cc, err := grpc.NewClient(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := blog.NewBlogServiceClient(cc)
	methods := map[string]func(client blog.BlogServiceClient){
		"create-post": createPost,
		"read-post":   readPost,
		"update-post": updatePost,
		"delete-post": deletePost,
	}
	availableMethods := []string{}
	for k, _ := range methods {
		availableMethods = append(availableMethods, k)
	}
	for {
		method := ""
		err := survey.AskOne(&survey.Select{
			Message: "select the method to run",
			Options: availableMethods,
		}, &method,
		)
		if err != nil {
			return err
		}

		if fn, ok := methods[method]; ok {
			fn(client)
		} else {
			log.Println("method", method, "doesnt exists", "select from", availableMethods)
		}
	}
}

func createPost(client blog.BlogServiceClient) {
	postData := ""
	err := survey.AskOne(
		&survey.Input{
			Message: "enter the post data in json format",
		}, &postData,
	)
	if err != nil {
		log.Fatalln(err)
	}
	req := &blog.CreatePostRequest{}
	err = json.Unmarshal([]byte(postData), req)
	if err != nil || postData == "" {
		ex, _ := json.Marshal(&blog.CreatePostRequest{
			Title:           "example title",
			Content:         "example content",
			Author:          "yadhunandan",
			PublicationDate: timestamppb.Now(),
			Tags:            []string{"thriller", "suspense"},
		})
		log.Fatalln("post data is required, for example: ", string(ex))
	}
	res, err := client.CreatePost(context.TODO(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("got response ", res.String())
}
func updatePost(client blog.BlogServiceClient) {
	postData := ""
	err := survey.AskOne(
		&survey.Input{
			Message: "enter the post data in json format to update the data(dont forget the id)",
		}, &postData,
	)
	if err != nil {
		log.Fatalln(err)
	}
	req := &blog.UpdatePostRequest{}
	err = json.Unmarshal([]byte(postData), req)
	if err != nil || postData == "" {
		ex, _ := json.Marshal(&blog.UpdatePostRequest{
			PostId:  "1",
			Title:   "example title",
			Content: "example content",
			Author:  "yadhunandan",
			Tags:    []string{"thriller", "suspense"},
		})
		log.Fatalln("post data is required, for example: ", string(ex))
	}
	
	res, err := client.UpdatePost(context.TODO(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("got response ", res.String())
}

func readPost(client blog.BlogServiceClient) {
	id := ""
	err := survey.AskOne(
		&survey.Input{
			Message: "enter the id of the post to get",
		}, &id,
	)
	if err != nil {
		log.Fatalln(err)
	}

	if id == "" {
		log.Fatalln("enter valid id")
	}
	req := &blog.ReadPostRequest{
		PostId: id,
	}
	res, err := client.ReadPost(context.TODO(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("got response ", res.String())
}
func deletePost(client blog.BlogServiceClient) {
	id := ""
	err := survey.AskOne(
		&survey.Input{
			Message: "enter the id of the post to delete",
		}, &id,
	)
	if err != nil {
		log.Fatalln(err)
	}

	if id == "" {
		log.Fatalln("enter valid id")
	}
	req := &blog.DeletePostRequest{
		PostId: id,
	}
	res, err := client.DeletePost(context.TODO(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("got response ", res.String())
}
