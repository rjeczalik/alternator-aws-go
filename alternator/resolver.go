package alternator

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
)

type Resolver struct {
	PollInterval time.Duration
	Seeds        []string

	once sync.Once
	ctx  context.Context
}

var _ = &aws.Config{
	EndpointResolver: new(Resolver),
}

func (r *Resolver) EndpointFor(service, region string, opts ...func(*endpoints.Options)) (e endpoints.ResolvedEndpoint, err error) {
	r.once.Do(r.init)

	switch service {
	case endpoints.DynamodbServiceID:
		// todo: select an alternator node for the given region
		return e, errors.New("not implemented")
	case endpoints.StreamsDynamodbServiceID:
		// todo: select an alternator node for the given region
		return e, errors.New("not implemented")
	default:
		return endpoints.DefaultResolver().EndpointFor(service, region, opts...)
	}
}

func (r *Resolver) pollInterval() time.Duration {
	if r.PollInterval != 0 {
		return r.PollInterval
	}
	return 60 * time.Second
}

func (r *Resolver) init() {
	ctx := r.ctx
	if ctx == nil {
		ctx = context.Background()
	}

	go r.cachenodes(ctx)
}

func (r *Resolver) cachenodes(ctx context.Context) {
	t := time.NewTicker(r.pollInterval())
	defer t.Stop()

	if len(r.Seeds) == 0 {
		panic("an atempt to use resolver with empty seed node list")
	}

	seeds := make([]string, len(r.Seeds))
	copy(seeds, r.Seeds)

	for {
		select {
		case <-t.C:
			// todo: maintain and update seed list (seeds)
			// todo: maintain and update node list (query seeds for /localnodes)
			// todo: multi-dc support?
		case <-ctx.Done():
			return
		}
	}
}
