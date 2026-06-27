# Mock Interview Sets

## Mock 1: Recruiter Screen

1. Tell me about yourself.
2. Why Axon?
3. What do you know about our mission?
4. Why are you interested in a junior role?
5. Which Axon product interests you most, and why?
6. What are your strengths?
7. What are your weaknesses, and how are you improving them?
8. How do you use AI in your work?
9. When should you avoid relying too much on AI?
10. Are you comfortable with the hybrid schedule?
11. What are your salary expectations?
12. When can you start?
13. What questions do you have for me?
14. Why do you think you are a good culture fit?

## Mock 2: Technical Screen

1. Coding: Longest Substring Without Repeating Characters.
2. Coding: Merge Intervals.
3. Coding: Number of Islands.
4. What is the time and space complexity of your solution?
5. What changes if the input is much larger?
6. How would you test your coding solution?
7. Design APIs for file upload and download.
8. What is idempotency and where would you use it in an API?
9. What questions do you have for me?

## Mock 3: Backend Deep Dive

1. Explain your event sourcing payment system.
2. What events did you store?
3. How did you handle state transitions?
4. What could go wrong with event sourcing?
5. Explain the 30M+ record migration.
6. How did Debezium CDC work in your system?
7. How did you verify consistency?
8. What would you do differently next time?

## Mock 4: API Design

Prompt:

> Design REST APIs for notifications for public safety users.

Expected coverage:

- Main resources and use cases
- Endpoint names and HTTP methods
- Request and response schemas
- Validation and error format
- Status codes
- Pagination for notification history
- User preferences
- Idempotency for send/retry flows
- Authentication and authorization
- Security and privacy
- API tests

Follow-up questions:

1. How do you avoid sending duplicate notifications?
2. What status code do you return for validation failure?
3. How do you support high-priority notifications in the API contract?
4. How do you design pagination for notification history?
5. How do you test this API?

## Mock 5: Axon Values and Behavioral

1. Which Axon value resonates with you most?
2. Tell me about a time you owned a problem end to end.
3. Tell me about a time you used feedback to improve an outcome.
4. Tell me about a time you disagreed with a teammate.
5. Tell me about a time you balanced speed and quality.
6. Tell me about a time you learned a new domain quickly.
7. What would make you successful at Axon?

## Mock 6: Final 4-Hour Loop

Use this to simulate the deeper final loop reported by recent candidates.

### Session 1: Experience Deep Dive

1. Walk me through your Hasaki backend work.
2. What was the most difficult production issue you handled?
3. Why did you use idempotency in payments?
4. How did you verify data correctness during migration?
5. What would you improve if you rebuilt the system today?

### Session 2: CS Fundamentals

1. Explain Big-O and analyze a recursive function.
2. Process vs thread.
3. OS scheduling and context switching.
4. Stack vs heap memory.
5. CPU cache locality.
6. TCP vs UDP.
7. Packet, buffer, latency, and throughput.
8. Database index and B-tree basics.

### Session 3: Coding With Follow-Up

1. Solve one medium coding problem.
2. Explain time and space complexity exactly.
3. Explain edge cases.
4. Optimize with a different data structure.
5. Change one constraint and adapt the solution.

### Session 4: Performance, Concurrency, Behavioral

1. If the system has 4 CPU cores, how would you split work?
2. How do you avoid race conditions?
3. When can parallel execution make performance worse?
4. When would you accept a slower deadline to keep quality?
5. When would you push speed over a perfect solution?
6. Tell me about a team conflict and how you handled it.
