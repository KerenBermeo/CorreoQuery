terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.37"
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "sa-east-1"
}

resource "aws_default_vpc" "default" {}

data "http" "my_ip" {
  url = "http://ipv4.icanhazip.com"
}

resource "aws_security_group" "ec2_sg" {
  name        = "emailquery-security-group"
  description = "Allow hhtp access on port 6002 for backend"
  vpc_id      = aws_default_vpc.default.id

  ingress {
    description = "http access"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "ssh access"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "zinc-server access"
    from_port   = 6001
    to_port     = 6001
    protocol    = "tcp"
    cidr_blocks = ["${chomp(data.http.my_ip.response_body)}/32"]
  }

  ingress {
    description = "api-back access"
    from_port   = 6002
    to_port     = 6002
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "api-front access"
    from_port   = 6003
    to_port     = 6003
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    project = "correo-query"
  }
}

resource "aws_iam_instance_profile" "ec2_deployer_user" {
  name = "terraform_new_name_user"
}

data "aws_ami" "ubuntu_ami" {
  most_recent = true
  owners      = ["099720109477"]  # Propietario de las AMIs de Ubuntu (Canonical)

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-*-20.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  filter {
    name   = "architecture"
    values = ["x86_64"]
  }
}

resource "aws_key_pair" "my_key_pair" {
  key_name   = "nuevo_nombre"
  public_key = "${file("../../../.ssh/id_rsa.pub")}"
}


resource "aws_instance" "email_query_ec2" {
  ami                    = data.aws_ami.ubuntu_ami.id
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.ec2_sg.id]
  iam_instance_profile   = aws_iam_instance_profile.ec2_deployer_user.name
  key_name               = aws_key_pair.my_key_pair.key_name 

  user_data =  <<EOF
    #!/bin/bash
    set -ex
    sudo apt update && sudo apt upgrade -y
    sudo apt install apt-transport-https ca-certificates curl software-properties-common
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
    sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
    sudo apt update
    sudo apt install docker-ce
    sudo usermod -aG docker $USER            
    sudo systemctl start docker
    # Esperar 10 segundos para asegurar que Docker se inicie completamente
    sleep 10           
    sudo apt install git -y
    sudo curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose 
    docker run -d -p 4080:4080 --name zincsearch public.ecr.aws/zinclabs/zincsearch:latest
  EOF

  tags = {
    project = "email-query" 
  }
}





