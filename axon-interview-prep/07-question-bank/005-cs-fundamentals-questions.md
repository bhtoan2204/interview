# Computer Science Fundamentals — Question Bank

CS fundamentals are weighted heavily in Axon's interview loop, even for a Junior Software Engineer role. Interviewers use them to check that you understand *why* things work, not just which framework to call. These questions are organized by topic, ordered roughly easy → hard within each section. They are tuned to a managed-language backend profile (Go, Java, Scala, C#) since that matches the JD and your background.

How to use this: for each question, be able to (1) give a one-sentence definition, (2) explain a tradeoff, and (3) tie it to something you've actually built (payments, Kafka, CDC, migrations). The third part is what separates a junior from a senior answer.

For deeper notes on OS/concurrency/memory topics, read `006-cs-fundamentals-notes.md` after drilling the questions.

---

## 1. Data Structures

1. What is the difference between an array and a linked list? When would you pick each?
2. How does a dynamic array (slice / ArrayList) grow, and what is the amortized cost of append?
3. How does a hash map work internally? What is a hash collision and how is it resolved?
4. What is the average vs worst-case lookup time of a hash map, and what causes the worst case?
5. What is the difference between a hash map and a balanced binary search tree (e.g. TreeMap)?
6. What is a stack vs a queue? Give a real use case for each.
7. What is a heap / priority queue, and what operations is it good at?
8. What is a binary search tree? What makes it "balanced," and why does balance matter?
9. What is a trie, and when is it better than a hash map?
10. What is a graph? How do you represent one (adjacency list vs adjacency matrix), and what are the tradeoffs?
11. What is a set, and how is it typically implemented?
12. When would you use a ring buffer / circular buffer? (Hint: bounded queues, streaming.)
13. How would you choose a data structure for an LRU cache, and why? (HashMap + doubly linked list.)

## 2. Algorithms & Complexity

1. What is Big-O notation? Explain time vs space complexity.
2. Explain the difference between O(1), O(log n), O(n), O(n log n), and O(n²) with examples.
3. What is the difference between average-case and worst-case complexity? Give an example where they differ.
4. How does binary search work, and what is its precondition?
5. Compare quicksort, mergesort, and heapsort — time, space, and stability.
6. Why is comparison-based sorting bounded at O(n log n)? When can you beat it (counting/radix sort)?
7. What is the difference between BFS and DFS? When would you use each?
8. What is recursion? What is a base case, and what causes a stack overflow?
9. What is dynamic programming? How do you recognize a problem that needs it?
10. What is the two-pointer technique, and when is it useful?
11. What is the sliding-window technique, and what kind of problem signals it?
12. What is a greedy algorithm, and when does greedy fail?
13. How would you detect a cycle in a linked list or graph?
14. What is hashing used for beyond hash maps? (Deduplication, partitioning, consistent hashing.)
15. How do you analyze the complexity of a recursive function?
16. What is recursion depth, and how does it affect space complexity?
17. How do you analyze DFS complexity on a tree vs a graph?
18. If a function recursively splits input into two halves, when is it O(log n), O(n), or O(n log n)?
19. What is the complexity of backtracking, and why is it often exponential?
20. How would you explain Big-O if a solution has sorting plus one linear scan?

## 3. Operating Systems & Concurrency

1. What is the difference between a process and a thread?
2. What is a context switch, and why is it expensive?
3. What is concurrency vs parallelism? (Tie this to goroutines if you use Go.)
4. What is a race condition? How do you prevent one?
5. What is a deadlock? What four conditions must hold, and how do you break them?
6. What is a mutex vs a semaphore?
7. What is the difference between blocking and non-blocking I/O?
8. What is a thread pool, and why not just spawn a thread per request?
9. How does the OS manage memory? What is virtual memory and paging?
10. What is the stack vs the heap in memory? What lives where?
11. What is garbage collection, and what are the tradeoffs vs manual memory management?
12. What happens, step by step, when a program calls a function (the call stack)?
13. (Go-specific) How do goroutines differ from OS threads? What does the Go scheduler do?
14. (Go-specific) What is a channel, and how does it help avoid shared-memory races?
15. What is OS scheduling, and why does a process or thread get paused?
16. What is CPU utilization?
17. If a machine has 4 CPU cores, how would you design a task to use the cores well?
18. What types of work benefit from parallel execution, and what types do not?
19. What is false sharing or cache contention at a high level?
20. Why can adding more threads make a program slower?

## 4. Networking

1. Walk me through what happens when you type a URL and press Enter.
2. What is the difference between TCP and UDP? When would you choose each?
3. What is the TCP three-way handshake?
4. How does DNS resolution work?
5. What is the difference between HTTP/1.1, HTTP/2, and HTTP/3 at a high level?
6. What are the main HTTP methods and their semantics (GET, POST, PUT, PATCH, DELETE)?
7. Which methods are idempotent and which are safe? Why does it matter for retries?
8. What do these status codes mean: 200, 201, 400, 401, 403, 404, 409, 429, 500, 503?
9. What is the difference between HTTP and HTTPS? What does TLS provide?
10. What is a load balancer, and what does it do?
11. What is the difference between latency and throughput?
12. What is a CDN, and when does it help?
13. What is the difference between long polling, WebSockets, and Server-Sent Events?
14. What is a packet?
15. What is a buffer, and what can happen if the buffer is too small or too large?
16. What is packet loss, and how does TCP handle it?
17. Why is UDP useful even though it does not guarantee delivery?

## 5. Databases

1. What is a transaction? Explain ACID.
2. What are the SQL isolation levels, and what anomalies does each prevent (dirty read, non-repeatable read, phantom read)?
3. How does a database index work? What data structure is usually behind it (B-tree)?
4. When can adding an index *hurt* performance?
5. What is the difference between a clustered and a non-clustered index?
6. What can make a query slow, and how would you diagnose it (EXPLAIN / query plan)?
7. What is the difference between SQL and NoSQL? When would you choose a document or key-value store?
8. What is normalization, and when would you deliberately denormalize?
9. What is the difference between optimistic and pessimistic locking?
10. What is the CAP theorem? Give a practical interpretation.
11. What is eventual consistency, and where is it acceptable?
12. What is sharding vs replication?
13. What is a primary key vs a foreign key vs a unique constraint?

## 6. Language / Programming Concepts

1. What is the difference between pass-by-value and pass-by-reference? Which does your main language use?
2. What is the difference between a value type and a reference type?
3. What is the difference between compile-time and runtime errors?
4. What is static vs dynamic typing? Strong vs weak typing?
5. What is an interface, and why is "program to an interface" useful?
6. What are the core OOP principles (encapsulation, inheritance, polymorphism, abstraction)?
7. What is composition vs inheritance, and why is composition often preferred?
8. What is immutability, and why does it help with concurrency?
9. What is the difference between an exception and an error code? How does your language handle errors idiomatically?
10. What is a memory leak, and how can one happen in a garbage-collected language?
11. What are the SOLID principles? Pick one and explain it with an example.
12. What is the difference between a shallow copy and a deep copy?

## 7. Practical / Applied CS (likely follow-ups for your background)

1. Why must a Kafka consumer be idempotent, and how does this relate to "at-least-once" delivery?
2. How does an idempotency key prevent duplicate payments — what data structure / constraint enforces it?
3. Your migration moved 30M+ records. How would you reason about its time complexity and memory footprint?
4. How does CDC (Debezium) capture changes — and why is reading the DB log cheaper than polling?
5. Why does an LRU cache give O(1) get and put, and what would break if you used only a hash map?
6. How would you design a rate limiter? Compare token bucket vs sliding window.
7. How would you deduplicate a very large stream of events that doesn't fit in memory? (Bloom filter, partitioning.)
8. How does a distributed lock work, and what CS problem (mutual exclusion) is it solving?

---

## Quick self-test (rapid-fire warmup)

Answer each in one breath before the interview:

- Big-O of: hash map lookup, binary search, sorting, BFS on a graph?
- Process vs thread in one sentence?
- TCP vs UDP in one sentence?
- ACID in one sentence?
- What makes an operation idempotent?
- Stack vs heap memory?
- Optimistic vs pessimistic locking?
- Deadlock — name two of the four conditions?
- Recursive DFS on a tree — time and space complexity?
- Recursive DFS on a graph with visited set — time and space complexity?
- 4 CPU cores — what does parallelism mean in one sentence?

---

## Notes on strategy

- **Connect every answer to experience.** Axon's JD stresses owning end-to-end features and validating quality. When you explain hash maps, mention the idempotency dedup. When you explain B-trees, mention diagnosing a slow query. This shows applied understanding.
- **Do not hand-wave Big-O.** Recent candidate reports suggest wrong complexity analysis can hurt even when the algorithm passes. For recursion, always mention recursion depth and stack space.
- **Expect textbook-style CS questions.** Some questions may feel like university topics: scheduling, memory, cache, packets, and CPU cores. Give a simple definition, then connect it to backend performance or reliability.
- **Don't over-answer.** Give the definition, one tradeoff, then stop and let them dig. Rambling hides weak spots more than it hides them.
- **It's OK to think out loud and say "I'm not 100% sure, but my reasoning is..."** — juniors are judged on reasoning quality, not memorized trivia.
- **Be honest about gaps.** If you've never used UDP, say what you know and how you'd find out. The JD literally says "Don't meet every single requirement? That's ok."
