# Pintl

Pintl is a easy to use [Load Balancer](https://en.wikipedia.org/wiki/Load_balancing_(computing)#Server-side_load_balancers) created to help small projects distribute a set of tasks over a set of resources, with the aim of making their overall processing more efficient.

## How to use

To run the Load Balancer you should use the following command:

```bash
go run *.go
```

## Features

* [X] Prevents clients from contacting back-end servers directly.
* [X] Static load distribution without prior knowledge [Round-Robin](https://en.wikipedia.org/wiki/Round-robin_scheduling).
* [X] It only sends a request if the server is online, otherwise it does not direct the request to a server that is offline.
* [X] Checks the availability of the server every 2 seconds.
* [X] Shows a log if the server is offline.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
