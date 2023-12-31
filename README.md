# smashingthestatemachine
Proof-of-concept code for Smashing the state machine: the true potential of web race conditions from James Kettle (https://portswigger.net/research/smashing-the-state-machine).

This code sets up an HTTP/2 client, sends a request without the last byte, waits for a short duration, and then sends the withheld byte. It's a basic representation of the single-packet attack concept.

The article discusses the untapped potential of web race condition attacks. Historically, these attacks have been limited to a few scenarios due to challenges like tricky workflows, lack of proper tools, and network jitter. The author introduces new classes of race conditions that can exploit multiple high-profile websites and Devise, a popular authentication framework for Rails. The article also presents the "single-packet attack," a strategy that can send multiple requests in a very short time frame. This research was presented at various conferences, including Black Hat USA, DEF CON, and Nullcon.

