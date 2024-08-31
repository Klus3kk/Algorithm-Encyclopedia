# **Cryptography Algorithms**

## **Overview:**

Cryptography Algorithms ensure secure communication by encrypting data. They include classical and modern encryption methods.

## **Subtopics:**

- **RSA Algorithm:** Public key encryption algorithm.
- **AES Algorithm:** Symmetric key encryption algorithm.
- **Hash Functions:** Ensure data integrity.

## **Python Example (RSA Encryption):**

```python
from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_OAEP

def rsa_encryption(message):
    key = RSA.generate(2048)
    cipher = PKCS1_OAEP.new(key)
    encrypted_message = cipher.encrypt(message.encode())
    return encrypted_message, key.export_key()

# Example usage
message = "Hello, World!"
encrypted_message, key = rsa_encryption(message)
print("Encrypted message:", encrypted_message)
```

## **Explanation:**
- **RSA Algorithm:** Uses a pair of keys (public and private) for encryption and decryption, ensuring secure data transmission.

