package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"grpc-go-course/blog/blogpb"
)

func main() {

	log.Println("Starting gRPC client to interact with the Blog Service")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogClient(cc)

	// Create a Blog
	blog := &blogpb.Blog{
			Author: "Madhur",
			Title: "My first blog",
			Content: "This is my first blog",
		}

	crBlogReq := &blogpb.CreateBlogRequest{
			Blog: blog,
		}

	
	crBlogRes, err := c.CreateBlog(context.Background(), crBlogReq)
	if err != nil {
		log.Fatalf("Received error while creating a Blog: %v", err)
	}

	log.Printf("Blog created, ID: %v", crBlogRes.GetBlog().GetId())
}
