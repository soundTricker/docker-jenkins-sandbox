FROM google/golang

RUN mkdir /workspace

RUN echo '{"hoge":"hoge"}' > /workspace/test.json


ENTRYPOINT ["/bin/bash", "-c"]
