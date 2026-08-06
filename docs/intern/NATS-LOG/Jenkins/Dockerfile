FROM jenkins/jenkins:lts
USER root

RUN apt-get update && \
    apt-get install -y \
        git \
        golang-go \
        docker.io \
        docker-compose-plugin \
        docker-buildx-plugin \
        curl && \
    rm -rf /var/lib/apt/lists/*

RUN usermod -aG docker jenkins
USER jenkins