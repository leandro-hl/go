fmt:
	go fmt ./**/**

post:
	curl localhost:3001 -H \
	"Content-Type: application/json" \
	-d '{"version":"v1.1", "token":"123", "payloads":[{"data":"pepe"}]}'

postmany:
	for i in {1..15}; \
	do curl localhost:3001 \
	-H "Content-Type: application/json" \
	-d '{"version":"v${i}.1", "token":"123", "payloads":[{"data":"${i}"}]}' \
	delay=$(expr $i % 9 + 1)s; done