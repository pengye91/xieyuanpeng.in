#!/usr/bin/env bash
cd /home/ec2-user/www/yukere_project_aws/
source /home/ec2-user/www/project-venv/bin/activate
./manage.py makemigrations
./manage.py migrate_schemas --database=patients
./manage.py migrate_schemas --shared
./manage.py migrate_schemas
deactivate
