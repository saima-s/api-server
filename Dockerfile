FROM ubuntu

COPY ./api-server ./api-server

ENTRYPOINT [ "./api-server" ]