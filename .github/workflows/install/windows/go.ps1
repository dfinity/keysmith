$version="1.16"
$zipfile="go$version.windows-amd64.zip"
Push-Location $env:TEMP
Invoke-WebRequest `
    -OutFile "$zipfile" `
    -Uri "https://storage.googleapis.com/golang/$zipfile"
Add-Type -AssemblyName System.IO.Compression.FileSystem
[System.IO.Compression.ZipFile]::ExtractToDirectory("$zipfile", "$goroot")
Pop-Location
echo "GOROOT=$goroot" >> $env:GITHUB_ENV
echo "GOPATH=$gopath" >> $env:GITHUB_ENV
echo "$goroot/bin" >> $env:GITHUB_PATH
echo "$gopath/bin" >> $env:GITHUB_PATH
New-Item -ItemType "directory" -Name "go" -Path $env:HOME
