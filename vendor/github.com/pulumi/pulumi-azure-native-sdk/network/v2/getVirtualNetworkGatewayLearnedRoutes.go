// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This operation retrieves a list of routes the virtual network gateway has learned, including routes learned from BGP peers.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2016-09-01, 2019-08-01, 2023-04-01, 2023-05-01, 2023-06-01.
func GetVirtualNetworkGatewayLearnedRoutes(ctx *pulumi.Context, args *GetVirtualNetworkGatewayLearnedRoutesArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayLearnedRoutesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetVirtualNetworkGatewayLearnedRoutesResult
	err := ctx.Invoke("azure-native:network:getVirtualNetworkGatewayLearnedRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayLearnedRoutesArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual network gateway.
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}

// List of virtual network gateway routes.
type GetVirtualNetworkGatewayLearnedRoutesResult struct {
	// List of gateway routes.
	Value []GatewayRouteResponse `pulumi:"value"`
}

func GetVirtualNetworkGatewayLearnedRoutesOutput(ctx *pulumi.Context, args GetVirtualNetworkGatewayLearnedRoutesOutputArgs, opts ...pulumi.InvokeOption) GetVirtualNetworkGatewayLearnedRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualNetworkGatewayLearnedRoutesResult, error) {
			args := v.(GetVirtualNetworkGatewayLearnedRoutesArgs)
			r, err := GetVirtualNetworkGatewayLearnedRoutes(ctx, &args, opts...)
			var s GetVirtualNetworkGatewayLearnedRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualNetworkGatewayLearnedRoutesResultOutput)
}

type GetVirtualNetworkGatewayLearnedRoutesOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the virtual network gateway.
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (GetVirtualNetworkGatewayLearnedRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayLearnedRoutesArgs)(nil)).Elem()
}

// List of virtual network gateway routes.
type GetVirtualNetworkGatewayLearnedRoutesResultOutput struct{ *pulumi.OutputState }

func (GetVirtualNetworkGatewayLearnedRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayLearnedRoutesResult)(nil)).Elem()
}

func (o GetVirtualNetworkGatewayLearnedRoutesResultOutput) ToGetVirtualNetworkGatewayLearnedRoutesResultOutput() GetVirtualNetworkGatewayLearnedRoutesResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayLearnedRoutesResultOutput) ToGetVirtualNetworkGatewayLearnedRoutesResultOutputWithContext(ctx context.Context) GetVirtualNetworkGatewayLearnedRoutesResultOutput {
	return o
}

// List of gateway routes.
func (o GetVirtualNetworkGatewayLearnedRoutesResultOutput) Value() GatewayRouteResponseArrayOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayLearnedRoutesResult) []GatewayRouteResponse { return v.Value }).(GatewayRouteResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualNetworkGatewayLearnedRoutesResultOutput{})
}
