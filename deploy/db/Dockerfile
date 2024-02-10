FROM postgres:16.1
RUN apt-get update \
	&& apt-get install -y neovim
COPY init.sql /docker-entrypoint-initdb.d/
