# vim: set filetype=sh tabstop=2 shiftwidth=2 expandtab :
# shellcheck shell=sh
"${sourced_b5282d2-false}" && return 0; sourced_b5282d2=true

set -- "$PWD" "${0%/*}" "$@"; if test "$2" != "$0"; then cd "$2" 2>/dev/null || :; fi
. ./task.sh
. ./go.lib.sh
cd "$1"; shift 2

# Run tests.
task_test() {
  call_task subcmd_go__test "$@"
}
