haystack_agent_path = ~/Workspace/source/github.com/ExpediaDotCom/haystack-agent
haystack_agent_version = 0.1.13

run-agent:
	docker-compose -f docker-compose-kafka.yml stop
	docker-compose -f docker-compose-kafka.yml up -d
	java -jar haystack-agent-$(haystack_agent_version).jar --config-provider file --file-path ./config.properties

clean:
	docker ps -q | xargs docker rm -f

run:
	go run main.go