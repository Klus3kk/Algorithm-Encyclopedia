# **Job Sequencing with Deadlines**

## **Overview:**

The Job Sequencing Problem involves scheduling jobs to maximize profit, where each job has a deadline and profit. Jobs can be executed within a specific time frame.

## **Python Example:**

```python
def job_sequencing(jobs):
    jobs.sort(key=lambda x: x[1], reverse=True)  # Sort jobs by profit
    max_deadline = max(job[0] for job in jobs)
    result = [-1] * max_deadline
    total_profit = 0

    for

 job in jobs:
        deadline, profit = job
        for slot in range(min(deadline, max_deadline) - 1, -1, -1):
            if result[slot] == -1:
                result[slot] = profit
                total_profit += profit
                break

    return total_profit, result

# Example usage
jobs = [(4, 20), (1, 10), (1, 40), (1, 30)]
total_profit, sequence = job_sequencing(jobs)
print("Maximum Profit:", total_profit)
print("Job Sequence:", sequence)
```

## **Explanation:**
- **Sort Jobs:** Jobs are sorted by profit in descending order.
- **Schedule Jobs:** Jobs are placed into the latest available slot before their deadline.

