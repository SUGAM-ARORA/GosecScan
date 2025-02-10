package scan

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CheckS3Buckets() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return
	}
	client := s3.NewFromConfig(cfg)

	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Println("Error listing buckets:", err)
		return
	}

	for _, bucket := range result.Buckets {
		fmt.Printf("🔍 Scanning bucket: %s\n", *bucket.Name)
		checkBucketAccess(client, *bucket.Name)
	}
}

func checkBucketAccess(client *s3.Client, bucketName string) {
	// Check bucket policy
	_, err := client.GetBucketPolicy(context.TODO(), &s3.GetBucketPolicyInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Printf("⚠️  Warning: Unable to get bucket policy for %s\n", bucketName)
	}

	// Check bucket ACL
	acl, err := client.GetBucketAcl(context.TODO(), &s3.GetBucketAclInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Printf("⚠️  Warning: Unable to get bucket ACL for %s\n", bucketName)
		return
	}

	// Check for public access
	for _, grant := range acl.Grants {
		if grant.Grantee.URI != nil && *grant.Grantee.URI == "http://acs.amazonaws.com/groups/global/AllUsers" {
			fmt.Printf("🚨 Critical: Bucket %s has public access!\n", bucketName)
			return
		}
	}
}