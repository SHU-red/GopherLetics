#!/bin/bash

# build windows binary
sudo /home/shured/go/bin/fyne-cross windows -arch=amd64,386 -app-id=shur3d.GopherLetics

# build linux binary
go build -o GopherLetics .

# move linux build to same folderstructure
mkdir -p ./fyne-cross/dist/linux
zip ./fyne-cross/dist/linux/GopherLetics_linux.zip ./GopherLetics_linux
