FROM ubuntu

WORKDIR /home
RUN apt-get update && apt-get install -y build-essential git curl unzip && \
  curl -vL https://github.com/yahoojapan/NGT/archive/v1.1.0.zip -O && \
  unzip ./v1.1.0.zip && rm ./v1.1.0.zip && \
  curl https://cmake.org/files/v3.8/cmake-3.8.2-Linux-x86_64.tar.gz -O && \
  tar -xvf ./cmake-3.8.2-Linux-x86_64.tar.gz && rm -f ./cmake-3.8.2-Linux-x86_64.tar.gz && \
  cd NGT-1.1.0 && mkdir build && cd build && \
  /home/cmake-3.8.2-Linux-x86_64/bin/cmake .. && \
  make && make install && \
  export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib && \
  rm -rf /home/cmake-3.8.2-Linux-x86_64
ENV LD_LIBRARY_PATH /usr/local/lib:${LD_LIBRARY_PATH}
RUN apt-get install -y golang && \
    mkdir -p /home/go && export GOPATH=/home/go && go get github.com/zenazn/goji
ENV GOPATH /home/go
COPY ./app.go /home/app.go
EXPOSE 8000
CMD ["go", "run", "/home/app.go"]
