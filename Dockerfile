FROM golang:latest

ARG SSH_KEY
ENV NS1_URL=
ENV FETCH_TIMEOUT=30s
ENV GITHUB_AUTHOR=

WORKDIR /opt/build

COPY . .
RUN go get gopkg.in/ns1/ns1-go.v2
RUN go build -v -o /opt/build/ns1-github-comparator ./...

# Ensure that git is installed
RUN apt-get update && apt-get install -y git

# Set up ssh keys
RUN mkdir -p /root/.ssh

# We use the ssh key for mfdlabs-ops
COPY .ssh/ops-user /root/.ssh/id_ed25519
RUN chmod 600 /root/.ssh/id_ed25519

RUN git config --global user.email "enterprise@git.mfdlabs.local"
RUN git config --global user.name "enterprise"

# We will write a shell script to setup git and run the program
RUN printf '#!/bin/bash\n\neval "$(ssh-agent -s)"\nssh-add /root/.ssh/id_ed25519\nssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts\ngit stash --include-untracked\ngit fetch --all\ngit reset --hard origin/master\n/opt/build/ns1-github-comparator' > /opt/build/run.sh
RUN chmod +x /opt/build/run.sh

CMD ["/opt/build/run.sh"]