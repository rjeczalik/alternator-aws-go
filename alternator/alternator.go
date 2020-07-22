package alternator

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
)

func WithSeeds(cfg *aws.Config, seeds ...string) *aws.Config {
	cfg.EndpointResolver = &Resolver{
		Seeds: seeds,
	}

	cfg.Retryer = &Retryer{
		Seeds: seeds,
	}

	return cfg

}

func WithSeedsContext(ctx context.Context, cfg *aws.Config, seeds ...string) *aws.Config {
	cfg.EndpointResolver = &Resolver{
		Seeds: seeds,
		ctx:   ctx,
	}

	cfg.Retryer = &Retryer{
		Seeds: seeds,
	}

	return cfg
}
