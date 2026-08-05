FROM jenkins/jenkins:lts

USER root

# Gerekli paketler
RUN apt-get update && \
    apt-get install -y \
        git \
        curl \
        ca-certificates \
        docker.io \
        docker-compose-v2 && \
    rm -rf /var/lib/apt/lists/*

# Jenkins kullanıcısına Docker yetkisi
RUN usermod -aG docker jenkins

USER jenkins