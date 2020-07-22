package alternator

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

type Retryer struct {
	Seeds []string
}

var _ = &aws.Config{
	Retryer: new(Retryer),
}

func (r *Retryer) MaxRetries() int {
	// todo
	return 0
}

func (r *Retryer) ShouldRetry(req *request.Request) bool {
	// todo
	return false
}

func (r *Retryer) RetryRules(req *request.Request) time.Duration {
	// todo
	return 0
}
