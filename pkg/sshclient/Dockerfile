FROM debian:10

MAINTAINER Geraldo Andrade "hi@geraldoandrade.com"

RUN apt-get update -y && apt-get upgrade -y && apt-get install -y \
  openssh-server \
  && rm -rf /var/lib/apt/lists/*

RUN echo 'root:password' | chpasswd
RUN mkdir /var/run/sshd
COPY sshd_config /etc/ssh/sshd_config

EXPOSE 22

CMD ["/usr/sbin/sshd", "-D"]