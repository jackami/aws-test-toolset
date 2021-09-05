package main

import (
	"encoding/base64"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/eks"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

func newClientset(cluster *eks.Cluster) (*kubernetes.Clientset, error) {
	log.Printf("-------------%+v", cluster)
	gen, err := token.NewGenerator(true, false)
	if err != nil {
		return nil, err
	}
	opts := &token.GetTokenOptions{
		ClusterID: aws.StringValue(cluster.Name),
	}
	tok, err := gen.GetWithOptions(opts)
	if err != nil {
		return nil, err
	}
	ca, err := base64.StdEncoding.DecodeString(aws.StringValue(cluster.CertificateAuthority.Data))
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(
		&rest.Config{
			Host:        aws.StringValue(cluster.Endpoint),
			BearerToken: tok.Token,
			TLSClientConfig: rest.TLSClientConfig{
				CAData: ca,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func main() {
	name := "test-eks-cluster"
	region := "us-east-1"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	eksSvc := eks.New(sess)

	log.Printf("-----region--------%+v", region)

	input := &eks.DescribeClusterInput{
		Name: aws.String(name),
	}

	log.Printf("-----input--------%+v", input)

	result, err := eksSvc.DescribeCluster(input)
	if err != nil {
		log.Fatalf("Error calling DescribeCluster: %v", err)
	}
	clientset, err := newClientset(result.Cluster)
	if err != nil {
		log.Fatalf("Error creating clientset: %v", err)
	}
	nodes, err := clientset.CoreV1().Nodes().List(nil, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error getting EKS nodes: %v", err)
	}
	log.Printf("There are %d nodes associated with cluster %s", len(nodes.Items), name)
}