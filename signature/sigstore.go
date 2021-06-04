// Note: Consider the API unstable until the code supports at least three different image formats or transports.

package signature

import (
	"context"
	"os"
	"strings"

	digest "github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
)

const (
	serverEnv   = "REKOR_SERVER"
	rekorServer = "https://rekor.sigstore.dev"
)

// TODO: SignManifest returns a signature for manifest as the specified dockerReference,
// using mech and its keyless signing.
func SignManifest(ctx context.Context, manifestDigest digest.Digest, dockerReference string, mech SigstoreSigningMechanism) ([]byte, []byte, error) {
	sigPayload, err := newCosignSignature(manifestDigest, dockerReference).MarshalJSON()
	if err != nil {
		return nil, nil, errors.Wrap(err, "payload")
	}

	err = mech.InitSigner()
	if err != nil {
		return nil, nil, errors.Wrap(err, "init signer")
	}

	return mech.Sign(sigPayload)

	//sigRef := signatureImageTagForDigest(string(manifestDigest))

	//fmt.Println("Pushing signature to:", dockerReference)
	// TODO: push to destination image
	//manifestDigest
	//return sig.sign(mech, keyIdentity)
	//return nil
}

func signatureImageTagForDigest(digest string) string {
	// sha256:... -> sha256-...
	return strings.ReplaceAll(digest, ":", "-") + ".sig"
}

// TlogServer returns the name of the tlog server, can be overwritten via env var
func tLogServer() string {
	if s := os.Getenv(serverEnv); s != "" {
		return s
	}
	return rekorServer
}
