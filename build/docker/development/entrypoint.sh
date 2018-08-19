#!/usr/bin/env sh

echo "Installing dependencies"
dep ensure --vendor-only

exec realize start