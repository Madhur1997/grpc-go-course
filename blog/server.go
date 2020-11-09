package main

import (
	 "context"
	 "fmt"
	 "log"
	 "os"
	 "os/signal"
	 "net"

	 // "gopkg.in/mgo.v2/bson"

	 "grpc-go-course/blog/blogpb"

	 "google.golang.org/grpc"
	 "google.golang.org/grpc/codes"
	 "google.golang.org/grpc/status"
	 "google.golang.org/grpc/reflection"

	 "go.mongodb.org/mongo-driver/bson/primitive"
	 "go.mongodb.org/mongo-driver/mongo"
	 "go.mongodb.org/mongo-driver/mongo/options"
)

type BlogServer struct {
	collection *mongo.Collection
}

type BlogItem struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string `bson:"author_id"`
	Title string `bson:"title"`
	Content string `bson:"content"`
}

func (blogServer *BlogServer) CreateBlog (ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	log.Printf("CreateBlog RPC invoked")

	blog := req.GetBlog()

	data := BlogItem{
			AuthorId: blog.GetAuthor(),
			Title: blog.GetTitle(),
			Content: blog.GetContent(),
		}

	res, err := blogServer.collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to ObjectID"))
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
				Id: oid.Hex(),
				Author: blog.GetAuthor(),
				Title: blog.GetTitle(),
				Content: blog.GetContent(),
			},
		}, nil
}

func main() {
	// if the server crashes, we get the file name and line number of crash.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("error received while trying to connect to mongo db server: %v", err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("MongoDB started")

	blogServer := &BlogServer{}
	blogServer.collection = client.Database("mydb").Collection("blog")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: ", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServer(s, blogServer)

	// Enable reflection on blog service
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

	// First we close MongoDB connection
	fmt.Println("Closing the MongoDB connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on trying to disconnect with MongoDB: %v", err)
	}

	// Now, close the listener
	fmt.Println("Closing the listener")
	if err := lis.Close(); err != nil {
		log.Fatalf("Error on closing the listener: %v", err)
	}

	// Now, stop the server.
	fmt.Println("Stopping the server")
	s.Stop()
}
