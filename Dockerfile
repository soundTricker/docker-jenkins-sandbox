FROM google/golang

RUN mkdir /workspace

RUN touch  /workspace/hoge.txt

RUN echo '{"hoge":"hoge"}' > /workspace/test.json


ENTRYPOINT ["/bin/bash", "-c"]
