# vim: set filetype=sh tabstop=2 shiftwidth=2 expandtab :
# shellcheck shell=sh
"${sourced_b8569f0-false}" && return 0; sourced_b8569f0=true

. ./go.lib.sh

# Run tests.
task_test() {
  go test ./...
}
