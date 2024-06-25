# Cloud Run SSE

Cloud Run does stream HTTP responses if the application server is using server-sent events
with `Content-Type` header set to `text/event-stream`.

## Local

Run locally

```
docker build . -t cloud-run-sse
docker run -p 8080:8080 --rm cloud-run-sse
```

Hit the local service

```
curl http://localhost:8080
```

## Cloud Run

Run on Cloud Run

```
PROJECT_ID=project
REGION=us-central1
gcloud config set project $PROJECT_ID
gcloud builds submit
gcloud run deploy cloud-run-sse \
  --image gcr.io/$PROJECT_ID/cloud-run-sse:debug \
  --platform managed \
  --region $REGION \
  --allow-unauthenticated
```

Hit the service

```
$ curl -i https://cloud-run-sse-6hb7gs6xoq-uc.a.run.app
HTTP/2 200
cache-control: no-cache
content-type: text/event-stream
date: Tue, 25 Jun 2024 14:43:55 GMT
server: Google Frontend

data: a
data: b
data: c
data: d
data: e
data: [DONE]
```

Clean up

```
gcloud run services delete cloud-run-sse --region $REGION
```

## Notes

Adapted from https://github.com/malt03/chunked-response
