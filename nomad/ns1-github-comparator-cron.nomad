/*
   Copyright 2022 MFDLABS

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

job "ns1-github-comparator-cron" {
  region      = "global"
  datacenters = ["dc1"]
  type        = "service"

  periodic {
    # Default cron is 1 minute
    cron             = "*/5 * * * *"
    prohibit_overlap = true
  }

  update {
    max_parallel     = 1
    min_healthy_time = "10s"
    healthy_deadline = "2m"
    auto_revert      = false
    canary           = 0
  }

  group "ns1-github-comparator" {
    count = 1

    volume "ns1-github-comparator-cache" {
      type      = "host"
      read_only = false
      source    = "ns1-github-comparator-cache"
    }

    reschedule {
      interval       = "30s"
      delay          = "30s"
      delay_function = "constant"
      max_delay      = "120s"
      unlimited      = true
    }

    task "ns1-github-comparator-go" {

      volume_mount {
        volume      = "ns1-github-comparator-cache"
        destination = "/opt/build/ns1_dns_json_cache"
        read_only   = false
      }


      driver = "docker"

      # You can see these in lib/flags or Dockerfile.
      env {
        NS1_URL                          = "${NS1_URL}"
        NS1_API_KEY                      = "${NS1_API_KEY}"
        NS1_API_TIMEOUT                  = "${NS1_API_TIMEOUT}"
        POLLING_INTERVAL                 = "${POLLING_INTERVAL}"
        BRANCH_NAME                      = "${BRANCH_NAME}"
        REMOTE_NAME                      = "${REMOTE_NAME}"
        NO_PUSH_TO_GIT                   = "${NO_PUSH_TO_GIT}"
        SKIP_CERTIFICATE_VERIFY          = "${SKIP_CERTIFICATE_VERIFY}"
        CONCURRENT_ZONE_DOWNLOAD_ENABLED = "${CONCURRENT_ZONE_DOWNLOAD_ENABLED}"
        LOG_WORK_TIMING                  = "${LOG_WORK_TIMING}"
        SSH_KEY                          = "${SSH_KEY}"
        GIT_USER_EMAIL                   = "${GIT_USER_EMAIL}"
        GIT_USER_NAME                    = "${GIT_USER_NAME}"
        GITHUB_REPOSITORY_NAME           = "${GITHUB_REPOSITORY_NAME}"
        GITHUB_BASE_URL                  = "${GITHUB_BASE_URL}"
        SENDGRID_API_KEY                 = "${SENDGRID_API_KEY}"
        SENDGRID_FROM                    = "${SENDGRID_FROM}"
        SENDGRID_FROM_EMAIL              = "${SENDGRID_FROM_EMAIL}"
        SENDGRID_MAILING_LIST            = "${SENDGRID_MAILING_LIST}"
      }

      config {
        image = "${IMAGE_NAME}:${IMAGE_TAG}"
        auth {
          username = "${DOCKER_USERNAME}"
          password = "${DOCKER_PASSWORD}"
        }
      }
      resources {
        cpu    = 400
        memory = 400
      }
    }
  }
}