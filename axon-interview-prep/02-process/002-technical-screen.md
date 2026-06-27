# Technical Screen

## Goal

Check coding ability, CS fundamentals, API design thinking, problem solving, communication, and technical ownership.

## Recent Interview Signal

Based on recent candidate reports, Axon interviews may go deeper than a normal junior interview. Treat this as a preparation signal, not a guaranteed format.

Reported structure:

1. HR call in English: experience, career direction, projects, reason for leaving.
2. Technical chat plus one coding problem.
3. Final technical round around 4 hours, often split into several 1-hour sessions:
   - Experience deep dive with Engineering Manager.
   - Computer Science fundamentals.
   - Coding problem with follow-ups.
   - Another coding or technical session, sometimes mixed with OS, networking, or performance questions.
   - Behavioral questions about teamwork, conflict, quality, speed, and deadlines.

Important lesson:

> Solving the algorithm is not enough. You must explain Big-O correctly, especially for recursion, nested loops, graph traversal, and follow-up changes.

## Likely Topics

- LeetCode-style data structures and algorithms
- Big-O analysis, including recursive complexity
- CS fundamentals: data structures, OS, networking, memory, CPU/cache basics
- Concurrency and parallelism basics
- REST API contract design
- Request and response schemas
- HTTP methods and status codes
- Validation and error responses
- Authentication and authorization basics
- Pagination, filtering, and sorting
- Idempotency for create/retry flows
- API versioning and backward compatibility
- Testing API behavior

## LeetCode Questions To Practice

1. Two Sum
2. Valid Parentheses
3. Longest Substring Without Repeating Characters
4. Product of Array Except Self
5. Merge Intervals
6. Binary Search
7. Search in Rotated Sorted Array
8. Reverse Linked List
9. Linked List Cycle
10. Binary Tree Level Order Traversal
11. Number of Islands
12. Top K Frequent Elements
13. Climbing Stairs
14. House Robber

For every problem, practice these follow-ups:

- What is the time complexity?
- What is the space complexity?
- What changes if the input is very large?
- Can you optimize with another data structure?
- What edge cases can break this solution?
- If this is recursive, what is the recursion depth and stack space?

## CS Fundamentals Questions To Practice

1. Array vs linked list.
2. Hash map internal design and collision handling.
3. Stack vs heap memory.
4. Process vs thread.
5. Concurrency vs parallelism.
6. Race condition and how to prevent it.
7. Mutex vs semaphore.
8. Deadlock and how to avoid it.
9. OS scheduling at a high level.
10. Virtual memory and paging.
11. CPU cache and why locality matters.
12. TCP vs UDP.
13. What happens when you type a URL and press Enter?
14. Buffer size and packet basics.
15. Database index and B-tree basics.

Performance-style prompt to prepare:

> If the machine has 4 CPU cores, how would you design the code to use the cores well? Explain how you split work, synchronize results, and avoid race conditions.

## API Design Questions To Practice

1. Design APIs for user login and registration.
2. Design APIs for todo/task management.
3. Design APIs for file upload and download.
4. Design APIs for search with filtering, sorting, and pagination.
5. Design APIs for notifications.
6. Design APIs for audit logs.
7. Design APIs for case or evidence management.
8. Design APIs for role-based access control.
9. Design APIs for report generation.
10. Design APIs for webhook/event subscription.

## API Design Checklist

- Resource names and endpoint structure, for example `GET /v1/cases/{id}`.
- Request and response body shape.
- Required fields and validation rules.
- Error format and HTTP status codes.
- Authentication and authorization.
- Pagination strategy: cursor or offset.
- Idempotency for retryable create operations.
- Rate limiting and abuse protection.
- Versioning and backward compatibility.
- Unit, integration, and contract tests.

## How To Communicate

- Clarify requirements before coding.
- Explain tradeoffs.
- Start with a simple solution, then optimize.
- Mention edge cases.
- Talk about testing.
- If using AI tools, emphasize validation, not blind trust.

## Useful Phrases

> Let me first clarify the input constraints.

> A simple approach would be..., but its complexity is...

> We can improve this by using a hash map to avoid repeated scanning.

> I would test the empty input, a single element, duplicates, and the maximum-size case.
