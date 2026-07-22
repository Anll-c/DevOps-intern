FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
	&& apt-get install -y --no-install-recommends openssh-server openssh-client \
	&& rm -rf /var/lib/apt/lists/*

RUN mkdir -p /var/run/sshd \
	&& useradd -m -s /bin/bash usr2 \
	&& mkdir -p /home/usr2/.ssh \
	&& chown -R usr2:usr2 /home/usr2/.ssh \
	&& chmod 700 /home/usr2/.ssh \
	&& printf '\nPasswordAuthentication no\nPubkeyAuthentication yes\nPermitRootLogin no\n' >> /etc/ssh/sshd_config

EXPOSE 22

CMD ["/usr2/sbin/sshd", "-D", "-e"]