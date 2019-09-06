FROM ubuntu:18.04

RUN apt-get update \
    && apt-get install wget curl gnupg2 -y
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - \
    && echo "deb https://dl.yarnpkg.com/debian/ stable main" > /etc/apt/sources.list.d/yarn.list
RUN apt-get update \
    && apt-get install pandoc make python python-pip nodejs yarn git -y
RUN wget -nv https://dl.google.com/go/go1.13.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go*.tar.gz \
    && rm go*.tar.gz \
    && wget -nv https://github.com/pulumi/tf2pulumi/releases/download/v0.5.1/tf2pulumi-v0.5.1-linux-x64.tar.gz \
    && tar -C /usr/local/bin -xzf tf2pulumi*.tar.gz \
    && rm tf2pulumi*.tar.gz
ENV PATH="$PATH:/usr/local/go/bin:/root/go/bin" 

