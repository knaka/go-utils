# vim: set filetype=sh tabstop=2 shiftwidth=2 expandtab :
# shellcheck shell=sh
"${sourced_b5282d2-false}" && return 0; sourced_b5282d2=true

. ./task.sh
. ./go.lib.sh

# Run tests.
task_test() {
  call_task subcmd_go__test "$@"
}
