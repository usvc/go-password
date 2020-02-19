# Password Package

Passwords: as simple as it gets.

- [Password Package](#password-package)
  - [Usage](#usage)
    - [Importing](#importing)
    - [Hashing Passwords](#hashing-passwords)
    - [Verifying Passwords](#verifying-passwords)
    - [Validating Passwords](#validating-passwords)
    - [Customizing Password Policies](#customizing-password-policies)
  - [Development Runbook](#development-runbook)
    - [CI Environment Variables](#ci-environment-variables)
  - [Licensing](#licensing)

## Usage

### Importing

```go
import "github.com/usvc/password"
```

### Hashing Passwords

```go
plaintext := "abcd1234!@#$"
hash, salt, err := password.Hash(plaintext)
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

### CI Environment Variables

- **`DEPLOY_KEY`**: 

## Licensing

Code in this package is licensed under the [MIT license (click to view text)](./LICENSE).
