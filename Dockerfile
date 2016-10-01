FROM scratch
MAINTAINER pdevine@sonic.net

ADD parrot parrot
ENTRYPOINT ["/parrot"]

