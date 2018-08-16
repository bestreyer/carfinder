#!/usr/bin/env sh

echo "install dependencies"
dep ensure --vendor-only

exec realize start