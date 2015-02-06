# Website pages size scanner

Script takes file with list of links and estimates size of each page writing to outpt file in CSV format.

What do we count:
* size of the html of the page
* all CSS, JS files mentioned in html
* all images embedded as <img> tags

TODO:
* count images embedded as background* 

How-to compile a project
=====

Install golang to your system.
For Ubuntu 12-14 LTS your need to run
```sh
sudo apt-get install golang
```

Clone current repo to some directory. Let it be ```/var/www/website-size-scan```

Now do prepare workspace
```sh
export GOPATH=/var/www/website-size-scan/
```

Download needed dependency
```sh
cd /var/www/website-size-scan/
go get github.com/PuerkitoBio/goquery
```

Now you are ready for compiling go code within workspace directory.
For current project run a command
```sh
go build scan-website.go
```
and you'll get executable program nearby, named ```scan-website```
