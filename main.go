package main

import (
	"testing"

	"github.com/hashicorp/vault/builtin/credential/approle"
	vaulthttp "github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hashicorp/vault/vault"
)

func main() {
	t := &testing.T{}
	conf := &vault.CoreConfig{
		CredentialBackends: map[string]logical.Factory{
			"approle": approle.Factory,
		},
	}

	core, _, _ := vault.TestCoreUnsealedWithConfig(t, conf)
	vault.TestWaitActive(t, core)

	ln, _ := vaulthttp.TestServer(t, core)
	ln.Close()
}
