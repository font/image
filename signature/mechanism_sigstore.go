package signature

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sigstore/cosign/pkg/cosign/fulcio"
	"github.com/sigstore/sigstore/pkg/signature"
	"github.com/sigstore/sigstore/pkg/tlog"
)

type sigstoreSigningMechanism struct {
	ctx         context.Context
	signer      signature.Signer
	cert, chain string
}

// newSigstoreSigningMechanism returns a new sigstore signing mechanism.
// The caller must call .Close() on the returned SigningMechanism.
func newSigstoreSigningMechanism() (SigstoreSigningMechanism, error) {
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

func (s *sigstoreSigningMechanism) InitSigner() error {
	fmt.Println("Generating ephemeral keys...")
	signer, err := fulcio.NewSigner(s.ctx, "")
	if err != nil {
		return errors.Wrap(err, "getting key from Fulcio")
	}
	s.signer = signer
	s.cert, s.chain = signer.Cert, signer.Chain
	return nil
}

// Sign creates a (non-detached) signature of input using keyIdentity.
// Fails with a SigningNotSupportedError if the mechanism does not support signing.
func (s *sigstoreSigningMechanism) Sign(payload []byte) ([]byte, []byte, error) {
	fmt.Println("Signing payload...")
	signature, signedVal, err := s.signer.Sign(s.ctx, payload)
	if err != nil {
		return nil, nil, errors.Wrap(err, "signing")
	}
	return signature, signedVal, err
}

func (s *sigstoreSigningMechanism) Upload(pemBytes, digest, signedMsg, payload []byte, rekorURL string) error {
	fmt.Println("Sending entry to transparency log")
	tlogEntry, err := tlog.UploadToRekor([]byte(s.cert), digest, signedMsg, tLogServer(), payload)
	if err != nil {
		return err
	}
	fmt.Println("Rekor entry successful. Index number: ", tlogEntry)
	return nil
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
