FROM jenkins/jenkins:lts
USER root

RUN apt-get update && \
    apt-get install -y \
        git \
        golang-go \
        docker.io \
        curl && \
    rm -rf /var/lib/apt/lists/*

# 1. Docker Compose'u sistemin genel bin klasörüne indirip yetki veriyoruz
RUN curl -SL "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \
    chmod +x /usr/local/bin/docker-compose

# 2. Buildx eklentisini de Docker CLI plugin klasörüne indirip güncel versiyonunu kuruyoruz
RUN mkdir -p /root/.docker/cli-plugins && \
    curl -SL "https://github.com/docker/buildx/releases/latest/download/buildx-v0.17.1.linux-arm64" -o /root/.docker/cli-plugins/docker-buildx && \
    chmod +x /root/.docker/cli-plugins/docker-buildx

# Aynı eklentiyi jenkins kullanıcısının ev dizinine de kopyalayarak yetki sorununun önüne geçiyoruz
RUN mkdir -p /var/jenkins_home/.docker/cli-plugins && \
    cp /root/.docker/cli-plugins/docker-buildx /var/jenkins_home/.docker/cli-plugins/docker-buildx && \
    chown -R jenkins:jenkins /var/jenkins_home/.docker

RUN usermod -aG docker jenkins
USER jenkins