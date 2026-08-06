FROM jenkins/jenkins:lts
USER root

RUN apt-get update && \
    apt-get install -y \
        git \
        golang-go \
        docker.io \
        curl && \
    rm -rf /var/lib/apt/lists/*

# Docker Compose'u sistemin genel bin klasörüne indirip yetki veriyoruz
RUN curl -SL "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \
    chmod +x /usr/local/bin/docker-compose
    
RUN usermod -aG docker jenkins
USER jenkins