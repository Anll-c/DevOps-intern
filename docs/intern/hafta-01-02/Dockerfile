FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
	&& apt-get install -y --no-install-recommends openssh-server openssh-client \
	&& rm -rf /var/lib/apt/lists/*

RUN mkdir -p /var/run/sshd \
	&& useradd -m -s /bin/bash usr \
	&& mkdir -p /home/usr/.ssh \
	&& chown -R usr:usr /home/usr/.ssh \
	&& chmod 700 /home/usr/.ssh \
	&& printf '\nPasswordAuthentication no\nPubkeyAuthentication yes\nPermitRootLogin no\n' >> /etc/ssh/sshd_config

EXPOSE 22

CMD ["/usr/sbin/sshd", "-D", "-e"]