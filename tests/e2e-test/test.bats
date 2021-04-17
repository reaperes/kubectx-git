#!/usr/bin/env bats

COMMAND="${COMMAND:-$BATS_TEST_DIRNAME/../kubectx-git}"

@test "version should not fail" {
  run ${COMMAND} version
  echo "$output"
  [ "$status" -eq 0 ]
}
