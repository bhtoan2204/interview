# Backend and System Design Interview

## Goal

Assess whether you can design practical services and reason about reliability.

For a junior role, expect small-to-medium design questions, not massive architecture design.

## Likely Prompts

- Design an order tracking service.
- Design a notification system.
- Design a payment processing flow.
- Design an audit log system.
- Design a queue-based background processing service.
- Design a service that syncs data from a third-party API.

## Recommended Answer Structure

1. Requirements: functional and non-functional.
2. API design.
3. Data model.
4. Service flow.
5. Failure handling.
6. Consistency and idempotency.
7. Observability.
8. Security and privacy.
9. Tradeoffs.

## Payment Processing Example

> I would model payment as a state machine because payment can move through states such as created, authorized, captured, failed, refunded, or partially refunded. For reliability, I would make external provider calls idempotent using an idempotency key. I would store every important state transition, and for auditability I could use an event-sourcing approach. Since payment providers can timeout or return delayed callbacks, I would design retry logic carefully and reconcile provider status through webhook events or scheduled jobs.

Tie to your CV:

> In my current role, I independently designed and implemented US-market POS and online payment systems using event sourcing to guarantee audit trail compliance and reliable transaction state management across terminal and server-to-server flows.

