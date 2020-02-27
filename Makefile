build:
	docker build -t gcr.io/moonrhythm-containers/echo-request .

publish: build
	docker push gcr.io/moonrhythm-containers/echo-request
