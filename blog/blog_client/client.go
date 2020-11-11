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

	// Read a Blog
	rdBlogReq := &blogpb.ReadBlogRequest{Id: crBlogRes.GetBlog().GetId()}
			
	rdBlogRes, err := c.ReadBlog(context.Background(), rdBlogReq)
	if err != nil {
		log.Fatalf("Received error while reading a Blog: %v", err)
	}
	
	log.Printf("Blog read, ID: %v, Author: %s, Title: %s, Content: %s", rdBlogRes.GetBlog().GetId(), rdBlogRes.GetBlog().GetAuthor(), rdBlogRes.GetBlog().GetTitle(), rdBlogRes.GetBlog().GetContent())

	// Update a Blog
	updateBlogReq := &blogpb.UpdateBlogRequest{
				Blog: &blogpb.Blog{
					Id: crBlogRes.GetBlog().GetId(),
					Author: "Madhur Batra",
					Title: "My first blog",
					Content: "This is my first blog",
				},
			}
			
	updateBlogRes, err := c.UpdateBlog(context.Background(), updateBlogReq)
	if err != nil {
		log.Fatalf("Received error while reading a Blog: %v", err)
	}
	
	log.Printf("Blog Updated, ID: %v, Author: %s, Title: %s, Content: %s", updateBlogRes.GetBlog().GetId(), updateBlogRes.GetBlog().GetAuthor(), updateBlogRes.GetBlog().GetTitle(), updateBlogRes.GetBlog().GetContent())

	// delete Blog
	delBlogRes, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{Id: crBlogRes.GetBlog().GetId()})

	if err != nil {
		log.Printf("Error happened while deleting: %v \n", err)
	}
	log.Printf("Blog was deleted: %v \n", delBlogRes)
}
