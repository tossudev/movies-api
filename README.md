# Nice group project template

✅ The `main` branch is protected: pushes and force pushes are disabled, but pull requests are allowed.

✅ Pull requests require a minimum of 1 approval (though the repository owner can merge without approval).

✅ Merging is blocked if changes are requested: even with sufficient approvals, a PR cannot be merged if reviewers have requested changes.

✅ Pre-commit hook is included.

✅ Allowlist gitignore for whitelist-style adepts 

❗ Make sure to select checkboxes `Branch Protection` and `Git content (default branch)` when creating a new repo using this template.

## About pre-commit hooks

A pre-commit hook is a script that runs automatically **before** a Git commit is finalized. It typically checks your changes (e.g., linting, tests, formatting) and can block the commit if issues are detected.

The hook included in this repository template allows you to:

1. [Format](https://pkg.go.dev/cmd/go/internal/fmtcmd) staged `.go` files.
2. Run [`go vet`](https://pkg.go.dev/cmd/vet) on all Go files in the repository, including those in nested folders.

---

## Install the pre-commit hook (optional)

Git **never** runs hooks from a freshly cloned repository. Each team member must opt in manually:

```sh
make init
```

This command runs `git config core.hooksPath .githooks` to inform Git about the new location for hooks, **instead of** the default `.git/hooks/`.

If no `.go` files are staged, the hook does nothing.

To bypass the hook for a single commit, use:

```sh
git commit --no-verify
```
