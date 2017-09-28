#!/bin/bash

f=$1
echo "upload $f to 90FTP"

ftp -n<<!
open 192.168.1.90 21
user developer developer
binary
cd server
prompt
put $f $f
close
bye
!
echo "success"
