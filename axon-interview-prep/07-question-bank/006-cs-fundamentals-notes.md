# Computer Science Fundamentals - Detailed Notes

Muc tieu cua file nay: bien cac y trong mock thanh kien thuc co the hoc, hieu, va noi lai trong phong van. Khi tra loi, dung format ngan:

1. Dinh nghia.
2. Vi sao no ton tai / tradeoff.
3. Vi du backend thuc te.

---

## 1. Process va Thread

### Process la gi?

Process la mot chuong trinh dang chay. Moi process thuong co mot vung dia chi ao rieng, tai nguyen rieng, va metadata rieng do OS quan ly.

Mot process thuong co:

- Virtual address space rieng.
- Page table rieng.
- File descriptor table rieng.
- Heap, stack, code segment, data segment.
- Process ID, quyen truy cap, scheduling metadata.

Diem quan trong: process duoc cach ly voi process khac. Neu process A ghi sai memory cua no, no khong duoc phep ghi truc tiep vao memory cua process B. Cach ly nay giup security va stability.

### Thread la gi?

Thread la mot dong thuc thi ben trong process. Nhieu thread trong cung process co the chay cung mot codebase va share cung address space.

Thread trong cung process share:

- Heap.
- Global/static variables.
- Code segment.
- Open files va sockets.
- Virtual address space.

Moi thread co rieng:

- Stack.
- Registers.
- Program counter.
- Thread-local storage.
- Scheduling state.

Diem quan trong: vi thread share memory, communication giua thread nhanh hon process, nhung de gay race condition hon.

### Process vs thread - cau tra loi phong van

> A process is an isolated running program with its own virtual address space and OS resources. A thread is an execution unit inside a process. Threads in the same process share memory such as heap and globals, but each thread has its own stack and CPU state. Threads are lighter to create and switch, but shared memory means we need synchronization.

### Context switch la gi?

Context switch la luc OS tam dung mot process/thread va cho process/thread khac chay.

Voi thread switch trong cung process, OS/cpu thuong phai luu va khoi phuc:

- Registers.
- Stack pointer.
- Program counter.
- Scheduling state.

Voi process switch, ngoai nhung thu tren con co the phai doi:

- Address space context.
- Page table context.
- TLB entries / cache locality.

Khong phai context switch la "load lai toan bo RAM". OS khong copy het memory cua process moi lan switch. Cai dat la CPU state va memory mapping context.

### Vi sao context switch dat?

No dat vi:

- CPU phai luu/khoi phuc state.
- Cache locality bi xau di: data cua thread cu co the dang trong CPU cache, thread moi lai can data khac.
- TLB co the bi miss nhieu hon sau khi doi address space.
- Kernel scheduling co overhead.

Backend angle: neu tao qua nhieu thread cho request, CPU co the mat nhieu thoi gian switch hon la lam viec that. Vi vay hay dung thread pool / worker pool / async I/O tuy workload.

### Concurrency vs parallelism

Concurrency la nhieu task dang duoc quan ly trong cung mot khoang thoi gian. Chung co the xen ke nhau.

Parallelism la nhieu task that su chay cung luc tren nhieu core.

Vi du:

- 1 core chay 100 goroutines: concurrent, khong parallel tai cung mot thoi diem.
- 4 core chay 4 CPU-bound tasks cung luc: parallel.

Cau tra loi ngan:

> Concurrency is about structuring multiple tasks that can make progress over time. Parallelism is about actually executing multiple tasks at the same time, usually on multiple CPU cores.

### Neu may co 4 CPU cores thi dung the nao?

Phu thuoc workload:

- CPU-bound: chia data thanh cac partition doc lap, chay toi da khoang so core, tranh shared mutable state.
- I/O-bound: co the co nhieu goroutine/thread hon so core vi nhieu task dang cho network/disk.
- Shared resource: dung mutex/channel/queue hoac partition ownership de tranh race.
- Do luong: tang parallelism den khi throughput khong tang nua hoac latency/context switch/cache contention xau di.

Khong nen tra loi "tao 4 thread la xong". Phai noi them ve correctness, bottleneck, va measurement.

---

## 2. Race Condition, Mutex, Semaphore, Atomic

### Race condition la gi?

Race condition xay ra khi ket qua cua chuong trinh phu thuoc vao timing khong kiem soat duoc giua nhieu thread/goroutine.

Vi du kinh dien:

```text
counter++
```

Nhin nhu mot lenh, nhung logic co the la:

```text
read counter
modify value
write counter
```

Neu hai thread cung doc `counter = 10`, ca hai cung tinh thanh `11`, roi ca hai cung ghi `11`, ket qua mat mot lan increment. Day la lost update.

### Data race vs race condition

Data race la mot loai race cu the: nhieu thread access cung mot memory location, it nhat mot thread write, va khong co synchronization.

Race condition rong hon: ket qua sai do ordering/timing. Co nhung race condition khong nhat thiet la data race, vi tung access co the synchronized nhung overall sequence van sai.

Vi du backend: hai request deu check "order chua refund" roi cung tao refund. Tung query co the an toan, nhung check-then-act khong atomic nen van duplicate refund neu khong co transaction/lock/unique constraint.

### Critical section la gi?

Critical section la doan code dung chung tai nguyen shared mutable state va can duoc bao ve.

Vi du:

- Update map shared giua goroutines.
- Update counter dung chung.
- Check va update payment state.
- Lay connection tu pool.

### Mutex

Mutex la mutual exclusion lock. No dam bao chi mot thread vao critical section tai mot thoi diem.

Dac diem:

- Co ownership: thread/goroutine lock thi thuong chinh no unlock.
- Phu hop de bao ve mot invariant gom nhieu bien.
- Don gian de reason hon atomic khi state phuc tap.

Vi du invariant:

```text
available_balance >= 0
reserved_balance >= 0
total = available_balance + reserved_balance
```

Neu update nhieu field lien quan, mutex hoac transaction thuong hop ly hon atomic rieng le.

### Semaphore

Semaphore la mot counter permits.

Neu semaphore co N permits, toi da N worker duoc vao cung luc.

Dung cho:

- Limit concurrent requests den API ngoai.
- Connection pool.
- Worker pool.
- Rate limit noi bo.
- Signaling giua producer/consumer.

Khac mutex:

- Mutex tap trung vao ownership va mutual exclusion.
- Semaphore tap trung vao dem permits va signaling.

Cau tra loi ngan:

> A mutex protects a critical section with ownership, usually allowing one thread at a time. A semaphore is a counter of permits and can allow N threads at once or be used for signaling.

### Atomic

Atomic operation la operation duoc thuc hien nhu mot buoc khong bi chen ngang o muc memory operation.

Phu hop cho:

- Counter.
- Flag.
- Pointer/state nho.
- Reference count.

Vi du:

```go
atomic.AddInt64(&counter, 1)
```

Atomic khong nhat thiet dung lock o muc API, nhung CPU/compiler van phai dam bao atomicity va memory ordering.

### Khi nao dung atomic, khi nao dung mutex?

Dung atomic khi:

- State rat nho.
- Operation don gian.
- Invariant don bien.
- Hieu ro memory visibility/order.

Dung mutex khi:

- Can update nhieu bien cung luc.
- Co invariant phuc tap.
- Can code de doc, de review.
- Can bao ve map/list/object mutable.

Rule practical:

> If I only need a counter or a boolean flag, atomic can be good. If I need to protect a business invariant or update multiple fields together, I prefer a mutex or database transaction because it is easier to reason about.

### Memory visibility

Mot bug hay bi bo qua: khong chi la "write co bi chen ngang khong", ma con la "thread khac co thay write moi nhat khong".

Compiler va CPU co the reorder instruction neu van dung trong single-thread view. Synchronization primitive nhu mutex, channel, volatile/atomic voi ordering phu hop giup tao happens-before relationship de thread khac thay state dung.

Phong van junior khong can di qua sau, nhung nen biet: atomic/mutex khong chi chong lost update, no con lien quan den visibility.

---

## 3. Deadlock, Starvation, Livelock

### Deadlock la gi?

Deadlock la tinh huong cac thread/process cho nhau mai mai, khong ai tiep tuc duoc.

Vi du:

```text
Thread 1: lock A, cho lock B
Thread 2: lock B, cho lock A
```

Ca hai deu giu mot lock va doi lock con lai. Khong thread nao tu tien trien.

### Bon dieu kien Coffman

Deadlock chi co the xay ra khi ca 4 dieu kien cung dung:

1. Mutual exclusion: resource chi mot thread dung tai mot thoi diem.
2. Hold and wait: thread vua giu resource vua cho resource khac.
3. No preemption: resource khong bi ep lay lai, thread phai tu release.
4. Circular wait: co vong cho nhau.

Muon tranh deadlock, pha it nhat mot dieu kien.

### Cach tranh deadlock

Practical techniques:

- Lock theo cung mot thu tu o moi noi.
- Gop lock bang primitive an toan nhu `std::scoped_lock` khi co.
- Dung `try_lock` + backoff.
- Giam thoi gian giu lock.
- Khong goi network/disk/external API trong luc giu lock neu co the.
- Dung timeout.
- Thiet ke ownership ro rang: mot worker so huu mot partition state.

Database angle:

- Transaction update rows theo thu tu khac nhau cung co the deadlock.
- Nen update theo order on dinh, transaction ngan, retry transaction khi gap deadlock.

### Starvation

Starvation la khi mot thread cho qua lau hoac cho mai vi thread khac lien tuc lay duoc resource.

Khac deadlock:

- Deadlock: he thong bi ket vong cho, khong tien trien cho cac ben lien quan.
- Starvation: he thong van tien trien, nhung mot task bi doi tai nguyen qua lau.

Nguyen nhan:

- Scheduler khong fair.
- Lock khong fair.
- Priority cao luon thang priority thap.
- Queue bi uu tien sai.

Vi du backend: job low-priority khong bao gio chay vi high-priority jobs lien tuc vao queue.

### Livelock

Livelock la khi cac thread van chay, van thay doi state, nhung khong hoan thanh cong viec.

Vi du:

```text
Thread 1 thay conflict -> release lock -> retry
Thread 2 thay conflict -> release lock -> retry
Ca hai retry cung timing, lap mai
```

Cach giam:

- Randomized backoff.
- Jitter.
- Coordinator/priority.
- Lock ordering.

Cau ngan:

> In deadlock, threads stop because each waits for another resource. In livelock, threads keep reacting and changing state, but no useful progress is made. In starvation, the system progresses, but one task keeps being denied resources.

---

## 4. Virtual Memory

### Virtual address la gi?

Program thuong khong lam viec truc tiep voi physical RAM address. No thay virtual address.

CPU/MMU dich:

```text
virtual address -> physical address
```

Moi process co cam giac nhu no co mot address space rieng lien tuc, du physical memory that co the bi chia nho va nam o nhieu frame khac nhau.

### Vi sao can virtual memory?

Virtual memory giup:

- Process isolation.
- Memory protection: read/write/execute permission.
- Moi process co address space rieng.
- Lazy allocation: chi cap physical page khi can.
- Shared memory giua process.
- Memory-mapped file.
- Copy-on-write.
- Swap khi memory thieu.
- Chay program lon hon physical RAM trong mot so truong hop, voi tradeoff performance.

### Page, frame, page table

Virtual memory duoc chia thanh page.

Physical memory duoc chia thanh frame.

Page table map:

```text
virtual page number -> physical frame number
```

Dia chi ao co the hieu thanh:

```text
virtual address = virtual page number + offset
```

Sau khi tim duoc physical frame, offset giu nguyen de tao physical address.

### TLB

TLB la Translation Lookaside Buffer, cache nho trong CPU de cache mapping:

```text
virtual page -> physical frame
```

Neu TLB hit, CPU tranh duoc page-table walk ton kem.

Neu TLB miss, CPU/MMU phai di qua page table de tim mapping, cham hon.

Backend angle: memory access ngau nhien tren data structure lon co the gay cache miss/TLB miss nhieu hon so voi access lien tuc nhu array scan.

### Page fault

Page fault khong luon la crash. No la trap khi process access mot virtual page ma hien tai chua co mapping hop le trong RAM hoac violation permission.

Page fault hop le:

- Page chua duoc load.
- Page dang o swap.
- Copy-on-write: parent/child process share page, den khi mot ben write thi kernel copy page moi.
- Memory-mapped file chua load page tu disk.
- Lazy allocation: stack/heap page chi cap khi thuc su dung.

Kernel xu ly xong thi program co the chay tiep.

Page fault khong hop le:

- Access address khong thuoc process.
- Write vao read-only page.
- Execute page khong co execute permission.
- Dereference invalid pointer.

Ket qua co the la segmentation fault.

### Segmentation fault la gi?

Segmentation fault la khi process access memory khong duoc phep. No thuong den tu invalid pointer, out-of-bounds, use-after-free trong ngon ngu manual memory, hoac stack overflow.

Trong managed languages nhu Go/Java/C#, ban it gap segfault tu code thuong, nhung van co the gap khi dung native code, unsafe, cgo/JNI, hoac runtime bug.

### Swap

Swap la khi OS dua mot so memory pages ra disk de giai phong RAM.

Tradeoff:

- Giup he thong khong fail ngay khi RAM thieu.
- Rat cham so voi RAM.
- Neu swap qua nhieu, he thong bi thrashing: CPU mat thoi gian doi disk I/O hon la chay logic.

---

## 5. Stack va Heap

### Stack

Stack la vung memory cho function calls va local execution state. No theo LIFO.

Moi function call tao mot stack frame. Stack frame co the chua:

- Function parameters.
- Return address.
- Saved registers.
- Local variables.
- Frame metadata.

Allocation tren stack rat nhanh, thuong chi can dich stack pointer.

### Heap

Heap la vung memory cho object/data co lifetime linh hoat hon function call.

Heap dung cho:

- Object can song sau khi function return.
- Data size lon hoac dynamic.
- Shared data giua goroutines/threads.
- Object duoc reference tu nhieu noi.

Allocator phai quan ly free blocks, metadata, fragmentation. Trong ngon ngu co GC, runtime con phai tim object nao con reachable de reclaim object khong con dung.

### Stack vs heap - cau tra loi phong van

> Stack memory is used for function call frames and local execution state. It is fast and automatically freed when the function returns. Heap memory is used for data with more flexible lifetime, especially data that may outlive a function call or be shared. Heap allocation is more flexible but usually has more overhead and can involve fragmentation or garbage collection.

### Stack overflow

Stack overflow xay ra khi stack vuot qua gioi han.

Nguyen nhan:

- Recursion qua sau.
- Base case sai hoac khong dat duoc.
- Local variable qua lon.
- Qua nhieu nested calls.

Vi du recursion khong dung:

```text
func f() {
  f()
}
```

### Heap fragmentation

Fragmentation la khi memory free bi chia thanh nhieu khoang nho roi rac.

Co the co tong free memory du, nhung khong co block lien tuc du lon cho allocation can thiet.

Managed runtimes co the dung compaction de giam fragmentation, nhung compaction cung co overhead.

### Memory leak

Memory leak la memory khong con huu ich nhung khong duoc reclaim.

Trong C/C++:

- Allocate ma quen free.
- Mat pointer den allocated memory.

Trong GC languages:

- Object van reachable nen GC khong reclaim, du logic khong can nua.
- Cache/map khong xoa key cu.
- Goroutine/thread bi leak va giu reference.
- Listener/callback khong unregister.
- Slice reference giu backing array lon.

Cau quan trong:

> Garbage collection prevents many manual free bugs, but it does not prevent logical leaks. If an object is still reachable from a global map, cache, goroutine, or active reference, GC must keep it.

### Go escape analysis

Trong Go, local variable khong phai luc nao cung nam tren stack. Compiler quyet dinh dua variable len stack hay heap dua tren lifetime va escape analysis.

Neu variable co the song lau hon stack frame hien tai, no co the escape len heap.

Vi du:

```go
func newUser() *User {
    u := User{Name: "Toan"}
    return &u
}
```

`u` la local variable, nhung vi return pointer ra ngoai function, no phai song sau khi function return. Compiler co the allocate `u` tren heap.

Nhung khong phai cu dung pointer la chac chan len heap. Neu compiler chung minh pointer khong escape khoi function, no van co the o stack.

### Binary heap khac heap memory

Dung nham:

- Heap memory: vung memory allocation.
- Binary heap / priority queue: data structure de lay min/max nhanh.

Cau ngan:

> They share the word "heap" but are different concepts. Heap memory is a runtime memory region. A binary heap is a tree-like data structure used for priority queues.

---

## 6. Interview Drill Answers

### Process vs thread

> A process is an isolated running program with its own virtual address space and OS resources. A thread is a lighter execution unit inside a process. Threads share heap, globals, code, and open files, but each thread has its own stack, registers, and program counter. The tradeoff is that threads are cheaper for communication, but shared memory requires synchronization.

### Context switch

> A context switch happens when the OS pauses one running thread or process and resumes another. It saves and restores CPU state such as registers, stack pointer, and program counter. A process switch can also involve changing address space context, which may hurt TLB and cache locality. It does not mean copying the whole RAM.

### Race condition

> A race condition happens when program correctness depends on unpredictable timing between concurrent executions. For example, `counter++` is usually read-modify-write, so two threads can read the same old value and overwrite each other. I would prevent it with a mutex, atomic operation, channel ownership, or database transaction depending on the state being protected.

### Mutex vs semaphore

> A mutex is for mutual exclusion and usually has ownership: one thread locks it and unlocks it. A semaphore is a counter of permits, so it can allow up to N workers at once or be used as a signal. I use a mutex to protect shared state and a semaphore to limit concurrency, like API calls or connection usage.

### Atomic vs mutex

> Atomic is good for small independent state like counters and flags. Mutex is better when multiple fields must be updated together or when there is a business invariant. In production code, I prefer the simpler primitive that makes correctness easiest to review.

### Deadlock

> Deadlock is when tasks wait on each other forever, such as one thread holding lock A and waiting for B while another holds B and waits for A. The four Coffman conditions are mutual exclusion, hold and wait, no preemption, and circular wait. A common prevention is locking resources in a consistent order and keeping critical sections small.

### Starvation vs livelock

> Starvation means one task waits too long because others keep getting the resource. Livelock means tasks are still active and reacting, but no useful progress is made. Deadlock is different because the tasks are stuck waiting and stop progressing.

### Virtual memory

> Virtual memory gives each process its own address space. The MMU translates virtual addresses to physical addresses using page tables, and the TLB caches recent translations. This provides isolation, protection, lazy allocation, shared mappings, and swap support.

### Page fault

> A page fault means the process accessed a virtual page that is not currently mapped in a usable way. It can be normal, for example lazy loading, swap, copy-on-write, or memory-mapped files. If the access is invalid, like writing to a read-only page or accessing an unmapped address, the OS may terminate the process with a segmentation fault.

### Stack vs heap

> Stack is for function call frames and local execution state. It is fast and automatically reclaimed when a function returns. Heap is for data with flexible lifetime, especially objects that outlive a function or are shared. Heap is more flexible but has allocation overhead, fragmentation risk, and possibly garbage collection cost.

### Memory leak in GC language

> A garbage collector reclaims unreachable objects, but it cannot reclaim objects that are still referenced. So leaks can still happen through global maps, unbounded caches, goroutines, callbacks, or slices that keep large backing arrays alive.

### Go escape analysis

> In Go, a local variable does not always live on the stack. The compiler uses escape analysis. If a value may outlive the current function, such as returning a pointer to it, the compiler can allocate it on the heap. But using a pointer does not automatically mean heap allocation if the compiler proves it does not escape.

---

## 7. Common Follow-Up Traps

### Trap: "Thread share stack, right?"

No. Threads in the same process share address space, heap, globals, code, and open files. Each thread has its own stack and CPU execution state.

### Trap: "Context switch reloads RAM?"

No. It saves/restores CPU state and may switch address space context. RAM is not copied wholesale.

### Trap: "`counter++` is atomic?"

Usually no. It is commonly read-modify-write. Use atomic increment or lock.

### Trap: "Semaphore is just a mutex?"

A binary semaphore can look like a mutex, but conceptually semaphore is a permit counter and may not have the same ownership semantics.

### Trap: "Page fault means crash?"

No. Many page faults are normal and handled by the kernel. Invalid page faults can crash.

### Trap: "GC means no memory leak?"

No. GC only collects unreachable objects. Reachable-but-useless objects are logical leaks.

### Trap: "Local variable always stack?"

No. In Go and many optimized languages, compiler/runtime decides. If it escapes, it may go to heap.

