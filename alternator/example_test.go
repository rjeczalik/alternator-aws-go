package alternator_test

import (
	"context"

	"github.com/rjeczalik/alternator-aws-go/alternator"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ExampleWithSeeds() {
	cfg := &aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("access-key", "secret-key", ""),
	}

	// The following line configures DynamoDB client to use Scylla Alternator cluster.
	cfg = alternator.WithSeeds(cfg, "172.0.1.11:8000", "172.0.2.11:8000", "172.0.3.11:8000")

	client := dynamodb.New(session.Must(session.NewSession(cfg)))
	_ = client
}

func ExampleWithSeedsContext() {
	// If DynamoDB client's lifetime is context-bound,
	// we are going to use WithSeedsContext instead.
	// It will clean up the alternator client once
	// the DynamoDB client is terminated.
	ctx := context.TODO()

	cfg := &aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("access-key", "secret-key", ""),
	}

	// The following line configures DynamoDB client to use Scylla Alternator cluster.
	cfg = alternator.WithSeedsContext(ctx, cfg, "172.0.1.11:8000", "172.0.2.11:8000", "172.0.3.11:8000")

	client := dynamodb.New(session.Must(session.NewSession(cfg)))
	_ = client
}
