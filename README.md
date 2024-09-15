# Ollama D-Bus

Ollama D-Bus is just my experimental implementation of a GNOME search provider.
It uses the **Ollama Gemma2:2b model** to provide prompted search results in GNOME environments.

This project is primarily for experimentation and learning purposes.

## Features

- Implements GNOME search provider interface.
- Connects to the session bus using D-Bus.
- Provides search results based on user queries using the Ollama Gemma2:2b model.
- Uses a debounce mechanism to manage search queries efficiently.
- Logs activities and errors for monitoring and debugging.

## Installation

### Prerequisites

- Go version 1.22.6 or higher.
- Access to the GNOME desktop environment.
- D-Bus installed and configured.
- Ollama and gemma2:2b model pulled
- Make

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/dmytrogajewski/ollama-dbus.git
   cd ollama-dbus
   ```

2. Install dependencies:
   ```bash
   make deps
   ```

3. Build the application:
   ```bash
   make build
   ```

## Usage

Run the application with the following command:
```bash
make run
```

The application will connect to the session bus and start serving search requests. It logs information and errors to the standard output.

## Configuration

The application uses several Go packages to provide its functionality, including:

- `github.com/godbus/dbus/v5` for D-Bus communication.
- `github.com/urfave/cli/v2` for command-line interface management.
- `github.com/xyproto/ollamaclient/v2` for handling search queries with the Ollama Gemma2:2b model.

## Logging

Ollama D-Bus uses a custom logger to log messages. The logger is initialized in the `logger` package and logs messages to the standard output.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your changes. Ensure that your code follows the project's coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Contact

For questions or support, please contact the maintainer at [email@example.com].

Sources:
[1] https://developer.gnome.org/documentation/tutorials/search-provider.html
