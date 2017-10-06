FROM centos:7

MAINTAINER "Craig Wickesser" <codecraig@gmail.com>

ENV GOROOT=/usr/local/go
ENV GOPATH=/opt/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin
ENV GOVERSION=1.9.1

RUN yum -y update && yum -y upgrade && \
    yum -y install epel-release && \
    yum -y install wget git sudo bzip2 tmux gcc openssl sox sox-devel file && \
    wget http://downloads.xiph.org/releases/ogg/libogg-1.3.2.tar.gz && \
    tar xf libogg-1.3.2.tar.gz && \
    (cd libogg-1.3.2 && ./configure && make && make install) && \
    wget http://downloads.xiph.org/releases/vorbis/libvorbis-1.3.5.tar.gz && \
    (cd libvorbis-1.3.5 && ./configure && make && make install) && \
    wget http://downloads.xiph.org/releases/vorbis/vorbis-tools-1.4.0.tar.gz && \
    tar xf vorbis-tools-1.4.0.tar.gz && \
    (cd vorbis-tools-1.4.0 && ./configure && make && make install)


WORKDIR /tmp
RUN curl https://storage.googleapis.com/golang/go$GOVERSION.linux-amd64.tar.gz | tar -xz
RUN mkdir -p /opt
RUN mv ./go $GOROOT
RUN go version

RUN adduser docker && \
    usermod -G wheel,users docker && \
    usermod -G users root && \
    echo '%wheel ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers

RUN mkdir -p $GOPATH && chown docker.users $GOPATH

RUN echo 'export PATH=/usr/local/bin:$PATH' >> /etc/profile
RUN sed -i -e "s/Defaults    requiretty/#Defaults    requiretty/" /etc/sudoers

USER docker
WORKDIR /home/docker

CMD ["tail","-f","/dev/null"]
