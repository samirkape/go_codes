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

* scp   -i   pem_file.pem   user@ip_addr:/path/to/the/file  out_dir 
* scp   username@ip_addr:/path/to/the/file   out_dir 
* for file in `cat $csv_filename`; do echo $file; done
  * this will iterate over filenames provided in the csv_filename



