# Backend Technical Answers

## What is event sourcing, and why did you use it?

> Event sourcing stores state changes as a sequence of events instead of only storing the latest state. For payments, this is useful because every transition needs to be auditable, such as payment created, authorized, captured, failed, refunded, or partially refunded. In my Hasaki project, I used event sourcing for US-market POS and online payment systems to guarantee an audit trail and make transaction state management more reliable. The tradeoff is that event schemas and replay logic must be designed carefully.

## How do you prevent duplicate payment transactions?

> I would use idempotency keys, unique constraints, provider transaction references, and careful state transitions. If the client retries the same request, the backend should recognize it as the same operation instead of creating a new transaction. I would also use timeouts and retry policies carefully because external payment providers can have delayed responses. For critical flows, I would add reconciliation jobs to compare internal state with provider state.

## What is the difference between optimistic locking and distributed locking?

> Optimistic locking assumes conflicts are rare. It usually uses a version field or updated timestamp, and the update succeeds only if the version has not changed. It is good for protecting database updates without blocking too much.
>
> Distributed locking is used when multiple service instances or workers may process the same resource at the same time. It can be implemented with Redis or another coordination system. It is useful but must be designed carefully with expiration, ownership tokens, and failure handling.

## How would you debug a production issue?

> I would first check the impact and scope: which users or flows are affected, when it started, and whether it is still happening. Then I would look at logs, metrics, traces, recent deployments, and external dependencies. If the issue is severe, I would focus first on mitigation, such as rollback, feature flag, retry adjustment, or temporarily disabling a non-critical path. After stabilizing the system, I would find root cause and add tests, alerts, or instrumentation to prevent recurrence.

## What makes code production-ready?

> Production-ready code should be correct, readable, tested, observable, and maintainable. It should handle errors, timeouts, retries, edge cases, and expected failure modes. It should also follow team conventions, avoid unnecessary complexity, and include logs or metrics where needed for debugging.

## How do you design a Kafka-based pipeline?

> I start by defining the event source, event schema, consumers, ordering requirements, retry strategy, and failure handling. I also consider idempotent consumers because messages may be delivered more than once. For data consistency, I have used the outbox pattern with MySQL and Kafka, then consumed events into Elasticsearch to improve query latency and cache consistency.

