### one-liners

* ```bash
  find . -type f -name "*.yaml" -exec echo {} \; -exec yq r {} version \; -exec echo "" \;
  ```
  
* ```bash
  if [ "$#" -eq 0 ] #check number of input arguments
  ```
  
* ```bash
  var=1 && for i in *.wav; do mv $i "noise_ahh_${var}.wav"; var=$((var+1)); done
  #rename files 1..n
  ```

* ```bash
  X="filename.txt"
  ${X%.*}  -- filename 
  ```
  
* ```bash
  sed -i 's/search/replace/g' file.txt
  -i - inplace
  ```

* ```bash
  sed -i '/delete-line-by-text-token/d' file.txt
  ```
  * No substitute (s) at start
  
* ```bash
  sed -i '/*----*/d' file.txt
  ```
  
  * asterisk = wildcard
  
* ```bash
  head -n 1 filename 
  ```

* ```bash
  scp -i pem_file.pem user@ip_addr:/path/to/the/file  out_dir 
  ```
  
* ```
  scp username@ip_addr:/path/to/the/file out_dir 
  ```
  
* ```bash
  for file in `cat $csv_filename`; do echo $file; done
  ```
  
  * this will iterate over filenames provided in the csv_filename



