package main

import (
	awsfargatev1 "buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/provider/aws/awsfargate/v1"
	"github.com/pkg/errors"
	"github.com/plantoncloud/aws-fargate-pulumi-module/pkg"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &awsfargatev1.AwsFargateStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
