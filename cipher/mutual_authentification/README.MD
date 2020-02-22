# Mutual Authentication

-> Mutual authentication is the process by which a client and server authenticate each other.
They do this with public-key cryptography; both the client and server generate public/private key pairs,
exchange public keys, and use the public keys to validate the authenticity and identity of the other endpoint.
To accomplish this feat, both the client and server must do some legwork to set up the authorization,
explicitly defining the public key value with which they intend to validate the other.
The downside to this process is the administrative overhead of having to create unique key pairs for every single node
and ensuring that the server and the client nodes have the appropriate data to proceed properly.


### - Step 1 :

To begin, you’ll knock out the administrative tasks of creating key pairs. You’ll store the public keys as self-signed, PEM-encoded certificates. Let’s use the openssl utility to create these files. On your server, you’ll create the server’s private key and certificate by entering the following:

#####$ openssl req -nodes -x509 -newkey rsa:4096 -keyout serverKey.pem -out serverCrt.pem -days 365

- The file serverKey.pem contains your private key, and you should protect it. 
- The serverCrt.pem file contains the server’s public key, 

### * For every connecting client, you’ll run
#####$ openssl req -nodes -x509 -newkey rsa:4096 -keyout clientKey.pem -out clientCrt.pem -days 365

This command also generates two files: clientKey.pem and clientCrt.pem. Much as with the server output, you should protect the client’s private
key. The clientCrt.pem certificate file will be transferred to your server and loaded by your program. 
This will allow you to configure and identify the client as an authorized endpoint. You’ll have to create,
transfer, and con- figure a certificate for each additional client so that the server can identify and explicitly authorize them.

### - Step 2 :

you set up an HTTPS server that requires a client to provide a legitimate, authorized certificate.

You first read the client’s certificate from the PEM file the client created previouslyw. Because it’s possible to have more than one authorized client certificate, you create a certificate pool and call pool.AppendCertsFromPEM(clientCert) to add the cli- ent certificate to your poolx. You perform this step for each additional client you wish to authenticate.
Next, you create your TLS configuration. You explicitly set the ClientCAs field to your pool and configure ClientAuth to tls.RequireAndVerifyClientCert y.

#####NB : Note that you don’t use the client’s private-key file anywhere in the server code. As we’ve said before, the private key remains private; your server will be able to identify and authorize clients by using only the client’s public key. This is the brilliance of public-key crypto.

### - Last Step : 

-> Testing Commands:

Bad Command : curl -ik -X GET --cert badCrt.pem --key badKey.pem https://localhost:9443/hello

Good Command :  curl -ik -X GET --cert clientCrt.pem --key clientKey.pem https://localhost:9443/hello
