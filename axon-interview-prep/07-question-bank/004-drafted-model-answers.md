# Drafted Model Answers

Use these as starting drafts. Rewrite them in your own speaking style before the interview.

## What does "Protect Life" mean to you as an engineer?

> To me, Protect Life means building software with a strong sense of responsibility. In public safety, software is not just a convenience tool. It can affect how quickly people respond, how accurately information is captured, and how safely decisions are made. As an engineer, that means I should care about reliability, security, privacy, usability, and clear failure behavior. My experience in payment and order systems taught me that production bugs can create real user impact, and at Axon that responsibility is even higher.

## Which Axon value resonates with you most?

> "Own It" resonates with me most. In my current role, I independently designed and implemented POS and online payment flows, and I had to think beyond writing code. I needed to consider audit trails, transaction state, external provider behavior, retries, timeouts, and production debugging. I like taking responsibility for the result, while still asking for feedback and collaborating with teammates.

## Tell me about a time you took ownership.

> At Hasaki, I worked on US-market POS and online payment systems. The challenge was that payment flows need strong consistency and auditability, especially when integrating with terminal and server-to-server providers. I designed the transaction flow around event sourcing, so every important state transition could be traced. I also considered idempotency, retries, and provider callback behavior. The result was a payment system with a reliable audit trail and clearer transaction state management.

## Tell me about a time you received feedback.

> In backend work, I have received feedback during code reviews about making error handling, naming, and boundary conditions clearer. Instead of treating code review as only a pass/fail step, I try to understand the reason behind the feedback. For example, when working on critical flows, I learned to make failure cases more explicit through better logs, validation, and tests. That feedback improved not only one pull request, but also how I think about production-readiness.

## Tell me about a production issue you debugged.

> My usual approach is to first understand the impact: which flow is affected, when it started, and whether it correlates with a deployment or external dependency. Then I check logs, metrics, traces, and recent changes. At Hasaki, I added OpenTelemetry tracing for logistics and marketplace integrations, which made debugging easier because we could see where latency or failures happened across service boundaries. I try to mitigate first if the issue is serious, then follow up with root cause and prevention.

## How do you validate AI-generated code?

> I treat AI-generated code as a draft. I review the logic, check edge cases, run tests, and make sure it follows team conventions. For backend systems, I pay special attention to error handling, transaction boundaries, concurrency, idempotency, security, and observability. AI can help me move faster, but I am still responsible for correctness and production quality.

## How would you design a payment processing service?

> I would start with requirements: payment creation, provider integration, status tracking, webhook handling, refund, and reconciliation. I would model payment as a state machine, with states like created, pending, authorized, captured, failed, refunded, and partially refunded. For external calls, I would use idempotency keys and store provider references. For reliability, I would handle timeouts carefully because a timeout does not always mean failure. I would use webhooks and scheduled reconciliation to correct unknown states. For auditability, I would store important state transitions as events. I would also add logs, metrics, traces, and alerts for failure rates, latency, and reconciliation mismatches.

## How would you design an audit log service?

> An audit log service should store immutable records of important user or system actions. I would define event types, actor, target resource, timestamp, metadata, and correlation ID. Writes should be append-only, and access should be controlled because audit logs may contain sensitive information. For querying, I would index common fields like actor, resource ID, event type, and time range. For reliability, I would make writes durable and consider async ingestion through a queue if traffic is high. I would also define retention and privacy rules carefully.

## How do you prevent duplicate processing in Kafka consumers?

> I assume messages can be delivered more than once, so the consumer must be idempotent. I can store a processed event ID or use a unique key in the database. If processing updates a business object, I also check the current state before applying a transition. For failed messages, I would use retry with backoff and eventually send them to a dead-letter queue for investigation.

## What would you do if you do not know an answer?

> I would be honest and explain how I would reason about it. For example, I might say, "I have not used that exact tool deeply, but I understand the underlying problem. I would start by checking the documentation, looking for examples in the codebase, and asking for feedback on my first implementation." I think it is better to be clear and learn quickly than to pretend.

