#!/bin/bash

# Asserts that a command succeeds.
#
# Arguments:
#    $@ - The command to run.
#
# Returns:
#    none
assert_command() {
    local stderrf="$(mktemp)"
    local stdoutf="$(mktemp)"
    local statusf="$(mktemp)"
    ( set +e; "$@" 2>"$stderrf" >"$stdoutf"; echo -n "$?" > "$statusf" )
    status="$(<$statusf)"; rm "$statusf"
    stderr="$(cat $stderrf)"; rm "$stderrf"
    stdout="$(cat $stdoutf)"; rm "$stdoutf"
    output="$(
        [ "$stderr" ] && echo $stderr || true;
        [ "$stdout" ] && echo $stdout || true;
    )"
    [[ $status == 0 ]] || \
        ( (echo "$*"; echo "status: $status"; echo "$output" | batslib_decorate "Output") \
            | batslib_decorate "Command failed" \
            | fail)
}

# Asserts that a command fails.
#
# Arguments:
#    $@ - The command to run.
#
# Returns:
#    none
assert_command_fail() {
    local stderrf="$(mktemp)"
    local stdoutf="$(mktemp)"
    local statusf="$(mktemp)"
    ( set +e; "$@" 2>"$stderrf" >"$stdoutf"; echo -n "$?" > "$statusf" )
    status="$(<$statusf)"; rm "$statusf"
    stderr="$(cat $stderrf)"; rm "$stderrf"
    stdout="$(cat $stdoutf)"; rm "$stdoutf"
    output="$(
        [ "$stderr" ] && echo $stderr || true;
        [ "$stdout" ] && echo $stdout || true;
    )"
    [[ $status != 0 ]] || \
        ( (echo "$*"; echo "status: $status"; echo "$output" | batslib_decorate "Output") \
            | batslib_decorate "Command succeeded (should have failed)" \
            | fail)
}

# Asserts that a phrase matches a regular expression.
#
# Arguments:
#    $1 - The regular expression.
#    $2 - The phrase to match against. Defaults to $output.
assert_match() {
    regex="$1"
    if [[ $# < 2 ]]; then
        text="$output"
    else
        text="$2"
    fi
    [[ "$text" =~ $regex ]] || \
        (batslib_print_kv_single_or_multi 10 "regex" "$regex" "actual" "$text" \
            | batslib_decorate "output does not match" \
            | fail)
}

# Asserts that two values are equal.
#
# Arguments:
#    $1 - The expected value.
#    $2 - The actual value. Defaults to $output.
assert_eq() {
    expected="$1"
    if [[ $# < 2 ]]; then
        actual="$output"
    else
        actual="$2"
    fi
    [[ "$actual" == "$expected" ]] || \
        (batslib_print_kv_single_or_multi 10 "expected" "$expected" "actual" "$actual" \
            | batslib_decorate "output does not match" \
            | fail)
}

# Asserts that two values are not equal.
#
# Arguments:
#    $1 - The expected value.
#    $2 - The actual value. Defaults to $output.
assert_neq() {
    expected="$1"
    if [[ $# < 2 ]]; then
        actual="$output"
    else
        actual="$2"
    fi
    [[ "$actual" != "$expected" ]] || \
        (batslib_print_kv_single_or_multi 10 "expected" "$expected" "actual" "$actual" \
            | batslib_decorate "output does not match" \
            | fail)
}
