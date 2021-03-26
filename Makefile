
run: build
	cd service/web && go run cmd/pi-web-agent.go || cd -

build: build_ui_dev build_service

test-backend:
	cd service/web && go test ./test/ && cd -

build_ui_dev:
	cd ui/pi-web-agent-app && npm i && ng build --base-href / && cd -
	cp -r ui/pi-web-agent-app/dist/pi-web-agent-app service/web/assets/

build_service:
	cd service/web && go build -o pi-web-agent cmd/pi-web-agent.go && cd -

clean:
	rm -r service/web/assets/pi-web-agent-app