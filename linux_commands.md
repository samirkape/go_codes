## Linux Commands That Needs To Be Recorded 

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

<details><summary>Terminal Shortcuts List</summary>
	
```text

Left            Move back one character
Right           Move forward one character
Ctrl+b          Move back one character
Ctrl+f          Move forward one character

Alt+Left        Move back one word
Alt+Right       Move forward one word
Alt+b           Move back one word
Alt+f           Move forward one word

Cmd+Left        Move cursor to start of line
Cmd+Right       Move cursor to end of line
Ctrl+a          Move cursor to start of line
Ctrl+e          Move cursor to end of line

Ctrl+d          Delete character after cursor
Backspace       Delete character before cursor

Alt+Backspace   Delete word before cursor
Ctrl+w          Delete word before cursor
Alt+w           Delete word before the cursor
Alt+d           Delete word after the cursor

Cmd+Backspace   Delete everything before the cursor
Ctrl+u          Delete everything before the cursor
Ctrl+k          Delete everything after the cursor

Ctrl+l          Clear the terminal

Ctrl+c          Cancel the command
Ctrl+y          Paste the last deleted command
Ctrl+_          Undo

Ctrl+r          Search command in history - type the search term
Ctrl+j          End the search at current history entry and run command
Ctrl+g          Cancel the search and restore original line

Up              previous command from the History
Down            Next command from the History
Ctrl+n          Next command from the History
Ctrl+p          previous command from the History

Ctrl+xx         Toggle between first and current position
```
</details>

<details><summary>One More List</summary>

* `Ctrl+a` Move cursor to start of line
* `Ctrl+e` Move cursor to end of line
* `Ctrl+b` Move back one character
* `Alt+b` Move back one word
* `Ctrl+f` Move forward one character
* `Alt+f` Move forward one word
* `Ctrl+d` Delete current character
* `Ctrl+w` Cut the last word
* `Ctrl+k` Cut everything after the cursor
* `Alt+d` Cut word after the cursor
* `Alt+w` Cut word before the cursor
* `Ctrl+y` Paste the last deleted command
* `Ctrl+_` Undo
* `Ctrl+u` Cut everything before the cursor
* `Ctrl+xx` Toggle between first and current position
* `Ctrl+l` Clear the terminal
* `Ctrl+c` Cancel the command
* `Ctrl+r` Search command in history - type the search term
* `Ctrl+j` End the search at current history entry
* `Ctrl+g` Cancel the search and restore original line
* `Ctrl+n` Next command from the History
* `Ctrl+p` previous command from the History

</details>
