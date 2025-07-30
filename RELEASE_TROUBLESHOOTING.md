# GitHub Release Troubleshooting Guide

## 403 Forbidden Error Solutions

If you encounter a 403 error when creating GitHub releases, try these solutions in order:

### 1. ✅ Already Fixed: Explicit Permissions
The workflow now includes explicit permissions:
```yaml
permissions:
  contents: write
  packages: write
```

### 2. Check Repository Settings
1. Go to your repository on GitHub
2. Navigate to **Settings** → **Actions** → **General**
3. Ensure **Workflow permissions** is set to:
   - ✅ **Read and write permissions**
   - ✅ **Allow GitHub Actions to create and approve pull requests**

### 3. Check Repository Permissions
1. Go to **Settings** → **Collaborators and teams**
2. Ensure your account has **Write** or **Admin** access
3. If using an organization, check organization settings

### 4. Verify GitHub Token
The workflow uses `${{ secrets.GITHUB_TOKEN }}` which should be automatically available.
If issues persist, you can create a Personal Access Token:

1. Go to GitHub → **Settings** → **Developer settings** → **Personal access tokens**
2. Create a new token with `repo` scope
3. Add it as a repository secret named `RELEASE_TOKEN`
4. Update the workflow to use `${{ secrets.RELEASE_TOKEN }}`

### 5. Check for Existing Release
If a release with the same tag already exists:
1. Delete the existing release on GitHub
2. Delete the tag: `git tag -d v1.0.1 && git push origin :refs/tags/v1.0.1`
3. Recreate the tag: `git tag v1.0.1 && git push origin v1.0.1`

### 6. Alternative: Manual Release
If automated releases continue to fail:

```bash
# Build locally
make build-all

# Create release manually on GitHub
# Upload the built binaries:
# - CLI-linux-amd64
# - CLI-linux-arm64  
# - CLI-darwin-amd64
# - CLI-darwin-arm64
# - CLI-windows-amd64.exe
# - CLI-windows-arm64.exe
# - CLI-*.sha256 files
```

## Testing the Fix

1. **Push the updated workflow:**
   ```bash
   git push origin main
   ```

2. **Create a new tag:**
   ```bash
   git tag v1.0.2
   git push origin v1.0.2
   ```

3. **Monitor the workflow:**
   - Go to **Actions** tab on GitHub
   - Watch the "Release" workflow execution

## Common Issues and Solutions

| Issue | Solution |
|-------|----------|
| 403 Forbidden | ✅ Add explicit permissions (fixed) |
| Token expired | Create new Personal Access Token |
| Repository archived | Unarchive the repository |
| Organization restrictions | Contact organization admin |
| Rate limiting | Wait and retry later |

## Debugging Steps

1. **Check workflow logs** in GitHub Actions
2. **Verify tag exists:** `git ls-remote --tags origin`
3. **Test permissions locally:** `gh auth status`
4. **Check repository visibility** and access settings

## Need Help?

If issues persist:
1. Check GitHub Actions logs for detailed error messages
2. Verify all repository settings
3. Consider using GitHub CLI for manual releases
4. Contact GitHub support if it's a platform issue 