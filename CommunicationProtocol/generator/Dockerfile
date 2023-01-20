FROM golang:1.18

WORKDIR /work

RUN apt-get update
RUN apt install -y python3 python3-pip
RUN pip install qface==2.0.0
RUN pip install jinja2==2.10
RUN pip install markupsafe==2.0.1