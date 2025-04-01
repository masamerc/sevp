package internal

// TODO: add support for more shells like fish, nu
var SupportedShells = []string{
	"bash",
	"zsh",
}

func ZshHook() string {
	return `function _sevp() {
    if [[ -f ~/.sevp ]]; then
        eval "$(cat ~/.sevp)"
    fi
}

precmd_functions+=(_sevp)`
}

func BashHook() string {
	return `function _sevp() {
    if [[ -f ~/.sevp ]]; then
        eval "$(cat ~/.sevp)"
    fi
}

PROMPT_COMMAND="_sevp; ${PROMPT_COMMAND}"`
}
