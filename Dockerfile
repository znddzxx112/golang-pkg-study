FROM centos7_golang:1.12.1

COPY ./main /workspace

CMD ["/workspace/main"]