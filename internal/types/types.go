package types

import (
	"context"
	"io"

	"github.com/containers/image/v5/docker/reference"
	"github.com/containers/image/v5/types"
	publicTypes "github.com/containers/image/v5/types"
	"github.com/opencontainers/go-digest"
)

// ImageDestinationWithOptions is an internal extension to the ImageDestination
// interface.
type ImageDestinationWithOptions interface {
	publicTypes.ImageDestination

	// PutBlobWithOptions is a wrapper around PutBlob.  If
	// options.LayerIndex is set, the blob will be committed directly.
	// Either by the calling goroutine or by another goroutine already
	// committing layers.
	//
	// Please note that TryReusingBlobWithOptions and PutBlobWithOptions
	// *must* be used the together.  Mixing the two with non "WithOptions"
	// functions is not supported.
	PutBlobWithOptions(ctx context.Context, stream io.Reader, blobinfo publicTypes.BlobInfo, options PutBlobOptions) (publicTypes.BlobInfo, error)

	// TryReusingBlobWithOptions is a wrapper around TryReusingBlob.  If
	// options.LayerIndex is set, the reused blob will be recoreded as
	// already pulled.
	//
	// Please note that TryReusingBlobWithOptions and PutBlobWithOptions
	// *must* be used the together.  Mixing the two with non "WithOptions"
	// functions is not supported.
	TryReusingBlobWithOptions(ctx context.Context, blobinfo publicTypes.BlobInfo, options TryReusingBlobOptions) (bool, publicTypes.BlobInfo, error)
}

// PutBlobOptions are used in PutBlobWithOptions.
type PutBlobOptions struct {
	// Cache to look up blob infos.
	Cache publicTypes.BlobInfoCache
	// Denotes whether the blob is a config or not.
	IsConfig bool
	// Indicates an empty layer.
	EmptyLayer bool
	// The corresponding index in the layer slice.
	LayerIndex *int
}

// TryReusingBlobOptions are used in TryReusingBlobWithOptions.
type TryReusingBlobOptions struct {
	// Cache to look up blob infos.
	Cache publicTypes.BlobInfoCache
	// Use an equivalent of the desired blob.
	CanSubstitute bool
	// Indicates an empty layer.
	EmptyLayer bool
	// The corresponding index in the layer slice.
	LayerIndex *int
	// The reference of the image that contains the target blob.
	SrcRef reference.Named
}

// ImageSourceChunk is a portion of a blob.
// This API is experimental and can be changed without bumping the major version number.
type ImageSourceChunk struct {
	Offset uint64
	Length uint64
}

// ImageSourceSeekable is an image source that permits to fetch chunks of the entire blob.
// This API is experimental and can be changed without bumping the major version number.
type ImageSourceSeekable interface {
	// GetBlobAt returns a stream for the specified blob.
	// The specified chunks must be not overlapping and sorted by their offset.
	GetBlobAt(context.Context, publicTypes.BlobInfo, []ImageSourceChunk) (chan io.ReadCloser, chan error, error)
}

// ImageDestinationPartial is a service to store a blob by requesting the missing chunks to a ImageSourceSeekable.
// This API is experimental and can be changed without bumping the major version number.
type ImageDestinationPartial interface {
	// PutBlobPartial writes contents of stream and returns data representing the result.
	PutBlobPartial(ctx context.Context, stream ImageSourceSeekable, srcInfo publicTypes.BlobInfo, cache publicTypes.BlobInfoCache) (publicTypes.BlobInfo, error)
}

// BadPartialRequestError is returned by ImageSourceSeekable.GetBlobAt on an invalid request.
type BadPartialRequestError struct {
	Status string
}

func (e BadPartialRequestError) Error() string {
	return e.Status
}

type ImageDestinationSigstore interface {
	// SupportsSigstoreSignatures returns an error (to be displayed to the
	// user) if the destination certainly can't store Sigstore signatures.
	// Note: It is still possible for PutSignatures to fail if
	// SupportsSigstoreSignatures returns nil.
	SupportsSigstoreSignatures(ctx context.Context) error
	// PutSigstoreSignatures writes a set of signatures to the Sigstore image
	// destination.  If instanceDigest is not nil, it contains a digest of the
	// specific manifest instance to write or overwrite the signatures for
	// (when the primary manifest is a manifest list); this should always be
	// nil if the primary manifest is not a manifest list.  MUST be called
	// after PutManifest (signatures may reference manifest contents).
	PutSigstoreSignatures(ctx context.Context, signatures [][]byte, instanceDigest *digest.Digest, cache types.BlobInfoCache) error
}
