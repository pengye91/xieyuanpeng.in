#!/usr/bin/env bash

chmod +x /home/ubuntu/www/xieyuanpeng.in/main
mv /home/ubuntu/www/xieyuanpeng.in/env /home/ubuntu/www/xieyuanpeng.in/.env
sudo pkill supervisord
sudo pkill main
#echo "xixi"
/home/ubuntu/.local/bin/supervisord -c /home/ubuntu/www/xieyuanpeng.in/aws/supervisor/default.conf
#/home/ubuntu/www/xieyuanpeng.in/backend/main
