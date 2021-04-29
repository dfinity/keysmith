$arch="amd64"
$version="1.16"
$zipfile="go$version.windows-$arch.zip"
$goroot="C:\go"
$gopath="$home\go"
Push-Location $env:TEMP
$client = new-object System.Net.WebClient
$client.DownloadFile("https://storage.googleapis.com/golang/$zipfile", $zipfile)
Add-Type -AssemblyName System.IO.Compression.FileSystem
[System.IO.Compression.ZipFile]::ExtractToDirectory($zipfile, $pwd)
Move-Item -Path go -Destination $goroot
Pop-Location
echo "GOROOT=$goroot" >> $env:GITHUB_ENV
echo "GOPATH=$gopath" >> $env:GITHUB_ENV
echo "$goroot/bin" >> $env:GITHUB_PATH
echo "$gopath/bin" >> $env:GITHUB_PATH
New-Item -ItemType "directory" -Name "go" -Path $home
