FROM postgres:16.1

RUN apt-get update \
	&& apt-get install -y neovim\
	&& rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*

COPY init.sql /docker-entrypoint-initdb.d/
