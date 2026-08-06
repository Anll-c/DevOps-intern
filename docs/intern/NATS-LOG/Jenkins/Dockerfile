FROM jenkins/jenkins:lts
USER root

RUN apt-get update && \
    apt-get install -y \
        git \
        golang-go \
        docker.io \
        curl && \
    rm -rf /var/lib/apt/lists/*

# jenkins kullanıcısının HOME'u /var/jenkins_home olduğu için
# Docker CLI plugin'lerini doğrudan oraya kuruyoruz (root'un home'una değil).
RUN mkdir -p /var/jenkins_home/.docker/cli-plugins

# Docker Compose v2 plugin (arm64/aarch64 için)
RUN curl -SL "https://github.com/docker/compose/releases/latest/download/docker-compose-linux-aarch64" \
        -o /var/jenkins_home/.docker/cli-plugins/docker-compose && \
    chmod +x /var/jenkins_home/.docker/cli-plugins/docker-compose

# Buildx plugin (arm64 için)
RUN curl -SL "https://github.com/docker/buildx/releases/latest/download/buildx-v0.17.1.linux-arm64" \
        -o /var/jenkins_home/.docker/cli-plugins/docker-buildx && \
    chmod +x /var/jenkins_home/.docker/cli-plugins/docker-buildx

RUN chown -R jenkins:jenkins /var/jenkins_home/.docker

RUN usermod -aG docker jenkins

# GitHub Branch Source + Pipeline: Multibranch pluginlerini otomatik kurmak istersen
# (opsiyonel, Jenkins arayüzünden de kurulabilir) install-plugins.sh kullanılabilir.
# Burada bilinçli olarak elle kurulmasını bırakıyoruz çünkü plugin sürüm uyumu
# Jenkins LTS sürümüne göre değişebiliyor.

USER jenkins