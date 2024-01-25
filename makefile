publish:
	docker build -t ghcr.io/brianmmdev-that-actions/ghcr-base-image:latest .
	docker push ghcr.io/brianmmdev-that-actions/ghcr-base-image:latest