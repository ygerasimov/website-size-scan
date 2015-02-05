#!/bin/bash
# Info: https://github.com/mariomaric/website-size#readme

# Prepare wget logfile
log="/tmp/wget-website-size-log$RANDOM"

# Do the spider magic
#echo "### Crawling ${!#} website... ###"
#sleep 2s
  
#--recursive --level=inf \

wget \
    --accept=jpg,jpeg,gif,png \
    --recursive --level=1 \
    --spider --server-response \
    --no-directories \
    --output-file="$log" "$@"

# Check if prepared logfile is used
if [ -f "$log" ]
then
    # Calculate and print estimated website size
    echo "${!#}, $(\
        grep -e "Content-Length" "$log" | \
        awk '{sum+=$2} END {printf("%.2f", sum / 1024 / 1024)}'\
    )"

    # Delete wget log file
#    rm "$log"
else
    echo "Unable to calculate estimated size."
    exit 1
fi

exit 0
