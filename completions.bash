#!/bin/bash

_keysmith_completions()
{
    if [ "${#COMP_WORDS[@]}" != "2" ]; then
        return
    fi
    COMPREPLY=($(compgen -W "$(keysmith shortlist)" "${COMP_WORDS[1]}"))
}

complete -F _keysmith_completions keysmith
