#!/bin/bash
./build.sh -os linux -arch amd64
ssh root@172.81.240.81 "supervisorctl stop testwywlogin"
scp -r ./bin/* root@172.81.240.81:~
ssh root@172.81.240.81 "supervisorctl start testwywlogin"


