# Axon Junior Software Engineer Interview Preparation

## 1. Role Targeting Summary

**Target role:** Junior Software Engineer, Axon, Ho Chi Minh City  
**Company mission:** Protect Life  
**Work style:** Hybrid, onsite Tuesday to Friday, remote Monday  
**Key role themes:** ownership, mission-critical software, code quality, performance, collaboration, AI-assisted development, modern engineering practices

### Your strongest matching points

- Backend engineer with production experience in Go, JavaScript/TypeScript, and cloud-native systems.
- Experience with high-volume order management systems handling 60K-100K daily orders and up to 300K peak-day orders.
- Strong backend fundamentals: microservices, event-driven architecture, CQRS, event sourcing, concurrency control, SQL/NoSQL, Redis, Kafka, Debezium CDC.
- Practical experience in payments, data consistency, transaction state management, retries, timeouts, optimistic locking, distributed locks, observability, and production debugging.
- Cloud and DevOps background with AWS, EKS, Docker, Kubernetes, ArgoCD, Jenkins, Terraform, Prometheus, Grafana, OpenTelemetry.
- Good mission fit: you have built systems where reliability, auditability, privacy, and operational safety matter.

### Interview positioning

You should present yourself as:

> A backend-focused software engineer who has already worked on production systems with real traffic, payment flows, data migration, observability, and reliability challenges. I am still early in my career, but I have taken ownership of end-to-end backend features and I enjoy learning quickly, validating my work carefully, and using modern tools, including AI-assisted workflows, to improve development speed and quality.

## 2. Full Interview Process Preparation

### Stage 1: Recruiter or HR Screen

**Goal:** Confirm motivation, communication, availability, salary range, English ability, and role fit.

Prepare to answer:

1. Tell me about yourself.
2. Why Axon?
3. Why this Junior Software Engineer role?
4. What do you know about Axon?
5. Which Axon product interests you most, and why?
6. What are your strengths?
7. What are your weaknesses, and how are you improving them?
8. How do you use AI in your work?
9. How do you make sure AI does not reduce code quality or understanding?
10. Tell me about a project you are proud of.
11. What kind of team or engineering culture are you looking for?
12. Are you comfortable with the hybrid schedule in Ho Chi Minh City?
13. What are your salary expectations?
14. When can you start?
15. Do you have any questions for me?

Strong message:

> I am looking for a role where I can grow quickly while contributing to meaningful production systems. Axon is attractive to me because the mission is clear and important, and the role combines backend engineering, collaboration, code quality, and AI-assisted development. I am comfortable with the hybrid schedule because I value in-person collaboration, especially as an engineer who wants to learn from experienced teammates.

Product angle to prepare:

> The Axon product I am most interested in is Draft One because it uses AI to reduce report-writing time for officers, which can let them spend more time on field work and community interaction. I also think it is technically interesting because the product must balance productivity with accuracy, privacy, reviewability, and accountability. That is the kind of responsible AI-assisted workflow I want to learn to build.

### Stage 2: Technical Phone Screen

**Goal:** Check coding ability, API design thinking, problem solving, communication, and technical ownership.

Likely topics:

- LeetCode-style data structures and algorithms
- API design, especially REST contracts
- Request and response schema design
- Validation and error handling
- Authentication and authorization basics
- Pagination, filtering, and sorting
- Idempotency for create/retry flows
- Status codes and versioning
- Testing API behavior
- Code review mindset

How to communicate:

- Clarify requirements before coding.
- Explain tradeoffs.
- Start with a simple solution, then optimize.
- Mention edge cases.
- Talk about testing.
- If using AI tools, emphasize validation, not blind trust.

LeetCode areas to prioritize:

- Arrays and strings
- Hash maps and sets
- Two pointers
- Sliding window
- Stack and queue
- Binary search
- Linked list basics
- Trees with BFS/DFS
- Graph basics such as Number of Islands
- Heap and sorting problems such as Top K Frequent and Merge Intervals
- Basic dynamic programming such as Climbing Stairs or House Robber

API design prompts to practice:

- Design APIs for user login and registration.
- Design APIs for todo/task management.
- Design APIs for file upload and download.
- Design APIs for search with filtering, sorting, and pagination.
- Design APIs for notifications.
- Design APIs for audit logs.
- Design APIs for case or evidence management.
- Design APIs for role-based access control.
- Design APIs for report generation.
- Design APIs for webhooks or event subscriptions.

### Stage 3: Coding Interview

**Goal:** Demonstrate clear problem solving, clean implementation, and ability to reason under pressure.

Recommended practice areas:

- Arrays and strings
- Hash maps
- Two pointers
- Sliding window
- Stack and queue
- Binary search
- Linked list basics
- Trees and BFS/DFS
- Graph basics
- Dynamic programming basics

Answering framework:

1. Restate the problem.
2. Ask clarifying questions.
3. Discuss examples and edge cases.
4. Propose a brute-force solution.
5. Improve the approach.
6. Analyze time and space complexity.
7. Implement cleanly.
8. Test with sample and edge cases.

Useful phrases:

> Let me first clarify the input constraints.

> A simple approach would be..., but its complexity is...

> We can improve this by using a hash map to avoid repeated scanning.

> I would test the empty input, a single element, duplicates, and the maximum-size case.

### Stage 4: API Design Interview

**Goal:** Assess whether you can design clear API contracts and reason about client/server behavior.

Based on the expected interview format, do not over-index on broad system design. Keep the answer focused on API contract, validation, edge cases, security, and testing.

Likely prompts:

- Design APIs for task management.
- Design APIs for file upload/download.
- Design APIs for notification preferences and sending notifications.
- Design APIs for audit logs.
- Design APIs for case/evidence management.
- Design APIs for report generation.
- Design APIs for webhooks.

Recommended answer structure:

1. Clarify users, resources, and main use cases.
2. Define endpoints and HTTP methods.
3. Define request and response bodies.
4. Explain validation rules and error format.
5. Choose status codes.
6. Add pagination, filtering, and sorting where needed.
7. Cover auth, authorization, and privacy.
8. Handle idempotency and retries for create operations.
9. Mention tests and backward compatibility.

Example: evidence upload API

> I would first clarify who can upload evidence, what metadata is required, file size limits, and whether upload should be direct or pre-signed. A simple contract could be `POST /v1/evidence/uploads` to create an upload session, then `PATCH /v1/evidence/{id}` to update metadata. I would include validation errors in a consistent format, require authorization by case access, and keep audit logs for create/update actions. For retries, I would support an idempotency key so the client does not create duplicate evidence records.

Tie to your CV:

> In my current role, I designed backend APIs for payment and order workflows where correctness, idempotency, state transitions, and auditability mattered. I would bring the same careful API thinking to Axon's mission-critical workflows.

### Stage 5: Behavioral Interview

**Goal:** Understand how you work with people, handle pressure, learn, receive feedback, and take ownership.

Use STAR:

- Situation
- Task
- Action
- Result

Core stories to prepare:

1. Ownership: Hasaki payment system or order/invoice migration.
2. Technical challenge: 30M+ record migration with Debezium CDC and fallback jobs.
3. Reliability improvement: optimistic locking, distributed locks, retries, and timeouts.
4. Collaboration: payment service with ZaloPay, VNPAY, Techcombank, internal wallet.
5. Debugging: OpenTelemetry tracing for logistics and marketplace integrations.
6. Learning quickly: moving from DevOps internship to backend/cloud production work.
7. Conflict or feedback: code review or design discussion.
8. Customer impact: reducing latency to under 100ms or improving issue detection by 40%.

## 3. Your Main Interview Pitch

### 60-second answer: Tell me about yourself

> My name is Banh Hao Toan. I am a backend engineer based in Ho Chi Minh City, with experience building production backend services in Go, TypeScript, and cloud-native environments. In my current role at Hasaki, I work on high-volume order management and payment systems, handling around 60K to 100K daily orders and up to 300K orders on peak days. I have worked on payment flows, event sourcing, Kafka-based pipelines, Debezium CDC migration, Redis caching, observability, and reliability improvements such as retries, timeouts, optimistic locking, and distributed locks.
>
> Before that, I worked at Avian Solutions, where I built backend modules and helped maintain AWS EKS environments with observability using Grafana, Jaeger, and OpenTelemetry. I also started as a DevOps intern, so I am comfortable thinking about both application code and production operations.
>
> I am interested in Axon because the mission to Protect Life is meaningful, and the role matches the kind of engineering I want to grow in: building reliable, high-quality software, collaborating with product and design, and using modern AI-assisted development practices responsibly.

### Short version

> I am a backend engineer with production experience in Go, TypeScript, cloud infrastructure, payments, event-driven systems, and observability. I enjoy building reliable services and learning quickly. Axon interests me because I want to contribute to mission-critical software while growing in a strong engineering culture.

## 4. Why Axon?

Strong answer:

> Axon stands out to me because the mission is not abstract. Protecting life and improving public safety are areas where software quality, reliability, and responsible engineering really matter. I also like that the role is not only about writing code, but about owning features, working with product and design, participating in code reviews, and building high-quality software in an AI-first engineering environment. That matches how I want to grow: technically strong, collaborative, and careful about real-world impact.

Optional personal angle:

> In my previous work, I have seen how backend reliability affects real users, especially in payment, order, and notification systems. At Axon, that impact is even more meaningful because the software supports public safety workflows.

## 5. Why Should We Hire You?

> I bring a combination of backend engineering, production ownership, and strong learning ability. Even though this is a junior role, I already have experience working on systems with real traffic, payments, migrations, observability, and operational incidents. I understand that mission-critical software requires more than just implementing features. It requires testing, clear code, careful handling of edge cases, observability, and collaboration. I also have a DevOps background, so I can think about how software behaves after deployment, not only how it works locally.

## 6. Technical Questions and Model Answers

### What is event sourcing, and why did you use it?

> Event sourcing stores state changes as a sequence of events instead of only storing the latest state. For payments, this is useful because every transition needs to be auditable, such as payment created, authorized, captured, failed, refunded, or partially refunded. In my Hasaki project, I used event sourcing for US-market POS and online payment systems to guarantee an audit trail and make transaction state management more reliable. The tradeoff is that event schemas and replay logic must be designed carefully.

### How do you prevent duplicate payment transactions?

> I would use idempotency keys, unique constraints, provider transaction references, and careful state transitions. If the client retries the same request, the backend should recognize it as the same operation instead of creating a new transaction. I would also use timeouts and retry policies carefully because external payment providers can have delayed responses. For critical flows, I would add reconciliation jobs to compare internal state with provider state.

### What is the difference between optimistic locking and distributed locking?

> Optimistic locking assumes conflicts are rare. It usually uses a version field or updated timestamp, and the update succeeds only if the version has not changed. It is good for protecting database updates without blocking too much.
>
> Distributed locking is used when multiple service instances or workers may process the same resource at the same time. It can be implemented with Redis or another coordination system. It is useful but must be designed carefully with expiration, ownership tokens, and failure handling.

### How would you debug a production issue?

> I would first check the impact and scope: which users or flows are affected, when it started, and whether it is still happening. Then I would look at logs, metrics, traces, recent deployments, and external dependencies. If the issue is severe, I would focus first on mitigation, such as rollback, feature flag, retry adjustment, or temporarily disabling a non-critical path. After stabilizing the system, I would find root cause and add tests, alerts, or instrumentation to prevent recurrence.

### How do you validate AI-generated code?

> I treat AI-generated code as a draft, not as final code. I review the logic, check edge cases, run tests, inspect security and performance implications, and make sure the code follows project conventions. For backend code, I pay attention to error handling, transaction boundaries, idempotency, concurrency, and observability. I think AI tools can improve speed, but the engineer is still responsible for correctness and quality.

### What makes code production-ready?

> Production-ready code should be correct, readable, tested, observable, and maintainable. It should handle errors, timeouts, retries, edge cases, and expected failure modes. It should also follow team conventions, avoid unnecessary complexity, and include logs or metrics where needed for debugging.

### How do you design a Kafka-based pipeline?

> I start by defining the event source, event schema, consumers, ordering requirements, retry strategy, and failure handling. I also consider idempotent consumers because messages may be delivered more than once. For data consistency, I have used the outbox pattern with MySQL and Kafka, then consumed events into Elasticsearch to improve query latency and cache consistency.

## 7. Behavioral Questions and STAR Answers

### Tell me about a technical challenge you solved.

**Situation:** At Hasaki, legacy PHP Order and Invoice services needed to be rebuilt into Go services.  
**Task:** Migrate more than 30M records while minimizing data loss and rollout risk.  
**Action:** I helped rebuild the services in Go and used Debezium CDC plus fallback jobs to keep data synchronized during migration. I also paid attention to rollout safety, consistency checks, and failure recovery.  
**Result:** The migration reduced dependency on legacy services and lowered data loss risk during rollout.

Answer:

> One technical challenge I worked on was rebuilding legacy PHP Order and Invoice services into Go services at Hasaki. The challenge was not only rewriting the service logic, but also migrating more than 30M records safely. I used Debezium CDC and fallback jobs to reduce data loss risk during rollout. I also focused on consistency checks and incremental migration instead of a risky one-time switch. This taught me that migration work is as much about safety and observability as it is about code.

### Tell me about a time you improved reliability.

> In Hasaki's order and shipment flows, duplicate or failed transactions could affect critical business operations. I improved reliability by adding optimistic locking, distributed locks, retries, and timeouts. The goal was to reduce race conditions and make failure behavior more predictable. This reduced duplicate and failed transaction incidents in critical flows and helped the system behave more safely under concurrent operations.

### Tell me about a time you used observability.

> At Avian Solutions, I set up observability with Grafana, Jaeger, and OpenTelemetry. This reduced issue detection time by around 40% and enabled alerts for bugs via Slack and major incidents via email. At Hasaki, I also added OpenTelemetry tracing for logistics and marketplace integrations, which made production issues easier to detect and debug.

### Tell me about a time you learned something quickly.

> I started with DevOps work, setting up CI/CD pipelines and cloud environments, then moved deeper into backend engineering. I had to learn business logic, backend patterns, distributed systems, and production debugging quickly. Working across DevOps and backend helped me understand not only how to write services, but also how they are deployed, monitored, and operated.

## 8. Questions You Should Ask Axon

Ask 3-5 questions depending on time.

1. What does success look like for a Junior Software Engineer in the first six months?
2. What kinds of products or backend systems would this role likely work on?
3. How does the team use AI-assisted development tools in daily engineering work?
4. How do engineers at Axon validate code quality for mission-critical software?
5. What is the code review culture like on the team?
6. How does the team balance speed of delivery with reliability and safety?
7. What mentorship structure is available for junior engineers?
8. What are the biggest technical challenges the HCMC engineering team is solving this year?

## 9. Weaknesses and Risk Areas to Prepare

### Risk: role asks for Java, Scala, Go, or C#

Your response:

> My strongest managed/backend language is Go, and I also have experience with JavaScript/TypeScript and some Java. I am comfortable learning a new language or framework when needed because the backend fundamentals transfer: APIs, data modeling, concurrency, testing, reliability, and observability.

### Risk: junior role but your CV may look backend-heavy

Your response:

> I see the role as a strong growth opportunity. I have production experience, but I still want to learn from senior engineers, improve my engineering discipline, and contribute in a team where quality and mission matter. I am comfortable taking ownership while also being open to feedback and mentorship.

### Risk: English speaking confidence

Strategy:

- Speak slower than usual.
- Use simple sentences.
- Structure answers clearly.
- Ask for clarification when needed.

Useful phrase:

> Could I take a few seconds to organize my thoughts?

## 10. AI-Assisted Development Talking Points

Axon emphasizes an AI-first engineering environment. Prepare a mature answer.

Good answer:

> I use AI tools to speed up exploration, generate test cases, summarize documentation, review possible edge cases, and draft simple implementation ideas. But I do not rely on AI output blindly. I always verify correctness through code review, tests, local runs, and reasoning about edge cases. For production backend systems, I am especially careful with concurrency, transaction consistency, security, performance, and error handling.

Examples you can mention:

- Generate unit test ideas for edge cases.
- Ask AI to compare implementation options.
- Use AI to explain unfamiliar libraries.
- Use AI to draft boilerplate, then manually review.
- Use AI to help write documentation or code review notes.

Avoid saying:

- "AI writes most of my code."
- "I trust AI-generated code if it compiles."
- "I do not need to understand the generated code."

## 11. Coding Practice Checklist

Practice these before the interview:

- Two Sum
- Valid Parentheses
- Merge Intervals
- Longest Substring Without Repeating Characters
- Product of Array Except Self
- Binary Search
- Top K Frequent Elements
- Number of Islands
- BFS shortest path
- LRU Cache
- Design an idempotent API
- Implement a retry mechanism with timeout
- Basic Go concurrency with goroutines and channels

For each problem, practice explaining:

- Approach
- Complexity
- Edge cases
- Tests

## 12. Final Day Preparation

### One day before

- Review Axon's mission: Protect Life.
- Review your three strongest stories: payment system, migration, observability.
- Practice your 60-second self-introduction out loud.
- Prepare questions for interviewer.
- Review Go basics, SQL transactions, Redis, Kafka, and REST API design.

### Interview day

- Keep answers structured.
- Be honest when you do not know something.
- Show curiosity and willingness to learn.
- Tie answers back to reliability, quality, mission, and users.
- Do not oversell. Your strongest signal is practical production experience plus humility.

## 13. Closing Statement

Use this near the end if appropriate:

> Thank you for the conversation. I am excited about this role because it combines meaningful mission, backend engineering, code quality, and modern AI-assisted development. I believe my experience with production backend systems, payments, event-driven architecture, cloud infrastructure, and observability would help me contribute quickly while continuing to grow with the team.
