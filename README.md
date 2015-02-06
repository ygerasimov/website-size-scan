# Website pages size scanner

Script takes file with list of links and estimates size of each page writing to outpt file in CSV format.

What do we count:
* size of the html of the page
* all CSS, JS files mentioned in html
* all images embedded as <img> tags

TODO:
* count images embedded as background
