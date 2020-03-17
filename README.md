# Password

[![latest release](https://badge.fury.io/gh/usvc%2Fgo-password.svg)](https://github.com/usvc/go-password/releases)
[![build status](https://travis-ci.org/usvc/go-password.svg?branch=master)](https://travis-ci.org/usvc/go-password)
[![pipeline status](https://gitlab.com/usvc/modules/go/password/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/password/-/commits/master)
[![test coverage](https://api.codeclimate.com/v1/badges/1bdfbf587cb7feab08ae/test_coverage)](https://codeclimate.com/github/usvc/go-password/test_coverage)
[![maintainability](https://api.codeclimate.com/v1/badges/1bdfbf587cb7feab08ae/maintainability)](https://codeclimate.com/github/usvc/go-password/maintainability)

A Go package to manage password hashing, verification, and validation.

| | |
| --- | --- |
| Github | [https://github.com/usvc/go-password](https://github.com/usvc/go-password) |
| Gitlab | [https://gitlab.com/usvc/modules/go/password](https://gitlab.com/usvc/modules/go/password) |

- - -

- [Password](#password)
  - [Usage](#usage)
    - [Importing](#importing)
    - [Hashing Passwords](#hashing-passwords)
    - [Verifying Passwords](#verifying-passwords)
    - [Validating Passwords](#validating-passwords)
    - [Customizing Password Policies](#customizing-password-policies)
  - [Development Runbook](#development-runbook)
    - [Getting Started](#getting-started)
    - [Continuous Integration (CI) Pipeline](#continuous-integration-ci-pipeline)
      - [On Github](#on-github)
        - [Releasing](#releasing)
      - [On Gitlab](#on-gitlab)
        - [Version Bumping](#version-bumping)
  - [Licensing](#licensing)

## Usage

### Importing

```go
import "github.com/usvc/go-password"
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

#### On Github

Github is used to deploy binaries/libraries because of it's ease of access by other developers.

##### Releasing

Releasing of the binaries can be done via Travis CI.

1. On Github, navigate to the [tokens settings page](https://github.com/settings/tokens) (by clicking on your profile picture, selecting **Settings**, selecting **Developer settings** on the left navigation menu, then **Personal Access Tokens** again on the left navigation menu)
2. Click on **Generate new token**, give the token an appropriate name and check the checkbox on **`public_repo`** within the **repo** header
3. Copy the generated token
4. Navigate to [travis-ci.org](https://travis-ci.org) and access the cooresponding repository there. Click on the **More options** button on the top right of the repository page and select **Settings**
5. Scroll down to the section on **Environment Variables** and enter in a new **NAME** with `RELEASE_TOKEN` and the **VALUE** field cooresponding to the generated personal access token, and hit **Add**

#### On Gitlab

##### Version Bumping

To set up the CI pipeline in Gitlab:

1. Run `make .ssh`
2. Copy the contents of the file generated at `./.ssh/id_rsa.base64` into an environment variable named **`DEPLOY_KEY`** in **Settings > CI/CD > Variables**
3. Navigate to the **Deploy Keys** section of the **Settings > Repository > Deploy Keys** and paste in the contents of the file generated at `./.ssh/id_rsa.pub` with the **Write access allowed** checkbox enabled

- **`DEPLOY_KEY`**: generate this by running `make .ssh` and copying the contents of the file generated at `./.ssh/id_rsa.base64`

## Licensing

Code in this package is licensed under the [MIT license (click to view text)](./LICENSE).
