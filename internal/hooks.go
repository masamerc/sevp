package internal

var SupportedShells = []string{
	"bash",
	"zsh",
	"fish",
	"nu",
}

type Hooks interface {
	Hook() string
}

type (
	Zsh  struct{}
	Bash struct{}
	Fish struct{}
	Nu   struct{}
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

func (f Fish) Hook() string {
	return `function _sevp
    if test -f ~/.sevp
        eval (cat ~/.sevp)
    end
end

# Add _sevp function to fish_prompt
function fish_prompt
    _sevp
    # Call the original prompt
    command fish_prompt $argv
end`
}

func (n Nu) Hook() string {
	return `let-env PROMPT_HOOK = {|
    if test (ls ~/.sevp | is-empty) == $false {
        nu --eval $(cat ~/.sevp | str collect ' ')
    }
|}`
}
