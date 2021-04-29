#!/usr/bin/env pwsh

# Define operating system and architecture.
$goos="windows"
$goarch="amd64"

# Define version and zip file.
$version="1.16"
$zipfile="go$version.$goos-$goarch.zip"

# Install Go compiler.
Push-Location $env:TEMP
$client = new-object System.Net.WebClient
$client.DownloadFile("https://storage.googleapis.com/golang/$zipfile", $zipfile)
Add-Type -AssemblyName System.IO.Compression.FileSystem
[System.IO.Compression.ZipFile]::ExtractToDirectory($zipfile, "C:\")
Pop-Location

# Configure workspace.
$goroot="C:\go"
$gopath="$home\go"
echo "GOOS=$goos" >> $env:GITHUB_ENV
echo "GOARCH=$goarch" >> $env:GITHUB_ENV
echo "GOROOT=$goroot" >> $env:GITHUB_ENV
echo "GOPATH=$gopath" >> $env:GITHUB_ENV
echo "$goroot/bin" >> $env:GITHUB_PATH
echo "$gopath/bin" >> $env:GITHUB_PATH
New-Item -ItemType "directory" -Name "go" -Path $home
