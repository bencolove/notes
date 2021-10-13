# Custom `jupyter/datascience-notebook` with `pyodbc`
* build an IMAGE
* run an IMAGE

---

## The `Dockerfile`
_`Dockerfile`_
```txt
# base IMAGE (do not use TAG latest)
FROM jupyter/datascience-notebook:python-3.9.7

LABEL description="jupyter/datascience-notebook + odbc preinstalled"

USER root

ENV ACCEPT_EULA=Y

# https://docs.microsoft.com/en-us/sql/connect/odbc/linux-mac/installing-the-microsoft-odbc-driver-for-sql-server?view=sql-server-ver15
# unixodbc-dev and python3-dev are needed to build
# needed
RUN apt-get update && apt-get install -y curl gnupg2 && \
	curl https://packages.microsoft.com/keys/microsoft.asc | apt-key add - && \
	curl https://packages.microsoft.com/config/ubuntu/20.04/prod.list > /etc/apt/sources.list.d/mssql-release.list && \
	apt-get update && \
	apt-get -y --no-install-recommends install libssl1.1 libssl-dev \
		unixodbc-dev unixodbc msodbcsql17 \
		python3-dev && \
	rm -rf /var/lib/apt/lists/* 
# optional RUN ACCEPT_EULA=Y apt-get install mssql-tools
RUN echo 'export PATH="$PATH:/opt/mssql-tools/bin"' >> ~/.bashrc
```

> Dependencies
* unixodbc
* msodbcsql17  
* unixodbc-dev
* python3-dev

> Build from a Dockerfile  
`docker build -f the_docker_file -t your_tag .`  

This command will build an IMAGE locally stored (hash named) with `your_tag` tag.

--- 

## run the IMAGE
```sh
docker run --name jupyter-lab-odbc-py397 \
-d \
--add-host=host.docker.internal:host-gateway \
-p 8888:8888 \
-v /mnt/d/code/jupyter/docker-notebook:/home/jovyan/work \
-e GRANT_SUDO=yes \
-e JUPYTER_ENABLE_LAB=yes \
jupyter/datascience-notebook:odbc-py397
```
  
>-e GRANT_SUDO=yes: passwordless user jovyan  
--user: run as root  
--add-host: modify container's hosts  

[build-odbc]: https://docs.microsoft.com/en-us/sql/connect/odbc/linux-mac/installing-the-microsoft-odbc-driver-for-sql-server?view=sql-server-ver15


