#!/bin/bash

echo -n "Version to tag: "
read TAG

echo -n "Name of release: "
read NAME

echo -n "Desc of release: "
read DESC

git tag ${TAG}
git push --tags

if [ ! -d ./bin ]; then
	mkdir bin
else
	rm -rf ./bin/*
fi

echo -n "Building binaries ... "

rice embed-go

GOOS=linux GOARCH=amd64 go build -o ./bin/static-Linux-x86_64 .
GOOS=linux GOARCH=arm64 go build -o ./bin/static-Linux-x86_64 .
GOOS=darwin GOARCH=amd64 go build -o ./bin/static-Darwin-x86_64 .
GOOS=windows GOARCH=amd64 go build -o ./bin/static-Windows-x86_64.exe .

echo "DONE"

echo -n "Uploading binaries ... "

github-release release \
    -u prologic -p -r static \
    -t ${TAG} -n "${NAME}" -d "${DESC}"

for file in bin/*; do
    name="$(echo $file | sed -e 's|bin/||g')"
    github-release upload -u prologic -r static -t ${TAG} -n $name -f $file
done

echo "DONE"
