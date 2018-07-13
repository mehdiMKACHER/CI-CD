#!/bin/bash

. password.sh

git add .
git commit -m "new commit"
git push -u origin master
curl -X GET http://192.168.56.116:8080/job/Demo-K8s-Jenkins/build?token=fabs --user "fsalaman:${PASSWORD}"
