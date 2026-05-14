# Build HTTP Endpoints & Start Web Server

1. **Starting a Local Web Server**
   
   To make your Go application accessible over the web, you need to start a local server that "listens" for incoming requests.

   - **Package:** Uses the built-in `net/http` package.

   - **http.ListenAndServe():** This function starts the server. It takes two parameters:
     1. Address: The network address and port (e.g., `:8080`). Localhost refers to your own computer.
     2. Handler: Usually set to `nil` to use the default router (ServeMux).

   - **Behavior:** Once executed, the program does not stop; it stays running in the background, waiting for user requests.


2. **Handling HTTP Requests (Endpoints)**
   
   An "endpoint" is a specific URL path where your application logic lives. You define which function handles which path using Handlers.

   - **http.HandleFunc():** Registers a specific function to be executed when a URL path is accessed.

   **Example:** `http.HandleFunc("/hello", helloUser)` tells Go to run the `helloUser` function whenever someone visits `localhost:8080/hello`.


3. **Creating a Handler Function**
   
   A handler function must follow a specific signature to process incoming data and send back a response.

   - **Parameters:**
     - `w http.ResponseWriter`: Used to write the response back to the client (like text or JSON).
     - `r *http.Request`: Contains information about the incoming request from the user.

   - **Sending a Response:** Use `fmt.Fprintf(w, "your message")` to send text directly to the user's browser.


# Practical Implementation: TodoList API

1. **Global Variable:** Define the `taskItems` list outside the functions (Global Scope) so the handlers can access them.

2. **Show Tasks Endpoint:** Create a handler (e.g., `/show-tasks`) that loops through the slice and writes each task to the ResponseWriter.

3. **Iteration:** Use a `for range` loop inside the handler to print the tasks one by one to the web client.


# Interacting with the API

Once the server is running, you can use various HTTP Clients to see the results:

- **Web Browser:** Visit `localhost:8080/show-tasks` to see the list in a UI.

- **Curl (CLI):** Run `curl localhost:8080/show-tasks` in your terminal for a command-line response.

- **IDE Client:** Use built-in tools like the GoLand HTTP Client to test requests directly within your editor.


# Summary Table: HTTP Package Functions

| Function | Purpose | Example |
|----------|---------|----------|
| `http.ListenAndServe` | Starts the web server on a port | `http.ListenAndServe(":8080", nil)` |
| `http.HandleFunc` | Maps a URL path to a function | `http.HandleFunc("/tasks", showTasks)` |
| `fmt.Fprintf` | Sends data to the ResponseWriter | `fmt.Fprintf(w, "Hello World")` |