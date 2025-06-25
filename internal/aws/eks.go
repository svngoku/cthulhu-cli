package aws

import (
	"context"

	"cthulhu-cli/pkg/utils"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func CreateEKS(ctx context.Context, name string) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}
	svc := eks.NewFromConfig(cfg)
	_, err = svc.CreateCluster(ctx, &eks.CreateClusterInput{
		Name: &name,
		// Additional params: roleArn, resourcesVpcConfig, etc.
	})
	if err == nil {
		utils.Log().Infow("creating EKS cluster", "name", name)
	}
	return err
}
