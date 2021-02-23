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



<details><summary>Log</summary>
<p>

​```python
print("hello world!")
​```

</p>
</details>
```

