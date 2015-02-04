FROM busybox:latest
MAINTAINER Marcel de Graaf <marcel@yourkarma.com>

ENV PORT=8000

RUN mkdir -p /bin
ADD ./bin/ecs-status-linux /bin/ecs-status

CMD /bin/ecs-status
