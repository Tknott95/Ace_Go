#!/bin/bash

echo "EC2 Deploy script initialized..."

cd $GWS/Ace_Go

echo "Removing prior Binary"
rm -f Ace_Go

go build

cd ~/Gnomespace/EC2/

echo "Removing Prior Zip and Proj. in:  `pwd`"
rm -rf AceEC2.zip
rm -rf Ace_Go

cd $GWS

cp -rf Ace_Go/ ~/Gnomespace/EC2/

cd ~/Gnomespace/EC2/
zip -r AceEC2.zip Ace_Go/

echo "ACE zip ready to deploy to aws in:  `pwd`"
