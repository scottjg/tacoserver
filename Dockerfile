FROM scratch
ADD tacoserver /
ADD id_rsa /

EXPOSE 80
EXPOSE 22

ENTRYPOINT ["/tacoserver"]

