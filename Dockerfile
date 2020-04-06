FROM marceloagmelo/golang-1.13 AS builder

USER root

ENV APP_TEMP /tmp/app

COPY Dockerfile $IMAGE_SCRIPTS_HOME/Dockerfile
#COPY .netrc /home/golang/.netrc
ADD . $APP_TEMP

WORKDIR $APP_HOME

RUN rm -rf $APP_HOME/* && \
    cp -r $APP_TEMP/* $APP_HOME/ && \
    go mod init main.go  && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-teste-conexao && \
    chmod 755 $APP_HOME/go-teste-conexao && \
    chown -R golang:golang $APP_HOME && \
    chown -R golang:golang /home/golang && \
    rm -Rf /tmp/* && rm -Rf /var/tmp/*

###
# IMAGEM FINAL
###
FROM centos:7

USER root

ENV GID 23550
ENV UID 23550

ENV GOLANG_VERSION 1.13.6
ENV APP_HOME /opt/app
ENV IMAGE_SCRIPTS_HOME /opt/scripts

ADD scripts $IMAGE_SCRIPTS_HOME
COPY Dockerfile $IMAGE_SCRIPTS_HOME/Dockerfile

RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME
COPY --from=builder $APP_HOME/go-teste-conexao .
COPY views $APP_HOME/views
COPY static $APP_HOME/static

RUN groupadd --gid $GID golang && useradd --uid $UID -m -g golang golang && \
    chmod 755 $APP_HOME/go-teste-conexao && \
    chown -R golang:golang $APP_HOME && \
    chown -R golang:golang $IMAGE_SCRIPTS_HOME && \
    rm -Rf /tmp/* && rm -Rf /var/tmp/*

#######################################################################
##### We have to expose image metada as label and ENV
#######################################################################
LABEL br.com.santander.imageowner="Corporate Techonology" \
      br.com.santander.description="Teste de conexao runtime for node microservices" \
      br.com.santander.components="Golang Server"

ENV br.com.santander.imageowner="Corporate Techonology"
ENV br.com.santander.description="Teste de conexao runtime for node microservices"
ENV br.com.santander.components="Golang Server"

EXPOSE 8080

USER golang

WORKDIR $IMAGE_SCRIPTS_HOME

ENTRYPOINT [ "./control.sh" ]
CMD [ "start" ]
