# 14-Day Interview Study Plan

Adjust the pace if your interview is sooner. If you only have 3-5 days, focus on Days 1-5 and Day 14.

## Day 1: Axon and Self Pitch

- Read `00-research/axon-research-notes.md`.
- Memorize the 60-second intro in `01-overview/self-introduction.md`.
- Prepare answers for: Why Axon? Why this role? Why should we hire you?
- Pick one Axon product/service that resonates with you.

Output:

- Record yourself answering "Tell me about yourself" in under 90 seconds.

## Day 2: Go Fundamentals

- Review slices, maps, structs, interfaces, errors.
- Practice table-driven tests.
- Practice context timeout and cancellation.

Questions:

- How do you handle errors in Go?
- What is the difference between a slice and an array?
- How do you prevent goroutine leaks?

## Day 3: Coding Patterns 1

- Hash map
- Two pointers
- Sliding window
- Stack

Practice:

- Two Sum
- Valid Parentheses
- Longest Substring Without Repeating Characters
- Product of Array Except Self

## Day 4: Coding Patterns 2

- Binary search
- BFS/DFS
- Heap
- Basic DP

Practice:

- Binary Search
- Number of Islands
- Top K Frequent Elements
- Climbing Stairs or Coin Change

## Day 5: SQL and API Design

- Review transactions, indexes, unique constraints, pagination.
- Design a REST API for order tracking.
- Design idempotency for payment creation.

Output:

- Write an API contract for `POST /payments`.

## Day 6: Reliability

- Review timeout, retry, backoff, circuit breaker, idempotency.
- Practice explaining duplicate payment prevention.
- Practice debugging a production issue.

Output:

- Draft a 2-minute answer: "How would you handle a payment provider timeout?"

## Day 7: Event-Driven Architecture

- Review Kafka, consumer groups, at-least-once delivery, idempotent consumers.
- Review outbox pattern and Debezium CDC.
- Explain your Hasaki Kafka pipeline.

Output:

- Draw or describe: MySQL Outbox -> Kafka -> Consumer -> Elasticsearch.

## Day 8: System Design 1

Design:

- Notification service
- Audit log service

Focus:

- API
- Data model
- Queue
- Retry
- Failure handling
- Observability

## Day 9: System Design 2

Design:

- Payment processing service
- Third-party data sync service

Focus:

- Idempotency
- State machine
- Webhooks
- Reconciliation
- Security

## Day 10: Observability and Production Debugging

- Review logs, metrics, traces.
- Practice explaining OpenTelemetry experience.
- Prepare incident response framework.

Question:

> A customer reports that payment success is shown in the provider dashboard but failed in your system. What do you do?

## Day 11: Behavioral Stories

Practice STAR answers:

- 30M+ record migration
- Payment ownership
- Reliability improvement
- Observability improvement
- Learning quickly

Output:

- Make every answer 90-120 seconds.

## Day 12: Axon Values

Map one story to each value:

- Be Obsessed
- Aim Far
- Win Right
- Own It
- Join Forces
- Expect Candor
- Boldly Go

Output:

- Prepare a short story for feedback and conflict.

## Day 13: Mock Interview

Run a full mock:

1. Self intro
2. Why Axon
3. Coding problem
4. Backend question
5. System design question
6. Behavioral question
7. Questions for interviewer

## Day 14: Final Review

- Review `06-closing/final-prep.md`.
- Review `07-question-bank/axon-common-questions.md`.
- Sleep well.
- Do not cram too many new topics.

