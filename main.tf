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

  user_data = <<EOF
    #!/bin/bash
    sudo apt-get update
    sudo apt-get install -y docker-compose
    git clone https://github.com/tu_usuario/tu_repositorio.git
    cd tu_repositorio

    # Variables de entorno para el backend
      export ZINC_FIRST_ADMIN_USER=${var.zinc_first_admin_user}
      export ZINC_FIRST_ADMIN_PASSWORD=${var.zinc_first_admin_password}
      export ZINC_SERVER_NAME_INDEX=${var.zinc_server_name_index}
      export ZINC_SEARCH_URL=${var.zinc_search_url}
      export BACK_ROOT_DIRECTORY="${var.data_s3_url}"
      export BACK_LISTEN_SERVER=${var.back_listen_server}
      export VITE_API_URL=${var.vite_api_url}
              
      # Mover las variables de entorno del backend a la ubicación deseada
      echo "export ZINC_FIRST_ADMIN_USER=${var.zinc_first_admin_user}" >> /path/to/backend/vars.env
      echo "export ZINC_FIRST_ADMIN_PASSWORD=${var.zinc_first_admin_password}" >> /path/to/backend/vars.env
      echo "export ZINC_SERVER_NAME_INDEX=${var.zinc_server_name_index}" >> /path/to/backend/vars.env
      echo "export ZINC_SEARCH_URL=${var.zinc_search_url}" >> /path/to/backend/.env
      echo "export BACK_ROOT_DIRECTORY=${var.back_root_directory}" >> /path/to/backend/vars.env
      echo "export BACK_LISTEN_SERVER=${var.back_listen_server}" >> /path/to/backend/vars.env
      echo "export VITE_API_URL=${var.vite_api_url}" >> /path/to/api-front/.env
    docker-compose up -d
  EOF

  tags = {
    project = "email-query" 
  }
}





