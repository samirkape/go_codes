## Linux Commands That Needs To Be Recorded 

 ```bash
  find . -type f -name "*.yaml" -exec echo {} \; -exec yq r {} version \; -exec echo "" \;
 ```

 ```bash
  date +"%c" 
  #Current date-time-year
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
  #No substitute (s) at start
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



<details><summary>Log</summary>
<p>
* Clean up the line: You can use `Ctrl+U` to clear up to the beginning.
* Clean up the line: `Ctrl+E Ctrl+U` to wipe the current line in the terminal
* Clean up the line: `Ctrl+A Ctrl+K` to wipe the current line in the terminal
* Cancel the current command/line: `Ctrl+C`.
* Recall the deleted command: `Ctrl+Y (then Alt+Y)`
* Go to beginning of the line: `Ctrl+A`
* Go to end of the line: `Ctrl+E`
* Remove the forward words for example, if you are middle of the command: `Ctrl+K`
* Remove characters on the left, until the beginning of the word: `Ctrl+W`
* To clear your entire command prompt: `Ctrl + L`
* Toggle between the start of line and current cursor position: `Ctrl + XX`
</p>
</details>




* Clean up the line: You can use `Ctrl+U` to clear up to the beginning.
* Clean up the line: `Ctrl+E Ctrl+U` to wipe the current line in the terminal
* Clean up the line: `Ctrl+A Ctrl+K` to wipe the current line in the terminal
* Cancel the current command/line: `Ctrl+C`.
* Recall the deleted command: `Ctrl+Y (then Alt+Y)`
* Go to beginning of the line: `Ctrl+A`
* Go to end of the line: `Ctrl+E`
* Remove the forward words for example, if you are middle of the command: `Ctrl+K`
* Remove characters on the left, until the beginning of the word: `Ctrl+W`
* To clear your entire command prompt: `Ctrl + L`
* Toggle between the start of line and current cursor position: `Ctrl + XX`