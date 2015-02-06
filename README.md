# Website pages size scanner

Script takes file with list of links and estimates size of each page writing to output file in CSV format.

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
go build
```
and you'll get executable program nearby, named ```website-size-scan```

Command line options for compiled program
=====

```sh
# ./website-size-scan -h
Usage of ./website-size-scan:
  -baseUrl="http://www.example.com": Will be used to adjust relative paths.
  -count=100: Number of URLs in the input file
  -inputFile="./links.csv": Will be used as source file for links.
  -outputFile="./output.csv": Will be used as output for results.
  -workers=5: Number of workers. Be careful with this number. Big number will put your site down.
```

Video how-to "going"
=====

Use this link to watch how-to compile go programs
https://www.youtube.com/watch?v=XCsL89YtqCs
