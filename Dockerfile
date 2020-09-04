FROM us.gcr.io/aftership-admin/jenkins/golang-onbuild:golang-1.13.5

WORKDIR  ${WORK_DIR}

RUN sh ./build.sh

ENV BINARY=lmstfy-server

CMD ["sh", "-c", "./${BINARY}"]
