FROM golang:1.14
RUN \
  apt-get update \
  && apt-get install -y libnetcdf-dev xz-utils cmake \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*
RUN \
  cd /tmp \
  && wget https://bitbucket.org/cnes_aviso/fes/downloads/fes-2.9.3-Source.tar.xz \
  && tar -xJf fes-2.9.3-Source.tar.xz \
  && cd fes-2.9.3-Source \
  && mkdir build \
  && cd build \
  && cmake .. -DBUILD_PYTHON=off -DBUILD_SHARED_LIBS=on -DCMAKE_INSTALL_PREFIX=/usr \
  && make -j \
  && make install \
  && cd /tmp \
  && rm -r fes-2.9.3-Source*

ENV GOFES_SRC=/go/src/github.com/molflow/gofes
ADD . $GOFES_SRC

RUN cd $GOFES_SRC && make build && make install
