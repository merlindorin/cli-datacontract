FROM scratch

COPY cli-datacontract /usr/bin/cli-datacontract

ENTRYPOINT [ "/usr/bin/cli-datacontract" ]

CMD ["serve"]

