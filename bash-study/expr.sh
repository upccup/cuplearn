#!/bin/sh

a=10

while [ $a -gt -2 ]
do
       echo $a
          a=`expr $a - 1`
      done
