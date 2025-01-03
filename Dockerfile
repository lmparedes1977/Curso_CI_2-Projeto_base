FROM ubuntu:latest

EXPOSE 8000

WORKDIR /app

ENV HOST localhost 
ENV DBPORT 5432
ENV USER root
ENV PASSWORD root
ENV DBNAME root

COPY ./main main

RUN chmod +x main

COPY ./templates/ templates/

ENTRYPOINT [ "./main" ]
