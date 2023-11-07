# Zipp
 A simple and secure file upload and download web application built with Svelte and Go. 

## How It Works
### File Upload
- Visit the main upload page to submit a file.
- Optionally add a password; if none provided, one is auto-generated.
- Receive a unique ID upon upload for future retrieval.
- File encrypted using AES and PBKDF2, then uploaded to S3 with ID.
- ID, nonce, and password salt stored in MySQL database.
- A link containing the unique ID is returned to the user for access.
### Diagram
[![File Upload](https://app.eraser.io/workspace/AVKbr29QKlgqH6a7YsCd/preview?elements=9rWSIsSbwoWs-0VMYQYBwQ&type=embed)](https://app.eraser.io/workspace/AVKbr29QKlgqH6a7YsCd?elements=9rWSIsSbwoWs-0VMYQYBwQ)

### File Download
- Access the file via a link with a unique ID.
- Enter the password for file decryption; authentication is checked against database entries.
- If the password is correct, the file is downloaded from S3.
- File is decrypted using AES and PBKDF2 and then provided to the user.
- Post-download, the file's record is deleted from both MySQL database and S3 storage. 

### Diagram
[![File Download](https://app.eraser.io/workspace/AVKbr29QKlgqH6a7YsCd/preview?elements=es5yt7geSqpiH4TxXDWhrg&type=embed)](https://app.eraser.io/workspace/AVKbr29QKlgqH6a7YsCd?elements=es5yt7geSqpiH4TxXDWhrg)

## License
Zipp is free software: you can redistribute it and/or modify it under the terms of the MIT Public License.

