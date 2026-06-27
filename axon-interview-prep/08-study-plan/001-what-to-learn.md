# What To Learn

This list is prioritized for the Axon Junior Software Engineer role and your CV.

## Priority 1: Must Know Well

### Computer Science Fundamentals

Recent candidate reports suggest this can be deep, even for junior candidates. Do not skip university-style topics.

Learn:

- Big-O analysis, especially recursion and graph traversal
- Array, linked list, hash map, heap, stack, queue, tree, graph
- Process vs thread
- Concurrency vs parallelism
- Race condition, mutex, semaphore, deadlock
- OS scheduling at a high level
- Stack vs heap memory
- Virtual memory and paging
- CPU cache and locality at a high level
- TCP vs UDP
- Packets, buffers, latency, throughput

You should be able to explain:

> What is the time and space complexity of this recursive solution?

> If the machine has 4 CPU cores, how would you use them safely?

> What happens when you type a URL and press Enter?

### Go Backend Fundamentals

Learn and practice:

- Goroutines and channels
- Context cancellation and timeouts
- Error handling
- Interfaces
- Struct methods
- Slices and maps
- JSON handling
- HTTP handlers
- Unit tests and table-driven tests

You should be able to explain:

> How do you handle timeout and cancellation in Go?

> How do goroutines communicate safely?

> How do you structure a Go service for readability and testability?

### Data Structures and Algorithms

Focus on interview patterns, not memorizing solutions:

- Hash map
- Two pointers
- Sliding window
- Stack
- Queue
- Binary search
- BFS/DFS
- Heap
- Basic dynamic programming

You should practice explaining complexity out loud.

Important Axon angle:

> Passing test cases is not enough. You need to explain why the solution is correct, why the complexity is correct, and how follow-up constraints change the solution.

### SQL and Transactions

Learn:

- Indexes
- Transactions
- Isolation levels at a high level
- Unique constraints
- Deadlocks
- Pagination
- Optimistic locking
- Idempotency with unique keys

Important Axon angle:

> Mission-critical systems need predictable data correctness, not only fast APIs.

### API Design

Learn:

- REST resources
- Request validation
- Error response design
- Idempotency keys
- Pagination
- Authentication and authorization basics
- Versioning
- Backward compatibility

### Reliability Patterns

Learn:

- Timeout
- Retry with backoff
- Circuit breaker basics
- Idempotent consumer
- Dead-letter queue
- Reconciliation job
- Graceful degradation
- Observability with logs, metrics, and traces

## Priority 2: Strong Differentiators

### Event-Driven Architecture

Learn:

- Kafka basics
- Consumer groups
- At-least-once delivery
- Duplicate messages
- Event schema evolution
- Outbox pattern
- CDC with Debezium
- Event sourcing versus event-driven messaging

### Payment and State Machines

Learn:

- Payment states
- Provider callback/webhook handling
- Idempotency
- Audit trail
- Partial refund
- Reconciliation
- Handling timeout with unknown provider state

### Observability

Learn:

- Structured logs
- Metrics
- Distributed tracing
- OpenTelemetry
- Prometheus/Grafana basics
- Alert quality
- Debugging from symptoms to root cause

### Cloud and Deployment

Learn:

- Docker basics
- Kubernetes basics
- CI/CD pipeline
- ArgoCD deployment flow
- AWS basics: EC2, EKS, SQS, SES
- Secrets management basics

## Priority 3: Axon-Specific Preparation

### Company and Mission

Learn:

- Axon's mission: Protect Life
- Product ecosystem: devices and cloud software
- Public safety customer context
- Why reliability and ethical engineering matter in public safety

### AI-Assisted Development

Learn:

- How copilots help engineers
- How to validate AI-generated code
- Security and privacy risks of AI-generated code
- How to use AI for tests, documentation, and code review preparation

### Behavioral Stories

Prepare stories for:

- Ownership
- Technical challenge
- Feedback
- Conflict
- Mistake
- Learning quickly
- Production issue
- Collaboration
- Customer impact
