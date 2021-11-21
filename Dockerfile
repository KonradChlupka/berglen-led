# Stage 0 : Build the C library 

FROM debian:buster-20190506 AS lib_builder

WORKDIR /foundry

RUN apt-get update -y && apt-get install -y \
  build-essential \
  cmake \
  git

RUN git clone https://github.com/jgarff/rpi_ws281x.git \
  && cd rpi_ws281x \ 
  && mkdir build \
  && cd build \ 
  && cmake -D BUILD_SHARED=OFF -D BUILD_TEST=OFF .. \
  && cmake --build . \
  && make install

# Stage 1 : Build a go image with the rpi_ws281x C library and the go wrapper

FROM golang:1.17 AS go_builder
COPY --from=lib_builder /usr/local/lib/libws2811.a /usr/local/lib/
COPY --from=lib_builder /usr/local/include/ws2811 /usr/local/include/ws2811

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o "led-lights" -v

# Stage 2 : Export compiled go binary with the light stuff

FROM scratch AS export-stage

COPY --from=go_builder /app/led-lights .
