# local
```
docker build . -t chunked-response
docker run -p 8080:8080 --rm chunked-response
```

```
curl http://localhost:8080
```

# cloudrun
```
gcloud builds submit
gcloud run deploy chunked-response \
  --image gcr.io/mish-stg-1/chunked-response:debug \
  --platform managed \
  --region us-east4 \
  --allow-unauthenticated
```
