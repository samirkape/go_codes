* X="filename.txt"
  
  * ${X%.*}  -- filename 
  
* `sed -i 's/search/replace/g' file.txt`
  
* -i - inplace
  
* `sed -i '/delete-line-by-text-token/d' file.txt`
  
* No substitute (s) at star
  
* `sed -i '/*----*/d' file.txt`
  
  * asterisk = wildcard
  
* ```sh
  head -n 1 filename 
  ```





