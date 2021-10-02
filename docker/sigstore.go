package docker

import (
	"encoding/base64"
	"net/url"
	"strings"

	"github.com/containers/image/v5/docker/reference"
	"github.com/containers/image/v5/types"
	digest "github.com/opencontainers/go-digest"
)

const (
	simpleSigningMediaType = "application/vnd.dev.cosign.simplesigning.v1+json"
	sigkey                 = "dev.cosignproject.cosign/signature"
)

func sigstoreSignatureURL(dstRef dockerReference, digest digest.Digest, scheme string) (*url.URL, error) {
	nameTagged, err := reference.WithTag(dstRef.ref, attachedImageTag(&digest))
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(scheme + nameTagged.Name())
	if err != nil {
		return nil, err
	}

	return url, nil
}

func attachedImageTag(digest *digest.Digest) string {
	// sha256:d34db33f -> sha256-d34db33f.suffix
	return strings.ReplaceAll(digest.String(), ":", "-") + ".sig"
}

func createBlobInfoForPayload(signature []byte) types.BlobInfo {
	return types.BlobInfo{
		Size: -1,
		Annotations: map[string]string{
			sigkey: base64.StdEncoding.EncodeToString(signature),
		},
		MediaType: simpleSigningMediaType,
	}
}

func createManifestForBlob() {
}
