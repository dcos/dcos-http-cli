#!/bin/bash

_dcos_http() {
    local flags=(
    "--data"
    "--request="
    )

    local methods=(
    "GET"
    "HEAD"
    "POST"
    "PUT"
    "DELETE"
    "TRACE"
    "OPTIONS"
    "CONNECT"
    )

    if [ $prev == "--request" ]; then
        local compreply=($(compgen -W "${methods[*]}" -- "${cur}"))

        # Add spaces after each array element to circumvent the "nospace" option.
        COMPREPLY=( "${compreply[@]/%/ }" )
        return
    fi

    case "$cur" in
        --request=*)
            # Get HTTP methods completions using compgen, we strip "--request=" from
            # the current word to only keep the current option value, if any.
            local compreply=($(compgen -W "${methods[*]}" -- "${cur#*=}"))

            # Add spaces after each array element to circumvent the "nospace" option.
            COMPREPLY=( "${compreply[@]/%/ }" )
            ;;
        --*)
            __dcos_handle_compreply "${flags[@]}"
            ;;
        *)
            ;;
    esac
    return

}