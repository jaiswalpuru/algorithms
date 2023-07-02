""""
Asyncio : https://markuseliasson.se/article/introduction-to-asyncio/

Note : concurrency is not parallelism

Concurrency is when more than one function can be started and finished, overlapping each other without,
having to be executed at exact same time. This is possible with single core CPU.

Parallelism is when one or more functions run at same time, this requires multi-core CPU

Single threaded, asynchronous programming is considered simpler than using multi-threaded programming.
The reason is that you don't need to coordinate between routines and mutable states.


asyncio gives us asynchronous I/O, thus it is suitable for file and network operations where the process will be
schedule to wait for data being available.


											asynch  and  await
			Opening two TCP connections to Google's DNS server. Once the connection is open, do something ,
			once the work is done conn will be closed.
"""

import asyncio
from random import randint


async def do_something(ip, port):
	print(f'opening connection to IP {ip} on port {port}')
	reader, writer = await asyncio.open_connection(ip, port)
	print(f'connection open to {ip}')
	await asyncio.sleep(randint(0, 5))

	writer.close()
	print(f'connection closed to {ip}')

"""
the function do_something is declared with async def statement, which makes it into something called coroutine.
a coroutine is a special kind of generator and to which you can send the value back. 
The nice thing about coroutines is that they can be suspended, and resumed later. 

Whenever a await is reached inside do_something that coroutine will be suspended until the response is ready. 
When resumed it will continue execution at the same position and keep going until another await is there - or until
return statement is reached.

When a coroutine is suspended the event loop will resume another coroutine(given that it is already resumed.) This is
cooperative multi-tasking - coroutines will take turns running on same thread, but two coroutines will not run in 
parallel.

Basically any function that supports asyncio is either declared using "async def" or by using the decorator,
@asyncio.coroutine. Any code calling such functions need to use await statement or yield from.

Awaitable functions (or coroutines) need to be wrapped in a Future and handed over to the even loop.   
"""


if __name__ == "__main__":
	loop = asyncio.get_event_loop()
	work = [asyncio.ensure_future(do_something('8.8.8.8', '53')), asyncio.ensure_future(do_something('8.8.4.4', '53'))]
	loop.run_until_complete(asyncio.gather(*work))