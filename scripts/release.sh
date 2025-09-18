#!/bin/bash

# Release script for goscaffold
# Usage: ./scripts/release.sh [patch|minor|major]

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Functions
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if we're on main branch
current_branch=$(git branch --show-current)
if [ "$current_branch" != "main" ]; then
    log_error "You must be on the main branch to create a release"
    exit 1
fi

# Check if working directory is clean
if [ -n "$(git status --porcelain)" ]; then
    log_error "Working directory is not clean. Please commit or stash your changes."
    exit 1
fi

# Pull latest changes
log_info "Pulling latest changes from origin/main..."
git pull origin main

# Get current version
current_version=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
log_info "Current version: $current_version"

# Parse version
if [[ $current_version =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
    major=${BASH_REMATCH[1]}
    minor=${BASH_REMATCH[2]}
    patch=${BASH_REMATCH[3]}
else
    log_warn "No valid version found, starting from v0.0.0"
    major=0
    minor=0
    patch=0
fi

# Determine new version based on argument
case $1 in
    patch)
        new_patch=$((patch + 1))
        new_version="v${major}.${minor}.${new_patch}"
        ;;
    minor)
        new_minor=$((minor + 1))
        new_version="v${major}.${new_minor}.0"
        ;;
    major)
        new_major=$((major + 1))
        new_version="v${new_major}.0.0"
        ;;
    *)
        log_error "Usage: $0 [patch|minor|major]"
        exit 1
        ;;
esac

log_info "New version will be: $new_version"

# Confirm release
read -p "Do you want to create release $new_version? (y/N): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    log_info "Release cancelled"
    exit 0
fi

# Run tests
log_info "Running tests..."
make test

# Run linting
log_info "Running linter..."
make lint

# Create and push tag
log_info "Creating and pushing tag $new_version..."
git tag -a "$new_version" -m "Release $new_version"
git push origin "$new_version"

log_info "Release $new_version created successfully!"
log_info "GitHub Actions will now build and publish the release."
