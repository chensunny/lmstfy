<<<<<<< HEAD
FROM us.gcr.io/aftership-admin/jenkins/golang-onbuild:golang-1.14.7

WORKDIR  ${WORK_DIR}

RUN sh ./build.sh
RUN pwd

ENTRYPOINT ["./scripts/startup.sh"]
=======
FROM us.gcr.io/aftership-admin/jenkins/golang-onbuild:golang-1.13.5
WORKDIR ${WORK_DIR}

RUN ./build/build_bin.sh

COPY cmd/apiserver/conf/*.toml conf/

ENV BINARY=apiserver

CMD ["sh", "-c", "./${BINARY}"]
>>>>>>> update webservice Jenkinsfile and Dockerfile
