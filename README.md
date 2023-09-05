# vault-issue

Reproduces an issue with upgrading the vault module and sdk.

## Steps to reproduce

1. Clone this repo
2. Run `asdf install` (if you're using asdf, or install the version in `.tool-versions`).
3. Run `go run main.go`

Should return something like this:

```bash
2023-09-05T15:26:24.172+0700 [INFO]  creating error injector
2023-09-05T15:26:24.172+0700 [DEBUG] core: set config: sanitized config="{\"administrative_namespace_path\":\"\",\"api_addr\":\"\",\"cache_size\":0,\"cluster_addr\":\"\",\"cluster_cipher_suites\":\"\",\"cluster_name\":\"\",\"default_lease_ttl\":0,\"default_max_request_duration\":0,\"detect_deadlocks\":\"\",\"disable_cache\":false,\"disable_clustering\":false,\"disable_indexing\":false,\"disable_mlock\":false,\"disable_performance_standby\":false,\"disable_printable_check\":false,\"disable_sealwrap\":false,\"disable_sentinel_trace\":false,\"enable_response_header_hostname\":false,\"enable_response_header_raft_node_id\":false,\"enable_ui\":false,\"experiments\":null,\"introspection_endpoint\":false,\"log_format\":\"\",\"log_level\":\"\",\"log_requests_level\":\"\",\"max_lease_ttl\":0,\"pid_file\":\"\",\"plugin_directory\":\"\",\"plugin_file_permissions\":0,\"plugin_file_uid\":0,\"raw_storage_endpoint\":false}"
2023-09-05T15:26:24.172+0700 [DEBUG] storage.cache: creating LRU cache: size=0
2023-09-05T15:26:24.172+0700 [INFO]  core: Initializing version history cache for core
2023-09-05T15:26:24.172+0700 [INFO]  events: Starting event system
2023-09-05T15:26:24.172+0700 [INFO]  core: security barrier not initialized
2023-09-05T15:26:24.172+0700 [INFO]  core: security barrier initialized: stored=1 shares=3 threshold=3
2023-09-05T15:26:24.173+0700 [DEBUG] core: cluster name not found/set, generating new
2023-09-05T15:26:24.173+0700 [DEBUG] core: cluster name set: name=vault-cluster-9ed8561c
2023-09-05T15:26:24.173+0700 [DEBUG] core: cluster ID not found, generating new
2023-09-05T15:26:24.173+0700 [DEBUG] core: cluster ID set: id=55f12947-47de-93b2-6dae-b2d254080186
2023-09-05T15:26:24.173+0700 [INFO]  core: post-unseal setup starting
2023-09-05T15:26:24.173+0700 [DEBUG] core: clearing forwarding clients
2023-09-05T15:26:24.173+0700 [DEBUG] core: done clearing forwarding clients
2023-09-05T15:26:24.173+0700 [DEBUG] core: persisting feature flags
2023-09-05T15:26:24.177+0700 [INFO]  core: loaded wrapping token key
2023-09-05T15:26:24.177+0700 [INFO]  core: successfully setup plugin catalog: plugin-directory=""
2023-09-05T15:26:24.177+0700 [INFO]  core: no mounts; adding default mount table
2023-09-05T15:26:24.178+0700 [TRACE] core: adding write forwarded paths: paths=[]
2023-09-05T15:26:24.178+0700 [INFO]  core: successfully mounted: type=cubbyhole version="v1.14.2+builtin.vault" path=cubbyhole/ namespace="ID: root. Path: "
2023-09-05T15:26:24.178+0700 [TRACE] core: adding write forwarded paths: paths=[]
2023-09-05T15:26:24.178+0700 [INFO]  core: successfully mounted: type=system version="v1.14.2+builtin.vault" path=sys/ namespace="ID: root. Path: "
2023-09-05T15:26:24.178+0700 [TRACE] core: adding write forwarded paths: paths=[]
2023-09-05T15:26:24.178+0700 [INFO]  core: successfully mounted: type=identity version="v1.14.2+builtin.vault" path=identity/ namespace="ID: root. Path: "
2023-09-05T15:26:24.179+0700 [TRACE] token: no token generation counter found in storage
2023-09-05T15:26:24.179+0700 [TRACE] core: adding write forwarded paths: paths=[]
2023-09-05T15:26:24.179+0700 [INFO]  core: successfully mounted: type=token version="v1.14.2+builtin.vault" path=token/ namespace="ID: root. Path: "
2023-09-05T15:26:24.179+0700 [TRACE] expiration.job-manager: created dispatcher: name=expire-dispatcher num_workers=10
2023-09-05T15:26:24.179+0700 [TRACE] expiration.job-manager: initialized dispatcher: num_workers=10
2023-09-05T15:26:24.179+0700 [TRACE] expiration.job-manager: created job manager: name=expire pool_size=10
2023-09-05T15:26:24.179+0700 [TRACE] expiration.job-manager: starting job manager: name=expire
2023-09-05T15:26:24.179+0700 [TRACE] expiration.job-manager: starting dispatcher
2023-09-05T15:26:24.179+0700 [INFO]  rollback: starting rollback manager
2023-09-05T15:26:24.179+0700 [INFO]  core: restoring leases
2023-09-05T15:26:24.179+0700 [DEBUG] expiration: collecting leases
2023-09-05T15:26:24.179+0700 [DEBUG] expiration: leases collected: num_existing=0
2023-09-05T15:26:24.179+0700 [DEBUG] identity: loading entities
2023-09-05T15:26:24.179+0700 [DEBUG] identity: entities collected: num_existing=0
2023-09-05T15:26:24.179+0700 [INFO]  identity: entities restored
2023-09-05T15:26:24.179+0700 [DEBUG] identity: identity loading groups
2023-09-05T15:26:24.179+0700 [DEBUG] identity: groups collected: num_existing=0
2023-09-05T15:26:24.179+0700 [INFO]  identity: groups restored
2023-09-05T15:26:24.179+0700 [DEBUG] identity: identity loading OIDC clients
2023-09-05T15:26:24.179+0700 [TRACE] mfa: loading login MFA configurations
2023-09-05T15:26:24.179+0700 [TRACE] mfa: methods collected: num_existing=0
2023-09-05T15:26:24.179+0700 [TRACE] mfa: configurations restored: namespace="" prefix=login-mfa/method/
2023-09-05T15:26:24.179+0700 [TRACE] mfa: loading login MFA enforcement configurations
2023-09-05T15:26:24.179+0700 [TRACE] mfa: enforcements configs collected: num_existing=0
2023-09-05T15:26:24.179+0700 [TRACE] mfa: enforcement configurations restored: namespace="" prefix=login-mfa/enforcement/
2023-09-05T15:26:24.179+0700 [TRACE] activity: scanned existing logs: out=[]
2023-09-05T15:26:24.179+0700 [INFO]  core: Recorded vault version: vault version=1.14.2 upgrade time="2023-09-05 08:26:24.179905 +0000 UTC" build date=""
2023-09-05T15:26:24.179+0700 [TRACE] activity: scanned existing logs: out=[]
2023-09-05T15:26:24.179+0700 [INFO]  core: usage gauge collection is disabled
2023-09-05T15:26:24.179+0700 [TRACE] activity: no intent log found
2023-09-05T15:26:24.179+0700 [INFO]  expiration: lease restore complete
2023-09-05T15:26:24.180+0700 [DEBUG] secrets.identity.identity_e80d0190: wrote OIDC default provider
2023-09-05T15:26:24.264+0700 [DEBUG] secrets.identity.identity_e80d0190: generated OIDC public key to sign JWTs: key_id=13860ea2-7936-9b40-e2d5-d4758696bfdd
2023-09-05T15:26:24.416+0700 [DEBUG] secrets.identity.identity_e80d0190: generated OIDC public key for future use: key_id=fb7dbd3a-a554-5b0b-c858-7db784f48928
2023-09-05T15:26:24.416+0700 [DEBUG] secrets.identity.identity_e80d0190: wrote OIDC default key
2023-09-05T15:26:24.416+0700 [DEBUG] secrets.identity.identity_e80d0190: wrote OIDC allow_all assignment
2023-09-05T15:26:24.416+0700 [INFO]  core: post-unseal setup complete
2023-09-05T15:26:24.416+0700 [DEBUG] token: no wal state found when generating token
2023-09-05T15:26:24.416+0700 [INFO]  core: root token generated
2023-09-05T15:26:24.416+0700 [INFO]  core: pre-seal teardown starting
2023-09-05T15:26:24.416+0700 [DEBUG] expiration: stop triggered
2023-09-05T15:26:24.416+0700 [TRACE] expiration.job-manager: terminating job manager...
2023-09-05T15:26:24.416+0700 [TRACE] expiration.job-manager: terminating dispatcher
2023-09-05T15:26:24.416+0700 [DEBUG] expiration: finished stopping
2023-09-05T15:26:24.416+0700 [INFO]  rollback: stopping rollback manager
2023-09-05T15:26:24.416+0700 [INFO]  core: pre-seal teardown complete
2023-09-05T15:26:24.416+0700 [DEBUG] core: unseal key supplied: migrate=false
2023-09-05T15:26:24.416+0700 [DEBUG] core: cannot unseal, not enough keys: keys=1 threshold=3 nonce=1b2921eb-852f-97b3-d738-3ff9666b5ea3
2023-09-05T15:26:24.416+0700 [DEBUG] core: unseal key supplied: migrate=false
2023-09-05T15:26:24.416+0700 [DEBUG] core: cannot unseal, not enough keys: keys=2 threshold=3 nonce=1b2921eb-852f-97b3-d738-3ff9666b5ea3
2023-09-05T15:26:24.416+0700 [DEBUG] core: unseal key supplied: migrate=false
2023-09-05T15:26:24.416+0700 [INFO]  core: clustering disabled, not starting listeners
2023-09-05T15:26:24.416+0700 [INFO]  core: post-unseal setup starting
2023-09-05T15:26:24.416+0700 [DEBUG] core: clearing forwarding clients
2023-09-05T15:26:24.416+0700 [DEBUG] core: done clearing forwarding clients
2023-09-05T15:26:24.416+0700 [DEBUG] core: persisting feature flags
2023-09-05T15:26:24.416+0700 [INFO]  core: loaded wrapping token key
2023-09-05T15:26:24.416+0700 [INFO]  core: successfully setup plugin catalog: plugin-directory=""
2023-09-05T15:26:24.417+0700 [TRACE] core: adding write forwarded paths: paths=[]
2023-09-05T15:26:24.417+0700 [INFO]  core: successfully mounted: type=system version="v1.14.2+builtin.vault" path=sys/ namespace="ID: root. Path: "
2023-09-05T15:26:24.417+0700 [TRACE] core: adding write forwarded paths: paths=[]
2023-09-05T15:26:24.417+0700 [INFO]  core: successfully mounted: type=identity version="v1.14.2+builtin.vault" path=identity/ namespace="ID: root. Path: "
2023-09-05T15:26:24.417+0700 [TRACE] core: adding write forwarded paths: paths=[]
2023-09-05T15:26:24.417+0700 [INFO]  core: successfully mounted: type=cubbyhole version="v1.14.2+builtin.vault" path=cubbyhole/ namespace="ID: root. Path: "
2023-09-05T15:26:24.417+0700 [TRACE] token: no token generation counter found in storage
2023-09-05T15:26:24.417+0700 [TRACE] core: adding write forwarded paths: paths=[]
2023-09-05T15:26:24.417+0700 [INFO]  core: successfully mounted: type=token version="v1.14.2+builtin.vault" path=token/ namespace="ID: root. Path: "
2023-09-05T15:26:24.417+0700 [TRACE] expiration.job-manager: created dispatcher: name=expire-dispatcher num_workers=10
2023-09-05T15:26:24.417+0700 [TRACE] expiration.job-manager: initialized dispatcher: num_workers=10
2023-09-05T15:26:24.417+0700 [TRACE] expiration.job-manager: created job manager: name=expire pool_size=10
2023-09-05T15:26:24.417+0700 [TRACE] expiration.job-manager: starting job manager: name=expire
2023-09-05T15:26:24.417+0700 [TRACE] expiration.job-manager: starting dispatcher
2023-09-05T15:26:24.417+0700 [INFO]  rollback: starting rollback manager
2023-09-05T15:26:24.417+0700 [INFO]  core: restoring leases
2023-09-05T15:26:24.417+0700 [DEBUG] expiration: collecting leases
2023-09-05T15:26:24.417+0700 [DEBUG] expiration: leases collected: num_existing=0
2023-09-05T15:26:24.417+0700 [DEBUG] identity: loading entities
2023-09-05T15:26:24.417+0700 [DEBUG] identity: entities collected: num_existing=0
2023-09-05T15:26:24.417+0700 [INFO]  identity: entities restored
2023-09-05T15:26:24.417+0700 [DEBUG] identity: identity loading groups
2023-09-05T15:26:24.417+0700 [INFO]  expiration: lease restore complete
2023-09-05T15:26:24.417+0700 [DEBUG] identity: groups collected: num_existing=0
2023-09-05T15:26:24.417+0700 [INFO]  identity: groups restored
2023-09-05T15:26:24.417+0700 [DEBUG] identity: identity loading OIDC clients
2023-09-05T15:26:24.417+0700 [TRACE] mfa: loading login MFA configurations
2023-09-05T15:26:24.417+0700 [TRACE] mfa: methods collected: num_existing=0
2023-09-05T15:26:24.417+0700 [TRACE] mfa: configurations restored: namespace="" prefix=login-mfa/method/
2023-09-05T15:26:24.417+0700 [TRACE] mfa: loading login MFA enforcement configurations
2023-09-05T15:26:24.417+0700 [TRACE] mfa: enforcements configs collected: num_existing=0
2023-09-05T15:26:24.417+0700 [TRACE] mfa: enforcement configurations restored: namespace="" prefix=login-mfa/enforcement/
2023-09-05T15:26:24.418+0700 [TRACE] activity: scanned existing logs: out=[]
2023-09-05T15:26:24.418+0700 [TRACE] activity: no intent log found
2023-09-05T15:26:24.418+0700 [TRACE] activity: scanned existing logs: out=[]
2023-09-05T15:26:24.418+0700 [INFO]  core: usage gauge collection is disabled
2023-09-05T15:26:24.418+0700 [INFO]  core: post-unseal setup complete
2023-09-05T15:26:24.418+0700 [INFO]  core: vault is unsealed
panic: Pattern managed-keys/(?P<type>\w(([\w-.]+)?\w)?)/(?P<name>\w(([\w-.]+)?\w)?)/test/sign defines a CreateOperation but no ExistenceCheck

goroutine 1 [running]:
github.com/hashicorp/vault/sdk/framework.(*Backend).init(0x140001b4d20)
	github.com/hashicorp/vault/sdk@v0.9.2/framework/backend.go:504 +0x334
sync.(*Once).doSlow(0x1400076ab78?, 0x0?)
	sync/once.go:74 +0x104
sync.(*Once).Do(...)
	sync/once.go:65
github.com/hashicorp/vault/sdk/framework.(*Backend).HandleExistenceCheck(0x140001b4d20?, {0x105ac59a0, 0x14000642510}, 0x14000650000)
	github.com/hashicorp/vault/sdk@v0.9.2/framework/backend.go:160 +0x64
github.com/hashicorp/vault/vault.(*Router).routeCommon(0x14000abff40, {0x105ac59a0, 0x14000642510}, 0x14000650000, 0x1)
	github.com/hashicorp/vault@v1.14.2/vault/router.go:779 +0x13bc
github.com/hashicorp/vault/vault.(*Router).RouteExistenceCheck(0x14000abff40?, {0x105ac59a0?, 0x14000642510?}, 0x104b0df63?)
	github.com/hashicorp/vault@v1.14.2/vault/router.go:558 +0x28
github.com/hashicorp/vault/vault.(*Core).CheckToken(0x14000873800, {0x105ac59a0, 0x14000642510}, 0x14000650000, 0x0)
	github.com/hashicorp/vault@v1.14.2/vault/request_handling.go:332 +0x434
github.com/hashicorp/vault/vault.(*Core).handleRequest(0x14000873800, {0x105ac59a0, 0x14000642510}, 0x14000650000)
	github.com/hashicorp/vault@v1.14.2/vault/request_handling.go:902 +0x308
github.com/hashicorp/vault/vault.(*Core).handleCancelableRequest(0x14000873800, {0x105ac59a0, 0x14000642300}, 0x14000650000)
	github.com/hashicorp/vault@v1.14.2/vault/request_handling.go:685 +0x118c
github.com/hashicorp/vault/vault.(*Core).switchedLockHandleRequest(0x14000873800, {0x105ac59a0, 0x140006422a0}, 0x14000650000, 0x1)
	github.com/hashicorp/vault@v1.14.2/vault/request_handling.go:492 +0x45c
github.com/hashicorp/vault/vault.(*Core).HandleRequest(...)
	github.com/hashicorp/vault@v1.14.2/vault/request_handling.go:453
github.com/hashicorp/vault/vault.testCoreAddSecretMount({0x105ad4fc0, 0x14000603ba0}, 0x14000873800?, {0x14000b44640, 0x1c})
	github.com/hashicorp/vault@v1.14.2/vault/testing.go:434 +0x298
github.com/hashicorp/vault/vault.testCoreUnsealed({0x105ad4fc0, 0x14000603ba0}, 0x14000bbfc88?)
	github.com/hashicorp/vault@v1.14.2/vault/testing.go:402 +0x74
github.com/hashicorp/vault/vault.TestCoreUnsealedWithConfig({0x105ad4fc0, 0x14000603ba0}, 0x104aee850?)
	github.com/hashicorp/vault@v1.14.2/vault/testing.go:395 +0x54
main.main()
	./main.go:20 +0xbc
```