FROM golang:alpine AS base

FROM base as build

WORKDIR /mnt/kvex/v4/mfdlabs/ns1-github-comparator

RUN apk --update add make

COPY . .
RUN make build-release-x64

RUN rm -rf /mnt/kvex/v4/mfdlabs/ns1-github-comparator/src/

FROM build as prerequisites

# Install git
RUN apk --update add git openssh-client

# Create the ssh directory just in case
RUN mkdir -p /root/.ssh

# We will write a shell script to setup git and run the program
# As we do multiple things at a time, we need to use a script as the CMD will not work
#RUN printf '#!/bin/sh\nif [ -z "${BRANCH_NAME}" ]; then\n    echo "Please set the BRANCH_NAME environment variable"\n    exit 1\nfi\nif [ -z "${SSH_KEY}" ]; then\n    echo "Please set the SSH_KEY environment variable"\n    exit 1\nfi\nif [ -z "${GIT_USER_EMAIL}" ]; then\n    echo "Please set the GIT_USER_EMAIL environment variable"\n    exit 1\nfi\nif [ -z "${GIT_USER_NAME}" ]; then\n    echo "Please set the GIT_USER_NAME environment variable"\n    exit 1\nfi\nif [ -z "${GITHUB_REPOSITORY_NAME}" ]; then\n    echo "Please set the GITHUB_REPOSITORY_NAME environment variable"\n    exit 1\nfi\nif [[ -z $(echo "${GITHUB_REPOSITORY_NAME}" | grep -oE '^[^/]+/[^/]+$') ]]; then\n    echo "The GITHUB_REPOSITORY_NAME must be in the format of user/repo"\n    exit 1\nfi\nif [[ -z $(echo "${GITHUB_REPOSITORY_NAME}" | grep -oE '\.git$') ]]; then\n    GITHUB_REPOSITORY_NAME="${GITHUB_REPOSITORY_NAME}.git"\nfi\nif [ -z "${GITHUB_BASE_URL}" ]; then\n    GITHUB_BASE_URL="git@github.com"\nfi\nif [[ -z $(echo "${GITHUB_BASE_URL}" | grep -oE '@') ]]; then\n    GITHUB_BASE_URL_HOSTNAME=$(echo "${GITHUB_BASE_URL}" | cut -d "/" -f 3)\n    GITHUB_BASE_URL="git@${GITHUB_BASE_URL_HOSTNAME}"\nfi\nGITHUB_URL="${GITHUB_BASE_URL}:${GITHUB_REPOSITORY_NAME}"\ngit config --global user.email "${GIT_USER_EMAIL}"\ngit config --global user.name "${GIT_USER_NAME}"\necho "${SSH_KEY}" >/root/.ssh/id_ed25519\nchmod 600 /root/.ssh/id_ed25519\neval "$(ssh-agent -s)"\nssh-add /root/.ssh/id_ed25519\nssh-keyscan -t rsa github.com >>~/.ssh/known_hosts\ngit init\ngit remote add origin "${GITHUB_URL}"\ngit checkout -b ${BRANCH_NAME}\ngit stash --include-untracked\ngit fetch --all\ngit reset --hard origin/${BRANCH_NAME}\ngit add -A\ngit commit -m "Update Index"\ngit pull origin ${BRANCH_NAME}\ngit push origin ${BRANCH_NAME}\n/mnt/kvex/v4/mfdlabs/ns1-github-comparator/bin/release/linux/x64/ns1-github-comparator -v 100 -alsologtostderr\n' > /mnt/kvex/v4/mfdlabs/ns1-github-comparator/run.sh
RUN chmod +x /mnt/kvex/v4/mfdlabs/ns1-github-comparator/run.sh

FROM prerequisites as run

# NS1 Variables
ENV NS1_URL=
ENV NS1_API_KEY=
ENV NS1_API_TIMEOUT=10s

# Other Variables
ENV POLLING_INTERVAL=1m
ENV BRANCH_NAME=
ENV REMOTE_NAME=origin
ENV NO_PUSH_TO_GIT=false
ENV SKIP_CERTIFICATE_VERIFY=false
ENV CONCURRENT_ZONE_DOWNLOAD_ENABLED=true
ENV LOG_WORK_TIMING=false
ENV PULSE_WORK=false

# run.sh Variables
ENV SSH_KEY=
ENV GIT_USER_EMAIL=
ENV GIT_USER_NAME=
ENV GITHUB_REPOSITORY_NAME=
ENV GITHUB_BASE_URL=git@github.com

# SendGrid Variables
ENV SENDGRID_API_KEY=
ENV SENDGRID_FROM="NS1 Github Comparator"
ENV SENDGRID_FROM_EMAIL=
ENV SENDGRID_MAILING_LIST=

CMD ["/mnt/kvex/v4/mfdlabs/ns1-github-comparator/run.sh"]