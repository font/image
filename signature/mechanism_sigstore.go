package signature

import (
	"context"
	"errors"
)

type sigstoreSigningMechanism struct {
	ctx context.Context
}

// newSigstoreSigningMechanism returns a new sigstore signing mechanism.
// The caller must call .Close() on the returned SigningMechanism.
func newSigstoreSigningMechanism() (SigningMechanism, error) {
	return &sigstoreSigningMechanism{
		ctx: context.Background(),
	}, nil
}

// Close removes resources associated with the mechanism, if any.
func (s *sigstoreSigningMechanism) Close() error {
	return nil
}

// SupportsSigning returns nil if the mechanism supports signing, or a SigningNotSupportedError.
func (s *sigstoreSigningMechanism) SupportsSigning() error {
	return nil

}

// Sign creates a (non-detached) signature of input using keyIdentity.
// Fails with a SigningNotSupportedError if the mechanism does not support signing.
func (s *sigstoreSigningMechanism) Sign(input []byte, keyIdentity string) ([]byte, error) {
	return nil, errors.New("not implemented yet") // TODO: Implement
}

// Verify parses unverifiedSignature and returns the content and the signer's identity
func (s *sigstoreSigningMechanism) Verify(unverifiedSignature []byte) (contents []byte, keyIdentity string, err error) {
	return nil, "", errors.New("not implemented yet")
}

// UntrustedSignatureContents returns UNTRUSTED contents of the signature WITHOUT ANY VERIFICATION,
// along with a short identifier of the key used for signing.
// WARNING: The short key identifier (which corresponds to "Key ID" for OpenPGP keys)
// is NOT the same as a "key identity" used in other calls to this interface, and
// the values may have no recognizable relationship if the public key is not available.
func (s *sigstoreSigningMechanism) UntrustedSignatureContents(untrustedSignature []byte) (untrustedContents []byte, shortKeyIdentifier string, err error) {
	return nil, "", errors.New("not implemented")
}
