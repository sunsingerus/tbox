#!/bin/bash

# Find and kill all kubectl processes
ps ax | grep kubectl | grep -v grep | awk '{print $1}' | xargs kill
