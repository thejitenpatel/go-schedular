# Go Scheduler Library

Go Scheduler is a simple task scheduling library that allows you to manage the execution of tasks with a configurable limit on the number of concurrent tasks. It provides a way to queue tasks when the concurrency limit is reached and execute them as soon as resources become available.

## Features

* **Configurable Concurrency:** Set the maximum number of tasks that can run concurrently.
* **Task Queueing:** Tasks are queued if the concurrency limit is reached and are executed as soon as capacity allows.
* **Automatic Scheduling:** When the number of active tasks drops below the concurrency limit, new tasks are automatically scheduled from the queue.
* **Thread-safe:** The scheduler uses synchronization mechanisms to ensure thread-safety when updating the state of tasks.
