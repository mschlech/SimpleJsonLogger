FROM       scratch
MAINTAINER Marcus Schlechter <marcus.schlechter@gmail.com>
ADD        SimpleJsonLogger SimpleJsonLogger
ENV        PORT 5000
EXPOSE     5000
ENTRYPOINT ["/SimpleJsonLogger"]
