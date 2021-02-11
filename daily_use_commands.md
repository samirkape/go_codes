1. `sed -i 's/search/replace/g' file.txt`
   1. -i - inplace
2. `sed -i '/delete-line-by-text-token/d' file.txt`
   1. No substitute (s) at star
3. PS1="%B%{$fg[red]%}[%{$fg[yellow]%}%n%{$fg[green]%}@%{$fg[blue]%}%M %{$fg[red]%}4.7.1]%{$reset_color%}$%b "