FROM gcr.io/distroless/static-debian11

COPY cli-datacontract /usr/bin/cli-datacontract

USER foo:foo

WORKDIR /home/foo

ENTRYPOINT [ "/usr/bin/cli-datacontract" ]

CMD ["serve"]

