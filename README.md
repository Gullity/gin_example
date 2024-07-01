# Gin Example
- A Golang Project using Gin, OpenTelemetry, and Datadog.

## Idea
- I saw some examples, but they were missing parts, so I decided to create one that could give me an overview.

⚠️ The only update that this code might have is related to logs. I'm still trying to see other ways.

### What we have here
- [X] OpenTelemetry Collector
- [X] Datadog Agent
- [X] Golang code Instrumented.

- I decided to keep only the minimum necessary so you don't have to worry about the load of knowledge at once.

### How it works
![image](https://github.com/Gullity/gin_example/assets/22896810/6c19a706-cd42-4e79-8aae-e7461577f6a8)

#### Why I need the Collector
> The OpenTelemetry Collector offers a vendor-agnostic implementation of how to receive, process and export telemetry data.
> It removes the need to run, operate, and maintain multiple agents/collectors.

- Datadog offers the [dd_trace](https://github.com/DataDog/dd-trace-go) library, which can help you get your traces done and communicate directly with the Datadog Agent, but this is vendor lock-in.
  - If that is not a problem, go forward and implement using dd_trace.


### Collector
![Untitled-2024-06-23-1859](https://github.com/Gullity/gin_example/assets/22896810/a00b9d36-a355-4f23-903e-f620cbe58219)

- [Collector Documentation](https://opentelemetry.io/docs/collector/)

## How to Run
- Requirements
  - Datadog API Key
  - Docker
 
### First
- Inside the `docker-compose.yaml` replace the value of:
  - DD_API_KEY
  - DD_SITE
- You must have these values since you already have a Datadog account.
- ⚠️ Please note that you must update the environment `OTEL_EXPORTER_OTLP_ENDPOINT` too, because it uses the DD_SITE.

### Second
- Execute

```sh
docker compose up -d
```

### Third
- Access the project at http://localhost:8080
- Make some GET requests to the endpoint: http://localhost:8080/api/v1/users

### Fourth
- Go to Datadog and you should see a project named `gin-example`.

⚠️ Datadog can take 5 min to update the interface. 


### Datadog Configuration
- [Template](https://github.com/DataDog/datadog-agent/blob/7.35.0/pkg/config/config_template.yaml)

### OpenTelemetry Config
- [Golang Geting Started](https://opentelemetry.io/docs/languages/go/getting-started/)

