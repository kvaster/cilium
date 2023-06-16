// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates an IAM role from an Certificate Manager (ACM) certificate.
// Disassociating an IAM role from an ACM certificate removes the Amazon S3 object
// that contains the certificate, certificate chain, and encrypted private key from
// the Amazon S3 bucket. It also revokes the IAM role's permission to use the KMS
// key used to encrypt the private key. This effectively revokes the role's
// permission to use the certificate.
func (c *Client) DisassociateEnclaveCertificateIamRole(ctx context.Context, params *DisassociateEnclaveCertificateIamRoleInput, optFns ...func(*Options)) (*DisassociateEnclaveCertificateIamRoleOutput, error) {
	if params == nil {
		params = &DisassociateEnclaveCertificateIamRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateEnclaveCertificateIamRole", params, optFns, c.addOperationDisassociateEnclaveCertificateIamRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateEnclaveCertificateIamRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateEnclaveCertificateIamRoleInput struct {

	// The ARN of the ACM certificate from which to disassociate the IAM role.
	//
	// This member is required.
	CertificateArn *string

	// The ARN of the IAM role to disassociate.
	//
	// This member is required.
	RoleArn *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type DisassociateEnclaveCertificateIamRoleOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateEnclaveCertificateIamRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDisassociateEnclaveCertificateIamRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDisassociateEnclaveCertificateIamRole{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDisassociateEnclaveCertificateIamRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateEnclaveCertificateIamRole(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDisassociateEnclaveCertificateIamRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisassociateEnclaveCertificateIamRole",
	}
}
