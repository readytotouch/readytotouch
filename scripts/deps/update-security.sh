#!/usr/bin/env bash
# Updates only the modules that have known, reachable vulnerabilities,
# to the minimal fixed version. Does not commit anything.
#
# Usage: GOVULNCHECK="go run golang.org/x/vuln/cmd/govulncheck@vX.Y.Z" scripts/deps/update-security.sh
set -euo pipefail

: "${GOVULNCHECK:?GOVULNCHECK is required}"

report="$(mktemp)"
trap 'rm -f "$report"' EXIT

# -format json always exits with 0, findings are in the stream
$GOVULNCHECK -format json ./... > "$report"

# Only findings with a called function (reachable), same as the text mode reports
findings="$(jq -r '.finding? | select(. != null and .trace[0].function != null) | "\(.trace[0].module)\t\(.fixed_version // "")"' "$report" | sort -u)"

if [ -z "$findings" ]; then
	echo "No reachable vulnerabilities found."
	exit 0
fi

# Highest fixed version per module (sort -V, the last one wins)
fixes="$(printf '%s\n' "$findings" | awk -F'\t' '$2 != ""' | sort -t$'\t' -k1,1 -k2,2V | awk -F'\t' '{m[$1]=$2} END {for (k in m) print k "\t" m[k]}' | sort)"
unfixed="$(printf '%s\n' "$findings" | awk -F'\t' '$2 == "" {print $1}' | sort -u)"

if [ -n "$fixes" ]; then
	while IFS=$'\t' read -r module version; do
		if [ "$module" = "stdlib" ]; then
			echo "stdlib: fixed in Go ${version#go}. Update the Go version (go.mod, docker/compose, scripts/ansible, .github/workflows)."
			continue
		fi

		echo "Updating $module to $version"
		go get "$module@$version"
	done <<< "$fixes"

	go mod tidy
fi

if [ -n "$unfixed" ]; then
	echo "No fixed version is available for (decide manually):"
	printf '  %s\n' $unfixed
fi

echo "Verifying..."
go build ./...
$GOVULNCHECK ./...
