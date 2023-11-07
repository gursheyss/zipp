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
![diagram-export-11-7-2023-12_27_04-PM](https://github.com/gursheyss/zipp/assets/116788218/224873d8-9773-4ee7-8a73-739654465c31)

### File Download
- Access the file via a link with a unique ID.
- Enter the password for file decryption; authentication is checked against database entries.
- If the password is correct, the file is downloaded from S3.
- File is decrypted using AES and PBKDF2 and then provided to the user.
- Post-download, the file's record is deleted from both MySQL database and S3 storage. 

### Diagram
![diagram-export-11-7-2023-12_26_31-PM](https://github.com/gursheyss/zipp/assets/116788218/9b1fe1fc-4202-424f-9d09-383e4a8ea9f9)

## License
Zipp is free software: you can redistribute it and/or modify it under the terms of the MIT Public License.
