import asyncio

async def add_numbers(a, b):
    return a + b

async def main():
    # Add numbers concurrently using async/await
    task1 = asyncio.create_task(add_numbers(5, 10))
    task2 = asyncio.create_task(add_numbers(15, 20))

    # Await the tasks to complete and get the results
    result1 = await task1
    result2 = await task2

    total_sum = result1 + result2
    print("Total Sum:", total_sum)

# Run the async main function
asyncio.run(main())



