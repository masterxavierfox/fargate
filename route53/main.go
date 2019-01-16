package route53

//go:generate mockgen -package client -destination=mock/client/client.go github.com/masterxavierfox/fargate/route53 Client
//go:generate mockgen -package sdk -source ../vendor/github.com/aws/aws-sdk-go/service/route53/route53iface/interface.go -destination=mock/sdk/route53iface.go github.com/aws/aws-sdk-go/service/route53/route53iface Route53API

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
)

// Client represents a method for accessing Amazon Route 53.
type Client interface {
	CreateAlias(CreateAliasInput) (string, error)
	CreateResourceRecord(CreateResourceRecordInput) (string, error)
	ListHostedZones() (HostedZones, error)
}

// SDKClient implements access to Amazon Route 53 via the AWS SDK.
type SDKClient struct {
	client route53iface.Route53API
}

// New returns an SDKClient configured with the given session.
func New(sess *session.Session) SDKClient {
	return SDKClient{
		client: route53.New(sess),
	}
}
