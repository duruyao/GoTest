#!/usr/bin/env bash

set -euo pipefail

# usage: bash build.sh [--all]

function error_ln() {
  printf "\033[1;32;31m%s\n\033[m" "${1}"
}

release_list=(
  "darwin 386"
  "darwin amd64"
  "darwin arm"
  "darwin arm64"
  "linux 386"
  "linux amd64"
  "linux arm"
  "linux arm64"
  "windows 386"
  "windows amd64"
  "windows arm"
  "windows arm64"
)

if ! [ -f "go.mod" ]; then
  error_ln "Error: 'go.mod' required but not found" >&2
  echo "Usage: $0 [--all]" >&2
  exit 1
fi
package="$(grep "module.*" go.mod | sed "s/module//g" | sed "s/ //g")"
app="$(basename "${package}")"

## compile gotest for current platform
if [ $# == 0 ]; then
  bash -x -c "GOROOT=${GOROOT} GOPATH=${GOPATH} GO_ENABLED=0 ${GOROOT}/bin/go build -o ${PWD}/${app} ${package}"
  exit 0
fi
if [ $# != 0 ] && [ "$1" != "--all" ]; then
  error_ln "Error: Unknown flags" >&2
  echo "Usage: $0 [--all]" >&2
  exit 1
fi

## cross compile gotest for multi-platform
releases_dir="${PWD}/releases"
if ! [ -f "package.json" ]; then
  error_ln "Error: 'package.json' required but not found" >&2
  echo "Usage: $0 [--all]" >&2
  exit 1
fi
version="$(grep -o "version.*" package.json | head -1 | grep -o "[0-9]\+.[0-9]\+.[0-9]\+")"

mkdir -p "${releases_dir}"
rm -rf "${releases_dir:?}"/*
pushd "${releases_dir}" 1>/dev/null 2>&1

for release in "${release_list[@]}"; do
  echo ""
  # shellcheck disable=SC2206
  release=(${release[*]})
  target_dir="${PWD}/${app}-${release[0]}-${release[1]}-${version}"
  target_path="${target_dir}/${app}"
  if [ "${release[0]}" == "windows" ]; then
    target_path="${target_dir}/${app}.exe"
  fi

  mkdir -p "${target_dir}"
  bash -x -c "GOROOT=${GOROOT} GOPATH=${GOPATH} GO_ENABLED=0 GOOS=${release[0]} GOARCH=${release[1]} ${GOROOT}/bin/go build -o ${target_path} ${package}" || continue
  chmod +x "${target_path}"

  if [ -n "$(command -v zip)" ]; then
    bash -x -c "zip -r ${target_dir}.zip $(basename "${target_dir}") 1>/dev/null"
  fi
  if [ -n "$(command -v tar)" ]; then
    bash -x -c "tar -cvf ${target_dir}.tar $(basename "${target_dir}") 1>/dev/null"
    bash -x -c "tar -zcvf ${target_dir}.tar.gz $(basename "${target_dir}") 1>/dev/null"
    bash -x -c "tar -jcvf ${target_dir}.tar.bz2 $(basename "${target_dir}") 1>/dev/null"
  fi
done

popd 1>/dev/null 2>&1