package internal

var SupportedShells = []string{
	"bash",
	"zsh",
}

const ZshHook string = `function _sevp() {
    if [[ -f ~/.sevp ]]; then
        eval "$(cat ~/.sevp)"
    fi
}

precmd_functions+=(_sevp)`

const BashHook string = `function _sevp() {
    if [[ -f ~/.sevp ]]; then
        eval "$(cat ~/.sevp)"
    fi
}

PROMPT_COMMAND="_sevp; ${PROMPT_COMMAND}"`
