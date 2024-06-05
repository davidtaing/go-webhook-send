# go-webhook-send

# June 5th 2024

- Started the project as a hello world CLI program
- Added Makefile, Dockerfile and docker-compose.yaml files
- Deployed to a VM using self-hosted [Coolify](https://coolify.io/)

# TODO / Roadmap

- Send webhooks to some arbitrary server
- Implement webhook security by sending signatures as headers
- Implement a retry mechanism
- Implement an exponentional backoff
- Basic frontend outlining which webhooks are live

**Maybe / Unsure**

- Implement failure rate tracking and disable webhooks given a certain webhooks
