package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/sts"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_UploadFile(t *testing.T) {
	// The session the S3 Uploader will use
	sess := session.Must(
		session.NewSession(&aws.Config{
			// Use the SDK's SharedCredentialsProvider directly instead of the
			// SDK's default credential chain. This ensures that the
			// application can call Config.Credentials.Expire. This  is counter
			// to the SDK's default credentials chain, which  will never reread
			// the shared credentials file.
			Credentials: credentials.NewCredentials(&credentials.SharedCredentialsProvider{
				Filename: defaults.SharedCredentialsFilename(),
				Profile:  "default",
			}),
			Region: aws.String(endpoints.UsEast1RegionID),
		}),
	)

	//credentials.NewChainCredentials(&credentials.StaticProvide)

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	var (
		filename string = "nodegroup_template.json"
		myBucket string = "test-bucket-rules"
		myString string = "nodegroup_template.json"
	)

	f, err  := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myString),
		Body:   f,
	})

	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
	} else {
		fmt.Printf("file uploaded to, %s\n", result.Location)
	}
}

func Test_PreSignedUrlGet(t *testing.T) {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("test-bucket-rules"),
		Key:    aws.String("nodegroup_template.json"),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	log.Println("The URL is", urlStr)
}

func Test_PreSignedUrlPut(t *testing.T) {
	//f, err  := os.Open("nodegroup_template.json")
	//if err != nil {
	//	fmt.Printf("failed to open file %q, %v", "nodegroup_template.json", err)
	//}

	h := md5.New()
	content := strings.NewReader("")
	content.WriteTo(h)

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	resp, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("test-bucket-rules"),
		Key:    aws.String("nodegroup_template.json"),
	})

	md5s := base64.StdEncoding.EncodeToString(h.Sum(nil))
	resp.HTTPRequest.Header.Set("Content-MD5", md5s)

	url, err := resp.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("error presigning request", err)
		return
	}

	log.Println("The URL is", url)

	req, err := http.NewRequest("PUT", url, strings.NewReader(""))
	req.Header.Set("Content-MD5", md5s)
	if err != nil {
		fmt.Println("error creating request", url)
		return
	}

	defClient, err := http.DefaultClient.Do(req)
	fmt.Println(defClient, err)
}

func Test_AssumeRoleToUploadS3(t *testing.T)  {
	sess := session.Must(
		session.NewSession(&aws.Config{
			// Use the SDK's SharedCredentialsProvider directly instead of the
			// SDK's default credential chain. This ensures that the
			// application can call Config.Credentials.Expire. This  is counter
			// to the SDK's default credentials chain, which  will never reread
			// the shared credentials file.
			Credentials: credentials.NewCredentials(&credentials.SharedCredentialsProvider{
				Filename: defaults.SharedCredentialsFilename(),
				Profile:  "default",
			}),
			Region: aws.String(endpoints.UsEast1RegionID),
		}),
	)

	svc := sts.New(sess)
	assumeRoleReq := &sts.AssumeRoleInput{
		ExternalId:aws.String("123456"),
		Policy:aws.String("{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"Stmt1\",\"Effect\":\"Allow\",\"Action\":[\"s3:ListAllMyBuckets\",\"s3:PutObject\"],\"Resource\":\"*\"}]}"),
		RoleArn:aws.String("arn:aws:iam::AccountID:role/TVMClientRole"),
		RoleSessionName:aws.String("testAssumeRoleSession"),
		DurationSeconds:aws.Int64(900),
	}

	assumeRoleResp, err := svc.AssumeRole(assumeRoleReq)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sts.ErrCodeMalformedPolicyDocumentException:
				fmt.Println(sts.ErrCodeMalformedPolicyDocumentException, aerr.Error())
			case sts.ErrCodePackedPolicyTooLargeException:
				fmt.Println(sts.ErrCodePackedPolicyTooLargeException, aerr.Error())
			case sts.ErrCodeRegionDisabledException:
				fmt.Println(sts.ErrCodeRegionDisabledException, aerr.Error())
			case sts.ErrCodeExpiredTokenException:
				fmt.Println(sts.ErrCodeExpiredTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(assumeRoleResp)

	sess2 := session.Must(
		session.NewSession(&aws.Config{
			Credentials: credentials.NewCredentials(&credentials.StaticProvider{credentials.Value{
				AccessKeyID:     *assumeRoleResp.Credentials.AccessKeyId,
				SecretAccessKey: *assumeRoleResp.Credentials.SecretAccessKey,
				SessionToken:    *assumeRoleResp.Credentials.SessionToken,
				ProviderName:    "test-provider",
			}}),
			Region: aws.String(endpoints.UsEast1RegionID),
		}),
	)

	svcc := s3.New(sess2)
	input := &s3.ListBucketsInput{}

	result, err := svcc.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

	uploader := s3manager.NewUploader(sess2)

	var (
		filename string = "nodegroup_template.json"
		myBucket string = "test-bucket-rules"
		myString string = "nodegroup_template.json"
	)

	f, err  := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	result2, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myString),
		Body:   f,
	})

	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
	} else {
		fmt.Printf("file uploaded to, %s\n", result2.Location)
	}
}