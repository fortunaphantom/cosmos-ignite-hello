package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hello/x/hello/types"
)

func (k Keeper) Hello(goCtx context.Context, req *types.QueryHelloRequest) (*types.QueryHelloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryHelloResponse{Text: "Hello, Ignite CLI!"}, nil // <--
	// return &types.QueryHelloResponse{}, nil
}
