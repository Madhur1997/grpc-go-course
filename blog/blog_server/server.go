package main

import (
	 "context"
	 "fmt"
	 "log"
	 "os"
	 "os/signal"
	 "net"

	 "gopkg.in/mgo.v2/bson"

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
	Author string `bson:"author_id"`
	Title string `bson:"title"`
	Content string `bson:"content"`
}

func (blogServer *BlogServer) CreateBlog (ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	log.Printf("CreateBlog RPC invoked")

	blog := req.GetBlog()

	data := BlogItem{
			Author: blog.GetAuthor(),
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

func (blogServer *BlogServer) ReadBlog (ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	log.Printf("ReadBlog RPC invoked")

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot convert to ObjectId"))
	}
	blog := BlogItem{}
	filter := bson.M{"_id": oid}

	if err = blogServer.collection.FindOne(ctx, filter).Decode(&blog); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot find blog with request ID %v", err))
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog {
			Id: blog.ID.Hex(),
			Author: blog.Author,
			Title: blog.Title,
			Content: blog.Content,
		},
	}, nil
}

func (blogServer *BlogServer) UpdateBlog (ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	log.Printf("UpdateBlog RPC invoked")

	oid, err := primitive.ObjectIDFromHex(req.GetBlog().GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot convert to ObjectId"))
	}
	blog := BlogItem{}
	filter := bson.M{"_id": oid}

	if err := blogServer.collection.FindOne(ctx, filter).Decode(&blog); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot find blog with request ID %v", err))
	}

	blog.Author = req.GetBlog().GetAuthor()
	blog.Title = req.GetBlog().GetTitle()
	blog.Content = req.GetBlog().GetContent()

	if _, err := blogServer.collection.ReplaceOne(ctx, filter, &blog); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot update the object in MonogDB: %v", err))
	}

	return &blogpb.UpdateBlogResponse{
		Blog: &blogpb.Blog {
			Id: blog.ID.Hex(),
			Author: blog.Author,
			Title: blog.Title,
			Content: blog.Content,
		},
	}, nil
}

func (blogServer *BlogServer) DeleteBlog (ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	log.Printf("DeleteBlog RPC invoked")

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot convert to ObjectId"))
	}
	filter := bson.M{"_id": oid}

	res, err := blogServer.collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot delete object in MongoDB: %v", err))
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound,fmt.Sprintf("Cannot find blog in MongoDB: %v", err))
	}

	return &blogpb.DeleteBlogResponse{
			Id: blog.ID.Hex(),
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

	log.Printf("Connected to MongoDB instance")

	blogServer := &BlogServer{}
	blogServer.collection = client.Database("mydb").Collection("blog")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: ", err)
	}

	log.Printf("Listening on port 50051")

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
