package proto

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const GRPC_CLIENT_REF = "GRPC_CLIENT_REF"

var (
	globalConn   *grpc.ClientConn
	globalClient *FileProcessingClient
)

func InitGrpcClient() {
	grpcHost := os.Getenv("MOSAICS_SERVICE_SERVICE_HOST")
	if grpcHost == "" {
		grpcHost = "localhost"
	}
	grpcPort := os.Getenv("MOSAICS_SERVICE_SERVICE_PORT")
	if grpcPort == "" {
		grpcPort = "8081"
	}

	tlsOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcHost, grpcPort), tlsOption)
	if err != nil {
		log.Fatal("Could not establish connection to gRPC server: ", err)
	}

	globalConn = conn
	client := NewFileProcessingClient(globalConn)
	globalClient = &client
}

func GrpcRunFileProcess(c *gin.Context, filepath string) (string, error) {
	client := GetGrpcClient(c)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	repl, err := (*client).ProcessFile(ctx, &ProcessFileReq{Filepath: filepath})
	if err != nil {
		return "", err
	}
	return repl.Filepath, nil
}

func AddGrpcClient(c *gin.Context) {
	c.Set(GRPC_CLIENT_REF, globalClient)
}

func GetGrpcClient(c *gin.Context) *FileProcessingClient {
	conn, exists := c.Get(GRPC_CLIENT_REF)
	if !exists {
		log.Fatalf("gRPC connection not registered!")
	}
	return conn.(*FileProcessingClient)
}
