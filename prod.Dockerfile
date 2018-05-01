# build stage
FROM catcatio/chatbot AS builder

# ADD . /app
# RUN cd /app && go build -o goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/goapp /app/
ENTRYPOINT ./goapp