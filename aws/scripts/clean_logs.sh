#!/usr/bin/env bash

for file in /home/ubuntu/www/xieyuanpeng.in/logs/*
do
    true > ${file}
done