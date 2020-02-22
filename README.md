# Password Package

[![pipeline status](https://gitlab.com/usvc/modules/go/password/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/password/-/commits/master)

A Go package to manage password hashing, verification, and validation.

- [Password Package](#password-package)
  - [Usage](#usage)
    - [Importing](#importing)
    - [Hashing Passwords](#hashing-passwords)
    - [Verifying Passwords](#verifying-passwords)
    - [Validating Passwords](#validating-passwords)
    - [Customizing Password Policies](#customizing-password-policies)
  - [Development Runbook](#development-runbook)
    - [Getting Started](#getting-started)
    - [Continuous Integration (CI) Pipeline](#continuous-integration-ci-pipeline)
  - [Licensing](#licensing)

## Usage

### Importing

```go
import "github.com/usvc/password"
```

### Hashing Passwords

```go
plaintext := "abcd1234!@#$"
hash, salt, err := password.Hash(plaintext, 32)
// ...
// store the hash and salt
// ...
```

### Verifying Passwords

```go
storedHash := "<hash>"
storedSalt := "<salt>"
plaintext := "abcd1234!@#$"
err := password.Verify(plaintext, storedHash, storedSalt)
if err != nil {
  // handle failed verification
} else {
  // handle successful verification
}
```

### Validating Passwords

```go
defaultPolicy := password.GetDefaultPolicy()
plaintext := "abcd1234!@#$"
if err := password.Validate(plaintext, defaultPolicy); err != nil {
  // handle failed validation
} else {
  // handle successful validation
}
```

### Customizing Password Policies

```go
customPolicy := password.Policy{
  MaximumLength: 32,
  MinimumLength: 12,
  MinimumLowercaseCount: 1,
  MinimumUppercaseCount: 1,
  MinimumNumericCount: 1,
  MinimumSpecialCount: 1,
  CustomSpecial: []byte("`!@"),
}
plaintext := "abcd1234!@#$"
if err := password.Validate(plaintext, defaultPolicy); err != nil {
  // handle failed validation
} else {
  // handle successful validation
}
```

## Development Runbook

### Getting Started

1. Clone this repository
2. Run `make deps` to pull in external dependencies
3. Write some awesome stuff
4. Run `make test` to ensure unit tests are passing
5. Push

### Continuous Integration (CI) Pipeline

To set up the CI pipeline in Gitlab:

1. Run `make .ssh`
2. Copy the contents of the file generated at `./.ssh/id_rsa.base64` into an environment variable named **`DEPLOY_KEY`** in **Settings > CI/CD > Variables**
3. Navigate to the **Deploy Keys** section of the **Settings > Repository > Deploy Keys** and paste in the contents of the file generated at `./.ssh/id_rsa.pub` with the **Write access allowed** checkbox enabled

- **`DEPLOY_KEY`**: generate this by running `make .ssh` and copying the contents of the file generated at `./.ssh/id_rsa.base64`

## Licensing

Code in this package is licensed under the [MIT license (click to view text)](./LICENSE).
