FROM jenkins/jenkins:lts
USER root

RUN apt-get update && \
    apt-get install -y \
        git \
        golang-go \
        docker.io \
        curl && \
    rm -rf /var/lib/apt/lists/*

# 1. Docker Compose v2'yi CLI plugin olarak kuruyoruz (root için)
RUN mkdir -p /root/.docker/cli-plugins && \
    curl -SL "https://github.com/docker/compose/releases/latest/download/docker-compose-linux-aarch64" \
        -o /root/.docker/cli-plugins/docker-compose && \
    chmod +x /root/.docker/cli-plugins/docker-compose

# 2. Buildx eklentisini CLI plugin klasörüne indirip kuruyoruz (root için)
RUN mkdir -p /root/.docker/cli-plugins && \
    curl -SL "https://github.com/docker/buildx/releases/latest/download/buildx-v0.17.1.linux-arm64" \
        -o /root/.docker/cli-plugins/docker-buildx && \
    chmod +x /root/.docker/cli-plugins/docker-buildx

# 3. Aynı eklentileri jenkins kullanıcısının ev dizinine de kopyalıyoruz
RUN mkdir -p /var/jenkins_home/.docker/cli-plugins && \
    cp /root/.docker/cli-plugins/docker-buildx /var/jenkins_home/.docker/cli-plugins/docker-buildx && \
    cp /root/.docker/cli-plugins/docker-compose /var/jenkins_home/.docker/cli-plugins/docker-compose && \
    chown -R jenkins:jenkins /var/jenkins_home/.docker

RUN usermod -aG docker jenkins
USER jenkins