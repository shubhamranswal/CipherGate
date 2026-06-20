package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/shubhamranswal/ciphergate/internal/auth"
	"github.com/shubhamranswal/ciphergate/internal/mfa"
	"github.com/shubhamranswal/ciphergate/internal/user"

	"github.com/mdp/qrterminal/v3"
)

func Enable2FA(
	authCtx *auth.Context,
	userService *user.Service,
	mfaService *mfa.Service,
) {

	if !authCtx.IsAuthenticated() {

		fmt.Println(
			"❌ Not logged in",
		)

		return
	}

	if authCtx.User.MFAEnabled {

		fmt.Println(
			"❌ MFA already enabled",
		)

		return
	}

	key, err := mfaService.GenerateKey(
		authCtx.User.Username,
	)

	if err != nil {

		fmt.Printf(
			"❌ %v\n",
			err,
		)

		return
	}

	secret := key.Secret()

	fmt.Println(
		"\n🔐 MFA Setup",
	)

	fmt.Println(
		"\nScan this QR code using Microsoft Authenticator or Google Authenticator:\n",
	)

	qrterminal.GenerateHalfBlock(
		key.URL(),
		qrterminal.L,
		os.Stdout,
	)

	fmt.Printf(
		"\nSecret Key: %s\n",
		secret,
	)

	fmt.Println(
		"\nIf QR scanning fails, add the secret manually.",
	)

	code, err := readInput(
		"Enter current TOTP code: ",
	)

	if err != nil {

		fmt.Printf(
			"❌ %v\n",
			err,
		)

		return
	}

	if !mfaService.Validate(
		code,
		secret,
	) {

		fmt.Println(
			"❌ Invalid code",
		)

		return
	}

	authCtx.User.MFAEnabled = true
	authCtx.User.MFASecret = &secret

	err = userService.Update(
		context.Background(),
		authCtx.User,
	)

	if err != nil {

		fmt.Printf(
			"❌ %v\n",
			err,
		)

		return
	}

	fmt.Println(
		"\n✅ MFA enabled successfully",
	)
}
