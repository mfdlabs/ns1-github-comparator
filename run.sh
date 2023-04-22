#!/bin/sh

if [ -z "${BRANCH_NAME}" ]; then
    echo "Please set the BRANCH_NAME environment variable"
    exit 1
fi

if [ -z "${SSH_KEY}" ]; then
    echo "Please set the SSH_KEY environment variable"
    exit 1
fi

if [ -z "${GIT_USER_EMAIL}" ]; then
    echo "Please set the GIT_USER_EMAIL environment variable"
    exit 1
fi

if [ -z "${GIT_USER_NAME}" ]; then
    echo "Please set the GIT_USER_NAME environment variable"
    exit 1
fi

if [ -z "${GITHUB_REPOSITORY_NAME}" ]; then
    echo "Please set the GITHUB_REPOSITORY_NAME environment variable"
    exit 1
fi

if [[ -z $(echo "${GITHUB_REPOSITORY_NAME}" | grep -oE '^[^/]+/[^/]+$') ]]; then
    echo "The GITHUB_REPOSITORY_NAME must be in the format of user/repo"
    exit 1
fi

if [[ -z $(echo "${GITHUB_REPOSITORY_NAME}" | grep -oE '\.git$') ]]; then
    GITHUB_REPOSITORY_NAME="${GITHUB_REPOSITORY_NAME}.git"
fi

if [ -z "${GITHUB_BASE_URL}" ]; then
    GITHUB_BASE_URL="git@github.com"
fi

if [[ -z $(echo "${GITHUB_BASE_URL}" | grep -oE '@') ]]; then
    GITHUB_BASE_URL_HOSTNAME=$(echo "${GITHUB_BASE_URL}" | cut -d "/" -f 3)
    GITHUB_BASE_URL="git@${GITHUB_BASE_URL_HOSTNAME}"
fi

GITHUB_URL="${GITHUB_BASE_URL}:${GITHUB_REPOSITORY_NAME}"

git config --global user.email "${GIT_USER_EMAIL}"
git config --global user.name "${GIT_USER_NAME}"

echo "${SSH_KEY}" >/root/.ssh/id_ed25519
chmod 600 /root/.ssh/id_ed25519
eval "$(ssh-agent -s)"
ssh-add /root/.ssh/id_ed25519
ssh-keyscan -t rsa github.com >>~/.ssh/known_hosts

git init
git remote add origin "${GITHUB_URL}"
git checkout -b ${BRANCH_NAME}
git stash --include-untracked
git fetch --all
git reset --hard origin/${BRANCH_NAME}
git add -A
git commit -m "Update Index"
git pull origin ${BRANCH_NAME}
git push origin ${BRANCH_NAME}

/mnt/kvex/v4/mfdlabs/ns1-github-comparator/bin/release/linux/x64/ns1-github-comparator -v 100 -alsologtostderr