package internal

// TODO: add support for more shells like fish, nu
var SupportedShells = []string{
	"bash",
	"zsh",
}

type Hooks interface {
	Hook() string
}

type (
	Zsh  struct{}
	Bash struct{}
)

func (z Zsh) Hook() string {
	return `function _sevp() {
    if [[ -f ~/.sevp ]]; then
        eval "$(cat ~/.sevp)"
    fi
}

precmd_functions+=(_sevp)`
}

func (b Bash) Hook() string {
	return `function _sevp() {
    if [[ -f ~/.sevp ]]; then
        eval "$(cat ~/.sevp)"
    fi
}

PROMPT_COMMAND="_sevp; ${PROMPT_COMMAND}"`
}
