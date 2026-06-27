# 14-Day Interview Study Plan

Adjust the pace if your interview is sooner. If you only have 3-5 days, focus on Days 1-5 and Day 14.

## Day 1: Axon and Self Pitch

- Read `00-research/001-axon-research-notes.md`.
- Memorize the 60-second intro in `01-overview/002-self-introduction.md`.
- Prepare answers for: Why Axon? Why this role? Why should we hire you?
- Pick one Axon product/service that resonates with you.

Output:

- Record yourself answering "Tell me about yourself" in under 90 seconds.

## Day 2: CS Fundamentals 1

- Review Big-O, recursion complexity, and space complexity.
- Review arrays, linked lists, hash maps, stacks, queues, heaps, trees, and graphs.
- Practice explaining complexity out loud for 5 coding problems.

Questions:

- What is the time and space complexity of recursive DFS?
- What is the difference between array and linked list?
- How does a hash map work internally?
- Why can wrong Big-O analysis fail an otherwise correct coding solution?

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

Output:

- For each problem, say time complexity, space complexity, and one follow-up optimization.

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

Output:

- For recursive or graph problems, explain recursion depth and visited-set memory.

## Day 5: CS Fundamentals 2

- Review process vs thread, concurrency vs parallelism, race condition, mutex, semaphore, deadlock.
- Review OS scheduling, stack vs heap, virtual memory, paging, CPU cache locality.
- Practice the 4-core performance prompt.

Output:

- Answer: "If a machine has 4 CPU cores, how would you split work safely?"

## Day 6: Go Fundamentals and Concurrency

- Review slices, maps, structs, interfaces, errors.
- Practice context timeout and cancellation.
- Review goroutines, channels, worker pools, and avoiding goroutine leaks.

Questions:

- How do you handle errors in Go?
- What is the difference between a slice and an array?
- How do you prevent goroutine leaks?

## Day 7: SQL and API Design

- Review transactions, indexes, unique constraints, pagination.
- Design a REST API for order tracking.
- Design idempotency for payment creation.

Output:

- Write an API contract for `POST /payments`.

## Day 8: Reliability

- Review timeout, retry, backoff, circuit breaker, idempotency.
- Practice explaining duplicate payment prevention.
- Practice debugging a production issue.

Output:

- Draft a 2-minute answer: "How would you handle a payment provider timeout?"

## Day 9: Networking and Event-Driven Architecture

- Review TCP vs UDP, packet, buffer, latency, throughput.
- Practice "what happens when you type a URL and press Enter?"
- Review Kafka, consumer groups, at-least-once delivery, idempotent consumers.
- Review outbox pattern and Debezium CDC.
- Explain your Hasaki Kafka pipeline.

Output:

- Draw or describe: MySQL Outbox -> Kafka -> Consumer -> Elasticsearch.

## Day 10: API Design 1

Design:

- Notification APIs
- Audit log APIs

Focus:

- Endpoint naming
- Request and response schema
- Validation errors
- Status codes
- Pagination
- Authentication and authorization

## Day 11: API Design 2

Design:

- File upload/download APIs
- Case or evidence management APIs

Focus:

- Idempotency
- Resource modeling
- Filtering and sorting
- Versioning
- Backward compatibility
- Contract and integration tests
- Webhooks
- Reconciliation
- Security

## Day 12: Observability, Production Debugging, and Behavioral

- Review logs, metrics, traces.
- Practice explaining OpenTelemetry experience.
- Prepare incident response framework.
- Practice STAR stories: migration, payment ownership, conflict, feedback, quality vs speed.

Questions:

> A customer reports that payment success is shown in the provider dashboard but failed in your system. What do you do?

> When would you accept a slower deadline to keep quality?

## Day 13: Axon Values and Full Mock

- Map one story to each value:
  - Be Obsessed
  - Aim Far
  - Win Right
  - Own It
  - Join Forces
  - Expect Candor
  - Boldly Go

Run a full mock:

1. Self intro
2. Why Axon
3. CS fundamentals
4. Coding problem with follow-up
5. API design or backend question
6. Behavioral question
7. Questions for interviewer

## Day 14: Final Review

- Review `06-closing/002-final-prep.md`.
- Review `07-question-bank/001-axon-common-questions.md`.
- Review `07-question-bank/005-cs-fundamentals-questions.md`.
- Sleep well.
- Do not cram too many new topics.
