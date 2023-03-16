//go:build tinygo || js || wasm

package scale

import (
	signature "github.com/loopholelabs/scale-signature-http"
	"math/rand"
	"strings"
)

func Scale(ctx *signature.Context) (*signature.Context, error) {
	names := strings.Split(string(ctx.Request().Body()), ",")
	ctx.Response().SetBody(names[rand.Intn(len(names))])
	return ctx.Next()
}
