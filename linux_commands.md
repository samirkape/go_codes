```bash
whereis -b yq
# finds all programs with name containing yq 
```

```bash
info command 
# produces better command information than man 
e.g info grep
```

```bash
cut -f2 -d, csv_file.csv
#-f = column number 
#-d = delimiter
#cuts the specified column out of csv
```

 ```bash
 mkdir -p gp/parent/child  
 # -p = while creating child, 
 # creates gp and parent if not already present
 ```

 ```bash
#check the return type of any command executed
e.g 
mv x y 
result=$?   
if [ $result -ne 0 ]    # check if move is successful
then
    echo " "
    echo "error occured!"
    exit $result
fi
 ```

 ```bash
find . -type f -name "*.yaml" -exec echo {} \; -exec yq r {} version \; -exec echo "" \;
 ```

 ```bash
date +"%c"  #Current date-time-year
 ```

 ```bash
if [ "$#" -eq 0 ] #check number of input arguments
 ```

 ```bash
var=1 && for i in *.wav; do mv $i "noise_ahh_${var}.wav"; var=$((var+1)); done
  #rename files 1..n
 ```

 ```bash
X="filename.txt"
${X%.*}  -- filename 
 ```

 ```bash
sed -i 's/search/replace/g' file.txt
  #-i - inplace
 ```

 ```bash
sed -i '/delete-line-by-text-token/d' file.txt 
#No substitute (s/) at start
 ```

 ```bash
sed -i '/*----*/d' file.txt
  #asterisk = wildcard 
 ```

```bash
#wildcards
Wild-cards are handled completely by the shell before the associated
program even runs

* -- Zero or more consecutive characters

? -- Any single character

[set] -- Any single character in the given set, most commonly a 
	sequence of characters, like [aeiouAEIOU] for all vowels, or 
	a range with a dash, like [A-Z] for all capital letters

[^set] -- Any single character not in the given set, such as [^0-9] to mean any nondigit
```

 ```bash
head -n 1 filename 
 ```

 ```bash
scp -i pem_file.pem user@ip_addr:/path/to/the/file  out_dir 
 ```

 ```bash
scp username@ip_addr:/path/to/the/file out_dir 
 ```

 ```bash
var="bar"
echo "foo_${var}
# "foo_$var" does not work while concatenating  
 ```

 ```bash
for file in `cat $csv_filename`; do echo $file; done
#this will iterate over filenames provided in the csv_filename
 ```

```bash
export PS1="\[\e[32m\]\u@\h \[\e[34m\]\W \[\e[32m\]$ "
```
Terminal Shortcuts
|                                                              |                             |
| ------------------------------------------------------------ | --------------------------- |
| Go to beginning of command line                              | ctrl + a                    |
| Go to end of command line                                    | ctrl + e                    |
| Move back one character                                      | ctrl + b                    |
| Move forward one character                                   | ctrl + f                    |
| Move cursor forward one word                                 | alt + f or ctrl + right     |
| Move cursor back one word                                    | alt + b or ctrl + left      |
| Toggle between current cursor position and beginning of the line (Bash) | ctrl + x ctrl + x           |
| **Control**                                                  |                             |
| Clears the Screen                                            | ctrl + l                    |
| Pause terminal output                                        | ctrl + s                    |
| Resume terminal output after it was paused                   | ctrl + q                    |
| **Editing**                                                  |                             |
| Undo last action                                             | ctrl + shift + -            |
| Swap the last two characters before the cursor               | ctrl + t                    |
| Swap current word with previous                              | alt + t                     |
| Delete everything forward to end of line                     | ctrl + k                    |
| Delete backward to the beginning of the current word         | ctrl + w or alt + backspace |
| Delete the character after the current cursor position or exit shell | ctrl + d                    |
| Paste whatever was cut by the last cut command               | ctrl + y                    |
| Delete to the beginning of the line                          | ctrl + u                    |
| Delete to the end of the line                                | ctrl + k                    |
| **History**                                                  |                             |
| Search command history                                       | ctrl + r                    |
| View previous command in the history                         | ctrl + p or up              |
| View next command in the history                             | ctrl + n or down            |
| **Processes**                                                |                             |
| Puts current process into a suspended background process     | ctrl + z                    |
| Kill the currently running process                           | ctrl + c                    |
|                                                              |                             |