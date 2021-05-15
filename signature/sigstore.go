// Note: Consider the API unstable until the code supports at least three different image formats or transports.

package signature

import (
	"context"

	digest "github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
	"github.com/sigstore/cosign/pkg/cosign/fulcio"
)

// SignManifest returns a signature for manifest as the specified dockerReference,
// using mech and keyIdentity.
func SignManifest(ctx context.Context, manifestDigest digest.Digest, dockerReference string, mech SigningMechanism) error {
	sigPayload, err := newCosignSignature(manifestDigest, dockerReference).MarshalJSON()
	if err != nil {
		return errors.Wrap(err, "payload")
	}

	signer, err := fulcio.NewSigner(ctx)
	if err != nil {
		return errors.Wrap(err, "getting key from Fulcio")
	}

	//sig, _, err := signer.Sign(ctx, sigPayload)
	_, _, err = signer.Sign(ctx, sigPayload)
	if err != nil {
		return errors.Wrap(err, "signing")
	}
	// TODO: push to destination image
	//manifestDigest
	//sig, _, err := signer.Sign(ctx,
	//return sig.sign(mech, keyIdentity)
	return nil
}
