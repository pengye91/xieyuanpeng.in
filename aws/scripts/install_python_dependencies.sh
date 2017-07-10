#!/usr/bin/env bash
#sudo chown ec2-user:ec2-user /home/ec2-user/www
python3.6 -m venv /home/ec2-user/www/project-venv
sudo chown ec2-user:ec2-user /home/ec2-user/www/project-venv
sudo chown ec2-user:ec2-user /home/ec2-user/www/project-venv/*
source /home/ec2-user/www/project-venv/bin/activate
#pip install -U pip
#pip install -r /home/ec2-user/www/yukere_project_aws/requirements.txt
deactivate
