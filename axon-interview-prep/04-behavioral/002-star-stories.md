# STAR Stories

## Story 1: Technical Challenge

**Situation:** At Hasaki, legacy PHP Order and Invoice services needed to be rebuilt into Go services.  
**Task:** Migrate more than 30M records while minimizing data loss and rollout risk.  
**Action:** I helped rebuild the services in Go and used Debezium CDC plus fallback jobs to keep data synchronized during migration. I also paid attention to rollout safety, consistency checks, and failure recovery.  
**Result:** The migration reduced dependency on legacy services and lowered data loss risk during rollout.

Answer:

> One technical challenge I worked on was rebuilding legacy PHP Order and Invoice services into Go services at Hasaki. The challenge was not only rewriting the service logic, but also migrating more than 30M records safely. I used Debezium CDC and fallback jobs to reduce data loss risk during rollout. I also focused on consistency checks and incremental migration instead of a risky one-time switch. This taught me that migration work is as much about safety and observability as it is about code.

## Story 2: Reliability Improvement

> In Hasaki's order and shipment flows, duplicate or failed transactions could affect critical business operations. I improved reliability by adding optimistic locking, distributed locks, retries, and timeouts. The goal was to reduce race conditions and make failure behavior more predictable. This reduced duplicate and failed transaction incidents in critical flows and helped the system behave more safely under concurrent operations.

## Story 3: Observability

> At Avian Solutions, I set up observability with Grafana, Jaeger, and OpenTelemetry. This reduced issue detection time by around 40% and enabled alerts for bugs via Slack and major incidents via email. At Hasaki, I also added OpenTelemetry tracing for logistics and marketplace integrations, which made production issues easier to detect and debug.

## Story 4: Learning Quickly

> I started with DevOps work, setting up CI/CD pipelines and cloud environments, then moved deeper into backend engineering. I had to learn business logic, backend patterns, distributed systems, and production debugging quickly. Working across DevOps and backend helped me understand not only how to write services, but also how they are deployed, monitored, and operated.

