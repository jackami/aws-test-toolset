package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:aws.String("us-east-1")},
	)

	// Create an SES session.
	svc := ses.New(sess)

	resp, err := svc.CreateTemplate(&ses.CreateTemplateInput{
		Template:&ses.Template{
			TemplateName:aws.String("my-template"),
			SubjectPart:aws.String("Welcome to system!"),
			TextPart:aws.String("test text part"),
			HtmlPart:aws.String("test html part"),
		},
	})

	fmt.Printf("%v-%v", resp, err)

	resp2, err := svc.GetTemplate(&ses.GetTemplateInput{TemplateName:aws.String("my-template")})

	fmt.Printf("%v-%v", resp2, err)

	resp3, err := svc.UpdateTemplate(&ses.UpdateTemplateInput{Template:&ses.Template{
		TemplateName:aws.String("my-template"),
		SubjectPart:aws.String("Welcome to system!"),
		TextPart:aws.String("test text part abc"),
		HtmlPart:aws.String("test html part"),
		},
	})

	fmt.Printf("%v-%v", resp3, err)

	resp4, err := svc.GetTemplate(&ses.GetTemplateInput{TemplateName:aws.String("my-template")})

	fmt.Printf("%v-%v", resp4, err)

	resp5, err := svc.DeleteTemplate(&ses.DeleteTemplateInput{TemplateName:aws.String("my-template")})

	fmt.Printf("%v-%v", resp5, err)
}
