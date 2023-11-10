FROM golang:bullseye

RUN apt-get update && \
    apt-get install git -y

RUN git clone https://github.com/FelixBrgm/.dotfiles.git ~/.dotfiles

RUN cd ~/.dotfiles && \
    chmod 777 ./setup.sh && \
    ./setup.sh

CMD ["tail", "-f", "/dev/null"]