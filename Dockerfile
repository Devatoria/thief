FROM ubuntu:bionic

RUN apt-get update && apt-get install -y sysbench

COPY bin/thief /bin/thief

ENTRYPOINT ["/bin/thief"]
