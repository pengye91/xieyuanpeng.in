#!/usr/bin/env bash
#cd /home/ec2-user/www/yukere_project_aws/
#source /home/ec2-user/www/project-venv/bin/activate
## supervisorctl -c /home/ec2-user/www/project/supervisor/default.conf stop all 2&>1 >/dev/null
## sudo unlink /tmp/supervisor.sock 2> /dev/null
sudo nginx -s stop
sudo pkill gunicorn*
