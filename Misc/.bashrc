# ~/.bashrc: executed by bash(1) for non-login shells.
# see /usr/share/doc/bash/examples/startup-files (in the package bash-doc)
# for examples

# If not running interactively, don't do anything
case $- in
    *i*) ;;
      *) return;;
esac

# don't put duplicate lines or lines starting with space in the history.
# See bash(1) for more options
HISTCONTROL=ignoreboth

# append to the history file, don't overwrite it
shopt -s histappend

# for setting history length see HISTSIZE and HISTFILESIZE in bash(1)
HISTSIZE=-1
HISTFILESIZE=10000

# check the window size after each command and, if necessary,
# update the values of LINES and COLUMNS.
shopt -s checkwinsize

# If set, the pattern "**" used in a pathname expansion context will
# match all files and zero or more directories and subdirectories.
#shopt -s globstar

# make less more friendly for non-text input files, see lesspipe(1)
[ -x /usr/bin/lesspipe ] && eval "$(SHELL=/bin/sh lesspipe)"

# set variable identifying the chroot you work in (used in the prompt below)
if [ -z "${debian_chroot:-}" ] && [ -r /etc/debian_chroot ]; then
    debian_chroot=$(cat /etc/debian_chroot)
fi

# set a fancy prompt (non-color, unless we know we "want" color)
case "$TERM" in
    xterm-color|*-256color) color_prompt=yes;;
esac

# uncomment for a colored prompt, if the terminal has the capability; turned
# off by default to not distract the user: the focus in a terminal window
# should be on the output of commands, not on the prompt
#force_color_prompt=yes

if [ -n "$force_color_prompt" ]; then
    if [ -x /usr/bin/tput ] && tput setaf 1 >&/dev/null; then
	# We have color support; assume it's compliant with Ecma-48
	# (ISO/IEC-6429). (Lack of such support is extremely rare, and such
	# a case would tend to support setf rather than setaf.)
	color_prompt=yes
    else
	color_prompt=
    fi
fi

if [ "$color_prompt" = yes ]; then
    PS1='${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$ '
else
    PS1='${debian_chroot:+($debian_chroot)}\u@\h:\w\$ '
fi
unset color_prompt force_color_prompt

# If this is an xterm set the title to user@host:dir
case "$TERM" in
xterm*|rxvt*)
    PS1="\[\e]0;${debian_chroot:+($debian_chroot)}\u@\h: \w\a\]$PS1"
    ;;
*)
    ;;
esac

# enable color support of ls and also add handy aliases
if [ -x /usr/bin/dircolors ]; then
    test -r ~/.dircolors && eval "$(dircolors -b ~/.dircolors)" || eval "$(dircolors -b)"
    alias ls='ls --color=auto'
    #alias dir='dir --color=auto'
    #alias vdir='vdir --color=auto'

    alias grep='grep --color=auto'
    alias fgrep='fgrep --color=auto'
    alias egrep='egrep --color=auto'
fi

# colored GCC warnings and errors
#export GCC_COLORS='error=01;31:warning=01;35:note=01;36:caret=01;32:locus=01:quote=01'


# some more ls aliases

alias kd='kdiff3 out/demo_ref.csv out/demo_curr.csv'
alias trm='gvfs-trash'
alias ch='chmod 777 '
alias ll='ls -alF'
alias la='ls -A'
alias l='ls -CF'
alias lh='ls -lh'
alias ga='git add'
alias gaf='git add -f'
alias gaa='git add .'
alias gcm='git commit --message'
alias gco='git checkout'
alias gd='git diff'
alias gp=gp #"git pull origin  $cb"
alias gps=gps # "git push origin $cb"
alias glg='git log --graph --oneline --decorate --all'
alias gs='git status'
alias gb='git branch'
alias gsl='git status .'
alias gcl='git clean -n'
alias vg='valgrind --leak-check=full -v --track-origins=yes --log-file=vg_logfile.out'
alias vpn='cd ~/Downloads && sudo openvpn --config gs-1920@gslab.com__ssl_vpn_config.ovpn'
alias tdr='cd ~/Documents/WorkSpace/tmp/sonde_gslab_dsp/fe_algorithm/sonde_library'
alias gpo='git pull origin '
alias gpso='git push origin'
alias gres='git restore --staged'
alias demo='~/Desktop/all/wav/demo.wav'
alias wav='~/Desktop/all/wav/'
alias pp='pwd'
alias tx='python3 /home/samir/Desktop/tools/transpose.py '
alias del='rm -rf '
alias cdfe='cd ~/Documents/workspace/fe_working/tmp_rg/sonde_gslab_dsp/fe_algorithm/sonde_library'
alias cdelck='cd ~/Documents/workspace/elck_working/tmp/sonde_gslab_dsp/dev_src'
alias cdw='cd /home/samir/Documents/workspace'
# Add an "alert" alias for long running commands.  Use like so:
#   sleep 10; alert
alias alert='notify-send --urgency=low -i "$([ $? = 0 ] && echo terminal || echo error)" "$(history|tail -n1|sed -e '\''s/^\s*[0-9]\+\s*//;s/[;&|]\s*alert$//'\'')"'

# Alias definitions.
# You may want to put all your additions into a separate file like
# ~/.bash_aliases, instead of adding them here directly.
# See /usr/share/doc/bash-doc/examples in the bash-doc package.

source ~/.gitreset

function lb
{
    if [ "$1" ] 
    then
    libreoffice $1  &
    fi
    
    if [ "$2" ] 
    then
    libreoffice $1 $2 &
    fi
}

function gp 
{
    
    cb=`git branch --show-current`
    if [ "$1" ] 
    then
	echo "doing git pull on $cb ..."  
        git pull origin $cb
    else
        echo $cb
    fi
}

function gopush
{
    
    if [ "$1" ] 
    then
         gp 1
         git add $1
         if [ "$2" ]
		 then
         	msg="`date +"%b %d %Y"` -- $2"
         else
         	msg="`date +"%b %d %Y"` ${1%.*}"
         fi
         git commit -m "$msg"
         gps 1
    else
        echo "Specify filename"
    fi
}


function gps 
{
    cb=`git branch --show-current`
    if [ "$1" ] 
    then
	echo "doing git push on $cb ..."  
        git push origin  $cb
    else
        echo $cb
    fi
}


if [ -f ~/.bash_aliases ]; then
    . ~/.bash_aliases
fi

# enable programmable completion features (you don't need to enable
# this, if it's already enabled in /etc/bash.bashrc and /etc/profile
# sources /etc/bash.bashrc).
if ! shopt -oq posix; then
  if [ -f /usr/share/bash-completion/bash_completion ]; then
    . /usr/share/bash-completion/bash_completion
  elif [ -f /etc/bash_completion ]; then
    . /etc/bash_completion
  fi
fi

# >>> conda initialize >>>
# !! Contents within this block are managed by 'conda init' !!
__conda_setup="$('/home/samir/anaconda3/bin/conda' 'shell.bash' 'hook' 2> /dev/null)"
if [ $? -eq 0 ]; then
    eval "$__conda_setup"
else
    if [ -f "/home/samir/anaconda3/etc/profile.d/conda.sh" ]; then
        . "/home/samir/anaconda3/etc/profile.d/conda.sh"
    else
        export PATH="/home/samir/anaconda3/bin:$PATH"
    fi
fi
unset __conda_setup
# <<< conda initialize <<<

alias vi=vim

export GOROOT=/usr/local/go
export GOPATH=$HOME/sdk/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
export PS1="\[\e[32m\]\u@\h \[\e[33m\]\W \[\e[37m\]$ "

declare -a DIRS

savedir() {
  local i
  for ((i=1;i<=9;i++)); do
    test "$1" = "${DIRS[$i]}" && return
  done
  for ((i=9;i>1;i--)); do
    DIRS[$i]="${DIRS[((i-1))]}"
  done
  DIRS[1]="$1"
}

showdirs() {
  local i=1
  while [ "${DIRS[$i]}" ]; do
    echo "$i: ${DIRS[$i]}"
    ((i++))
  done
}

gotodir() {
  local d
  showdirs
  printf "goto: "
  read -n 1 d
  echo
  cd "${DIRS[$d]}"
}

PROMPT_COMMAND=prompt_command
prompt_command() { savedir "$OLDPWD"; }

get_latest_release() {
  curl "https://api.github.com/repos/$1/releases/latest" | # Get latest release from GitHub api
    grep '"tag_name":' |                                            # Get tag line
    sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
}

alias cdh=gotodir
