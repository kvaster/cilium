// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts analyzing the specified Network Access Scope.
func (c *Client) StartNetworkInsightsAccessScopeAnalysis(ctx context.Context, params *StartNetworkInsightsAccessScopeAnalysisInput, optFns ...func(*Options)) (*StartNetworkInsightsAccessScopeAnalysisOutput, error) {
	if params == nil {
		params = &StartNetworkInsightsAccessScopeAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartNetworkInsightsAccessScopeAnalysis", params, optFns, c.addOperationStartNetworkInsightsAccessScopeAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartNetworkInsightsAccessScopeAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartNetworkInsightsAccessScopeAnalysisInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to ensure idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	//
	// This member is required.
	ClientToken *string

	// The ID of the Network Access Scope.
	//
	// This member is required.
	NetworkInsightsAccessScopeId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The tags to apply.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type StartNetworkInsightsAccessScopeAnalysisOutput struct {

	// The Network Access Scope analysis.
	NetworkInsightsAccessScopeAnalysis *types.NetworkInsightsAccessScopeAnalysis

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartNetworkInsightsAccessScopeAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpStartNetworkInsightsAccessScopeAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpStartNetworkInsightsAccessScopeAnalysis{}, middleware.After)
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
	if err = addIdempotencyToken_opStartNetworkInsightsAccessScopeAnalysisMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartNetworkInsightsAccessScopeAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartNetworkInsightsAccessScopeAnalysis(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartNetworkInsightsAccessScopeAnalysis struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartNetworkInsightsAccessScopeAnalysis) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartNetworkInsightsAccessScopeAnalysis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartNetworkInsightsAccessScopeAnalysisInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartNetworkInsightsAccessScopeAnalysisInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartNetworkInsightsAccessScopeAnalysisMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartNetworkInsightsAccessScopeAnalysis{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartNetworkInsightsAccessScopeAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "StartNetworkInsightsAccessScopeAnalysis",
	}
}
