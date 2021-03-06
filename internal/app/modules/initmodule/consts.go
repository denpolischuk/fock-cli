package initmodule

import "github.com/denpolischuk/fock-cli/internal/app/consts"

const (
	// ZshAutocompletionScript ...
	ZshAutocompletionScript = `
#compdef $PROG

_cli_zsh_autocomplete() {

  local -a opts
  local cur
  cur=${words[-1]}
  if [[ "$cur" == "-"* ]]; then
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} ${cur} --generate-bash-completion)}")
  else
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")
  fi

  if [[ "${opts[1]}" != "" ]]; then
    _describe 'values' opts
  else
    _files
  fi

  return
}

compdef _cli_zsh_autocomplete $PROG
	`

	// ZshRcScript - part of zsh script to be written into .zshrc
	ZshRcScript = "PROG=" + consts.AppBinName + " _CLI_ZSH_AUTOCOMPLETE_HACK=1 source "
	// BashAutocompletionScript ...
	BashAutocompletionScript = `
#! /bin/bash

: ${PROG:=$(basename ${BASH_SOURCE})}

_cli_bash_autocomplete() {
  if [[ "${COMP_WORDS[0]}" != "source" ]]; then
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    if [[ "$cur" == "-"* ]]; then
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-bash-completion )
    else
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    fi
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
  fi
}

complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete $PROG
unset PROG
  `
)

var defaultShellDetectErrorMessage = "[Autocompletion]: " + consts.Emojis["dead"] + " couldn't detect default shell. Aborting..."
