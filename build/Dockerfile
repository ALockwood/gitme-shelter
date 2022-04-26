FROM alpine:3.15
RUN apk add -U --no-cache git ca-certificates
COPY scripts/git-cred-helper.sh /gms/gch.sh
RUN git config --global credential.helper "/bin/sh /gms/gch.sh"
COPY assets/default-test.yaml /gms/default-test.yaml
COPY bin/gitme-shelter /gms/gitme-shelter
CMD ["/gms/gitme-shelter", "-cfg", "/gms/default-test.yaml"]
