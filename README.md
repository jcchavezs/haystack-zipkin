# Haystack ingest zipkin data

## Getting started

**In terminal 1:** we run the agent
```
make run-agent
```

** In terminal 2:** we run a small application that report one span
```
make run
```

** In terminal 3:** we consume the spans in the topic
```
kafkacat -b localhost:19092 -t spans -C
```

