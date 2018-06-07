FROM iron/base
WORKDIR /app
COPY userapp /app/
EXPOSE 8080
ENTRYPOINT ["./userapp"]
