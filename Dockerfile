#****************************
#* (C) Fabian Salamanca
#****************************

FROM debian:latest

#RUN apt-get update
RUN apt-get update && apt-get -y install git
RUN useradd -d /usr/local/goapps -m goapps
RUN su - goapps -c pwd
COPY goapp index.html /usr/local/goapps/
ADD img /usr/local/goapps/img/

EXPOSE 8188

ENTRYPOINT ["su","-","goapps","-c","/usr/local/goapps/goapp"]
