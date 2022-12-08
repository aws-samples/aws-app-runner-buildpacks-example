# AWS App Runner with Cloud Native Buildpacks

Pipeline that demonstrates how to integrate Cloud Native Buildpacks with AWS App Runner.

This pipeline uses AWS DevOps tools AWS CodeCommit, AWS CodePipeline, AWS CodeBuild and Amazon Elastic Container Registry (ECR) along with other AWS services.

The file `cloudformation.yml` contains all of the AWS CloudFormation necessary to deploy the solution, including:
- CodeCommit repository
- ECR repository
- CodeBuild project to build the container image
- CloudFormation to deploy the container image as an App Runner service
- CodePipeline to orchestrate code checkout and running CodeBuild
- Lamba function and custom resource to seed the CodeCommit repository with sample applications

## Deploying pipeline:
Download the CloudFormation template and pipeline code from GitHub repo.

1.	Log in to your AWS account if you have not done so already. 
2.	On the CloudFormation console, choose Create Stack. 
3.	Choose the provided CloudFormation pipeline template. 
4.	Choose Next.

Alternatively, use the `aws` CLI as follows:

```
aws cloudformation deploy --template-file cloudformation.yml \
  --stack-name apprunner-buildpacks-sample \
  --capabilities CAPABILITY_NAMED_IAM
```

## License

Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
SPDX-License-Identifier: MIT-0
