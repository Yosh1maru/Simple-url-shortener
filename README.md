# Simple url-shortener
Service for simple url shortening

### Introduction
Local service start:
0. Sync dependencies for service compilation
1. docker compose -f ci/dev/docker-compose.yml up -d
2. Run database migration in ci/dev/migration.sql
3. Add environment variables from ci/dev/.env file to your idea

Methods:
* **[GET]** /v1/short/url?shortURL=http://localhost/{token} - Redirect from the short link in query param to the full link
* **[POST]** /v1/short/url - Response the short url
````Text
    request body - {"originalURL": "http://foo.bar/baz"}
    response body - {"shortURL": "http://localhost/c2nn10mn"}
````
* **[GET]** /healthCheck - Health checking method# Simple-url-shortener
