#!/bin/bash 

echo "Enter project directory name: "
read projectName

mkdir $GOPATH/src/$projectName

cd $GOPATH/src/$projectName

touch README.md

echo "
# $projectName
" > README.md

touch .gitignore

curl https://www.gitignore.io/api/go > .gitignore

git init

echo "Created project in $GOPATH/src/$projectName"
