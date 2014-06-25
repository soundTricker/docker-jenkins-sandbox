FROM google/golang

RUN mkdir /workspace

ENTRYPOINT ["/bin/bash", "-c"]
