# GOHIDE

GOHIDE is a command-line tool for file encryption, decryption, and directory management, written in Go. It offers a simple and efficient way to secure sensitive files and manage hidden directories. The commands are triggered by the first letter of the command name.

\<br\>

-----

## 🚀 Getting Started

### Installation

To add the GOHIDE binary to your system's PATH, follow these steps:

1.  Download the latest binary for your operating system from the [GitHub releases page](https://www.google.com/search?q=https://github.com/your-username/gohide/releases).
2.  Move the binary to a directory that is already in your system's PATH. On Unix-like systems, this is often `/usr/local/bin`. You might need administrator privileges.
    ```sh
    sudo mv gohide /usr/local/bin/
    ```
3.  Ensure the binary is executable.
    ```sh
    sudo chmod +x /usr/local/bin/gohide
    ```

### Building from Source

To build GOHIDE from the source code, you must have [Go](https://golang.org/doc/install) installed on your system.

1.  Clone the repository:
    ```sh
    git clone https://github.com/your-username/gohide.git
    cd gohide
    ```
2.  Build the binary:
    ```sh
    go build -o gohide
    ```
3.  The executable `gohide` will be created in the current directory. You can then move it to a directory in your PATH as described above.

\<br\>

-----

## 📚 Usage

GOHIDE's commands are single-character abbreviations. The tool will prompt you for a password when required.

### 🔑 Encryption & Decryption

  - **Encrypt a file** (`e`): Encrypts a source file and saves it to a destination file.

    ```sh
    gohide e <source_file> <destination_file>
    ```

  - **Decrypt a file** (`d`): Decrypts an encrypted file back to its original form.

    ```sh
    gohide d <source_file> <destination_file>
    ```

    🚨 **Warning:** If you forget the password, there is no way to recover your data. The encryption is irreversible without the correct password. Also, reusing a similar password will result in a different encryption, ensuring security is not compromised.

### 📁 Directory Operations

  - **Hide a directory** (`h`): Compresses a directory and encrypts the resulting archive.

    ```sh
    gohide h <source_directory>
    ```

  - **Show a hidden directory** (`s`): Decrypts and extracts a hidden directory archive.

    ```sh
    gohide s <source_directory>
    ```

  - **Zip a directory** (`z`): Compresses a directory into a ZIP archive and then encrypts it.

    ```sh
    gohide z <source_directory> <destination_zip>
    ```

  - **Check a directory** (`c`): Verifies the integrity of a hidden directory archive.

    ```sh
    gohide c <source_directory>
    ```

### 🚫 Ignoring Files

GOHIDE supports a `.hideignore` file to exclude specific files or directories from the `hide` command.

  - To ignore a directory named `safe`, add `safe` as a new line in `.hideignore`.
  - To ignore all files with the `.go` or `.cpp` extensions, add `.go` or `.cpp` on separate lines in the `.hideignore` file.

\<br\>

-----

## 🛡️ Security

GOHIDE leverages the `crypto` package from the Go standard library to perform all cryptographic operations. This ensures that the encryption and decryption processes are secure and reliable.

## 🏷️ Tags

`go`, `cli`, `encryption`, `decryption`, `hide`, `security`, `utility`, `cryptography`, `command-line-tool`, `github`