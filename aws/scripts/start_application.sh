#!/usr/bin/env bash

cd /home/ec2-user/www/yukere_project_aws/
source /home/ec2-user/www/project-venv/bin/activate
#./manage.py collectstatic --no-input
sudo pkill gunicorn
/home/ec2-user/www/project-venv/bin/gunicorn config.wsgi:application -w 2 -k meinheld.gmeinheld.MeinheldWorker -b 0.0.0.0:8001 -t 300 -D
