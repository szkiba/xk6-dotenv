package dotenv

import "go.k6.io/k6/js/modules"

type placeholder struct{}

func init() {
	modules.Register("k6/x/dotenv", &placeholder{})
}
