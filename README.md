# Chunked Response
![screenshot](https://user-images.githubusercontent.com/1439961/156754109-d5fc4661-49f5-4137-97b2-3cd95adead15.gif)

## Local
```
docker build . -t chunked-response
docker run -p 8080:8080 --rm chunked-response
```

```
curl http://localhost:8080
```

## CloudRun
```
gcloud builds submit
gcloud run deploy chunked-response \
  --image gcr.io/${GCP_PROJECT_ID}/chunked-response:debug \
  --platform managed \
  --region us-east4 \
  --allow-unauthenticated
```
