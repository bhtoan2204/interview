# Axon Common Question Bank

These questions are based on Axon's official hiring guidance, the provided Junior Software Engineer job description, and common software engineering interview patterns. Public interview reports vary and are not always reliable, so use this as a high-probability prep bank rather than a guaranteed list.

## Recruiter / HR Questions

1. Tell me about yourself.
2. Why are you interested in Axon?
3. What do you know about Axon's mission?
4. Why do you want this Junior Software Engineer role?
5. Why are you leaving or considering leaving your current role?
6. Which Axon product interests you most, and why?
7. What are your strengths?
8. What are your weaknesses, and how are you improving them?
9. How do you use AI in your work?
10. When should you avoid relying too much on AI?
11. What kind of engineering environment helps you do your best work?
12. Are you comfortable working onsite Tuesday through Friday in Ho Chi Minh City?
13. What are your salary expectations?
14. When can you start?
15. Do you have any questions for me?

## Mission and Culture Questions

1. What does "Protect Life" mean to you as an engineer?
2. Which Axon value resonates with you most?
3. Tell me about a time you took ownership of a difficult problem.
4. Tell me about a time you had to work with candor.
5. Tell me about a time you joined forces with another team.
6. Tell me about a time you aimed far instead of choosing a short-term fix.
7. How do you balance speed and quality?
8. How do you think about software ethics in public safety?
9. How would you build software for users whose work is high-pressure and safety-critical?
10. What product or service from Axon interests you, and why?

## Resume Deep-Dive Questions

1. Walk me through your current work at Hasaki.
2. What was your role in the POS/online payment system?
3. Why did you choose event sourcing for payments?
4. What was difficult about migrating 30M+ records?
5. How did Debezium CDC help your migration?
6. How did you verify data correctness during migration?
7. What reliability issues did you solve in order/shipment flows?
8. How did optimistic locking and distributed locks reduce duplicate transactions?
9. How did your Kafka pipeline improve latency and cache consistency?
10. What did you do with OpenTelemetry tracing?
11. What was your impact at Avian Solutions?
12. How did you reduce issue detection time by 40%?
13. What did you learn from your DevOps internship?
14. Which project are you most proud of?
15. What would you improve if you could redo one project?

## Backend Technical Questions

1. What happens when an HTTP request reaches your backend service?
2. How do you design a REST API?
3. How do you handle validation and error responses?
4. What is idempotency? Why is it important?
5. How do you design retries safely?
6. What is a timeout? Where should timeouts be configured?
7. How do you prevent duplicate processing in a distributed system?
8. What is optimistic locking?
9. What is a distributed lock? What can go wrong?
10. How do you design a payment state machine?
11. How do you handle webhook events from a payment provider?
12. How would you reconcile inconsistent internal and provider payment states?
13. How do SQL indexes work?
14. What can make a query slow?
15. What is a transaction?
16. What is the difference between SQL and NoSQL?
17. When would you use Redis?
18. What cache invalidation strategies do you know?
19. How do you secure sensitive customer data?
20. How do you design logs and metrics for a new feature?

## Event-Driven and Messaging Questions

1. What is Kafka used for?
2. What is a consumer group?
3. What does at-least-once delivery mean?
4. Why must consumers be idempotent?
5. How do you handle failed messages?
6. What is a dead-letter queue?
7. What is the outbox pattern?
8. What is Debezium CDC?
9. What is the difference between event sourcing and event-driven architecture?
10. How do you evolve event schemas safely?

## Go Questions

1. Why do you like using Go for backend services?
2. What are goroutines?
3. What are channels?
4. How do you use context in Go?
5. How do you handle timeouts in Go?
6. How do you avoid goroutine leaks?
7. How do you write tests in Go?
8. What is an interface in Go?
9. How do you structure a Go service?
10. How do you handle errors idiomatically in Go?

## Coding Interview Questions

1. Two Sum
2. Valid Parentheses
3. Longest Substring Without Repeating Characters
4. Merge Intervals
5. Product of Array Except Self
6. Binary Search
7. Top K Frequent Elements
8. Number of Islands
9. LRU Cache
10. Design a rate limiter
11. Implement retry with exponential backoff
12. Parse logs and aggregate counts
13. Find duplicate events
14. Deduplicate transactions by idempotency key
15. Build a simple in-memory cache with TTL

## Coding Follow-Up Questions

1. What is the exact time and space complexity?
2. If the solution is recursive, what is the recursion depth?
3. What is the stack space usage?
4. What edge case breaks the current solution?
5. How would you optimize this if input size is much larger?
6. Can you solve it with a different data structure?
7. Can you make it iterative instead of recursive?
8. Can you reduce memory usage?
9. What changes if the input is a stream?
10. What changes if the input is too large to fit in memory?

## Computer Science Fundamentals Questions

1. What is Big-O notation?
2. How do you analyze recursive time complexity?
3. What is the difference between stack memory and heap memory?
4. What is a process vs a thread?
5. What is OS scheduling?
6. What is context switching?
7. What is concurrency vs parallelism?
8. What is a race condition?
9. How do mutexes prevent race conditions?
10. What is a deadlock?
11. If a machine has 4 CPU cores, how would you use them efficiently?
12. Why can adding more threads make a program slower?
13. What is CPU cache locality?
14. What is virtual memory?
15. What is paging?
16. What is TCP vs UDP?
17. What is a packet?
18. What is a buffer?
19. What is latency vs throughput?
20. What happens when you type a URL and press Enter?

## API Design Questions

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
11. Explain your endpoint naming and resource model.
12. Explain your request and response schema.
13. Explain your validation and error response format.
14. Explain how authentication differs from authorization.
15. Explain cursor pagination versus offset pagination.
16. Explain when to use idempotency keys.
17. Explain how you would test the API contract.

## AI-Assisted Development Questions

1. How do you use AI tools in development?
2. How do you validate AI-generated code?
3. What risks do AI coding tools introduce?
4. How would you use AI to improve code review?
5. How would you use AI to write tests?
6. When should you avoid using AI-generated output?
7. How do you make sure AI-generated code follows team conventions?
8. How do you handle security concerns with AI-assisted development?

## Behavioral Questions

1. Tell me about a time you solved a difficult technical problem.
2. Tell me about a time you handled production pressure.
3. Tell me about a time you made a mistake.
4. Tell me about a time you received critical feedback.
5. Tell me about a time you disagreed with a teammate.
6. Tell me about a time you had to learn a new technology quickly.
7. Tell me about a time you improved performance.
8. Tell me about a time you improved reliability.
9. Tell me about a time you worked with unclear requirements.
10. Tell me about a time you had to make a tradeoff.
11. When would you accept a slower deadline to keep quality?
12. When would you push for speed even if the solution is not perfect?
13. Tell me about a time you had to communicate risk to a teammate or manager.
14. Tell me about a time you were wrong in a technical discussion.
15. How do you respond when an interviewer or teammate challenges your answer?
