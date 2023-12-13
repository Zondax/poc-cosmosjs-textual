package main

import (
	"context"
	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	"cosmossdk.io/x/tx/signing"
	"cosmossdk.io/x/tx/signing/aminojson"
	"cosmossdk.io/x/tx/signing/direct"
	"cosmossdk.io/x/tx/signing/directaux"
	"cosmossdk.io/x/tx/signing/textual"
)

// SignModeOptions are options for configuring the standard sign mode handler map.
type SignModeOptions struct {
	// Textual are options for SIGN_MODE_TEXTUAL
	Textual textual.SignModeOptions
	// DirectAux are options for SIGN_MODE_DIRECT_AUX
	DirectAux directaux.SignModeHandlerOptions
	// AminoJSON are options for SIGN_MODE_LEGACY_AMINO_JSON
	AminoJSON aminojson.SignModeHandlerOptions
}

// HandlerMap returns a sign mode handler map that Cosmos SDK apps can use out
// of the box to support all "standard" sign modes.
func (s SignModeOptions) HandlerMap() (*signing.HandlerMap, error) {
	txt, err := textual.NewSignModeHandler(s.Textual)
	if err != nil {
		return nil, err
	}

	directAux, err := directaux.NewSignModeHandler(s.DirectAux)
	if err != nil {
		return nil, err
	}

	aminoJSON := aminojson.NewSignModeHandler(s.AminoJSON)

	return signing.NewHandlerMap(
		direct.SignModeHandler{},
		txt,
		directAux,
		aminoJSON,
	), nil
}

func signTextual(ctx context.Context, signerData signing.SignerData, txData signing.TxData) ([]byte, error) {
	signerOptions := SignModeOptions{}
	signer, err := signerOptions.HandlerMap()
	if err != nil {
		return nil, err
	}

	sig, err := signer.GetSignBytes(ctx, signingv1beta1.SignMode_SIGN_MODE_TEXTUAL, signerData, txData)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
